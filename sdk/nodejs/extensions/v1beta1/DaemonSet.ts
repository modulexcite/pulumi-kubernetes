// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * DaemonSet represents the configuration of a daemon set.
 */
export class DaemonSet extends pulumi.CustomResource {
    /**
     * Get an existing DaemonSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DaemonSet {
        pulumi.log.warn("DaemonSet is deprecated: extensions/v1beta1/DaemonSet is deprecated by apps/v1/DaemonSet and not supported by Kubernetes v1.16+ clusters.")
        return new DaemonSet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:extensions/v1beta1:DaemonSet';

    /**
     * Returns true if the given object is an instance of DaemonSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DaemonSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DaemonSet.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta | undefined>;
    /**
     * The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    public readonly spec!: pulumi.Output<outputs.extensions.v1beta1.DaemonSetSpec | undefined>;
    /**
     * The current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.extensions.v1beta1.DaemonSetStatus | undefined>;

    /**
     * Create a DaemonSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated extensions/v1beta1/DaemonSet is deprecated by apps/v1/DaemonSet and not supported by Kubernetes v1.16+ clusters. */
    constructor(name: string, args?: DaemonSetArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DaemonSet is deprecated: extensions/v1beta1/DaemonSet is deprecated by apps/v1/DaemonSet and not supported by Kubernetes v1.16+ clusters.")
        let inputs: pulumi.Inputs = {};
        inputs["apiVersion"] = (args ? args.apiVersion : undefined) || "extensions/v1beta1";
        inputs["kind"] = (args ? args.kind : undefined) || "DaemonSet";
        inputs["metadata"] = args ? args.metadata : undefined;
        inputs["spec"] = args ? args.spec : undefined;
        inputs["status"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:apps/v1:DaemonSet" }, { type: "kubernetes:apps/v1beta2:DaemonSet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DaemonSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DaemonSet resource.
 */
export interface DaemonSetArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * The desired behavior of this daemon set. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    readonly spec?: pulumi.Input<inputs.extensions.v1beta1.DaemonSetSpec>;
}
