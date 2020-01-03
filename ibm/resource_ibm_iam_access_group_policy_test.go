package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"

	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMIAMAccessGroupPolicy_Basic(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyUpdateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Service(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyService(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyUpdateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_ResourceInstance(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Resource_Group(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceGroup(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Resource_Type(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_import(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_access_group_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_account_management(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_access_group_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Attributese(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "is"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupPolicyDestroy(s *terraform.State) error {
	iampapClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := idParts(policyID)
		if err != nil {
			return err
		}

		accgrpPolicyID := parts[1]

		err = iampapClient.V1Policy().Delete(accgrpPolicyID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for access group policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAccessGroupPolicyExists(n string, obj iampapv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iampapClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID

		parts, err := idParts(policyID)
		if err != nil {
			return err
		}

		accgrpPolicyID := parts[1]

		policy, err := iampapClient.V1Policy().Get(accgrpPolicyID)
		obj = policy
		return nil
	}
}

func testAccCheckIBMIAMAccessGroupPolicyBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
  			name = "%s"
		}

		resource "ibm_iam_access_group_policy" "policy" {
  			access_group_id = ibm_iam_access_group.accgrp.id
  			roles           = ["Viewer"]
  			tags            = ["tag1"]
		}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyUpdateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer", "Manager"]
			tags            = ["tag1", "tag2"]
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyService(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
  		}

		resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles        = ["Viewer"]

			resources {
		  	service = "cloud-object-storage"
			}
		  }
		  
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyUpdateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer", "Manager"]
	  
			resources {
		 	 service = "kms"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		 	 service              = "kms"
		 	 resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceGroup(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			name = "Default"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
	  
			resources {
		 	 service           = "containers-kubernetes"
		 	 resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			name = "Default"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Administrator"]
	  
			resources {
		  		resource_type = "resource-group"
		  		resource      = data.ibm_resource_group.group.id
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyImport(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	 	 }
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
	  	}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyAccountManagement(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id    = ibm_iam_access_group.accgrp.id
			roles              = ["Administrator"]
			account_management = true
	  	}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyAttributes(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
	  
			resources {
		  	service = "is"
		  	attributes = {
				"vpcId" = "*"
		  	}
			}
	  	}

	`, name)
}
