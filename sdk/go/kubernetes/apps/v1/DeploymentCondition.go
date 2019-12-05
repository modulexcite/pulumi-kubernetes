// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// DeploymentCondition describes the state of a deployment at a certain point.
type DeploymentCondition struct {
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *string `pulumi:"lastTransitionTime"`

	// The last time this condition was updated.
	LastUpdateTime *string `pulumi:"lastUpdateTime"`

	// A human readable message indicating details about the transition.
	Message *string `pulumi:"message"`

	// The reason for the condition's last transition.
	Reason *string `pulumi:"reason"`

	// Status of the condition, one of True, False, Unknown.
	Status string `pulumi:"status"`

	// Type of deployment condition.
	Type string `pulumi:"type"`

}

var _DeploymentConditionType = reflect.TypeOf((*DeploymentCondition)(nil)).Elem()

// DeploymentConditionInput represents an input type that resolves to a DeploymentCondition.
type DeploymentConditionInput interface {
	ElementType() reflect.Type

	ToDeploymentConditionOutput() DeploymentConditionOutput
	ToDeploymentConditionOutputWithContext(ctx context.Context) DeploymentConditionOutput
}

// DeploymentConditionArgs is a DeploymentConditionInput whose fields are all Input types.
type DeploymentConditionArgs struct {
	// Status of the condition, one of True, False, Unknown.
	Status pulumi.StringInput `pulumi:"status"`

	// Type of deployment condition.
	Type pulumi.StringInput `pulumi:"type"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime pulumi.StringInput `pulumi:"lastTransitionTime"`

	// The last time this condition was updated.
	LastUpdateTime pulumi.StringInput `pulumi:"lastUpdateTime"`

	// A human readable message indicating details about the transition.
	Message pulumi.StringInput `pulumi:"message"`

	// The reason for the condition's last transition.
	Reason pulumi.StringInput `pulumi:"reason"`

}

func (a DeploymentConditionArgs) ElementType() reflect.Type {
	return _DeploymentConditionType
}

func (a DeploymentConditionArgs) ToDeploymentConditionOutput() DeploymentConditionOutput {
	return pulumi.ToOutput(a).(DeploymentConditionOutput)
}

func (a DeploymentConditionArgs) ToDeploymentConditionOutputWithContext(ctx context.Context) DeploymentConditionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentConditionOutput)
}

// DeploymentConditionOutput is an output type that resolves to a Input.
type DeploymentConditionOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DeploymentConditionOutput{}) }

func (DeploymentConditionOutput) ElementType() reflect.Type {
	return _DeploymentConditionType
}

func (o DeploymentConditionOutput) LastTransitionTime() pulumi.StringOutput {
	return o.Apply(func(v DeploymentCondition) *string {
		return v.LastTransitionTime
	}).(pulumi.StringOutput)
}

func (o DeploymentConditionOutput) LastUpdateTime() pulumi.StringOutput {
	return o.Apply(func(v DeploymentCondition) *string {
		return v.LastUpdateTime
	}).(pulumi.StringOutput)
}

func (o DeploymentConditionOutput) Message() pulumi.StringOutput {
	return o.Apply(func(v DeploymentCondition) *string {
		return v.Message
	}).(pulumi.StringOutput)
}

func (o DeploymentConditionOutput) Reason() pulumi.StringOutput {
	return o.Apply(func(v DeploymentCondition) *string {
		return v.Reason
	}).(pulumi.StringOutput)
}

func (o DeploymentConditionOutput) Status() pulumi.StringOutput {
	return o.Apply(func(v DeploymentCondition) string {
		return v.Status
	}).(pulumi.StringOutput)
}

func (o DeploymentConditionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v DeploymentCondition) string {
		return v.Type
	}).(pulumi.StringOutput)
}

var _DeploymentConditionArrayType = reflect.TypeOf((*[]DeploymentCondition)(nil)).Elem()

type DeploymentConditionArrayInput interface {
	ElementType() reflect.Type

	ToDeploymentConditionArrayOutput() DeploymentConditionArrayOutput
	ToDeploymentConditionArrayOutputWithContext(ctx context.Context) DeploymentConditionArrayOutput
}

type DeploymentConditionArray []DeploymentConditionInput

func (a DeploymentConditionArray) ElementType() reflect.Type {
	return _DeploymentConditionArrayType
}

func (a DeploymentConditionArray) ToDeploymentConditionArrayOutput() DeploymentConditionArrayOutput {
	return pulumi.ToOutput(a).(DeploymentConditionArrayOutput)
}

func (a DeploymentConditionArray) ToDeploymentConditionArrayOutputWithContext(ctx context.Context) DeploymentConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentConditionArrayOutput)
}

type DeploymentConditionArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(DeploymentConditionArrayOutput{}) }

func (DeploymentConditionArrayOutput) ElementType() reflect.Type {
	return _DeploymentConditionArrayType
}

func (o DeploymentConditionArrayOutput) Index(i pulumi.IntInput) DeploymentConditionOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DeploymentCondition {
		return vs[0].([]DeploymentCondition)[vs[1].(int)]
	}).(DeploymentConditionOutput)
}