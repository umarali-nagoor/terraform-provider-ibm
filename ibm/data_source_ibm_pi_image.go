package ibm

import (
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceIBMPIImage() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPIImagesRead,
		Schema: map[string]*schema.Schema{

			helpers.PIImageName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Imagename Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PICloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"imageid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operatingsystem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIImagesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()

	if err != nil {
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	imageC := instance.NewIBMPIImageClient(sess, powerinstanceid)
	imagedata, err := imageC.Get(d.Get(helpers.PIImageName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("imageid", imagedata.ImageID)
	d.Set("state", imagedata.State)
	d.Set("size", imagedata.Size)
	d.Set("architecture", imagedata.Specifications.Architecture)
	d.Set("hypervisor", imagedata.Specifications.HypervisorType)
	d.Set("operatingsystem", imagedata.Specifications.OperatingSystem)

	return nil

}
