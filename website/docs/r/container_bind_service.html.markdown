---
layout: "ibm"
page_title: "IBM: container_bind_service"
sidebar_current: "docs-ibm-resource-container-bind-service"
description: |-
  Manages IBM container cluster.
---

# ibm\_container_bind_service

Bind an IBM service to a Kubernetes namespace. With this resource, you can attach an existing service to an existing Kubernetes cluster.

## Example Usage

In the following example, you can bind a service to a cluster:

```hcl
resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id       = "cluster_name"
  service_instance_name = "service_name"
  namespace_id          = "default"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name_id` - (Required, string) The name or ID of the cluster.
* `service_instance_space_guid` - (Removed) The space GUID associated with the service instance. This attribute is removed.
* `service_instance_name_id` - (Removed) The name or ID of the service that you want to bind to the cluster. This attribute is removed, use `service_instance_id` or `service_instance_name` instead.
* `namespace_id` - (Required, string) The Kubernetes namespace.
* `service_instance_name` - (Optional, string) The name of the service that you want to bind to the cluster. Conflicts with `service_instance_id`.
* `service_instance_id` - (Optional, string) The ID of the service that you want to bind to the cluster. Conflicts with `service_instance_name`.
* `key` - (Optional, string) Specify an existing service key to use for the service binding.
* `role` - (Optional, string) Specify the IAM role for the service key. This flag does not work if you specify an existing key to use or for services that are not IAM-enabled, such as Cloud Foundry services.
* `org_guid` - (Deprecated, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
* `space_guid` - (Deprecated, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Deprecated, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
* `region` - (Deprecated, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `tags` - (Optional, array of strings) Tags associated with the container bind service instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the bind service resource. The id is composed of \<cluster_name_id\>/\<service_instance_name or service_instance_id\>/\<namespace_id/>
* `namespace_id` -  The Kubernetes namespace.
* `secret_name` - The secret name. This field is removed.

## Import

ibm_container_bind_service can be imported using cluster_name_id, service_instance_name or service_instance_id and namespace_id, eg

```
$ terraform import ibm_container_bind_service.example mycluster/myservice/default