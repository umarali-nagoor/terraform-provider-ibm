package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisDomainDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisDomainDataSourceConfig_basic1("test_acc", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis_domain.cis_domain", "status", "pending"),
					resource.TestCheckResourceAttr("data.ibm_cis_domain.cis_domain", "original_name_servers.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_cis_domain.cis_domain", "name_servers.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceName string, domain string) string {
	return fmt.Sprintf(`
	
	data "ibm_cis_domain" "cis_domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "test123.com"
	}
	  
	data "ibm_resource_group" "test_acc" {
		name = "Default"
	}
	  
	data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "test-domain"
	}
	`)
}
