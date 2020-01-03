package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/mitchellh/go-homedir"
)

func TestAccIBMContainerClusterConfigDataSource_basic(t *testing.T) {
	homeDir, err := homedir.Dir()
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	if err != nil {
		t.Fatalf("Error fetching homedir: %s", err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterDataSourceConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster_config.testacc_ds_cluster", "config_dir", homeDir),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "config_file_path"),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterCalicoConfigDataSource_basic(t *testing.T) {
	homeDir, err := homedir.Dir()
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	if err != nil {
		t.Fatalf("Error fetching homedir: %s", err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterCalicoConfigDataSource(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster_config.testacc_ds_cluster", "config_dir", homeDir),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "config_file_path"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "calico_config_file_path"),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterConfigDataSource_WithoutOptionalFields(t *testing.T) {
	homeDir, err := homedir.Dir()
	clusterName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	if err != nil {
		t.Fatalf("Error fetching homedir: %s", err)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterDataSourceConfigWithoutOptionalFields(clusterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_container_cluster_config.testacc_ds_cluster", "config_dir", homeDir),
					resource.TestCheckResourceAttrSet(
						"data.ibm_container_cluster_config.testacc_ds_cluster", "config_file_path"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterDataSourceConfigWithoutOptionalFields(clustername string) string {
	return fmt.Sprintf(`
	
resource "ibm_container_cluster" "testacc_cluster" {
    name = "%s"
    datacenter = "%s"
	machine_type = "%s"
	hardware       = "shared"
	public_vlan_id = "%s"
	private_vlan_id = "%s"
	region = "%s"
}
data "ibm_container_cluster_config" "testacc_ds_cluster" {
    cluster_name_id = "${ibm_container_cluster.testacc_cluster.id}"
    region = "%s"
}`, clustername, datacenter, machineType, publicVlanID, privateVlanID, csRegion, csRegion)
}

func testAccCheckIBMContainerClusterDataSourceConfig(clustername string) string {
	return fmt.Sprintf(`
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}

data "ibm_space" "testacc_ds_space" {
    org = "%s"
	space = "%s"
}

data "ibm_account" "testacc_acc" {
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
}


resource "ibm_container_cluster" "testacc_cluster" {
    name = "%s"
    datacenter = "%s"
	org_guid = "${data.ibm_org.testacc_ds_org.id}"
	space_guid = "${data.ibm_space.testacc_ds_space.id}"
	account_guid = "${data.ibm_account.testacc_acc.id}"

	machine_type = "%s"
	hardware       = "shared"
	public_vlan_id = "%s"
	private_vlan_id = "%s"
	region = "%s"
}
data "ibm_container_cluster_config" "testacc_ds_cluster" {
    cluster_name_id = "${ibm_container_cluster.testacc_cluster.id}"
	org_guid = "${data.ibm_org.testacc_ds_org.id}"
	space_guid = "${data.ibm_space.testacc_ds_space.id}"
	account_guid = "${data.ibm_account.testacc_acc.id}"
	region = "%s"
}`, cfOrganization, cfOrganization, cfSpace, clustername, datacenter, machineType, publicVlanID, privateVlanID, csRegion, csRegion)
}

func testAccCheckIBMContainerClusterCalicoConfigDataSource(clustername string) string {
	return fmt.Sprintf(`
resource "ibm_container_cluster" "testacc_cluster" {
    name = "%s"
    datacenter = "%s"
	machine_type = "%s"
	hardware       = "shared"
	public_vlan_id = "%s"
	private_vlan_id = "%s"
	region = "%s"
}
data "ibm_container_cluster_config" "testacc_ds_cluster" {
    cluster_name_id = "${ibm_container_cluster.testacc_cluster.id}"
	region = "%s"
	network = true
}`, clustername, datacenter, machineType, publicVlanID, privateVlanID, csRegion, csRegion)
}
