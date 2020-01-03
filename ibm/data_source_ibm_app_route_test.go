package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMAppRouteDataSource_basic(t *testing.T) {
	host := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppRouteDataSourceConfig(host),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_app_route.testacc_route", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppRouteDataSourceConfig(host string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		org   = "%s"
		space = "%s"
	  }
	  
	  data "ibm_app_domain_shared" "domain" {
		name = "mybluemix.net"
	  }
	  
	  resource "ibm_app_route" "route" {
		domain_guid = data.ibm_app_domain_shared.domain.id
		space_guid  = data.ibm_space.spacedata.id
		host        = "%s"
		path        = "/app"
	  }
	  
	  data "ibm_app_route" "testacc_route" {
		domain_guid = ibm_app_route.route.domain_guid
		space_guid  = ibm_app_route.route.space_guid
		host        = ibm_app_route.route.host
		path        = ibm_app_route.route.path
	  }
	`, cfOrganization, cfSpace, host)

}
