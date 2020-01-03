---
layout: "ibm"
page_title: "Provider: IBM"
sidebar_current: "docs-ibm-index"
description: |-
  The IBM Cloud provider is used to interact with IBM Cloud resources.
---

# IBM Cloud Provider

The IBM Cloud provider is used to manage IBM Cloud resources. The provider must be configured with the proper credentials before it can be used.

Use the navigation menu on the left to read about the available data sources and resources.

## Example Usage


```hcl
# Configure the IBM Cloud Provider
provider "ibm" {
  ibmcloud_api_key      = var.ibm_bmx_api_key
  iaas_classic_username = var.ibm_sl_username
  iaas_classic_api_key  = var.ibm_sl_api_key
}

# Create an IBM Cloud infrastructure SSH key. You can find the SSH key surfaces in the infrastructure console under Devices > Manage > SSH Keys
resource "ibm_compute_ssh_key" "test_key_1" {
  label      = "test_key_1"
  public_key = var.ssh_public_key
}

# Create a virtual server with the SSH key
resource "ibm_compute_vm_instance" "my_server_2" {
  hostname          = "host-b.example.com"
  domain            = "example.com"
  ssh_key_ids       = [123456, ibm_compute_ssh_key.test_key_1.id]
  os_reference_code = "CENTOS_6_64"
  datacenter        = "ams01"
  network_speed     = 10
  cores             = 1
  memory            = 1024
}

# Reference details of the IBM Cloud space
data "ibm_space" "space" {
  space = var.space
  org   = var.org
}

# Create an instance of an IBM service
resource "ibm_service_instance" "service" {
  name       = var.instance_name
  space_guid = data.ibm_space.space.id
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}

# Create a Cloud Functions action
resource "ibm_function_action" "nodehello" {
  name = "action-name"
  exec {
    kind = "nodejs:6"
    code = file("hellonode.js")
  }
}

# Create a IS VPC and instance
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc1"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet1"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh1"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance1"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "b-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
  }

  vpc       = ibm_is_vpc.testacc_vpc.id
  zone      = "us-south-1"
  keys      = [ibm_is_ssh_key.testacc_sshkey.id]
  user_data = file("nginx.sh")
}

resource "ibm_is_floating_ip" "testacc_floatingip" {
  name   = "testfip"
  target = ibm_is_instance.testacc_instance.primary_network_interface[0].id
}

```

## Authentication

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

### Static credentials ###

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:

```hcl
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```


### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```hcl
provider "ibm" {}
```

Usage:

```shell
export IC_API_KEY="bmx_api_key"
export IAAS_CLASSIC_USERNAME="sl_username"
export IAAS_CLASSIC_API_KEY="sl_api_key"
terraform plan
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `ibmcloud_api_key` - (optional) The IBM Cloud platform API key. You must either add it as a credential in the provider block or source it from the `IC_API_KEY` (higher precedence) or `IBMCLOUD_API_KEY` environment variable. The key is required to provision Cloud Foundry or IBM Cloud Container Service resources, such as any resource that begins with `ibm` or `ibm_container`. `ibmcloud_api_key` will have higher precedence than `bluemix_api_key`.

* `bluemix_api_key` - (deprecated, optional) The IBM Cloud platform API key. You must either add it as a credential in the provider block or source it from the `BM_API_KEY` (higher precedence) or `BLUEMIX_API_KEY` environment variable. The key is required to provision Cloud Foundry or IBM Cloud Container Service resources, such as any resource that begins with `ibm` or `ibm_container`.

* `ibmcloud_timeout` - (optional) The timeout, expressed in seconds, for interacting with IBM Cloud APIs. You can also source the timeout from the `IC_TIMEOUT` (higher precedence) or `IBMCLOUD_TIMEOUT` environment variable. The default value is `60`. `ibmcloud_timeout` will have higher precedence than `bluemix_timeout`.

* `bluemix_timeout` - (deprecated, optional) The timeout, expressed in seconds, for interacting with IBM Cloud APIs. You can also source the timeout from the `BM_TIMEOUT` (higher precedence) or `BLUEMIX_TIMEOUT` environment variable. The default value is `60`.

