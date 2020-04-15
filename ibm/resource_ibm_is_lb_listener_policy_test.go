package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func TestAccIBMISLBListenerPolicy_basic(t *testing.T) {
	var policyID string
	vpcname := fmt.Sprintf("terraformLBLisuat-vpc-%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformLBLisuat-subnet-%d", acctest.RandInt())
	lbname := fmt.Sprintf("tflblisuat%d", acctest.RandInt())
	lblistenerpolicyname1 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandInt())
	lblistenerpolicyname2 := fmt.Sprintf("tflblisuat-listener-policy-%d", acctest.RandInt())

	priority1 := "1"
	protocol := "http"
	port := "8080"
	action := "forward"
	priority2 := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBListenerPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port, protocol, lblistenerpolicyname1, action, priority1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener_policy.testacc_lb_listener_policy", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", lbname),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "priority", priority1),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBListenerPolicyConfigUpdate(vpcname, subnetname, ISZoneName, ISCIDR, lbname, port, protocol, lblistenerpolicyname2, priority2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBListenerPolicyExists("ibm_is_lb_listener.testacc_lb_listener", policyID),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "name", lblistenerpolicyname2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_listener_policy.testacc_lb_listener_policy", "proprity", priority2),
				),
			},
		},
	})
}

func testAccCheckIBMISLBListenerPolicyDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb_listener_policy" {
				continue
			}

			if rs.Primary.ID == "" {
				return errors.New("No Record ID is set")
			}

			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			lbID := parts[0]
			lbListenerID := parts[1]
			policyID := parts[2]

			getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
				LoadBalancerID: &lbID,
				ListenerID:     &lbListenerID,
				ID:             &policyID,
			}

			policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

			if err == nil {
				return fmt.Errorf("LBLIstenerPolicy still exists: %s %v", rs.Primary.ID, policy)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb_listener_policy" {
				continue
			}

			if rs.Primary.ID == "" {
				return errors.New("No Record ID is set")
			}

			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			lbID := parts[0]
			lbListenerID := parts[1]
			policyID := parts[2]

			getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
				LoadBalancerID: &lbID,
				ListenerID:     &lbListenerID,
				ID:             &policyID,
			}

			policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

			if err == nil {
				return fmt.Errorf("LBLIstenerPolicy still exists: %s %v", rs.Primary.ID, policy)
			}
		}
	}

	return nil
}

func testAccCheckIBMISLBListenerPolicyExists(n string, policyID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbListenerID := parts[1]
		policyID := parts[2]

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()

			getLbListenerPolicyOptions := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{
				LoadBalancerID: &lbID,
				ListenerID:     &lbListenerID,
				ID:             &policyID,
			}

			policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

			if err != nil {
				return err
			}

			policyID = *policy.ID

		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getLbListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
				LoadBalancerID: &lbID,
				ListenerID:     &lbListenerID,
				ID:             &policyID,
			}

			policy, _, err := sess.GetLoadBalancerListenerPolicy(getLbListenerPolicyOptions)

			if err != nil {
				return err
			}

			policyID = *policy.ID
		}
		return nil
	}
}

func testAccCheckIBMISLBListenerPolicyConfig(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, action, priority string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["ibm_is_subnet.testacc_subnet.id"]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = ibm_is_lb.testacc_LB.id
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "%s"
		priority = %s
		name = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, action, priority, lblistenerpolicyname)

}

func testAccCheckIBMISLBListenerPolicyConfigUpdate(vpcname, subnetname, zone, cidr, lbname, port, protocol, lblistenerpolicyname, priority string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_listener" "testacc_lb_listener" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		port = %s
		protocol = "%s"
	}

	resource "ibm_is_lb_listener_policy" "testacc_lb_listener_policy" {
        lb = ibm_is_lb.testacc_LB.id
        listener = ibm_is_lb_listener.testacc_lb_listener.listener_id
        action = "forward"
		priority = %s
		name = "%s"
}`, vpcname, subnetname, zone, cidr, lbname, port, protocol, priority, lblistenerpolicyname)

}
