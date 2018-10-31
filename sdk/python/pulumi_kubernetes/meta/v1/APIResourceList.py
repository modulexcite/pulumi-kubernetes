import pulumi
import pulumi.runtime

from ...tables import _CASING_FORWARD_TABLE, _CASING_BACKWARD_TABLE

class APIResourceList(pulumi.CustomResource):
    """
    APIResourceList is a list of APIResource, it is used to expose the name of the resources
    supported in a specific group and version, and if the resource is namespaced.
    """
    def __init__(self, __name__, __opts__=None, group_version=None, resources=None):
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'v1'
        self.apiVersion = 'v1'

        __props__['kind'] = 'APIResourceList'
        self.kind = 'APIResourceList'

        if not groupVersion:
            raise TypeError('Missing required property groupVersion')
        elif not isinstance(groupVersion, str):
            raise TypeError('Expected property aliases to be a str')
        self.group_version = group_version
        """
        groupVersion is the group and version this APIResourceList is for.
        """
        __props__['groupVersion'] = group_version

        if not resources:
            raise TypeError('Missing required property resources')
        elif not isinstance(resources, list):
            raise TypeError('Expected property aliases to be a list')
        self.resources = resources
        """
        resources contains the name of the resources and if they are namespaced.
        """
        __props__['resources'] = resources

        super(APIResourceList, self).__init__(
            "kubernetes:core/v1:APIResourceList",
            __name__,
            __props__,
            __opts__)

    def translate_output_property(self, prop: str) -> str:
        return _CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return _CASING_BACKWARD_TABLE.get(prop) or prop
