// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action
 */
export class SelfSubjectAccessReview extends pulumi.CustomResource {
    /**
     * Get an existing SelfSubjectAccessReview resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SelfSubjectAccessReview {
        return new SelfSubjectAccessReview(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReview';

    /**
     * Returns true if the given object is an instance of SelfSubjectAccessReview.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SelfSubjectAccessReview {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SelfSubjectAccessReview.__pulumiType;
    }

    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<string | undefined>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta | undefined>;
    /**
     * Spec holds information about the request being evaluated.  user and groups must be empty
     */
    public readonly spec!: pulumi.Output<outputs.authorization.v1.SelfSubjectAccessReviewSpec | undefined>;
    /**
     * Status is filled in by the server and indicates whether the request is allowed or not
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.authorization.v1.SubjectAccessReviewStatus | undefined>;

    /**
     * Create a SelfSubjectAccessReview resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SelfSubjectAccessReviewArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.spec === undefined) {
                throw new Error("Missing required property 'spec'");
            }
        inputs["apiVersion"] = (args ? args.apiVersion : undefined) || "authorization.k8s.io/v1";
        inputs["kind"] = (args ? args.kind : undefined) || "SelfSubjectAccessReview";
        inputs["metadata"] = args ? args.metadata : undefined;
        inputs["spec"] = args ? args.spec : undefined;
        inputs["status"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReview" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SelfSubjectAccessReview.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SelfSubjectAccessReview resource.
 */
export interface SelfSubjectAccessReviewArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Spec holds information about the request being evaluated.  user and groups must be empty
     */
    readonly spec: pulumi.Input<inputs.authorization.v1.SelfSubjectAccessReviewSpec>;
}
