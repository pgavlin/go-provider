package pulumiapi

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate/client"
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/provider"
)

// StackIdentifier contains the information necessary to identify a Pulumi stack.
type StackIdentifier struct {
	// The organization that owns the stack.
	Organization string `pulumi:"organization"`
	// The project associated with the stack.
	Project string `pulumi:"project"`
	// The name of the stack.
	Stack string `pulumi:"stack"`
}

func (i StackIdentifier) String() string {
	return fmt.Sprintf("%v/%v/%v", i.Organization, i.Project, i.Stack)
}

// StackOperation decribes an operation being performed on a Pulumi stack.
type StackOperation struct {
	// The kind of operation being performed.
	Kind string `pulumi:"kind"`
	// The user who initiated the operation.
	Author string `pulumi:"author"`
	// The time at which the operation started.
	Started int64 `pulumi:"started"`
}

type StackArgs struct {
	// Identifier is the identifier for the stack.
	Identifier StackIdentifier `pulumi:"identifier,immutable"`
	// Tags are the tags for the stack, if any.
	Tags map[apitype.StackTagName]string `pulumi:"tags,optional"`
	// ForceDelete indicates whether or not the stack should be deleted even if it contains resources.
	ForceDelete bool `pulumi:"forceDelete,optional"`
}

// This is a Pulumi resource.
//
//pulumi:resource
type Stack struct {
	// Identifier is the identifier for the stack.
	Identifier StackIdentifier `pulumi:"identifier"`
	// Tags are the tags for the stack, if any.
	Tags map[apitype.StackTagName]string `pulumi:"tags,optional"`
	// ForceDelete indicates whether or not the stack should be deleted even if it contains resources.
	ForceDelete bool `pulumi:"forceDelete"`

	// CurrentOperation provides information about the stack operation that is in progress, if any.
	CurrentOperation *StackOperation `pulumi:"currentOperation"`
	// Version is the current version of the stack.
	Version int `pulumi:"version"`
}

func (s *Stack) Args() *StackArgs {
	return &StackArgs{
		Identifier:  s.Identifier,
		Tags:        s.Tags,
		ForceDelete: s.ForceDelete,
	}
}

func (s *Stack) Create(ctx *provider.Context, p *Provider, args *StackArgs, options provider.CreateOptions) (provider.ID, error) {
	clientIdentifier := client.StackIdentifier{
		Owner:   args.Identifier.Organization,
		Project: args.Identifier.Project,
		Stack:   args.Identifier.Stack,
	}
	stack, err := p.client.CreateStack(ctx, clientIdentifier, args.Tags)
	if err != nil {
		return "", err
	}

	s.Identifier = StackIdentifier{
		Organization: stack.OrgName,
		Project:      stack.ProjectName,
		Stack:        string(stack.StackName),
	}
	s.ForceDelete = args.ForceDelete
	s.Tags = stack.Tags
	s.Version = stack.Version

	if stack.CurrentOperation != nil {
		s.CurrentOperation = &StackOperation{
			Kind:    string(stack.CurrentOperation.Kind),
			Author:  stack.CurrentOperation.Author,
			Started: stack.CurrentOperation.Started,
		}
	}

	return provider.ID(s.Identifier.String()), nil
}

func (s *Stack) Read(ctx *provider.Context, p *Provider, id provider.ID, options provider.ReadOptions) error {
	components := strings.Split(string(id), "/")
	if len(components) != 3 {
		return fmt.Errorf("stack ID must be of the form 'organization/project/stack'")
	}

	clientIdentifier := client.StackIdentifier{
		Owner:   components[0],
		Project: components[1],
		Stack:   components[2],
	}
	stack, err := p.client.GetStack(ctx, clientIdentifier)
	if err != nil {
		return err
	}

	s.Identifier = StackIdentifier{
		Organization: stack.OrgName,
		Project:      stack.ProjectName,
		Stack:        string(stack.StackName),
	}
	s.Tags = stack.Tags
	s.Version = stack.Version

	if stack.CurrentOperation != nil {
		s.CurrentOperation = &StackOperation{
			Kind:    string(stack.CurrentOperation.Kind),
			Author:  stack.CurrentOperation.Author,
			Started: stack.CurrentOperation.Started,
		}
	}

	return nil
}

func (s *Stack) Update(ctx *provider.Context, p *Provider, id provider.ID, args *StackArgs, options provider.UpdateOptions) error {
	// TODO: update tags if there was any change

	s.ForceDelete = args.ForceDelete
	return nil
}

func (s *Stack) Delete(ctx *provider.Context, p *Provider, id provider.ID, options provider.DeleteOptions) error {
	clientIdentifier := client.StackIdentifier{
		Owner:   s.Identifier.Organization,
		Project: s.Identifier.Project,
		Stack:   s.Identifier.Stack,
	}
	_, err := p.client.DeleteStack(ctx, clientIdentifier, s.ForceDelete)
	return err
}

type GetStackArgs struct {
	Identifier StackIdentifier `pulumi:"identifier"`
}

//pulumi:function
func GetStack(ctx *provider.Context, p *Provider, args *GetStackArgs, options provider.CallOptions) (*Stack, error) {
	var s Stack
	if err := s.Read(ctx, p, provider.ID(args.Identifier.String()), provider.ReadOptions{}); err != nil {
		return nil, err
	}
	return &s, nil
}
