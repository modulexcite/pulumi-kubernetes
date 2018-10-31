import pulumi
import pulumi.runtime

from ...tables import _CASING_FORWARD_TABLE, _CASING_BACKWARD_TABLE

class PodPreset(pulumi.CustomResource):
    """
    PodPreset is a policy resource that defines additional runtime requirements for a Pod.
    """
    def __init__(self, __name__, __opts__=None, metadata=None, spec=None):
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'settings.k8s.io/v1alpha1'
        self.apiVersion = 'settings.k8s.io/v1alpha1'

        __props__['kind'] = 'PodPreset'
        self.kind = 'PodPreset'

        if metadata and not isinstance(metadata, dict):
            raise TypeError('Expected property aliases to be a dict')
        self.metadata = metadata
        
        __props__['metadata'] = metadata

        if spec and not isinstance(spec, dict):
            raise TypeError('Expected property aliases to be a dict')
        self.spec = spec
        
        __props__['spec'] = spec

        super(PodPreset, self).__init__(
            "kubernetes:settings.k8s.io/v1alpha1:PodPreset",
            __name__,
            __props__,
            __opts__)

    def translate_output_property(self, prop: str) -> str:
        return _CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return _CASING_BACKWARD_TABLE.get(prop) or prop
