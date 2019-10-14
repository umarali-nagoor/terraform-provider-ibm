data "ibm_resource_group" "group" {
  //name = "test"
  is_default = "true"
}

resource "ibm_cis" "cis_instance" {
  name = "test"
  plan = "standard"

  //resource_group_id = "1"
  resource_group_id = data.ibm_resource_group.group.id
  tags              = ["tag1", "tag2"]
  location          = "global"
  timeouts {
    create = "65m"
    update = "65m"
    delete = "65m"
  }
}

