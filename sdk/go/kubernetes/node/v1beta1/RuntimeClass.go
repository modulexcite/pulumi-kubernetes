// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is
// used to determine which container runtime is used to run all containers in a pod. RuntimeClasses
// are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec.
// The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.
// For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClass struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Handler specifies the underlying runtime and configuration that the CRI implementation will use
	// to handle pods of this class. The possible values are specific to the node & CRI configuration.
	// It is assumed that all handlers are available on every node, and handlers of the same name are
	// equivalent on every node. For example, a handler called "runc" might specify that the runc OCI
	// runtime (using native Linux containers) will be used to run the containers in a pod. The Handler
	// must conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler pulumi.StringOutput `pulumi:"handler"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Overhead represents the resource overhead associated with running a pod for a given
	// RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level
	// as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	Overhead OverheadOutput `pulumi:"overhead"`

	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass
	// are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be
	// supported by all nodes.
	Scheduling SchedulingOutput `pulumi:"scheduling"`

}

// RuntimeClassArgs is the set of arguments needed to create a RuntimeClass resource.
type RuntimeClassArgs struct {
	// Handler specifies the underlying runtime and configuration that the CRI implementation will use
	// to handle pods of this class. The possible values are specific to the node & CRI configuration.
	// It is assumed that all handlers are available on every node, and handlers of the same name are
	// equivalent on every node. For example, a handler called "runc" might specify that the runc OCI
	// runtime (using native Linux containers) will be used to run the containers in a pod. The Handler
	// must conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler pulumi.StringInput `pulumi:"handler"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Overhead represents the resource overhead associated with running a pod for a given
	// RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level
	// as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	Overhead OverheadInput `pulumi:"overhead"`

	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass
	// are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be
	// supported by all nodes.
	Scheduling SchedulingInput `pulumi:"scheduling"`

}

// NewRuntimeClass creates a RuntimeClass resource with the given unique name, arguments, and options.
func NewRuntimeClass(ctx *pulumi.Context, name string, args *RuntimeClassArgs, opts ...pulumi.ResourceOption) (*RuntimeClass, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Handler == nil {
		return nil, errors.New("missing required argument 'Handler'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("node.k8s.io/v1beta1")
		args.Kind = pulumi.String("RuntimeClass")
		inputs["handler"] = args.Handler.ToStringOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Overhead; i != nil {
			inputs["overhead"] = i.ToOverheadOutput()
		}
		if i := args.Scheduling; i != nil {
			inputs["scheduling"] = i.ToSchedulingOutput()
		}
	}
	var resource RuntimeClass
	err := ctx.RegisterResource("kubernetes:node.k8s.io/v1beta1:RuntimeClass", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeClass gets an existing RuntimeClass resource's state with the given name and ID.
func GetRuntimeClass(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*RuntimeClass, error) {
	var resource RuntimeClass
	err := ctx.ReadResource("kubernetes:node.k8s.io/v1beta1:RuntimeClass", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is
// used to determine which container runtime is used to run all containers in a pod. RuntimeClasses
// are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec.
// The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.
// For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClassProperty struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Handler specifies the underlying runtime and configuration that the CRI implementation will use
	// to handle pods of this class. The possible values are specific to the node & CRI configuration.
	// It is assumed that all handlers are available on every node, and handlers of the same name are
	// equivalent on every node. For example, a handler called "runc" might specify that the runc OCI
	// runtime (using native Linux containers) will be used to run the containers in a pod. The Handler
	// must conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler string `pulumi:"handler"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Overhead represents the resource overhead associated with running a pod for a given
	// RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level
	// as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	Overhead *Overhead `pulumi:"overhead"`

	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass
	// are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be
	// supported by all nodes.
	Scheduling *Scheduling `pulumi:"scheduling"`

}

var _RuntimeClassPropertyType = reflect.TypeOf((*RuntimeClassProperty)(nil)).Elem()

// RuntimeClassPropertyInput represents an input type that resolves to a RuntimeClassProperty.
type RuntimeClassPropertyInput interface {
	ElementType() reflect.Type

	ToRuntimeClassPropertyOutput() RuntimeClassPropertyOutput
	ToRuntimeClassPropertyOutputWithContext(ctx context.Context) RuntimeClassPropertyOutput
}

// RuntimeClassPropertyArgs is a RuntimeClassPropertyInput whose fields are all Input types.
type RuntimeClassPropertyArgs struct {
	// Handler specifies the underlying runtime and configuration that the CRI implementation will use
	// to handle pods of this class. The possible values are specific to the node & CRI configuration.
	// It is assumed that all handlers are available on every node, and handlers of the same name are
	// equivalent on every node. For example, a handler called "runc" might specify that the runc OCI
	// runtime (using native Linux containers) will be used to run the containers in a pod. The Handler
	// must conform to the DNS Label (RFC 1123) requirements, and is immutable.
	Handler pulumi.StringInput `pulumi:"handler"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`

	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Overhead represents the resource overhead associated with running a pod for a given
	// RuntimeClass. For more details, see
	// https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level
	// as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	Overhead OverheadInput `pulumi:"overhead"`

	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass
	// are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be
	// supported by all nodes.
	Scheduling SchedulingInput `pulumi:"scheduling"`

}

func (a RuntimeClassPropertyArgs) ElementType() reflect.Type {
	return _RuntimeClassPropertyType
}

func (a RuntimeClassPropertyArgs) ToRuntimeClassPropertyOutput() RuntimeClassPropertyOutput {
	return pulumi.ToOutput(a).(RuntimeClassPropertyOutput)
}

func (a RuntimeClassPropertyArgs) ToRuntimeClassPropertyOutputWithContext(ctx context.Context) RuntimeClassPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RuntimeClassPropertyOutput)
}

// RuntimeClassPropertyOutput is an output type that resolves to a Input.
type RuntimeClassPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RuntimeClassPropertyOutput{}) }

