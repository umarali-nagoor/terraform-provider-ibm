package ibm

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isPublicGatewayName              = "name"
	isPublicGatewayFloatingIP        = "floating_ip"
	isPublicGatewayStatus            = "status"
	isPublicGatewayVPC               = "vpc"
	isPublicGatewayZone              = "zone"
	isPublicGatewayFloatingIPAddress = "address"

	isPublicGatewayProvisioning     = "provisioning"
	isPublicGatewayProvisioningDone = "available"
	isPublicGatewayDeleting         = "deleting"
	isPublicGatewayDeleted          = "done"
)

func resourceIBMISPublicGateway() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISPublicGatewayCreate,
		Read:     resourceIBMISPublicGatewayRead,
		Update:   resourceIBMISPublicGatewayUpdate,
		Delete:   resourceIBMISPublicGatewayDelete,
		Exists:   resourceIBMISPublicGatewayExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isPublicGatewayName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isPublicGatewayFloatingIP: {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						isPublicGatewayFloatingIPAddress: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			isPublicGatewayStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isPublicGatewayVPC: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			isPublicGatewayZone: {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISPublicGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	name := d.Get(isPublicGatewayName).(string)
	vpc := d.Get(isPublicGatewayVPC).(string)
	zone := d.Get(isPublicGatewayZone).(string)
	floatingipID := ""
	floatingipadd := ""
	if floatingipdataIntf, ok := d.GetOk(isPublicGatewayFloatingIP); ok {
		floatingipdata := floatingipdataIntf.(map[string]interface{})
		floatingipidintf, ispresent := floatingipdata["id"]
		if ispresent {
			floatingipID = floatingipidintf.(string)
		}
		floatingipaddintf, ispresentadd := floatingipdata[isPublicGatewayFloatingIPAddress]
		if ispresentadd {
			floatingipadd = floatingipaddintf.(string)
		}
	}

	publicgwC := network.NewPublicGatewayClient(sess)
	publicgw, err := publicgwC.Create(name, zone, vpc, floatingipID, floatingipadd)
	if err != nil {
		return err
	}

	d.SetId(publicgw.ID.String())
	log.Printf("[INFO] PublicGateway : %s", publicgw.ID.String())

	_, err = isWaitForPublicGatewayAvailable(publicgwC, d.Id(), d)
	if err != nil {
		return err
	}

	return resourceIBMISPublicGatewayRead(d, meta)
}

func isWaitForPublicGatewayAvailable(publicgwC *network.PublicGatewayClient, id string, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for public gateway (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isPublicGatewayProvisioning},
		Target:     []string{isPublicGatewayProvisioningDone},
		Refresh:    isPublicGatewayRefreshFunc(publicgwC, id),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPublicGatewayRefreshFunc(publicgwC *network.PublicGatewayClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		publicgw, err := publicgwC.Get(id)
		if err != nil {
			return nil, "", err
		}

		// if its still pending, returning provisioning
		if publicgw.Status == "pending" {
			return publicgw, isPublicGatewayProvisioning, nil
		}

		log.Printf("[Debug] state = %s", publicgw.Status)
		log.Printf("[Debug] gw = %s", publicgw.FloatingIP)
		return publicgw, isPublicGatewayProvisioningDone, nil
	}
}

func resourceIBMISPublicGatewayRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	publicgwC := network.NewPublicGatewayClient(sess)

	publicgw, err := publicgwC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set("id", publicgw.ID.String())
	d.Set(isPublicGatewayName, publicgw.Name)
	if publicgw.FloatingIP != nil {

		floatIP := map[string]interface{}{
			"id":                             publicgw.FloatingIP.ID.String(),
			isPublicGatewayFloatingIPAddress: publicgw.FloatingIP.Address,
		}
		d.Set(isPublicGatewayFloatingIP, floatIP)

	}

	d.Set(isPublicGatewayStatus, publicgw.Status)
	d.Set(isPublicGatewayZone, publicgw.Zone.Name)
	d.Set(isPublicGatewayVPC, publicgw.Vpc.ID.String())

	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/network/publicGateways")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/network/publicGateways")
	}
	d.Set(ResourceName, publicgw.Name)
	d.Set(ResourceCRN, publicgw.Crn)
	d.Set(ResourceStatus, publicgw.Status)
	if publicgw.ResourceGroup != nil {
		d.Set(ResourceGroupName, publicgw.ResourceGroup.Name)
	}
	return nil
}

func resourceIBMISPublicGatewayUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	publicgwC := network.NewPublicGatewayClient(sess)

	name := ""
	if d.HasChange(isPublicGatewayName) {
		name = d.Get(isPublicGatewayName).(string)
	}

	_, err = publicgwC.Update(d.Id(), name)
	if err != nil {
		return err
	}

	return resourceIBMISPublicGatewayRead(d, meta)
}

func resourceIBMISPublicGatewayDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	publicgwC := network.NewPublicGatewayClient(sess)
	err = publicgwC.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForPublicGatewayDeleted(publicgwC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForPublicGatewayDeleted(pg *network.PublicGatewayClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for public gateway (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isPublicGatewayDeleting},
		Target:     []string{},
		Refresh:    isPublicGatewayDeleteRefreshFunc(pg, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isPublicGatewayDeleteRefreshFunc(pg *network.PublicGatewayClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		publicGateway, err := pg.Get(id)
		if err == nil {
			return publicGateway, isPublicGatewayDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("[DEBUG] returning deleted")
				return nil, isPublicGatewayDeleted, nil
			}
		}
		log.Printf("[DEBUG] returning x")
		return nil, isPublicGatewayDeleting, err
	}
}

func resourceIBMISPublicGatewayExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	publicgwC := network.NewPublicGatewayClient(sess)

	_, err = publicgwC.Get(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
