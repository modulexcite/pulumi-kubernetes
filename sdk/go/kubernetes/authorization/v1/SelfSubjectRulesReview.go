// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	metaV1 "github.com/pulumi/pulumi-kubernetes/sdk/go/kubernetes/meta/v1"
)

// SelfSubjectRulesReview enumerates the set of actions the current user can perform within a
// namespace. The returned list of actions may be incomplete depending on the server's authorization
// mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by
// UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should
// NOT Be used by external systems to drive authorization decisions as this raises confused deputy,
// cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview
// are the correct way to defer authorization decisions to the API server.
type SelfSubjectRulesReview struct {
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

	
	Metadata metaV1.ObjectMetaOutput `pulumi:"metadata"`

	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpecOutput `pulumi:"spec"`

	// Status is filled in by the server and indicates the set of actions a user can perform.
	Status SubjectRulesReviewStatusOutput `pulumi:"status"`

}

// SelfSubjectRulesReviewArgs is the set of arguments needed to create a SelfSubjectRulesReview resource.
type SelfSubjectRulesReviewArgs struct {
	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpecInput `pulumi:"spec"`

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

	
	Metadata metaV1.ObjectMetaInput `pulumi:"metadata"`

}

// NewSelfSubjectRulesReview creates a SelfSubjectRulesReview resource with the given unique name, arguments, and options.
func NewSelfSubjectRulesReview(ctx *pulumi.Context, name string, args *SelfSubjectRulesReviewArgs, opts ...pulumi.ResourceOption) (*SelfSubjectRulesReview, error) {
	inputs := map[string]pulumi.Input{}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args != nil {
		args.ApiVersion = pulumi.String("authorization.k8s.io/v1")
		args.Kind = pulumi.String("SelfSubjectRulesReview")
		inputs["spec"] = args.Spec.ToSelfSubjectRulesReviewSpecOutput()
		if i := args.ApiVersion; i != nil {
			inputs["apiVersion"] = i.ToStringOutput()
		}
		if i := args.Kind; i != nil {
			inputs["kind"] = i.ToStringOutput()
		}
		if i := args.Metadata; i != nil {
			inputs["metadata"] = i.ToObjectMetaOutput()
		}
	}
	var resource SelfSubjectRulesReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectRulesReview gets an existing SelfSubjectRulesReview resource's state with the given name and ID.
func GetSelfSubjectRulesReview(ctx *pulumi.Context, name string, id pulumi.IDInput, opts ...pulumi.ResourceOption) (*SelfSubjectRulesReview, error) {
	var resource SelfSubjectRulesReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview", name, id, nil, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
