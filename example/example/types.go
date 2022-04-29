package example

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Tag describes a resource tag.
type Tag struct {
	// The key for the tag.
	Key string `pulumi:"key"`
	// The Value for the tag.
	Value string `pulumi:"value"`
}

type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
}

type TagArrayOutput struct{ *pulumi.OutputState }

func (TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (o TagArrayOutput) ToTagArrayOutput() TagArrayOutput {
	return o
}

type TagArrayArgs []TagInput

func (TagArrayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (i TagArrayArgs) ToTagArrayOutput() TagArrayOutput {
	return pulumi.ToOutput(i).(TagArrayOutput)
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

type TagArgs struct {
	Key   pulumi.StringInput
	Value pulumi.StringInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (i TagArgs) ToTagOutput() TagOutput {
	return pulumi.ToOutput(i).(TagOutput)
}
