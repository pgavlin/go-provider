package example

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ComponentArgs struct {
	// Here's an input arg.
	SomeValue pulumi.IntInput `pulumi:"someValue"`

	// Here are some tags.
	Tags TagArrayInput `pulumi:"tags,optional"`
}

// This is a Pulumi component.
type Component struct {
	pulumi.ResourceState

	// Here's an output property.
	SomeOtherValue pulumi.IntOutput `pulumi:"someOtherValue"`

	// Here are the tag outputs.
	Tags TagArrayOutput `pulumi:"tags,optional"`
}

//pulumi:constructor
func NewComponent(ctx *pulumi.Context, name string, args *ComponentArgs, options ...pulumi.ResourceOption) (*Component, error) {
	return nil, nil
}

func init() {
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagOutput{})
}
