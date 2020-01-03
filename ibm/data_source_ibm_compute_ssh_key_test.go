package ibm

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMComputeSSHKeyDataSource_basic(t *testing.T) {
	label := fmt.Sprintf("ssh_key_test_ds_label_%d", acctest.RandInt())
	notes := fmt.Sprintf("ssh_key_test_ds_notes_%d", acctest.RandInt())

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeSSHKeyDataSourceConfig(label, notes, publicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_compute_ssh_key.testacc_ds_ssh_key", "public_key", publicKey),
					resource.TestCheckResourceAttr("data.ibm_compute_ssh_key.testacc_ds_ssh_key", "notes", notes),
					resource.TestMatchResourceAttr("data.ibm_compute_ssh_key.testacc_ds_ssh_key", "fingerprint", regexp.MustCompile("^[0-9a-f]{2}:")),
				),
			},
		},
	})
}

func testAccCheckIBMComputeSSHKeyDataSourceConfig(label, notes, publicKey string) string {
	return fmt.Sprintf(`
resource "ibm_compute_ssh_key" "testacc_ds_ssh_key" {
    label = "%s"
    notes = "%s"
    public_key = "%s"
}
data "ibm_compute_ssh_key" "testacc_ds_ssh_key" {
    label = "${ibm_compute_ssh_key.testacc_ds_ssh_key.label}"
}`, label, notes, publicKey)
}
