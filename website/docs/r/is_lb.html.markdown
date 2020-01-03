---
layout: "ibm"
page_title: "IBM : load balancer"
sidebar_current: "docs-ibm-resource-is-lb"
description: |-
  Manages IBM load balancer.
---

# ibm\_is_lb

Provides a load balancer resource. This allows load balancer to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a load balancer:

```hcl
resource "ibm_is_lb" "lb" {
  name    = "loadbalancer1"
  subnets = ["04813493-15d6-4150-9948-6cc646cb67f2"]
}

```

## Timeouts

ibm_is_lb provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating Instance.
* `delete` - (Default 60 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the loadbalancer.
* `subnets` - (Required, list) ID of the subnets to provision this load balancer.
* `type` - (Optional, Forces new resource, string) The type of the load balancer. Default value `public`. Supported values `public` and  `private`.
* `resource_group` - (Optional, Forces new resource, string) The resource group where the load balancer to be created.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the load balancer.
* `public_ips` - The public IP addresses assigned to this load balancer.
* `private_ips` - The private IP addresses assigned to this load balancer.
* `status` - The status of load balancer.
* `operating_status` - The operating status of this load balancer.
* `hostname` - Fully qualified domain name assigned to this load balancer.
* `resource_controller_url` - The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance.


## Import

ibm_is_lb can be imported using lbID, eg

```
$ terraform import ibm_is_lb.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
