package example

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

type ResourceArgs struct {
	Tags []Tag `pulumi:"tags,optional"`
}

// This is a Pulumi resource.
//
//pulumi:resource
type Resource struct {
	// These are the tags for the resource.
	Tags []Tag `pulumi:"tags,optional"`
}

func (r *Resource) Create(p *Provider, args *ResourceArgs, options interface{}) (resource.ID, error) {
	r.Tags = args.Tags
	return "id", nil
}

func (*Resource) Read(p *Provider, id resource.ID, options interface{}) error {
	return nil
}

func (*Resource) Update(p *Provider, id resource.ID, args *ResourceArgs, options interface{}) error {
	return nil
}

func (*Resource) Delete(p *Provider, id resource.ID, options interface{}) error {
	return nil
}
