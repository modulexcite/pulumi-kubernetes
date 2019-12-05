// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A topology selector term represents the result of label queries. A null or empty topology
// selector term matches no objects. The requirements of them are ANDed. It provides a subset of
// functionality as NodeSelectorTerm. This is an alpha feature and may change in the future.
type TopologySelectorTerm struct {
	// A list of topology selector requirements by labels.
	MatchLabelExpressions []TopologySelectorLabelRequirement `pulumi:"matchLabelExpressions"`

}

var _TopologySelectorTermType = reflect.TypeOf((*TopologySelectorTerm)(nil)).Elem()

// TopologySelectorTermInput represents an input type that resolves to a TopologySelectorTerm.
type TopologySelectorTermInput interface {
	ElementType() reflect.Type

	ToTopologySelectorTermOutput() TopologySelectorTermOutput
	ToTopologySelectorTermOutputWithContext(ctx context.Context) TopologySelectorTermOutput
}

// TopologySelectorTermArgs is a TopologySelectorTermInput whose fields are all Input types.
type TopologySelectorTermArgs struct {
	// A list of topology selector requirements by labels.
	MatchLabelExpressions TopologySelectorLabelRequirementArrayInput `pulumi:"matchLabelExpressions"`

}

func (a TopologySelectorTermArgs) ElementType() reflect.Type {
	return _TopologySelectorTermType
}

func (a TopologySelectorTermArgs) ToTopologySelectorTermOutput() TopologySelectorTermOutput {
	return pulumi.ToOutput(a).(TopologySelectorTermOutput)
}

func (a TopologySelectorTermArgs) ToTopologySelectorTermOutputWithContext(ctx context.Context) TopologySelectorTermOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TopologySelectorTermOutput)
}

// TopologySelectorTermOutput is an output type that resolves to a Input.
type TopologySelectorTermOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(TopologySelectorTermOutput{}) }

func (TopologySelectorTermOutput) ElementType() reflect.Type {
	return _TopologySelectorTermType
}

func (o TopologySelectorTermOutput) MatchLabelExpressions() TopologySelectorLabelRequirementArrayOutput {
	return o.Apply(func(v TopologySelectorTerm) []TopologySelectorLabelRequirement {
		return v.MatchLabelExpressions
	}).(TopologySelectorLabelRequirementArrayOutput)
}

var _TopologySelectorTermArrayType = reflect.TypeOf((*[]TopologySelectorTerm)(nil)).Elem()

type TopologySelectorTermArrayInput interface {
	ElementType() reflect.Type

	ToTopologySelectorTermArrayOutput() TopologySelectorTermArrayOutput
	ToTopologySelectorTermArrayOutputWithContext(ctx context.Context) TopologySelectorTermArrayOutput
}

type TopologySelectorTermArray []TopologySelectorTermInput

func (a TopologySelectorTermArray) ElementType() reflect.Type {
	return _TopologySelectorTermArrayType
}

func (a TopologySelectorTermArray) ToTopologySelectorTermArrayOutput() TopologySelectorTermArrayOutput {
	return pulumi.ToOutput(a).(TopologySelectorTermArrayOutput)
}

func (a TopologySelectorTermArray) ToTopologySelectorTermArrayOutputWithContext(ctx context.Context) TopologySelectorTermArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TopologySelectorTermArrayOutput)
}

type TopologySelectorTermArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(TopologySelectorTermArrayOutput{}) }

func (TopologySelectorTermArrayOutput) ElementType() reflect.Type {
	return _TopologySelectorTermArrayType
}

func (o TopologySelectorTermArrayOutput) Index(i pulumi.IntInput) TopologySelectorTermOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TopologySelectorTerm {
		return vs[0].([]TopologySelectorTerm)[vs[1].(int)]
	}).(TopologySelectorTermOutput)
}