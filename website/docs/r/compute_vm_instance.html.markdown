---
layout: "ibm"
page_title: "IBM: compute_vm_instance"
sidebar_current: "docs-ibm-resource-compute-vm-instance"
description: |-
  Manages IBM VM instances.
---

# ibm\_compute_vm_instance

Provides a resource for VM instances. This allows VM instances to be created, updated, and deleted.

For additional details, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Virtual_Guest).

**NOTE**: Update is not supported when the `bulk_vms` parameter is used.

## Example Usage

In the following example, you can create a VM instance using a Debian image:

```hcl
resource "ibm_compute_vm_instance" "twc_terraform_sample" {
  hostname                   = "twc-terraform-sample-name"
  domain                     = "bar.example.com"
  os_reference_code          = "DEBIAN_8_64"
  datacenter                 = "wdc01"
  network_speed              = 10
  hourly_billing             = true
  private_network_only       = false
  cores                      = 1
  memory                     = 1024
  disks                      = [25, 10, 20]
  user_metadata              = "{\"value\":\"newvalue\"}"
  dedicated_acct_host_only   = true
  local_disk                 = false
  public_vlan_id             = 1391277
  private_vlan_id            = 7721931
  private_security_group_ids = [576973]
}
```

In the following example, you can create a VM instance using a block device template:

```hcl
resource "ibm_compute_vm_instance" "terraform-sample-BDTGroup" {
  hostname   = "terraform-sample-blockDeviceTemplateGroup"
  domain     = "bar.example.com"
  datacenter = "ams01"

  //public_network_speed = 10
  hourly_billing = false
  cores          = 1
  memory         = 1024
  local_disk     = false
  image_id       = 12345
  tags = [
    "collectd",
    "mesos-master",
  ]
  public_subnet  = "50.97.46.160/28"
  private_subnet = "10.56.109.128/26"
}
```

In the following example, you can create a VM instance using a flavor:

```hcl
resource "ibm_compute_vm_instance" "terraform-sample-flavor" {
  hostname             = "terraform-sample-flavor"
  domain               = "bar.example.com"
  os_reference_code    = "DEBIAN_8_64"
  datacenter           = "dal06"
  network_speed        = 10
  hourly_billing       = true
  local_disk           = false
  private_network_only = false
  flavor_key_name      = "B1_2X8X25"
  user_metadata        = "{\\\"value\\\":\\\"newvalue\\\"}"

  // provide disk 3, 4, 5 and so on
  disks = [10, 20, 30]
  tags  = ["collectd"]

  // It should be false
  dedicated_acct_host_only = false
  ipv6_enabled             = true
  secondary_ip_count       = 4
  notes                    = "VM notes"
}
```
In the following example, you can create multiple vm's

```hcl
resource "ibm_compute_vm_instance" "terraform-bulk-vms" {
  bulk_vms {
    hostname = "vm1"

    domain = "bar.example.com"
  }
  bulk_vms {
    hostname = "vm2"

    domain = "bar.example.com"
  }

  os_reference_code    = "CENTOS_7_64"
  datacenter           = "dal09"
  network_speed        = 100
  hourly_billing       = true
  private_network_only = true
  cores                = 1
  memory               = 1024
  disks                = [25]
  local_disk           = false
}

```

In the following example, you can retry to create a VM instance using a datacenter_choice. If VM fails to place order on first datacenter or vlans it retries to place order on subsequent datacenters and vlans untill place order is successfull:

```hcl
resource "ibm_compute_vm_instance" "terraform-retry" {
  hostname          = "vmretry"
  domain            = "example.com"
  network_speed     = 100
  hourly_billing    = true
  cores             = 1
  memory            = 1024
  local_disk        = false
  os_reference_code = "DEBIAN_7_64"
  disks             = [25]

  datacenter_choice = [
    {
      datacenter      = "dal09"
      public_vlan_id  = 123245
      private_vlan_id = 123255
    },
    {
      datacenter = "wdc54"
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 153345
      private_vlan_id = 123255
    },
    {
      datacenter = "dal06"
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 123245
      private_vlan_id = 123255
    },
    {
      datacenter      = "dal09"
      public_vlan_id  = 1232454
      private_vlan_id = 1234567
    },
  ]
}

```


