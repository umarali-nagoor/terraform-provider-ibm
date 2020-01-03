package ibm

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	token "github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam/token"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var bucketTypes = []string{"single_site_location", "region_location", "cross_region_location"}

func dataSourceIBMCosBucket() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCosBucketRead,

		Schema: map[string]*schema.Schema{
			"bucket_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(bucketTypes),
				Required:     true,
			},
			"bucket_region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of resource instance",
			},
			"key_protect": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of the key you want to use data at rest encryption",
			},
			"single_site_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"region_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cross_region_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_class": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"s3_endpoint_public": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public endpoint for the COS bucket",
			},
			"s3_endpoint_private": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Private endpoint for the COS bucket",
			},
		},
	}
}

func dataSourceIBMCosBucketRead(d *schema.ResourceData, meta interface{}) error {
	var s3Conf *aws.Config
	rsConClient, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	bucketName := d.Get("bucket_name").(string)
	serviceID := d.Get("resource_instance_id").(string)
	bucketType := d.Get("bucket_type").(string)
	bucketRegion := d.Get("bucket_region").(string)

	apiEndpoint, apiEndpointPrivate := selectCosApi(bucketLocationConvert(bucketType), bucketRegion)
	authEndpoint, err := rsConClient.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
	apiKey := rsConClient.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, serviceID)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := rsConClient.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  rsConClient.Config.IAMAccessToken,
				RefreshToken: rsConClient.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, serviceID)).WithS3ForcePathStyle(true)
	}
	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	headInput := &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	}
	err = s3Client.WaitUntilBucketExists(headInput)
	if err != nil {
		return fmt.Errorf("failed waiting for bucket %s to be created, %v",
			bucketName, err)
	}
	bucketLocationInput := &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	}
	bucketLocationConstraint, err := s3Client.GetBucketLocation(bucketLocationInput)
	if err != nil {
		return err
	}
	bLocationConstraint := *bucketLocationConstraint.LocationConstraint

	singleSiteLocationRegex, err := regexp.Compile("^[a-z]{3}[0-9][0-9]-[a-z]{4,8}$")
	if err != nil {
		return err
	}
	regionLocationRegex, err := regexp.Compile("^[a-z]{2}-[a-z]{2,5}-[a-z]{4,8}$")
	if err != nil {
		return err
	}
	crossRegionLocationRegex, err := regexp.Compile("^[a-z]{2}-[a-z]{4,8}$")
	if err != nil {
		return err
	}

	if singleSiteLocationRegex.MatchString(bLocationConstraint) {
		d.Set("single_site_location", strings.Split(bLocationConstraint, "-")[0])
		d.Set("storage_class", strings.Split(bLocationConstraint, "-")[1])
	}
	if regionLocationRegex.MatchString(bLocationConstraint) {
		d.Set("region_location", fmt.Sprintf("%s-%s", strings.Split(bLocationConstraint, "-")[0], strings.Split(bLocationConstraint, "-")[1]))
		d.Set("storage_class", strings.Split(bLocationConstraint, "-")[2])
	}
	if crossRegionLocationRegex.MatchString(bLocationConstraint) {
		d.Set("cross_region_location", strings.Split(bLocationConstraint, "-")[0])
		d.Set("storage_class", strings.Split(bLocationConstraint, "-")[1])
	}

	head, err := s3Client.HeadBucket(headInput)
	if err != nil {
		return err
	}
	bucketID := fmt.Sprintf("%s:%s:%s:meta:%s:%s", strings.Replace(serviceID, "::", "", -1), "bucket", bucketName, bucketLocationConvert(bucketType), bucketRegion)
	d.SetId(bucketID)
	d.Set("key_protect", head.IBMSSEKPCrkId)
	bucketCRN := fmt.Sprintf("%s:%s:%s", strings.Replace(serviceID, "::", "", -1), "bucket", bucketName)
	d.Set("crn", bucketCRN)
	d.Set("resource_instance_id", serviceID)
	d.Set("s3_endpoint_public", apiEndpoint)
	d.Set("s3_endpoint_private", apiEndpointPrivate)
	return nil
}

func bucketLocationConvert(locationtype string) string {
	if locationtype == "cross_region_location" {
		return "crl"
	}
	if locationtype == "region_location" {
		return "rl"
	}
	if locationtype == "single_site_location" {
		return "crl"
	}
	return ""
}
