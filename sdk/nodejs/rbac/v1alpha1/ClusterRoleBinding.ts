// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBinding, and will no longer be served in v1.20.
 */
export class ClusterRoleBinding extends pulumi.CustomResource {
    /**
     * Get an existing ClusterRoleBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClusterRoleBinding {
        return new ClusterRoleBinding(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRoleBinding';

    /**
     * Returns true if the given object is an instance of ClusterRoleBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterRoleBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterRoleBinding.__pulumiType;
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
     * Standard object's metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta | undefined>;
    /**
     * RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
     */
    public readonly roleRef!: pulumi.Output<outputs.rbac.v1alpha1.RoleRef | undefined>;
    /**
     * Subjects holds references to the objects the role applies to.
     */
    public readonly subjects!: pulumi.Output<outputs.rbac.v1alpha1.Subject[] | undefined>;

    /**
     * Create a ClusterRoleBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterRoleBindingArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.roleRef === undefined) {
                throw new Error("Missing required property 'roleRef'");
            }
        inputs["apiVersion"] = "rbac.authorization.k8s.io/v1alpha1";
        inputs["kind"] = "ClusterRoleBinding";
        inputs["metadata"] = args ? args.metadata : undefined;
        inputs["roleRef"] = args ? args.roleRef : undefined;
        inputs["subjects"] = args ? args.subjects : undefined;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "kubernetes:rbac.authorization.k8s.io/v1:ClusterRoleBinding" }, { type: "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBinding" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ClusterRoleBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClusterRoleBinding resource.
 */
export interface ClusterRoleBindingArgs {
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Standard object's metadata.
     */
    readonly metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
     */
    readonly roleRef: pulumi.Input<inputs.rbac.v1alpha1.RoleRef>;
    /**
     * Subjects holds references to the objects the role applies to.
     */
    readonly subjects?: pulumi.Input<pulumi.Input<inputs.rbac.v1alpha1.Subject>[]>;
}
