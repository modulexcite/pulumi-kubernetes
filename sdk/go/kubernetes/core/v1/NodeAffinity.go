// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Node affinity is a group of node affinity scheduling rules.
type NodeAffinity struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions
	// specified by this field, but it may choose a node that violates one or more of the expressions.
	// The node that is most preferred is the one with the greatest sum of weights, i.e. for each node
	// that meets all of the scheduling requirements (resource request, requiredDuringScheduling
	// affinity expressions, etc.), compute a sum by iterating through the elements of this field and
	// adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s)
	// with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTerm `pulumi:"preferredDuringSchedulingIgnoredDuringExecution"`

	// If the affinity requirements specified by this field are not met at scheduling time, the pod
	// will not be scheduled onto the node. If the affinity requirements specified by this field cease
	// to be met at some point during pod execution (e.g. due to an update), the system may or may not
	// try to eventually evict the pod from its node.
	RequiredDuringSchedulingIgnoredDuringExecution *NodeSelector `pulumi:"requiredDuringSchedulingIgnoredDuringExecution"`

}

var _NodeAffinityType = reflect.TypeOf((*NodeAffinity)(nil)).Elem()

// NodeAffinityInput represents an input type that resolves to a NodeAffinity.
type NodeAffinityInput interface {
	ElementType() reflect.Type

	ToNodeAffinityOutput() NodeAffinityOutput
	ToNodeAffinityOutputWithContext(ctx context.Context) NodeAffinityOutput
}

// NodeAffinityArgs is a NodeAffinityInput whose fields are all Input types.
type NodeAffinityArgs struct {
	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions
	// specified by this field, but it may choose a node that violates one or more of the expressions.
	// The node that is most preferred is the one with the greatest sum of weights, i.e. for each node
	// that meets all of the scheduling requirements (resource request, requiredDuringScheduling
	// affinity expressions, etc.), compute a sum by iterating through the elements of this field and
	// adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s)
	// with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution PreferredSchedulingTermArrayInput `pulumi:"preferredDuringSchedulingIgnoredDuringExecution"`

	// If the affinity requirements specified by this field are not met at scheduling time, the pod
	// will not be scheduled onto the node. If the affinity requirements specified by this field cease
	// to be met at some point during pod execution (e.g. due to an update), the system may or may not
	// try to eventually evict the pod from its node.
	RequiredDuringSchedulingIgnoredDuringExecution NodeSelectorInput `pulumi:"requiredDuringSchedulingIgnoredDuringExecution"`

}

func (a NodeAffinityArgs) ElementType() reflect.Type {
	return _NodeAffinityType
}

func (a NodeAffinityArgs) ToNodeAffinityOutput() NodeAffinityOutput {
	return pulumi.ToOutput(a).(NodeAffinityOutput)
}

func (a NodeAffinityArgs) ToNodeAffinityOutputWithContext(ctx context.Context) NodeAffinityOutput {
	return pulumi.ToOutputWithContext(ctx, a).(NodeAffinityOutput)
}

// NodeAffinityOutput is an output type that resolves to a Input.
type NodeAffinityOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(NodeAffinityOutput{}) }

func (NodeAffinityOutput) ElementType() reflect.Type {
	return _NodeAffinityType
}

func (o NodeAffinityOutput) PreferredDuringSchedulingIgnoredDuringExecution() PreferredSchedulingTermArrayOutput {
	return o.Apply(func(v NodeAffinity) []PreferredSchedulingTerm {
		return v.PreferredDuringSchedulingIgnoredDuringExecution
	}).(PreferredSchedulingTermArrayOutput)
}

func (o NodeAffinityOutput) RequiredDuringSchedulingIgnoredDuringExecution() NodeSelectorOutput {
	return o.Apply(func(v NodeAffinity) *NodeSelector {
		return v.RequiredDuringSchedulingIgnoredDuringExecution
	}).(NodeSelectorOutput)
}
