import pulumi
import pulumi.runtime

from ... import tables

class StorageClass(pulumi.CustomResource):
    """
    StorageClass describes the parameters for a class of storage for which PersistentVolumes can be
    dynamically provisioned.
    
    StorageClasses are non-namespaced; the name of the storage class according to etcd is in
    ObjectMeta.Name.
    """
    def __init__(self, __name__, __opts__=None, allow_volume_expansion=None, metadata=None, mount_options=None, parameters=None, provisioner=None, reclaim_policy=None, volume_binding_mode=None):
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'storage.k8s.io/v1beta1'
        __props__['kind'] = 'StorageClass'
        if not provisioner:
            raise TypeError('Missing required property provisioner')
        __props__['provisioner'] = provisioner
        __props__['allowVolumeExpansion'] = allow_volume_expansion
        __props__['metadata'] = metadata
        __props__['mountOptions'] = mount_options
        __props__['parameters'] = parameters
        __props__['reclaimPolicy'] = reclaim_policy
        __props__['volumeBindingMode'] = volume_binding_mode

        super(StorageClass, self).__init__(
            "kubernetes:storage.k8s.io/v1beta1:StorageClass",
            __name__,
            __props__,
            __opts__)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop