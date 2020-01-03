---
layout: "ibm"
page_title: "IBM: ibm_security_group"
sidebar_current: "docs-ibm-datasource-security-group"
description: |-
  Get information about an IBM Security Group.
---

# ibm\_security_group

Import the details of an existing security group as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_security_group" "allow_ssh" {
    name = "allow_ssh"
}
```

The following example shows how you can use this data source to reference the security group IDs in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```hcl
resource "ibm_compute_vm_instance" "vm1" {
  # TF-UPGRADE-TODO: In Terraform v0.10 and earlier, it was sometimes necessary to
  # force an interpolation expression to be interpreted as a list by wrapping it
  # in an extra set of list brackets. That form was supported for compatibility in
  # v0.11, but is no longer supported in Terraform v0.12.
  #
  # If the expression in the following list itself returns a list, remove the
  # brackets to avoid interpretation as a list of lists. If the expression
  # returns a single list item then leave it as-is and remove this TODO comment.
  private_security_group_ids = [data.ibm_security_group.allow_ssh.id]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the security group, as it was defined in IBM Cloud Classic Infrastructure (SoftLayer).
* `description` - (Optional, string) The description of the security group, as it was defined in IBM Cloud Classic Infrastructure (SoftLayer).
* `most_recent` - (Optional, boolean) If more than one security group has the same name or description, you can set this argument to `true` to import only the most recent security group.
  **NOTE**: The search must return only one match, otherwise Terraform fails. Ensure that your name and description combinations are specific enough to return a single security group key only, or set the `most_recent` argument to `true`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the security group.
* `description` - The description of the security group.
