// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// PersistentVolumeClaim is a user's request for and claim to a persistent volume
type PersistentVolumeClaim struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Spec defines the desired characteristics of a volume requested by a pod author. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecOutput `pulumi:"spec"`

	// Status represents the current information/status of a persistent volume claim. Read-only. More
	// info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Status PersistentVolumeClaimStatusOutput `pulumi:"status"`

}

// PersistentVolumeClaimArgs is the set of arguments needed to create a PersistentVolumeClaim resource.
type PersistentVolumeClaimArgs struct {
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

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Spec defines the desired characteristics of a volume requested by a pod author. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecInput `pulumi:"spec"`

}

// NewPersistentVolumeClaim creates a PersistentVolumeClaim resource with the given unique name, arguments, and options.
func NewPersistentVolumeClaim(ctx *pulumi.Context, name string, args *PersistentVolumeClaimArgs, opts ...pulumi.ResourceOption) (*PersistentVolumeClaim, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("v1")
		args.Kind = pulumi.String("PersistentVolumeClaim")
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
		if i := args.Spec; i != nil {
			inputs["spec"] = i.ToPersistentVolumeClaimSpecOutput()
		}
	}
	var resource PersistentVolumeClaim
	err := ctx.RegisterResource("kubernetes:core/v1:PersistentVolumeClaim", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistentVolumeClaim gets an existing PersistentVolumeClaim resource's state with the given name and ID.
func GetPersistentVolumeClaim(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*PersistentVolumeClaim, error) {
	var resource PersistentVolumeClaim
	err := ctx.ReadResource("kubernetes:core/v1:PersistentVolumeClaim", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// PersistentVolumeClaim is a user's request for and claim to a persistent volume
type PersistentVolumeClaimProperty struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should
	// convert recognized schemas to the latest internal value, and may reject unrecognized values.
	// More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer
	// this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
	// info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metaV1.ObjectMeta `pulumi:"metadata"`

	// Spec defines the desired characteristics of a volume requested by a pod author. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec *PersistentVolumeClaimSpec `pulumi:"spec"`

	// Status represents the current information/status of a persistent volume claim. Read-only. More
	// info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Status PersistentVolumeClaimStatus `pulumi:"status"`

}

var _PersistentVolumeClaimPropertyType = reflect.TypeOf((*PersistentVolumeClaimProperty)(nil)).Elem()

// PersistentVolumeClaimPropertyInput represents an input type that resolves to a PersistentVolumeClaimProperty.
type PersistentVolumeClaimPropertyInput interface {
	ElementType() reflect.Type

	ToPersistentVolumeClaimPropertyOutput() PersistentVolumeClaimPropertyOutput
	ToPersistentVolumeClaimPropertyOutputWithContext(ctx context.Context) PersistentVolumeClaimPropertyOutput
}

// PersistentVolumeClaimPropertyArgs is a PersistentVolumeClaimPropertyInput whose fields are all Input types.
type PersistentVolumeClaimPropertyArgs struct {
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

	// Standard object's metadata. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

	// Spec defines the desired characteristics of a volume requested by a pod author. More info:
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	Spec PersistentVolumeClaimSpecInput `pulumi:"spec"`

}

func (a PersistentVolumeClaimPropertyArgs) ElementType() reflect.Type {
	return _PersistentVolumeClaimPropertyType
}

func (a PersistentVolumeClaimPropertyArgs) ToPersistentVolumeClaimPropertyOutput() PersistentVolumeClaimPropertyOutput {
	return pulumi.ToOutput(a).(PersistentVolumeClaimPropertyOutput)
}

func (a PersistentVolumeClaimPropertyArgs) ToPersistentVolumeClaimPropertyOutputWithContext(ctx context.Context) PersistentVolumeClaimPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PersistentVolumeClaimPropertyOutput)
}

// PersistentVolumeClaimPropertyOutput is an output type that resolves to a Input.
type PersistentVolumeClaimPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PersistentVolumeClaimPropertyOutput{}) }

func (PersistentVolumeClaimPropertyOutput) ElementType() reflect.Type {
	return _PersistentVolumeClaimPropertyType
}

func (o PersistentVolumeClaimPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v PersistentVolumeClaimProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o PersistentVolumeClaimPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v PersistentVolumeClaimProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o PersistentVolumeClaimPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v PersistentVolumeClaimProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o PersistentVolumeClaimPropertyOutput) Spec() PersistentVolumeClaimSpecOutput {
	return o.Apply(func(v PersistentVolumeClaimProperty) *PersistentVolumeClaimSpec {
		return v.Spec
	}).(PersistentVolumeClaimSpecOutput)
}

func (o PersistentVolumeClaimPropertyOutput) Status() PersistentVolumeClaimStatusOutput {
	return o.Apply(func(v PersistentVolumeClaimProperty) PersistentVolumeClaimStatus {
		return v.Status
	}).(PersistentVolumeClaimStatusOutput)
}

var _PersistentVolumeClaimPropertyArrayType = reflect.TypeOf((*[]PersistentVolumeClaimProperty)(nil)).Elem()

type PersistentVolumeClaimPropertyArrayInput interface {
	ElementType() reflect.Type

	ToPersistentVolumeClaimPropertyArrayOutput() PersistentVolumeClaimPropertyArrayOutput
	ToPersistentVolumeClaimPropertyArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimPropertyArrayOutput
}

type PersistentVolumeClaimPropertyArray []PersistentVolumeClaimPropertyInput

func (a PersistentVolumeClaimPropertyArray) ElementType() reflect.Type {
	return _PersistentVolumeClaimPropertyArrayType
}

func (a PersistentVolumeClaimPropertyArray) ToPersistentVolumeClaimPropertyArrayOutput() PersistentVolumeClaimPropertyArrayOutput {
	return pulumi.ToOutput(a).(PersistentVolumeClaimPropertyArrayOutput)
}

func (a PersistentVolumeClaimPropertyArray) ToPersistentVolumeClaimPropertyArrayOutputWithContext(ctx context.Context) PersistentVolumeClaimPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PersistentVolumeClaimPropertyArrayOutput)
}

type PersistentVolumeClaimPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(PersistentVolumeClaimPropertyArrayOutput{}) }

func (PersistentVolumeClaimPropertyArrayOutput) ElementType() reflect.Type {
	return _PersistentVolumeClaimPropertyArrayType
}

func (o PersistentVolumeClaimPropertyArrayOutput) Index(i pulumi.IntInput) PersistentVolumeClaimPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) PersistentVolumeClaimProperty {
		return vs[0].([]PersistentVolumeClaimProperty)[vs[1].(int)]
	}).(PersistentVolumeClaimPropertyOutput)
}