## Argument Reference

The following arguments are supported:

* `hostname` - (Optional, string) The hostname for the computing instance.</br>
    **NOTE**: Conflicts with `bulk_vms`.
* `domain` - (Optional, string)  The domain for the computing instance.</br>
    **NOTE**: Conflicts with `bulk_vms`.
* `bulk_vms` - (Optional, Forces new resource, list) List of hostname and domain of the computing instance. The minimum number of vm's to be defined is 2. Nested `bulk_vms` blocks must have the following structure:
    * `hostname` - (Required, Forces new resource, string) The hostname for the computing instance.
    * `domain` - (Required, Forces new resource, string) The domain for the computing instance.</br>
    **NOTE**: Conflicts with `hostname` and `domain`.
* `cores` - (Optional, integer) The number of CPU cores that you want to allocate.  
    **NOTE**: Conflicts with `flavor_key_name`.
* `memory` - (Optional, integer) The amount of memory, expressed in megabytes, that you want to allocate.  
    **NOTE**: Conflicts with `flavor_key_name`.
* `flavor_key_name` - (Optional, string) The flavor key name that you want to use to provision the instance. To see available Flavor key name, log in to the [IBM Cloud Classic Infrastructure (SoftLayer) API](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest/getCreateObjectOptions.json), using your API key as the password.  
    **NOTE**: Conflicts with `cores` and `memory`.
* `datacenter` - (Optional, Forces new resource, string) The datacenter in which you want to provision the instance.  
    **NOTE**: If `dedicated_host_name` or `dedicated_host_id`
    is provided then the datacenter should be same as the dedicated host datacenter. 
    If `placement_group_name` or `placement_group_id`
    is provided then the datacenter should be same as the placement group datacenter.
    Conflicts with `datacenter_choice`. 
* `hourly_billing` - (Optional, Forces new resource, boolean) The billing type for the instance. When set to `true`, the computing instance is billed on hourly usage. Otherwise, the instance is billed on a monthly basis. The default value is `true`.
* `local_disk`- (Optional, Forces new resource, boolean) The disk type for the instance. When set to `true`, the disks for the computing instance are provisioned on the host that the instance runs. Otherwise, SAN disks are provisioned. The default value is `true`.
* `dedicated_acct_host_only` - (Optional, Forces new resource, boolean) Specifies whether the instance must only run on hosts with instances from the same account. The default value is `false`. If VM is provisioned using flavorKeyName, value should be set to `false`.  
     **NOTE**: Conflicts with `dedicated_host_name`, `dedicated_host_id`, `placement_group_name` and `placement_group_id`.
