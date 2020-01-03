---
layout: "ibm"
page_title: "IBM : Cloud Object Storage Bucket"
sidebar_current: "docs-ibm-resource-cos-bucket"
description: |-
  Manages IBM CloudObject Storage Bucket.
---

# ibm\_cos_bucket

Create or delete a bucket in a cloud object storage.

## Example Usage

In the following example, you can create a three buckets:

```hcl
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

resource "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

resource "ibm_cos_bucket" "standard-ams03" {
  bucket_name          = "a-standard-bucket-at-ams"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  single_site_location = "ams03"
  storage_class        = "standard"
}

resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "a-flex-bucket-at-us-south"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  region_location      = "us-south"
  storage_class        = "flex"
}

resource "ibm_cos_bucket" "cold-ap" {
  bucket_name           = "a-cold-bucket-at-ap"
  resource_instance_id  = ibm_resource_instance.cos_instance.id
  cross_region_location = "ap"
  storage_class         = "cold"
}
```

## Argument Reference

The following arguments are supported:

* `bucket_name` - (Required, string) The name of the bucket.
* `resource_instance_id` - (Required, string) The id of Cloud Object Storage instance.
* `key_protect` - (Optional, bool) CRN of the Key Protect instance where there a root key is already provisioned. Authorization required: [Docs](https://cloud.ibm.com/docs/services/cloud-object-storage?topic=cloud-object-storage-encryption#grant-service-authorization)
* `single_site_location` - (Optional,string) Location if single site bucket is desired. Accepted values: 'ams03', 'che01', 'hkg02', 'mel01', 'mex01', 'mil01', 'mon01', 'osl01', 'sjc04', 'sao01', 'seo01', 'tor01' Conflicts with: `region_location`, `cross_region_location`
* `region_location` - (Optional,string) Location if regional bucket is desired. Accepted values: 'au-syd', 'eu-de', 'eu-gb', 'jp-tok', 'us-east', 'us-south' Conflicts with: `single_site_location`, `cross_region_location`
* `cross_region_location` - (Optional,string) Location if cross regional bucket is desired. Accepted values: 'us', 'eu', 'ap' Conflicts with: `single_site_location`, `region_location`

* **Note** - One of the location option must be present.
* `storage_class` - (Required, string) Storage class of the bucket. Accepted values: 'standard', 'vault', 'cold', 'flex'

## Attribute Reference

The following attributes are exported:

* `id` - The ID of the bucket.
* `crn` - The CRN of the bucket.
* `resource_instance_id` - The id of Cloud Object Storage instance.
* `key_protect` - CRN of the Key Protect instance where there a root key is already provisioned.
* `single_site_location` - Location if single site bucket was created.
* `region_location` - Location if regional bucket was created.
* `cross_region_location` - Location if cross regional bucket was created.
* `storage_class` - Storage class of the bucket.

## Import

The `ibm_cos_bucket` resource can be imported using the `id`. The ID is formed from the `CRN` (Cloud Resource Name), the `bucket type` which must be `ssl` for single_site_location, `rl` for region_location or `crl` for cross_region_location, and the bucket location. The `CRN` and bucket location can be found on the portal.

id = $CRN:meta:$buckettype:$bucketlocation


```
$ terraform import ibm_cos_bucket.mybucket <crn>

$ terraform import ibm_cos_bucket.mybucket crn:v1:bluemix:public:cloud-object-storage:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3:bucket:mybucketname:meta:crl:eu
```
