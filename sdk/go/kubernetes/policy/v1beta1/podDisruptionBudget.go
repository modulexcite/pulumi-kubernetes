// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudget struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPtrOutput `pulumi:"spec"`
	// Most recently observed status of the PodDisruptionBudget.
	Status PodDisruptionBudgetStatusPtrOutput `pulumi:"status"`
}

// NewPodDisruptionBudget registers a new resource with the given unique name, arguments, and options.
func NewPodDisruptionBudget(ctx *pulumi.Context,
	name string, args *PodDisruptionBudgetArgs, opts ...pulumi.ResourceOption) (*PodDisruptionBudget, error) {
	if args == nil {
		args = &PodDisruptionBudgetArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("policy/v1beta1")
	args.Kind = pulumi.StringPtr("PodDisruptionBudget")
	var resource PodDisruptionBudget
	err := ctx.RegisterResource("kubernetes:policy/v1beta1:PodDisruptionBudget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodDisruptionBudget gets an existing PodDisruptionBudget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodDisruptionBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodDisruptionBudgetState, opts ...pulumi.ResourceOption) (*PodDisruptionBudget, error) {
	var resource PodDisruptionBudget
	err := ctx.ReadResource("kubernetes:policy/v1beta1:PodDisruptionBudget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodDisruptionBudget resources.
type podDisruptionBudgetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpec `pulumi:"spec"`
	// Most recently observed status of the PodDisruptionBudget.
	Status *PodDisruptionBudgetStatus `pulumi:"status"`
}

type PodDisruptionBudgetState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPtrInput
	// Most recently observed status of the PodDisruptionBudget.
	Status PodDisruptionBudgetStatusPtrInput
}

func (PodDisruptionBudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetState)(nil)).Elem()
}

type podDisruptionBudgetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpec `pulumi:"spec"`
}

// The set of arguments for constructing a PodDisruptionBudget resource.
type PodDisruptionBudgetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpecPtrInput
}

func (PodDisruptionBudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podDisruptionBudgetArgs)(nil)).Elem()
}