* `dedicated_host_id` - (Optional, Forces new resource, integer) Specifies [dedicated host](https://cloud.ibm.com/docs/vsi/vsi_dedicated.html) for the instance by its id.  
     **NOTE**: Conflicts with `dedicated_acct_host_only`, `dedicated_host_name`, `placement_group_name` and `placement_group_id`.
* `dedicated_host_name` - (Optional, Forces new resource, string) Specifies [dedicated host](https://cloud.ibm.com/docs/vsi/vsi_dedicated.html) for the instance by its name.  
     **NOTE**: Conflicts with `dedicated_acct_host_only`, `dedicated_host_id`, `placement_group_name` and `placement_group_id`.
* `placement_group_id` - (Optional, Forces new resource, integer) Specifies [placement group](https://cloud.ibm.com/docs/vsi/vsi_dedicated.html) for the instance by its id.  
     **NOTE**: Conflicts with `dedicated_acct_host_only`, `dedicated_host_name`, `dedicated_host_id` and `placement_group_name`.
* `placement_group_name` - (Optional, Forces new resource, string) Specifies [placement group](https://cloud.ibm.com/docs/vsi/vsi_dedicated.html) for the instance by its name.  
     **NOTE**: Conflicts with `dedicated_acct_host_only`, `dedicated_host_id`, `dedicated_host_name` and `placement_group_id`
* `transient` - (Optional, Forces new resource, boolean) Specifies whether to provision a transient virtual server. The default value is `false`. Transient instances cannot be upgraded or downgraded. Transient instances cannot use local storage.  
    **NOTE**: Conflicts with `dedicated_acct_host_only`, `dedicated_host_id`, `dedicated_host_name`, `cores`, `memory`, `public_bandwidth_limited` and `public_bandwidth_unlimited`
* `os_reference_code` - (Optional, Forces new resource, string) The operating system reference code that is used to provision the computing instance. To see available OS reference codes, log in to the [IBM Cloud Classic Infrastructure (SoftLayer) API](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest_Block_Device_Template_Group/getVhdImportSoftwareDescriptions.json?objectMask=referenceCode), using your API key as the password.  
    **NOTE**: Conflicts with `image_id`.
*   `image_id` - (Optional, Forces new resource, integer) The image template ID you want to use to provision the computing instance. This is not the global identifier (UUID), but the image template group ID that should point to a valid global identifier. To retrieve the image template ID from the IBM Cloud infrastructure customer portal, navigate to **Devices > Manage > Images**, click the desired image, and note the ID number in the resulting URL.  

    **NOTE**: Conflicts with `os_reference_code`. If you don't know the ID(s) of your image templates, you can [refer to an image template ID by name using a data source](../d/compute_image_template.html).
*  `network_speed` - (Optional, integer) The connection speed (in Mbps) for the instance's network components. The default value is `100`.
*  `private_network_only` - (Optional, Forces new resource, boolean) When set to `true`, a compute instance only has access to the private network. The default value is `false`.
*  `private_security_group_ids` - (Optional, Forces new resource, array of integers) The ids of security groups to apply on the private interface.
This attribute can't be updated. This is provided so that you can apply security groups to  your VSI right from the beginning, the first time it comes up. If you would like to add or remove security groups in the future to this VSI then you should consider using `ibm_network_interface_sg_attachment` resource. If you use this attribute in addition to `ibm_network_interface_sg_attachment` resource you might get some spurious diffs. So use one of these consistently for a particular VSI.
*  `public_vlan_id` - (Optional, Forces new resource, integer) The public VLAN ID for the public network interface of the instance. Accepted values are in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the desired VLAN and note the ID number in the browser URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).  
    **NOTE**: Conflicts with `datacenter_choice`.
* `private_vlan_id` - (Optional, Forces new resource, integer) The private VLAN ID for the private network interface of the instance. You can find accepted values in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the desired VLAN and note the ID number in the browser URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).  
    **NOTE**: Conflicts with `datacenter_choice`.
* `public_security_group_ids` - (Optional, Forces new resource, array of integers) The ids of security groups to apply on the public interface.
This attribute can't be updated. This is provided so that you can apply security groups to  your VSI right from the beginning, the first time it comes up. If you would like to add or remove security groups in the future to this VSI then you should consider using `ibm_network_interface_sg_attachment` resource. If you use this attribute in addition to `ibm_network_interface_sg_attachment` resource you might get some spurious diffs. So use one of these consistently for a particular VSI.
* `public_subnet` - (Optional, Forces new resource, string) The public subnet for the public network interface of the instance. Accepted values are primary public networks. You can find accepted values in the [subnets doc](https://cloud.ibm.com/classic/network/subnets).
* `private_subnet` - (Optional, Forces new resource, string) The private subnet for the private network interface of the instance. Accepted values are primary private networks. You can find accepted values in the [subnets doc](https://cloud.ibm.com/classic/network/subnets).
* `disks` - (Optional, array of integers) The numeric disk sizes (in GBs) for the instance's block device and disk image settings. The default value is the smallest available capacity for the primary disk. If you specify an image template, the template provides the disk capacity. If you specify the flavorKeyName, first disk is provided by the flavor.
* `user_metadata` - (Optional, Forces new resource, string) Arbitrary data to be made available to the computing instance.
*  `notes` - (Optional, string) Descriptive text of up to 1000 characters about the VM instance.
* `ssh_key_ids` - (Optional, array of numbers) The SSH key IDs to install on the computing instance when the instance provisions.  
    **NOTE:** If you don't know the ID(s) for your SSH keys, you can [reference your SSH keys by their labels](../d/compute_ssh_key.html).
* `file_storage_ids` - (Optional, array of numbers) File storage to which this computing instance should have access. File storage must be in the same data center as the bare metal server. If you use this argument to authorize access to file storage, then do not use the `allowed_virtual_guest_ids` argument in the `ibm_storage_file` resource in order to prevent the same storage being added twice.
* `block_storage_ids` - (Optional, array of numbers) File storage to which this computing instance should have access. File storage must be in the same data center as the bare metal server. If you use this argument to authorize access to file storage, then do not use the `allowed_virtual_guest_ids` argument in the `ibm_storage_block` resource in order to prevent the same storage being added twice.
* `post_install_script_uri` - (Optional, Forces new resource, string) The URI of the script to be downloaded and executed after installation is complete.
* `tags` - (Optional, array of strings) Tags associated with the VM instance. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.
* `ipv6_enabled` - (Optional, Forces new resource, boolean) The primary public IPv6 address. The default value is `false`.
* `ipv6_static_enabled` - (Optional, boolean) The public static IPv6 address block of `/64`. The default value is `false`.
*  `secondary_ip_count` - (Optional, Forces new resource, integer) Specifies secondary public IPv4 addresses. Accepted values are `4` and `8`.
*  `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the VM instance to become available before declaring it as created. It is also the same amount of time waited for no active transactions before proceeding with an update or deletion. The default value is `90`.
* `public_bandwidth_limited` - (Optional, Forces new resource, int). Allowed public network traffic(GB) per month. It can be greater than 0 when the server is a monthly based server. Defaults to the smallest available capacity for the public bandwidth are used.  
    **NOTE**: Conflicts with `private_network_only` and `public_bandwidth_unlimited`.
* `public_bandwidth_unlimited` - (Optional, Forces new resource, boolean). Allowed unlimited public network traffic(GB) per month for a monthly based server. The `network_speed` should be 100 Mbps. Default value: `false`.  
    **NOTE**: Conflicts with `private_network_only` and `public_bandwidth_limited`.
* `evault` - (Optional, Forces new resource, int). Allowed evault(GB) per month for monthly based servers.
* `datacenter_choice` - (Optional, list) A nested block to describe datacenter choice options to retry on different datacenters and vlans. Nested `datacenter_choice` blocks must have the following structure:
    * `datacenter` - (Required, string) The datacenter in which you want to provision the instance.
    * `public_vlan_id` - (Optional, string) The public VLAN ID for the public network interface of the instance. Accepted values are in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the desired VLAN and note the ID number in the browser URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).
    * `private_vlan_id` - (Optional, Forces new resource, string) The private VLAN ID for the private network interface of the instance. You can find accepted values in the [VLAN doc](https://cloud.ibm.com/classic/network/vlans). Click the desired VLAN and note the ID number in the browser URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).   
    **NOTE**: Conflicts with `datacenter`, `private_vlan_id`, `public_vlan_id`, `placement_group_name` and `placement_group_id`.



## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VM instance.
* `ipv4_address` - The public IPv4 address of the VM instance.
* `ip_address_id_private` - The unique identifier for the private IPv4 address assigned to the VM instance.
* `ipv4_address_private` - The private IPv4 address of the VM instance.
* `ip_address_id` - The unique identifier for the public IPv4 address assigned to the VM instance.
* `ipv6_address` - The public IPv6 address of the VM instance provided when `ipv6_enabled` is set to `true`.
* `ipv6_address_id` - The unique identifier for the public IPv6 address assigned to the VM instance provided when `ipv6_enabled` is set to `true`.
* `private_subnet_id` - The unique identifier of the subnet `ipv4_address_private` belongs to.
* `public_ipv6_subnet` - The public IPv6 subnet provided when `ipv6_enabled` is set to `true`.
* `public_ipv6_subnet_id` - The unique identifier of the subnet `ipv6_address` belongs to.
* `public_subnet_id` - The unique identifier of the subnet `ipv4_address` belongs to.
* `secondary_ip_addresses` - The public secondary IPv4 addresses of the VM instance.
* `public_interface_id` - The ID of the primary public interface.
* `private_interface_id` - The ID of the primary private interface.
