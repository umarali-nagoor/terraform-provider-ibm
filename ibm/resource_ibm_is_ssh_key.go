package ibm

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isKeyName          = "name"
	isKeyPublicKey     = "public_key"
	isKeyType          = "type"
	isKeyFingerprint   = "fingerprint"
	isKeyLength        = "length"
	isKeyResourceGroup = "resource_group"
)

func resourceIBMISSSHKey() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISSSHKeyCreate,
		Read:     resourceIBMISSSHKeyRead,
		Update:   resourceIBMISSSHKeyUpdate,
		Delete:   resourceIBMISSSHKeyDelete,
		Exists:   resourceIBMISSSHKeyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isKeyName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isKeyPublicKey: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			isKeyType: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isKeyFingerprint: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isKeyLength: {
				Type:     schema.TypeInt,
				Computed: true,
			},

			isKeyResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
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

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Key create")
	name := d.Get(isKeyName).(string)
	publickey := d.Get(isKeyPublicKey).(string)
	var rg string
	rg = ""
	if grp, ok := d.GetOk(isKeyResourceGroup); ok {
		rg = grp.(string)
	}

	keyC := compute.NewKeyClient(sess)
	key, err := keyC.Create(name, publickey, rg)
	if err != nil {
		log.Printf("[DEBUG] Key err %s", err)
		return err
	}

	d.SetId(key.ID.String())
	log.Printf("[INFO] Key : %s", key.ID.String())
	return resourceIBMISSSHKeyRead(d, meta)
}

func resourceIBMISSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	keyC := compute.NewKeyClient(sess)

	key, err := keyC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isKeyName, key.Name)
	d.Set(isKeyPublicKey, key.PublicKey)
	d.Set(isKeyType, key.Type)
	d.Set(isKeyFingerprint, key.Fingerprint)
	d.Set(isKeyLength, key.Length)
	d.Set(isKeyResourceGroup, key.ResourceGroup.ID)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/compute/sshKeys")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/compute/sshKeys")
	}
	d.Set(ResourceName, key.Name)
	d.Set(ResourceCRN, key.Crn)
	if key.ResourceGroup != nil {
		d.Set(ResourceGroupName, key.ResourceGroup.Name)
	}
	return nil
}

func resourceIBMISSSHKeyUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	keyC := compute.NewKeyClient(sess)

	if d.HasChange(isKeyName) {
		name := d.Get(isKeyName).(string)
		_, err := keyC.Update(d.Id(), name)
		if err != nil {
			return err
		}
	}

	return resourceIBMISSSHKeyRead(d, meta)
}

func resourceIBMISSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	keyC := compute.NewKeyClient(sess)
	err = keyC.Delete(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISSSHKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	keyC := compute.NewKeyClient(sess)

	_, err = keyC.Get(d.Id())
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
