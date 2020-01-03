package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMIAMServicePolicy_Basic(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyUpdateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Service(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyService(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyUpdateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ResourceInstance(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Group(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyResourceGroup(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Type(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyResourceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_import(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
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

func TestAccIBMIAMServicePolicy_account_management(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMServicePolicyDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_service_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := idParts(policyID)
		if err != nil {
			return err
		}
		servicePolicyID := parts[1]

		// Try to find the key
		err = rsContClient.V1Policy().Delete(servicePolicyID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for service policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMServicePolicyExists(n string, obj iampapv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID

		parts, err := idParts(policyID)
		if err != nil {
			return err
		}
		servicePolicyID := parts[1]

		// Try to find the key
		policy, err := rsContClient.V1Policy().Get(servicePolicyID)
		obj = policy
		return nil
	}
}

func testAccCheckIBMIAMServicePolicyBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			tags           = ["tag1"]
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer", "Manager"]
			tags           = ["tag1", "tag2"]
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyService(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
		 	 service = "cloud-object-storage"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer", "Manager"]
	  
			resources {
		  		service = "kms"
		  		region  = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		 		 service              = "kms"
		  		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, name)
}

func testAccCheckIBMIAMServicePolicyResourceGroup(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			name = "Default"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
		 		service           = "containers-kubernetes"
		  		resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			name = "Default"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Administrator"]
	  
			resources {
		  		resource_type = "resource-group"
		  		resource      = data.ibm_resource_group.group.id
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyImport(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyAccountManagement(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
			roles              = ["Viewer"]
			account_management = true
	  	}

	`, name)
}
