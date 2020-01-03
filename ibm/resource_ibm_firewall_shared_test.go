package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMFirewallShared_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMHardwareFirewallSharedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMFirewallShared_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_hardware_firewall_shared.test_firewall", "firewall_type", "10MBPS_HARDWARE_FIREWALL"),
					resource.TestCheckResourceAttr(
						"ibm_hardware_firewall_shared.test_firewall", "virtual_instance_id", "1234567"),
					resource.TestCheckResourceAttr(
						"ibm_hardware_firewall_shared.test_firewall", "guest_type", "virtual machine"),
				),
			},
		},
	})
}

func testAccCheckIBMHardwareFirewallSharedDestroy(s *terraform.State) error {
	service := services.GetNetworkComponentFirewallService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_hardware_firewall_shared" {
			continue
		}

		fwId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the domain
		_, err := service.Id(fwId).GetObject()

		if err == nil {
			return fmt.Errorf("Hardware Firewall Shared with id %d still exists", fwId)
		}
	}

	return nil
}

const testAccCheckIBMFirewallShared_basic = `
resource "ibm_compute_vm_instance" "fwvm1" {
    hostname = "testing"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "sjc01"
    network_speed = 100
    hourly_billing = false
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_hardware_firewall_shared" "test_firewall" {
	firewall_type = "100MBPS_HARDWARE_FIREWALL"
	virtual_instance_id = "${ibm_compute_vm_instance.fwvm1.id}"
	guest_type="virtual machine"}`
