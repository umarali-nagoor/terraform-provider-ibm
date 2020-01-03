---
layout: "ibm"
page_title: "IBM : VPN-gateway"
sidebar_current: "docs-ibm-resource-is-vpn-gateway"
description: |-
  Manages IBM VPN Gateway
---

# ibm\_is_vpn_gateway

Provides a VPN gateway resource. This allows VPN gateway to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a VPN gateway:

```hcl
resource "ibm_is_vpn_gateway" "testacc_vpn_gateway" {
  name   = "test"
  subnet = "a4ce411d-e118-4802-95ad-525e6ea0cfc9"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the VPN gateway.
* `subnet` - (Required, Forces new resource, string) The unique identifier for this subnet.
* `resource_group` - (Optional, Forces new resource, string) The resource group where the VPN gateway to be created.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPN gateway.
* `status` - The status of VPN gateway.
* `public_ip_address` -  The IP address assigned to this VPN gateway.
* `resource_controller_url` - The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance.


## Import

ibm_is_vpn_gateway can be imported using ID, eg

```
$ terraform import ibm_is_vpn_gateway.example d7bec597-4726-451f-8a63-e62e6f19c32c
```