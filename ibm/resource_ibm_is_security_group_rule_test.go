package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

func TestAccIBMISSecurityGroupRule_basic(t *testing.T) {
	var securityGroupRule *models.SecurityGroupRule

	vpcname := fmt.Sprintf("terraformsecurityruleuat-vpc-%d", acctest.RandInt())
	name1 := fmt.Sprintf("terraformsecurityruleuat-create-step-name-%d", acctest.RandInt())
	//name2 := fmt.Sprintf("terraformsecurityuat-update-step-name-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISSecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISsecurityGroupRuleConfig(vpcname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISSecurityGroupRuleExists("ibm_is_security_group_rule.testacc_security_group_rule_all", &securityGroupRule),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_security_group", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISSecurityGroupRuleDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	securityGroupC := network.NewSecurityGroupClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_security_group_rule" {
			continue
		}

		secgrpID, ruleID, err := parseISTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = securityGroupC.GetRule(secgrpID, ruleID)

		if err == nil {
			return fmt.Errorf("securitygrouprule still exists: %s", ruleID)
		}
	}

	return nil
}

func testAccCheckIBMISSecurityGroupRuleExists(n string, securityGroup **models.SecurityGroupRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		securityGroupC := network.NewSecurityGroupClient(sess)
		secgrpID, ruleID, err := parseISTerraformID(rs.Primary.ID)
		if err != nil {
			return err
		}

		foundsecurityGroup, err := securityGroupC.GetRule(secgrpID, ruleID)

		if err != nil {
			return err
		}

		*securityGroup = foundsecurityGroup
		return nil
	}
}

func testAccCheckIBMISsecurityGroupRuleConfig(vpcname, name string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "%s"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_all" {
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "inbound"
	remote = "127.0.0.1"
 }
 
 resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp" {
	depends_on =["ibm_is_security_group_rule.testacc_security_group_rule_all"]
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "inbound"
	remote = "127.0.0.1"
	icmp = {
		code = 20
		type = 30
	}

 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
	depends_on =["ibm_is_security_group_rule.testacc_security_group_rule_icmp"]
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "inbound"
	remote = "127.0.0.1"
	udp = {
		port_min = 805
		port_max = 807
	}
 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp" {
	depends_on =["ibm_is_security_group_rule.testacc_security_group_rule_udp"]
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "inbound"
	remote = "127.0.0.1"
	tcp = {
		port_min = 8080
		port_max = 8080
	}
 }
 `, vpcname, name)

}
