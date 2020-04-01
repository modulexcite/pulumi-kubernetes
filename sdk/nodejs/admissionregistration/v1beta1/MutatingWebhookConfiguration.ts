// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object. Deprecated in v1.16, planned for removal in v1.19. Use admissionregistration.k8s.io/v1 MutatingWebhookConfiguration instead.
 */
export class MutatingWebhookConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing MutatingWebhookConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MutatingWebhookConfiguration {
        return new MutatingWebhookConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration';

    /**
     * Returns true if the given object is an instance of MutatingWebhookConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MutatingWebhookConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MutatingWebhookConfiguration.__pulumiType;
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
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta | undefined>;
    /**
     * Webhooks is a list of webhooks and the affected resources and operations.
     */
    public readonly webhooks!: pulumi.Output<outputs.admissionregistration.v1beta1.MutatingWebhook[] | undefined>;

    /**
     * Create a MutatingWebhookConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MutatingWebhookConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        inputs["apiVersion"] = (args ? args.apiVersion : undefined) || "admissionregistration.k8s.io/v1beta1";
        inputs["kind"] = (args ? args.kind : undefined) || "MutatingWebhookConfiguration";
        inputs["metadata"] = args ? args.metadata : undefined;
        inputs["webhooks"] = args ? args.webhooks : undefined;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:admissionregistration.k8s.io/v1:MutatingWebhookConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(MutatingWebhookConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a MutatingWebhookConfiguration resource.
 */
export interface MutatingWebhookConfigurationArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Webhooks is a list of webhooks and the affected resources and operations.
     */
    readonly webhooks?: pulumi.Input<pulumi.Input<inputs.admissionregistration.v1beta1.MutatingWebhook>[]>;
}
