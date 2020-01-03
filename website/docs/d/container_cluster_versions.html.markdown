---
layout: "ibm"
page_title: "IBM: ibm_container_cluster_versions"
sidebar_current: "docs-ibm-datasource-container-cluster-versions"
description: |-
  List supported kubernetes versions on IBM Cloud.
---

# ibm\_container_cluster_versions

Get the list of supported kubernetes versions on IBM Cloud. Please refer to https://cloud.ibm.com/docs/containers/cs_versions.html#cs_versions for detail instructions.

## Example Usage

```hcl
data "ibm_container_cluster_versions" "cluster_versions" {
  region = "eu-de"
}
```

## Argument Reference

The following arguments are supported:

* `org_guid` - (Deprecated, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
* `space_guid` - (Deprecated, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Deprecated, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
* `region` - (Deprecated, string) The region to target. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster versions.
* `valid_kube_versions` - The supported kubernetes versions on IBM Cloud.
* `valid_openshift_versions` - The supported openshift versions on IBM Cloud.