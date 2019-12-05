// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// Job represents the configuration of a single job.
// 
// This resource waits until its status is ready before registering success
// for create/update, and populating output properties from the current state of the resource.
// The following conditions are used to determine whether the resource creation has
// succeeded or failed:
// 
// 1. The Job's '.status.startTime' is set, which indicates that the Job has started running.
// 2. The Job's '.status.conditions' has a status of type 'Complete', and a 'status' set
//    to 'True'.
// 3. The Job's '.status.conditions' do not have a status of type 'Failed', with a
// 	'status' set to 'True'. If this condition is set, we should fail the Job immediately.
// 
// If the Job has not reached a Ready state after 10 minutes, it will
// time out and mark the resource update as Failed. You can override the default timeout value
// by setting the 'customTimeouts' option on the resource.
type Job struct {
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

	// Specification of the desired behavior of a job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec JobSpecOutput `pulumi:"spec"`

	// Current status of a job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status JobStatusOutput `pulumi:"status"`

}

// JobArgs is the set of arguments needed to create a Job resource.
type JobArgs struct {
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

	// Specification of the desired behavior of a job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec JobSpecInput `pulumi:"spec"`

}

// NewJob creates a Job resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context, name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		args.ApiVersion = pulumi.String("batch/v1")
		args.Kind = pulumi.String("Job")
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
			inputs["spec"] = i.ToJobSpecOutput()
		}
	}
	var resource Job
	err := ctx.RegisterResource("kubernetes:batch/v1:Job", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name and ID.
func GetJob(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("kubernetes:batch/v1:Job", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Job represents the configuration of a single job.
type JobProperty struct {
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

	// Specification of the desired behavior of a job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *JobSpec `pulumi:"spec"`

	// Current status of a job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status JobStatus `pulumi:"status"`

}

var _JobPropertyType = reflect.TypeOf((*JobProperty)(nil)).Elem()

// JobPropertyInput represents an input type that resolves to a JobProperty.
type JobPropertyInput interface {
	ElementType() reflect.Type

	ToJobPropertyOutput() JobPropertyOutput
	ToJobPropertyOutputWithContext(ctx context.Context) JobPropertyOutput
}

// JobPropertyArgs is a JobPropertyInput whose fields are all Input types.
type JobPropertyArgs struct {
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

	// Specification of the desired behavior of a job. More info:
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec JobSpecInput `pulumi:"spec"`

}

func (a JobPropertyArgs) ElementType() reflect.Type {
	return _JobPropertyType
}

func (a JobPropertyArgs) ToJobPropertyOutput() JobPropertyOutput {
	return pulumi.ToOutput(a).(JobPropertyOutput)
}

func (a JobPropertyArgs) ToJobPropertyOutputWithContext(ctx context.Context) JobPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, a).(JobPropertyOutput)
}

// JobPropertyOutput is an output type that resolves to a Input.
type JobPropertyOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(JobPropertyOutput{}) }

func (JobPropertyOutput) ElementType() reflect.Type {
	return _JobPropertyType
}

func (o JobPropertyOutput) ApiVersion() pulumi.StringOutput {
	return o.Apply(func(v JobProperty) *string {
		return v.ApiVersion
	}).(pulumi.StringOutput)
}

func (o JobPropertyOutput) Kind() pulumi.StringOutput {
	return o.Apply(func(v JobProperty) *string {
		return v.Kind
	}).(pulumi.StringOutput)
}

func (o JobPropertyOutput) Metadata() metaV1.ObjectMetaOutput {
	return o.Apply(func(v JobProperty) *metaV1.ObjectMeta {
		return v.Metadata
	}).(metaV1.ObjectMetaOutput)
}

func (o JobPropertyOutput) Spec() JobSpecOutput {
	return o.Apply(func(v JobProperty) *JobSpec {
		return v.Spec
	}).(JobSpecOutput)
}

func (o JobPropertyOutput) Status() JobStatusOutput {
	return o.Apply(func(v JobProperty) JobStatus {
		return v.Status
	}).(JobStatusOutput)
}

var _JobPropertyArrayType = reflect.TypeOf((*[]JobProperty)(nil)).Elem()

type JobPropertyArrayInput interface {
	ElementType() reflect.Type

	ToJobPropertyArrayOutput() JobPropertyArrayOutput
	ToJobPropertyArrayOutputWithContext(ctx context.Context) JobPropertyArrayOutput
}

type JobPropertyArray []JobPropertyInput

func (a JobPropertyArray) ElementType() reflect.Type {
	return _JobPropertyArrayType
}

func (a JobPropertyArray) ToJobPropertyArrayOutput() JobPropertyArrayOutput {
	return pulumi.ToOutput(a).(JobPropertyArrayOutput)
}

func (a JobPropertyArray) ToJobPropertyArrayOutputWithContext(ctx context.Context) JobPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(JobPropertyArrayOutput)
}

type JobPropertyArrayOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(JobPropertyArrayOutput{}) }

func (JobPropertyArrayOutput) ElementType() reflect.Type {
	return _JobPropertyArrayType
}

func (o JobPropertyArrayOutput) Index(i pulumi.IntInput) JobPropertyOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) JobProperty {
		return vs[0].([]JobProperty)[vs[1].(int)]
	}).(JobPropertyOutput)
}