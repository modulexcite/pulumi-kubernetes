// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support
// ownership management or SELinux relabeling.
type GlusterfsPersistentVolumeSource struct {
	// EndpointsName is the endpoint name that details Glusterfs topology. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints string `pulumi:"endpoints"`

	// EndpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty,
	// the EndpointNamespace defaults to the same namespace as the bound PVC. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	EndpointsNamespace *string `pulumi:"endpointsNamespace"`

	// Path is the Glusterfs volume path. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path string `pulumi:"path"`

	// ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults
	// to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	ReadOnly *bool `pulumi:"readOnly"`

}

var _GlusterfsPersistentVolumeSourceType = reflect.TypeOf((*GlusterfsPersistentVolumeSource)(nil)).Elem()

// GlusterfsPersistentVolumeSourceInput represents an input type that resolves to a GlusterfsPersistentVolumeSource.
type GlusterfsPersistentVolumeSourceInput interface {
	ElementType() reflect.Type

	ToGlusterfsPersistentVolumeSourceOutput() GlusterfsPersistentVolumeSourceOutput
	ToGlusterfsPersistentVolumeSourceOutputWithContext(ctx context.Context) GlusterfsPersistentVolumeSourceOutput
}

// GlusterfsPersistentVolumeSourceArgs is a GlusterfsPersistentVolumeSourceInput whose fields are all Input types.
type GlusterfsPersistentVolumeSourceArgs struct {
	// EndpointsName is the endpoint name that details Glusterfs topology. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints pulumi.StringInput `pulumi:"endpoints"`

	// Path is the Glusterfs volume path. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path pulumi.StringInput `pulumi:"path"`

	// EndpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty,
	// the EndpointNamespace defaults to the same namespace as the bound PVC. More info:
	// https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	EndpointsNamespace pulumi.StringInput `pulumi:"endpointsNamespace"`

	// ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults
	// to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	ReadOnly pulumi.BoolInput `pulumi:"readOnly"`

}

func (a GlusterfsPersistentVolumeSourceArgs) ElementType() reflect.Type {
	return _GlusterfsPersistentVolumeSourceType
}

func (a GlusterfsPersistentVolumeSourceArgs) ToGlusterfsPersistentVolumeSourceOutput() GlusterfsPersistentVolumeSourceOutput {
	return pulumi.ToOutput(a).(GlusterfsPersistentVolumeSourceOutput)
}

func (a GlusterfsPersistentVolumeSourceArgs) ToGlusterfsPersistentVolumeSourceOutputWithContext(ctx context.Context) GlusterfsPersistentVolumeSourceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(GlusterfsPersistentVolumeSourceOutput)
}

// GlusterfsPersistentVolumeSourceOutput is an output type that resolves to a Input.
type GlusterfsPersistentVolumeSourceOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(GlusterfsPersistentVolumeSourceOutput{}) }

func (GlusterfsPersistentVolumeSourceOutput) ElementType() reflect.Type {
	return _GlusterfsPersistentVolumeSourceType
}

func (o GlusterfsPersistentVolumeSourceOutput) Endpoints() pulumi.StringOutput {
	return o.Apply(func(v GlusterfsPersistentVolumeSource) string {
		return v.Endpoints
	}).(pulumi.StringOutput)
}

func (o GlusterfsPersistentVolumeSourceOutput) EndpointsNamespace() pulumi.StringOutput {
	return o.Apply(func(v GlusterfsPersistentVolumeSource) *string {
		return v.EndpointsNamespace
	}).(pulumi.StringOutput)
}

func (o GlusterfsPersistentVolumeSourceOutput) Path() pulumi.StringOutput {
	return o.Apply(func(v GlusterfsPersistentVolumeSource) string {
		return v.Path
	}).(pulumi.StringOutput)
}

func (o GlusterfsPersistentVolumeSourceOutput) ReadOnly() pulumi.BoolOutput {
	return o.Apply(func(v GlusterfsPersistentVolumeSource) *bool {
		return v.ReadOnly
	}).(pulumi.BoolOutput)
}
