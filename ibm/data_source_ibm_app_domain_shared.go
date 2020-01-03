package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMAppDomainShared() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMAppDomainSharedRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description:  "The name of the shared domain",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateDomainName,
			},
		},
	}
}

func dataSourceIBMAppDomainSharedRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	domainName := d.Get("name").(string)
	shdomain, err := cfClient.SharedDomains().FindByName(domainName)
	if err != nil {
		return fmt.Errorf("Error retrieving shared domain: %s", err)
	}
	d.SetId(shdomain.GUID)
	return nil

}
