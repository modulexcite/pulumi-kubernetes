// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";

export class VpnConnection extends lumi.NamedResource implements VpnConnectionArgs {
    public readonly customerGatewayConfiguration?: string;
    public readonly customerGatewayId: string;
    public readonly routes?: { destinationCidrBlock?: string, source?: string, state?: string }[];
    public readonly staticRoutesOnly?: boolean;
    public readonly tags?: {[key: string]: any};
    public readonly tunnel1Address?: string;
    public readonly tunnel1BgpAsn?: string;
    public readonly tunnel1BgpHoldtime?: number;
    public readonly tunnel1CgwInsideAddress?: string;
    public readonly tunnel1PresharedKey?: string;
    public readonly tunnel1VgwInsideAddress?: string;
    public readonly tunnel2Address?: string;
    public readonly tunnel2BgpAsn?: string;
    public readonly tunnel2BgpHoldtime?: number;
    public readonly tunnel2CgwInsideAddress?: string;
    public readonly tunnel2PresharedKey?: string;
    public readonly tunnel2VgwInsideAddress?: string;
    public readonly type: string;
    public readonly vgwTelemetry?: { acceptedRouteCount?: number, lastStatusChange?: string, outsideIpAddress?: string, status?: string, statusMessage?: string }[];
    public readonly vpnGatewayId: string;

    constructor(name: string, args: VpnConnectionArgs) {
        super(name);
        this.customerGatewayConfiguration = args.customerGatewayConfiguration;
        this.customerGatewayId = args.customerGatewayId;
        this.routes = args.routes;
        this.staticRoutesOnly = args.staticRoutesOnly;
        this.tags = args.tags;
        this.tunnel1Address = args.tunnel1Address;
        this.tunnel1BgpAsn = args.tunnel1BgpAsn;
        this.tunnel1BgpHoldtime = args.tunnel1BgpHoldtime;
        this.tunnel1CgwInsideAddress = args.tunnel1CgwInsideAddress;
        this.tunnel1PresharedKey = args.tunnel1PresharedKey;
        this.tunnel1VgwInsideAddress = args.tunnel1VgwInsideAddress;
        this.tunnel2Address = args.tunnel2Address;
        this.tunnel2BgpAsn = args.tunnel2BgpAsn;
        this.tunnel2BgpHoldtime = args.tunnel2BgpHoldtime;
        this.tunnel2CgwInsideAddress = args.tunnel2CgwInsideAddress;
        this.tunnel2PresharedKey = args.tunnel2PresharedKey;
        this.tunnel2VgwInsideAddress = args.tunnel2VgwInsideAddress;
        this.type = args.type;
        this.vgwTelemetry = args.vgwTelemetry;
        this.vpnGatewayId = args.vpnGatewayId;
    }
}

export interface VpnConnectionArgs {
    readonly customerGatewayConfiguration?: string;
    readonly customerGatewayId: string;
    readonly routes?: { destinationCidrBlock?: string, source?: string, state?: string }[];
    readonly staticRoutesOnly?: boolean;
    readonly tags?: {[key: string]: any};
    readonly tunnel1Address?: string;
    readonly tunnel1BgpAsn?: string;
    readonly tunnel1BgpHoldtime?: number;
    readonly tunnel1CgwInsideAddress?: string;
    readonly tunnel1PresharedKey?: string;
    readonly tunnel1VgwInsideAddress?: string;
    readonly tunnel2Address?: string;
    readonly tunnel2BgpAsn?: string;
    readonly tunnel2BgpHoldtime?: number;
    readonly tunnel2CgwInsideAddress?: string;
    readonly tunnel2PresharedKey?: string;
    readonly tunnel2VgwInsideAddress?: string;
    readonly type: string;
    readonly vgwTelemetry?: { acceptedRouteCount?: number, lastStatusChange?: string, outsideIpAddress?: string, status?: string, statusMessage?: string }[];
    readonly vpnGatewayId: string;
}
