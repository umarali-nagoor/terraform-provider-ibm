package ibm

import (
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	gouuid "github.com/satori/go.uuid"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	networkc "github.ibm.com/Bluemix/riaas-go-client/riaas/client/network"
)

const (
	isSecurityGroupName          = "name"
	isSecurityGroupVPC           = "vpc"
	isSecurityGroupRules         = "rules"
	isSecurityGroupResourceGroup = "resource_group"
)

func resourceIBMISSecurityGroup() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMISSecurityGroupCreate,
		Read:     resourceIBMISSecurityGroupRead,
		Update:   resourceIBMISSecurityGroupUpdate,
		Delete:   resourceIBMISSecurityGroupDelete,
		Exists:   resourceIBMISSecurityGroupExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			isSecurityGroupName: {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Security group name",
			},
			isSecurityGroupVPC: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Security group's resource group id",
				ForceNew:    true,
			},

			isSecurityGroupRules: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Security Rules",
				Elem: &schema.Resource{
					Schema: makeIBMISSecurityRuleSchema(),
				},
			},

			isSecurityGroupResourceGroup: {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func resourceIBMISSecurityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	parsed, err := parseIBMISSecurityGroupDictionary(d, "create")
	if err != nil {
		return err
	}
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	sgdef, err := makeIBMISSecurityGroupCreateParams(parsed)
	group, err := sgC.Create(*sgdef)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	d.SetId(group.ID.String())
	return resourceIBMISSecurityGroupRead(d, meta)
}

func resourceIBMISSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	group, err := sgC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set(isSecurityGroupName, group.Name)
	d.Set(isSecurityGroupVPC, group.Vpc.ID.String())
	rules := make([]map[string]interface{}, 0)
	if len(group.Rules) > 0 {
		for _, rule := range group.Rules {
			r := make(map[string]interface{})
			if rule.Code != nil {
				r[isSecurityGroupRuleCode] = int(*rule.Code)
			}
			if rule.Type != nil {
				r[isSecurityGroupRuleType] = int(*rule.Type)
			}
			if rule.PortMin != nil {
				r[isSecurityGroupRulePortMin] = int(*rule.PortMin)
			}
			if rule.PortMax != nil {
				r[isSecurityGroupRulePortMax] = int(*rule.PortMax)
			}
			r[isSecurityGroupRuleDirection] = rule.Direction
			r[isSecurityGroupRuleIPVersion] = rule.IPVersion
			if rule.Protocol != nil {
				r[isSecurityGroupRuleProtocol] = *rule.Protocol
			}

			rules = append(rules, r)
		}
	}
	d.Set(isSecurityGroupRules, rules)
	d.SetId(group.ID.String())
	if group.ResourceGroup != nil {
		d.Set(isSecurityGroupResourceGroup, group.ResourceGroup.ID)
	}
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/network/securityGroups")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/network/securityGroups")
	}
	d.Set(ResourceName, group.Name)
	d.Set(ResourceCRN, group.Crn)
	if group.ResourceGroup != nil {
		d.Set(ResourceGroupName, group.ResourceGroup.Name)
	}
	return nil
}

func resourceIBMISSecurityGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)
	if !d.HasChange(isSecurityGroupName) {
		return resourceIBMISSecurityGroupRead(d, meta)
	}

	name := d.Get(isSecurityGroupName).(string)

	_, err = sgC.Update(d.Id(), name)
	if err != nil {
		return err
	}
	err = resourceIBMISSecurityGroupRead(d, meta)
	return err
}

func resourceIBMISSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	sgC := network.NewSecurityGroupClient(sess)

	err = sgC.Delete(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return nil
			}
		}
		return err
	}
	return err
}

func resourceIBMISSecurityGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	sgC := network.NewSecurityGroupClient(sess)

	_, err = sgC.Get(d.Id())
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

type parsedIBMISSecurityGroupRule struct {
	// After parsing, unused string fields are set to
	// "" and unused int64 fields will be set to -1.
	// This ("" for unused strings and -1 for unused int64s)
	// is expected by our riaas API client.
	secgrpID       string
	ruleID         string
	direction      string
	ipversion      string
	remote         string
	remoteAddress  string
	remoteCIDR     string
	remoteSecGrpID string
	protocol       string
	icmpType       int64
	icmpCode       int64
	portMin        int64
	portMax        int64
}

func newParsedIBMISSecurityGroupRule() *parsedIBMISSecurityGroupRule {
	return &parsedIBMISSecurityGroupRule{
		icmpType: -1,
		icmpCode: -1,
		portMin:  -1,
		portMax:  -1,
	}
}

type parsedIBMISSecurityGroupDictionary struct {
	name          string
	resourceGroup string
	vpc           string
}

func newParsedIBMISSecurityGroupDictionary() *parsedIBMISSecurityGroupDictionary {
	p := &parsedIBMISSecurityGroupDictionary{}
	return p
}

func parseIBMISSecurityGroupDictionary(d *schema.ResourceData, tag string) (*parsedIBMISSecurityGroupDictionary, error) {
	parsed := newParsedIBMISSecurityGroupDictionary()
	parsed.name = d.Get(isSecurityGroupName).(string)
	parsed.vpc = d.Get(isSecurityGroupVPC).(string)
	if rg, ok := d.GetOk(isSecurityGroupResourceGroup); ok {
		parsed.resourceGroup = rg.(string)

	}

	return parsed, nil
}

func makeStrfmtUUID(s string) (strfmt.UUID, error) {
	uuid, err := gouuid.FromString(s)
	if err != nil {
		return strfmt.UUID(""), err
	}
	return strfmt.UUID(uuid.String()), nil
}

func makeIBMISSecurityGroupCreateParams(parsed *parsedIBMISSecurityGroupDictionary) (*networkc.PostSecurityGroupsBody, error) {
	params := &networkc.PostSecurityGroupsBody{}
	params.Name = parsed.name
	if parsed.resourceGroup != "" {
		rgref := networkc.PostSecurityGroupsParamsBodyResourceGroup{
			ID: strfmt.UUID(parsed.resourceGroup),
		}
		params.ResourceGroup = &rgref
	}

	params.Vpc = &networkc.PostSecurityGroupsParamsBodyVpc{ID: strfmt.UUID(parsed.vpc)}
	return params, nil
}

func makeIBMISSecurityRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		isSecurityGroupRuleDirection: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Direction of traffic to enforce, either inbound or outbound",
		},

		isSecurityGroupRuleIPVersion: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IP version: ipv4 or ipv6",
		},

		isSecurityGroupRuleRemote: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Security group id: an IP address, a CIDR block, or a single security group identifier",
		},

		isSecurityGroupRuleType: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRuleCode: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRulePortMin: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRulePortMax: {
			Type:     schema.TypeInt,
			Computed: true,
		},

		isSecurityGroupRuleProtocol: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
