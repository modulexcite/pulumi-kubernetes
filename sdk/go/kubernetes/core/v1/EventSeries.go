// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// EventSeries contain information on series of events, i.e. thing that was/is happening
// continuously for some time.
type EventSeries struct {
	// Number of occurrences in this series up to the last heartbeat time
	Count *int `pulumi:"count"`

	// Time of the last occurrence observed
	LastObservedTime *string `pulumi:"lastObservedTime"`

	// State of this Series: Ongoing or Finished Deprecated. Planned removal for 1.18
	State *string `pulumi:"state"`

}

var _EventSeriesType = reflect.TypeOf((*EventSeries)(nil)).Elem()

// EventSeriesInput represents an input type that resolves to a EventSeries.
type EventSeriesInput interface {
	ElementType() reflect.Type

	ToEventSeriesOutput() EventSeriesOutput
	ToEventSeriesOutputWithContext(ctx context.Context) EventSeriesOutput
}

// EventSeriesArgs is a EventSeriesInput whose fields are all Input types.
type EventSeriesArgs struct {
	// Number of occurrences in this series up to the last heartbeat time
	Count pulumi.IntInput `pulumi:"count"`

	// Time of the last occurrence observed
	LastObservedTime pulumi.StringInput `pulumi:"lastObservedTime"`

	// State of this Series: Ongoing or Finished Deprecated. Planned removal for 1.18
	State pulumi.StringInput `pulumi:"state"`

}

func (a EventSeriesArgs) ElementType() reflect.Type {
	return _EventSeriesType
}

func (a EventSeriesArgs) ToEventSeriesOutput() EventSeriesOutput {
	return pulumi.ToOutput(a).(EventSeriesOutput)
}

func (a EventSeriesArgs) ToEventSeriesOutputWithContext(ctx context.Context) EventSeriesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventSeriesOutput)
}

// EventSeriesOutput is an output type that resolves to a Input.
type EventSeriesOutput struct { *pulumi.OutputState }

func init() { pulumi.RegisterOutputType(EventSeriesOutput{}) }

func (EventSeriesOutput) ElementType() reflect.Type {
	return _EventSeriesType
}

func (o EventSeriesOutput) Count() pulumi.IntOutput {
	return o.Apply(func(v EventSeries) *int {
		return v.Count
	}).(pulumi.IntOutput)
}

func (o EventSeriesOutput) LastObservedTime() pulumi.StringOutput {
	return o.Apply(func(v EventSeries) *string {
		return v.LastObservedTime
	}).(pulumi.StringOutput)
}

func (o EventSeriesOutput) State() pulumi.StringOutput {
	return o.Apply(func(v EventSeries) *string {
		return v.State
	}).(pulumi.StringOutput)
}
