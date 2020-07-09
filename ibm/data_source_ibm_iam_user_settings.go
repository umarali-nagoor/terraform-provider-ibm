package ibm

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMIAMUserSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMUserSettingsRead,

		Schema: map[string]*schema.Schema{

			"iam_id": {
				Description: "User's IAM ID or or email of user",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			"allowed_ip_addresses": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of allowed IPv4 or IPv6 addresses ",
			},
		},
	}
}

func dataSourceIBMIAMUserSettingsRead(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	userEmail := d.Get("iam_id").(string)

	accountID, err := getUserAccountID(d, meta)
	if err != nil {
		return err
	}

	user, err := getAccountUser(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	iamID := user.IbmUniqueId

	UserSettings, UserSettingError := Client.GetUserSettings(accountID, iamID)
	if UserSettingError != nil {
		return UserSettingError
	}

	iplist := strings.Split(UserSettings.AllowedIPAddresses, ",")
	d.Set("allowed_ip_addresses", iplist)
	d.SetId(userEmail)

	return nil
}
