package ibm

import (
	"fmt"
	"log"
	"strings"

	v2 "github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	iamUserSettingIamID              = "iam_id"
	iamUserSettingAllowedIPAddresses = "allowed_ip_addresses"
)

func resourceIBMUserSettings() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMUserSettingsCreate,
		Read:     resourceIBMIAMUserSettingsRead,
		Update:   resourceIBMIAMUserSettingsUpdate,
		Delete:   resourceIBMIAMUserSettingsDelete,
		Exists:   resourceIBMIAMUserSettingsExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			iamUserSettingIamID: {
				Description: "User's IAM ID or or email of user",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},

			iamUserSettingAllowedIPAddresses: {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    false,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of allowed IPv4 or IPv6 addresses ",
			},
		},
	}
}

func resourceIBMIAMUserSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	client := userManagement.UserInvite()

	userEmail := d.Get(iamUserSettingIamID).(string)

	//Read from Bluemix UserConfig
	accountID, err := getUserAccountID(d, meta)
	if err != nil {
		return err
	}

	user, err := getAccountUser(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	iamID := user.IbmUniqueId

	UserSettingsPayload := v2.UserSettingOptions{}

	if ip, ok := d.GetOk(iamUserSettingAllowedIPAddresses); ok && ip != nil {
		var ips = make([]string, 0)
		for _, i := range ip.([]interface{}) {
			ips = append(ips, i.(string))
		}
		ipStr := strings.Join(ips, ",")
		log.Printf("******* Create allowed ip string %s", ipStr)
		UserSettingsPayload.AllowedIPAddresses = ipStr
	}

	log.Printf("******* AccountId %s\n IamID %s\n Payload %v\n", accountID, iamID, UserSettingsPayload)
	_, UserSettingError := client.ManageUserSettings(accountID, iamID, UserSettingsPayload)

	/*if UserSettingError != nil {
		return UserSettingError
	}*/

	if UserSettingError != nil && !strings.Contains(UserSettingError.Error(), "EmptyResponseBody") {
		return fmt.Errorf("Received Empty body response %s", UserSettingError)
	}

	d.SetId(iamID)

	return resourceIBMIAMUserSettingsRead(d, meta)
}

func getUserAccountID(d *schema.ResourceData, meta interface{}) (string, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return "", err
	}
	return userDetails.userAccount, nil
}

func resourceIBMIAMUserSettingsRead(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	accountID, err := getUserAccountID(d, meta)
	if err != nil {
		return err
	}

	UserSettings, UserSettingError := Client.GetUserSettings(accountID, d.Id())
	if UserSettingError != nil {
		return UserSettingError
	}

	/*if UserSettingError != nil && !strings.Contains(UserSettingError.Error(), "EmptyResponseBody") {
		return fmt.Errorf("Received Empty body response %s", UserSettingError)
	}*/

	iplist := strings.Split(UserSettings.AllowedIPAddresses, ",")
	log.Printf("******* Read allowed ip list %v", iplist)
	d.Set(iamUserSettingAllowedIPAddresses, iplist)

	return nil

}

func resourceIBMIAMUserSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	// validate change
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	accountID, err := getUserAccountID(d, meta)
	if err != nil {
		return err
	}

	hasChanged := false

	userSettingPayload := v2.UserSettingOptions{}

	if d.HasChange(iamUserSettingAllowedIPAddresses) {
		if ip, ok := d.GetOk(iamUserSettingAllowedIPAddresses); ok && ip != nil {
			var ips = make([]string, 0)
			for _, i := range ip.([]interface{}) {
				ips = append(ips, i.(string))
			}
			ipStr := strings.Join(ips, ",")
			log.Printf("******* Update allowed ip string %s", ipStr)
			userSettingPayload.AllowedIPAddresses = ipStr
		}
		hasChanged = true
	}

	if hasChanged {
		_, UserSettingError := Client.ManageUserSettings(accountID, d.Id(), userSettingPayload)
		//Client.ManageUserSettings(accountID, d.Id(), userSettingPayload)
		if UserSettingError != nil && !strings.Contains(UserSettingError.Error(), "EmptyResponseBody") {
			return fmt.Errorf("Received Empty body response %s", UserSettingError)
		}
		/*if Error != nil {
			return Error
		}*/
	}

	return resourceIBMIAMUserSettingsRead(d, meta)
}

func resourceIBMIAMUserSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	accountID, err := getUserAccountID(d, meta)
	if err != nil {
		return err
	}
	userSettingPayload := v2.UserSettingOptions{}

	//Client.ManageUserSettings(accountID, d.Id(), userSettingPayload)

	_, UserSettingError := Client.ManageUserSettings(accountID, d.Id(), userSettingPayload)
	if UserSettingError != nil && !strings.Contains(UserSettingError.Error(), "EmptyResponseBody") {
		return fmt.Errorf("Received Empty body response %s", UserSettingError)
	}
	/*if Err != nil {
		return Err
	}*/

	return nil
}

func resourceIBMIAMUserSettingsExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return false, err
	}
	Client := userManagement.UserInvite()

	accountID, err := getUserAccountID(d, meta)
	if err != nil {
		return false, err
	}

	_, settingErr := Client.GetUserSettings(accountID, d.Id())

	if settingErr != nil {
		return false, settingErr
	}
	return true, nil
}