* `softlayer_username` - (deprecated, optional) The IBM Cloud Classic Infrastructure (SoftLayer) user name. You must either add it as a credential in the provider block or source it from the `SL_USERNAME` (higher precedence) or `SOFTLAYER_USERNAME` environment variable. `iaas_classic_username` will have higher precedence than `softlayer_username`.

* `iaas_classic_username` - (optional) The IBM Cloud Classic Infrastructure (SoftLayer) user name. You must either add it as a credential in the provider block or source it from the `IAAS_CLASSIC_USERNAME`  environment variable.

* `softlayer_api_key` - (deprecated, optional) The IBM Cloud Classic Infrastructure API key. You must either add it as a credential in the provider block or source it from the `SL_API_KEY` (higher precedence) or `SOFTLAYER_API_KEY` environment variable. The key is required to provision infrastructure resources, such as any resource that begins with `ibm_compute`. `iaas_classic_api_key` will have higher precedence than `softlayer_api_key`.

* `iaas_classic_api_key` - (optional) The IBM Cloud Classic Infrastructure API key. You must either add it as a credential in the provider block or source it from the `IAAS_CLASSIC_API_KEY` environment variable.

* `softlayer_endpoint_url` - (deprecated, optional) The IBM Cloud Classic Infrastructure endpoint url. You can also source it from the `SL_ENDPOINT_URL` (higher precedence) or `SOFTLAYER_ENDPOINT_URL` environment variable. `iaas_classic_endpoint_url` will have higher precedence than `softlayer_endpoint_url`.

* `iaas_classic_endpoint_url` - (optional) The IBM Cloud Classic Infrastructure endpoint url. You can also source it from the `IAAS_CLASSIC_ENDPOINT_URL` environment variable. The default value is `https://api.softlayer.com/rest/v3`.

* `softlayer_timeout` - (optional) The timeout, expressed in seconds, for the IBM Cloud Classic Infrastructure APIs. You can also source the timeout from the `SL_TIMEOUT` (higher precedence) or `SOFTLAYER_TIMEOUT` environment variable. `iaas_classic_timeout` will have higher precedence than `softlayer_timeout`.

* `iaas_classic_timeout` - (optional) The timeout, expressed in seconds, for the IBM Cloud Clasic Infrastructure APIs. You can also source the timeout from the `IAAS_CLASSIC_TIMEOUT` environment variable. The default value is `60`.

* `region` - (optional) The IBM Cloud region. You can also source it from the `IC_REGION` (higher precedence) or `IBMCLOUD_REGION` `BM_REGION` `BLUEMIX_REGION` environment variable. The default value is `us-south`.

* `resource_group` - (optional) The Resource Group ID. You can also source it from the `IC_RESOURCE_GROUP` (higher precedence) or `IBMCLOUD_RESOURCE_GROUP` `BM_RESOURCE_GROUP` `BLUEMIX_RESOURCE_GROUP` environment variable.

* `max_retries` - (Optional) This is the maximum number of times an IBM Cloud infrastructure API call is retried, in the case where requests are getting network related timeout and rate limit exceeded error code. You can also source it from the `MAX_RETRIES` environment variable. The default value is `10`.

* `function_namespace` - (Optional) Your Cloud Functions namespace is composed from your IBM Cloud org and space like \<org\>_\<space\>. This attribute is required only when creating a Cloud Functions resource. It must be provided when you are creating such resources in IBM Cloud. You can also source it from the FUNCTION_NAMESPACE environment variable.

* `riaas_endpoint` - (deprected, Optional) The next generation infrastructure service API endpoint . It can also be sourced from the `RIAAS_ENDPOINT`. Default value: `us-south.iaas.cloud.ibm.com`. 

* `generation` - (Optional) The generation of Virtual Private Cloud. It can also be sourced from the `IC_GENERATION` (higher precedence) or `IBMCLOUD_GENERATION` environment variable. Default value: `2`. `1` for VPC Classic and `2` for VPC NextGen.


***Note***
The CloudFoundry endpoint has been updated in this release of IBM Cloud Terraform provider v0.17.4.  If you are using an earlier version of IBM Cloud Terraform provider, export the `IBMCLOUD_UAA_ENDPOINT` to the new authentication endpoint, as illustrated below

```shell
export IBMCLOUD_UAA_ENDPOINT="https://iam.cloud.ibm.com/cloudfoundry/login/<region>/"
```