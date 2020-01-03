package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPITenantDataSource_basic(t *testing.T) {

	//name := "Trial Tenant"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPITenantDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_pi_tenant.testacc_ds_tenant", "pi_cloud_instance_id", pi_cloud_instance_id),
				),
			},
		},
	})
}

func testAccCheckIBMPITenantDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_pi_tenant" "testacc_ds_tenant" {
    pi_cloud_instance_id = "%s"
}`, pi_cloud_instance_id)

}