func (RuntimeClassPropertyOutput) ElementType() reflect.Type {
	return _RuntimeClassPropertyType
}

func (o RuntimeClassPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v RuntimeClassProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o RuntimeClassPropertyOutput) Handler() pulumi.StringOutput {
	return o.Apply(func(v RuntimeClassProperty) string {
		return v.Handler
	}).(pulumi.StringOutput)
}

func (o RuntimeClassPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v RuntimeClassProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o RuntimeClassPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v RuntimeClassProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o RuntimeClassPropertyOutput) Overhead() OverheadOutput {
	return o.Apply(func(v RuntimeClassProperty) *Overhead {
		return v.Overhead
	}).(OverheadOutput)
}

func (o RuntimeClassPropertyOutput) Scheduling() SchedulingOutput {
	return o.Apply(func(v RuntimeClassProperty) *Scheduling {
		return v.Scheduling
	}).(SchedulingOutput)
}

var _RuntimeClassPropertyArrayType = reflect.TypeOf((*[]RuntimeClassProperty)(nil)).Elem()

type RuntimeClassPropertyArrayInput interface {
	ElementType() reflect.Type

	ToRuntimeClassPropertyArrayOutput() RuntimeClassPropertyArrayOutput
	ToRuntimeClassPropertyArrayOutputWithContext(ctx context.Context) RuntimeClassPropertyArrayOutput
}

type RuntimeClassPropertyArray []RuntimeClassPropertyInput

func (a RuntimeClassPropertyArray) ElementType() reflect.Type {
	return _RuntimeClassPropertyArrayType
}

func (a RuntimeClassPropertyArray) ToRuntimeClassPropertyArrayOutput() RuntimeClassPropertyArrayOutput {
	return pulumi.ToOutput(a).(RuntimeClassPropertyArrayOutput)
}

func (a RuntimeClassPropertyArray) ToRuntimeClassPropertyArrayOutputWithContext(ctx context.Context) RuntimeClassPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RuntimeClassPropertyArrayOutput)
}

type RuntimeClassPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(RuntimeClassPropertyArrayOutput{}) }

func (RuntimeClassPropertyArrayOutput) ElementType() reflect.Type {
	return _RuntimeClassPropertyArrayType
}

func (o RuntimeClassPropertyArrayOutput) Index(i pulumi.IntInput) RuntimeClassPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) RuntimeClassProperty {
		return vs[0].([]RuntimeClassProperty)[vs[1].(int)]
	}).(RuntimeClassPropertyOutput)
}