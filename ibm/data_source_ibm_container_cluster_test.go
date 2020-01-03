package ibm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMContainerClusterDataSource_basic(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	serviceName := fmt.Sprintf("terraform-%d", acctest.RandInt())
	serviceKeyName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterDataSource(clusterName, serviceName, serviceKeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "is_trusted", "false"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "bounded_services.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "albs.#", "2"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster.testacc_ds_cluster", "resource_group_id"),
					testAccIBMClusterVlansCheck("data.ibm_container_cluster.testacc_ds_cluster"),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterDataSourceWithOutOrgSpaceAccount(t *testing.T) {
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterDataSourceWithOutOrgSpaceAccount(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "worker_count", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "worker_pools.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_container_cluster.testacc_ds_cluster", "albs.#", "2"),
					testAccIBMClusterVlansCheck("data.ibm_container_cluster.testacc_ds_cluster"),
				),
			},
		},
	})
}

func testAccIBMClusterVlansCheck(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		var (
			vlansSize int
			err       error
		)

		if vlansSize, err = strconv.Atoi(a["vlans.#"]); err != nil {
			return err
		}
		if vlansSize < 1 {
			return fmt.Errorf("No subnets found")
		}
		return nil
	}
}

func testAccCheckIBMContainerClusterDataSource(clusterName, serviceName, serviceKeyName string) string {
	return fmt.Sprintf(`
data "ibm_org" "testacc_ds_org" {
  org = "%s"
}

data "ibm_space" "testacc_ds_space" {
  org   = "%s"
  space = "%s"
}

data "ibm_account" "testacc_acc" {
  org_guid = data.ibm_org.testacc_ds_org.id
}

resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  org_guid        = data.ibm_org.testacc_ds_org.id
  space_guid      = data.ibm_space.testacc_ds_space.id
  account_guid    = data.ibm_account.testacc_acc.id
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  subnet_id       = ["%s"]
}

resource "ibm_service_instance" "service" {
  name       = "%s"
  space_guid = data.ibm_space.testacc_ds_space.id
  service    = "speech_to_text"
  plan       = "lite"
  tags       = ["cluster-service", "cluster-bind"]
}

resource "ibm_service_key" "serviceKey" {
  name                  = "%s"
  service_instance_guid = ibm_service_instance.service.id
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = ibm_container_cluster.testacc_cluster.name
  service_instance_id = ibm_service_instance.service.id
  namespace_id        = "default"
  org_guid            = data.ibm_org.testacc_ds_org.id
  space_guid          = data.ibm_space.testacc_ds_space.id
  account_guid        = data.ibm_account.testacc_acc.id
  region              = "%s"
}

data "ibm_container_cluster" "testacc_ds_cluster" {
  org_guid        = data.ibm_org.testacc_ds_org.id
  space_guid      = data.ibm_space.testacc_ds_space.id
  account_guid    = data.ibm_account.testacc_acc.id
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
  region          = "%s"
  depends_on      = [ibm_container_bind_service.bind_service]
}
`, cfOrganization, cfOrganization, cfSpace, clusterName, datacenter, machineType, publicVlanID, privateVlanID, subnetID, serviceName, serviceKeyName, csRegion, csRegion)
}

func testAccCheckIBMContainerClusterDataSourceWithOutOrgSpaceAccount(clusterName string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "%s"
  datacenter      = "%s"
  machine_type    = "%s"
  hardware        = "shared"
  public_vlan_id  = "%s"
  private_vlan_id = "%s"
  region          = "%s"
}

data "ibm_container_cluster" "testacc_ds_cluster" {
  cluster_name_id = ibm_container_cluster.testacc_cluster.id
  region          = "%s"
}
`, clusterName, datacenter, machineType, publicVlanID, privateVlanID, csRegion, csRegion)
}
