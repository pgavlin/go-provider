package pulumiapi

import (
	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate/client"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/provider"
)

type WhoAmIArgs struct{}

type WhoAmIResult struct {
	// The Identity of the current user.
	Identity string `pulumi:"identity"`
}

// WhoAmI returns the identity of the current user.
//
//pulumi:function
func WhoAmI(ctx *provider.Context, p *Provider, _ *WhoAmIArgs, options provider.CallOptions) (*WhoAmIResult, error) {
	identity, err := p.client.GetPulumiAccountName(ctx)
	if err != nil {
		return nil, err
	}
	return &WhoAmIResult{Identity: identity}, nil
}

// ListStacksTag describes a tag filter for ListStacks.
type ListStacksTag struct {
	// The name of the tag to match.
	Name string `pulumi:"name"`
	// The value of the tag to match.
	Value string `pulumi:"value"`
}

type ListStacksArgs struct {
	// Filter to stacks that belong to this organization.
	Organization string `pulumi:"organization,optional"`
	// Filter to stacks that belong to this project.
	Project string `pulumi:"project,optional"`
	// Filter to stacks with a matching tag.
	Tag *ListStacksTag `pulumi:"tag"`
}

type StackSummary struct {
	// The identity of the stack.
	Identifier StackIdentifier `pulumi:"identifier"`
	// The last update time of the stack, if any,
	LastUpdateTime *int64 `pulumi:"lastUpdateTime"`
	// The resource count of the stack, if any.
	ResourceCount *int `pulumi:"resourceCount"`
}

type ListStacksResult struct {
	Stacks []StackSummary `pulumi:"stacks"`
}

// ListStacks lists the stacks that match the given filters.
//
//pulumi:function
func ListStacks(ctx *provider.Context, p *Provider, args *ListStacksArgs, options provider.CallOptions) (*ListStacksResult, error) {
	filter := client.ListStacksFilter{}
	if args.Project != "" {
		filter.Project = &args.Project
	}
	if args.Organization != "" {
		filter.Organization = &args.Organization
	}
	if args.Tag != nil {
		filter.TagName, filter.TagValue = &args.Tag.Name, &args.Tag.Value
	}
	stacks, err := p.client.ListStacks(ctx, filter)
	if err != nil {
		return nil, err
	}

	result := make([]StackSummary, len(stacks))
	for i, s := range stacks {
		result[i] = StackSummary{
			Identifier: StackIdentifier{
				Organization: s.OrgName,
				Project:      s.ProjectName,
				Stack:        s.StackName,
			},
			LastUpdateTime: s.LastUpdate,
			ResourceCount:  s.ResourceCount,
		}
	}
	return &ListStacksResult{Stacks: result}, nil
}
