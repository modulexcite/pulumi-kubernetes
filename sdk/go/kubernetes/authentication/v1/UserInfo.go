// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// UserInfo holds the information about the user needed to implement the user.Info interface.
type UserInfo struct {
	// Any additional information provided by the authenticator.
	Extra map[string][]string `pulumi:"extra"`

	// The names of groups this user is a part of.
	Groups []string `pulumi:"groups"`

	// A unique value that identifies this user across time. If this user is deleted and another user
	// by the same name is added, they will have different UIDs.
	Uid *string `pulumi:"uid"`

	// The name that uniquely identifies this user among all active users.
	Username *string `pulumi:"username"`

}

var _UserInfoType = reflect.TypeOf((*UserInfo)(nil)).Elem()

// UserInfoInput represents an input type that resolves to a UserInfo.
type UserInfoInput interface {
	ElementType() reflect.Type

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput
}

// UserInfoArgs is a UserInfoInput whose fields are all Input types.
type UserInfoArgs struct {
	// Any additional information provided by the authenticator.
	Extra pulumi.StringArrayMapInput `pulumi:"extra"`

	// The names of groups this user is a part of.
	Groups pulumi.StringArrayInput `pulumi:"groups"`

	// A unique value that identifies this user across time. If this user is deleted and another user
	// by the same name is added, they will have different UIDs.
	Uid pulumi.StringInput `pulumi:"uid"`

	// The name that uniquely identifies this user among all active users.
	Username pulumi.StringInput `pulumi:"username"`

}

func (a UserInfoArgs) ElementType() reflect.Type {
	return _UserInfoType
}

func (a UserInfoArgs) ToUserInfoOutput() UserInfoOutput {
	return pulumi.ToOutput(a).(UserInfoOutput)
}

func (a UserInfoArgs) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, a).(UserInfoOutput)
}

// UserInfoOutput is an output type that resolves to a Input.
type UserInfoOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(UserInfoOutput{}) }

func (UserInfoOutput) ElementType() reflect.Type {
	return _UserInfoType
}

func (o UserInfoOutput) Extra() pulumi.StringArrayMapOutput {
	return o.Apply(func(v UserInfo) map[string][]string {
		return v.Extra
	}).(pulumi.StringArrayMapOutput)
}

func (o UserInfoOutput) Groups() pulumi.StringArrayOutput {
	return o.Apply(func(v UserInfo) []string {
		return v.Groups
	}).(pulumi.StringArrayOutput)
}

func (o UserInfoOutput) Uid() pulumi.StringOutput {
	return o.Apply(func(v UserInfo) *string {
		return v.Uid
	}).(pulumi.StringOutput)
}

func (o UserInfoOutput) Username() pulumi.StringOutput {
	return o.Apply(func(v UserInfo) *string {
		return v.Username
	}).(pulumi.StringOutput)
}
