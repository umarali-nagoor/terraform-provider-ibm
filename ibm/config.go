package ibm

import (
	"errors"
	"fmt"
	"log"
	gohttp "net/http"
	"os"
	"strings"
	"time"
	// Added code for the Power Colo Offering

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	jwt "github.com/dgrijalva/jwt-go"
	slsession "github.com/softlayer/softlayer-go/session"
	issession "github.ibm.com/Bluemix/riaas-go-client/session"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/certificatemanager"
	"github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/api/globalsearch/globalsearchv2"
	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/api/schematics"
	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	ibmpisession "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
)

//RetryDelay
const RetryAPIDelay = 5 * time.Second

//BluemixRegion ...
var BluemixRegion string

var (
	errEmptySoftLayerCredentials = errors.New("iaas_classic_username and iaas_classic_api_key must be provided. Please see the documentation on how to configure them")
	errEmptyBluemixCredentials   = errors.New("ibmcloud_api_key or bluemix_api_key or iam_token and iam_refresh_token must be provided. Please see the documentation on how to configure it")
)

//UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	userAccount string
	cloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
}

//Config stores user provider input
type Config struct {
	//BluemixAPIKey is the Bluemix api key
	BluemixAPIKey string
	//Bluemix region
	Region string
	//Resource group id
	ResourceGroup string
	//Bluemix API timeout
	BluemixTimeout time.Duration

	//Softlayer end point url
	SoftLayerEndpointURL string

	//Softlayer API timeout
	SoftLayerTimeout time.Duration

	// Softlayer User Name
	SoftLayerUserName string

	// Softlayer API Key
	SoftLayerAPIKey string

	//Retry Count for API calls
	//Unexposed in the schema at this point as they are used only during session creation for a few calls
	//When sdk implements it we an expose them for expected behaviour
	//https://github.com/softlayer/softlayer-go/issues/41
	RetryCount int
	//Constant Retry Delay for API calls
	RetryDelay time.Duration

	// FunctionNameSpace ...
	FunctionNameSpace string

	//Riaas End point
	RiaasEndPoint string

	//Generation
	Generation int

	//IAM Token
	IAMToken string

	//IAM Refresh Token
	IAMRefreshToken string

	// PowerService Instance
	PowerServiceInstance string
}

//Session stores the information required for communication with the SoftLayer and Bluemix API
type Session struct {
	// SoftLayerSesssion is the the SoftLayer session used to connect to the SoftLayer API
	SoftLayerSession *slsession.Session

	// BluemixSession is the the Bluemix session used to connect to the Bluemix API
	BluemixSession *bxsession.Session
}

// ClientSession ...
type ClientSession interface {
	BluemixSession() (*bxsession.Session, error)
	BluemixAcccountAPI() (accountv2.AccountServiceAPI, error)
	BluemixAcccountv1API() (accountv1.AccountServiceAPI, error)
	BluemixUserDetails() (*UserConfig, error)
	ContainerAPI() (containerv1.ContainerServiceAPI, error)
	VpcContainerAPI() (containerv2.ContainerServiceAPI, error)
	CisAPI() (cisv1.CisServiceAPI, error)
	FunctionClient() (*whisk.Client, error)
	GlobalSearchAPI() (globalsearchv2.GlobalSearchServiceAPI, error)
	GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error)
	ICDAPI() (icdv4.ICDServiceAPI, error)
	IAMAPI() (iamv1.IAMServiceAPI, error)
	IAMPAPAPI() (iampapv1.IAMPAPAPI, error)
	IAMUUMAPI() (iamuumv1.IAMUUMServiceAPI, error)
	ISSession() (*issession.Session, error)
	MccpAPI() (mccpv2.MccpServiceAPI, error)
	ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error)
	ResourceManagementAPI() (management.ResourceManagementAPI, error)
	ResourceControllerAPI() (controller.ResourceControllerAPI, error)
	SoftLayerSession() *slsession.Session
	IBMPISession() (*ibmpisession.IBMPISession, error)
	SchematicsAPI() (schematics.SchematicsServiceAPI, error)
	UserManagementAPI() (usermanagementv2.UserManagementAPI, error)
	CertificateManagerAPI() (certificatemanager.CertificateManagerServiceAPI, error)
}

type clientSession struct {
	session *Session

	accountConfigErr     error
	bmxAccountServiceAPI accountv2.AccountServiceAPI

	accountV1ConfigErr     error
	bmxAccountv1ServiceAPI accountv1.AccountServiceAPI

	bmxUserDetails  *UserConfig
	bmxUserFetchErr error

	csConfigErr  error
	csServiceAPI containerv1.ContainerServiceAPI

	csv2ConfigErr  error
	csv2ServiceAPI containerv2.ContainerServiceAPI

	stxConfigErr  error
	stxServiceAPI schematics.SchematicsServiceAPI

	certManagementErr error
	certManagementAPI certificatemanager.CertificateManagerServiceAPI

	cfConfigErr  error
	cfServiceAPI mccpv2.MccpServiceAPI

	cisConfigErr  error
	cisServiceAPI cisv1.CisServiceAPI

	functionConfigErr error
	functionClient    *whisk.Client

	globalSearchConfigErr  error
	globalSearchServiceAPI globalsearchv2.GlobalSearchServiceAPI

	globalTaggingConfigErr  error
	globalTaggingServiceAPI globaltaggingv3.GlobalTaggingServiceAPI

	iamPAPConfigErr  error
	iamPAPServiceAPI iampapv1.IAMPAPAPI

	iamUUMConfigErr  error
	iamUUMServiceAPI iamuumv1.IAMUUMServiceAPI

	iamConfigErr  error
	iamServiceAPI iamv1.IAMServiceAPI

	userManagementErr error
	userManagementAPI usermanagementv2.UserManagementAPI

	icdConfigErr  error
	icdServiceAPI icdv4.ICDServiceAPI

	isConfigErr error
	isSession   *issession.Session

	resourceControllerConfigErr  error
	resourceControllerServiceAPI controller.ResourceControllerAPI

	resourceManagementConfigErr  error
	resourceManagementServiceAPI management.ResourceManagementAPI

	resourceCatalogConfigErr  error
	resourceCatalogServiceAPI catalog.ResourceCatalogAPI

	powerConfigErr error
	ibmpiConfigErr error
	ibmpiSession   *ibmpisession.IBMPISession

	bluemixSessionErr error
}

// BluemixAcccountAPI ...
func (sess clientSession) BluemixAcccountAPI() (accountv2.AccountServiceAPI, error) {
	return sess.bmxAccountServiceAPI, sess.accountConfigErr
}

// BluemixAcccountAPI ...
func (sess clientSession) BluemixAcccountv1API() (accountv1.AccountServiceAPI, error) {
	return sess.bmxAccountv1ServiceAPI, sess.accountV1ConfigErr
}

// BluemixSession to provide the Bluemix Session
func (sess clientSession) BluemixSession() (*bxsession.Session, error) {
	return sess.session.BluemixSession, sess.bluemixSessionErr
}

// BluemixUserDetails ...
func (sess clientSession) BluemixUserDetails() (*UserConfig, error) {
	return sess.bmxUserDetails, sess.bmxUserFetchErr
}

// ContainerAPI provides Container Service APIs ...
func (sess clientSession) ContainerAPI() (containerv1.ContainerServiceAPI, error) {
	return sess.csServiceAPI, sess.csConfigErr
}

// VpcContainerAPI provides v2Container Service APIs ...
func (sess clientSession) VpcContainerAPI() (containerv2.ContainerServiceAPI, error) {
	return sess.csv2ServiceAPI, sess.csv2ConfigErr
}

// SchematicsAPI provides schematics Service APIs ...
func (sess clientSession) SchematicsAPI() (schematics.SchematicsServiceAPI, error) {
	return sess.stxServiceAPI, sess.stxConfigErr
}

// CisAPI provides Cloud Internet Services APIs ...
func (sess clientSession) CisAPI() (cisv1.CisServiceAPI, error) {
	return sess.cisServiceAPI, sess.cisConfigErr
}

// FunctionClient ...
func (sess clientSession) FunctionClient() (*whisk.Client, error) {
	return sess.functionClient, sess.functionConfigErr
}

// GlobalSearchAPI provides Global Search  APIs ...
func (sess clientSession) GlobalSearchAPI() (globalsearchv2.GlobalSearchServiceAPI, error) {
	return sess.globalSearchServiceAPI, sess.globalSearchConfigErr
}

// GlobalTaggingAPI provides Global Search  APIs ...
func (sess clientSession) GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error) {
	return sess.globalTaggingServiceAPI, sess.globalTaggingConfigErr
}

// IAMAPI provides IAM PAP APIs ...
func (sess clientSession) IAMAPI() (iamv1.IAMServiceAPI, error) {
	return sess.iamServiceAPI, sess.iamConfigErr
}

// UserManagementAPI provides User management APIs ...
func (sess clientSession) UserManagementAPI() (usermanagementv2.UserManagementAPI, error) {
	return sess.userManagementAPI, sess.userManagementErr
}

// IAMPAPAPI provides IAM PAP APIs ...
func (sess clientSession) IAMPAPAPI() (iampapv1.IAMPAPAPI, error) {
	return sess.iamPAPServiceAPI, sess.iamPAPConfigErr
}

// IAMUUMAPI provides IAM UUM APIs ...
func (sess clientSession) IAMUUMAPI() (iamuumv1.IAMUUMServiceAPI, error) {
	return sess.iamUUMServiceAPI, sess.iamUUMConfigErr
}

// IcdAPI provides IBM Cloud Databases APIs ...
func (sess clientSession) ICDAPI() (icdv4.ICDServiceAPI, error) {
	return sess.icdServiceAPI, sess.icdConfigErr
}

// ISSession to provide the IS RIASS Session
func (sess clientSession) ISSession() (*issession.Session, error) {
	return sess.isSession, sess.isConfigErr
}

// MccpAPI provides Multi Cloud Controller Proxy APIs ...
func (sess clientSession) MccpAPI() (mccpv2.MccpServiceAPI, error) {
	return sess.cfServiceAPI, sess.cfConfigErr
}

// ResourceCatalogAPI ...
func (sess clientSession) ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error) {
	return sess.resourceCatalogServiceAPI, sess.resourceCatalogConfigErr
}

// ResourceManagementAPI ...
func (sess clientSession) ResourceManagementAPI() (management.ResourceManagementAPI, error) {
	return sess.resourceManagementServiceAPI, sess.resourceManagementConfigErr
}

// ResourceControllerAPI ...
func (sess clientSession) ResourceControllerAPI() (controller.ResourceControllerAPI, error) {
	return sess.resourceControllerServiceAPI, sess.resourceControllerConfigErr
}

// SoftLayerSession providers SoftLayer Session
func (sess clientSession) SoftLayerSession() *slsession.Session {
	return sess.session.SoftLayerSession
}

// CertManagementAPI provides Certificate  management APIs ...
func (sess clientSession) CertificateManagerAPI() (certificatemanager.CertificateManagerServiceAPI, error) {
	return sess.certManagementAPI, sess.certManagementErr
}

// Session to the Power Colo Service

func (sess clientSession) IBMPISession() (*ibmpisession.IBMPISession, error) {
	return sess.ibmpiSession, sess.powerConfigErr
}

// ClientSession configures and returns a fully initialized ClientSession
func (c *Config) ClientSession() (interface{}, error) {
	sess, err := newSession(c)
	if err != nil {
		return nil, err
	}
	session := clientSession{
		session: sess,
	}

	if sess.BluemixSession == nil {
		//Can be nil only  if bluemix_api_key is not provided
		log.Println("Skipping Bluemix Clients configuration")
		session.bluemixSessionErr = errEmptyBluemixCredentials
		session.accountConfigErr = errEmptyBluemixCredentials
		session.accountV1ConfigErr = errEmptyBluemixCredentials
		session.csConfigErr = errEmptyBluemixCredentials
		session.csv2ConfigErr = errEmptyBluemixCredentials
		session.stxConfigErr = errEmptyBluemixCredentials
		session.cfConfigErr = errEmptyBluemixCredentials
		session.cisConfigErr = errEmptyBluemixCredentials
		session.functionConfigErr = errEmptyBluemixCredentials
		session.globalSearchConfigErr = errEmptyBluemixCredentials
		session.globalTaggingConfigErr = errEmptyBluemixCredentials
		session.iamConfigErr = errEmptyBluemixCredentials
		session.iamPAPConfigErr = errEmptyBluemixCredentials
		session.iamUUMConfigErr = errEmptyBluemixCredentials
		session.icdConfigErr = errEmptyBluemixCredentials
		session.resourceCatalogConfigErr = errEmptyBluemixCredentials
		session.resourceManagementConfigErr = errEmptyBluemixCredentials
		session.resourceControllerConfigErr = errEmptyBluemixCredentials
		session.isConfigErr = errEmptyBluemixCredentials
		session.powerConfigErr = errEmptyBluemixCredentials
		session.ibmpiConfigErr = errEmptyBluemixCredentials
		session.userManagementErr = errEmptyBluemixCredentials
		session.certManagementErr = errEmptyBluemixCredentials

		return session, nil
	}

	if sess.BluemixSession.Config.BluemixAPIKey != "" {
		err = authenticateAPIKey(sess.BluemixSession)
		if err != nil {
			session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
			session.functionConfigErr = fmt.Errorf("Error occured while fetching auth key for function: %q", err)
			session.isConfigErr = fmt.Errorf("Error occured while fetching auth key for vpc: %q", err)
			session.powerConfigErr = fmt.Errorf("Error occured while fetching the auth key for power iaas: %q", err)
			session.ibmpiConfigErr = fmt.Errorf("Error occured while fetching the auth key for power iaas: %q", err)
		}
		err = authenticateCF(sess.BluemixSession)
		if err != nil {
			session.functionConfigErr = fmt.Errorf("Error occured while fetching auth key for function: %q", err)
		}
	}

	if sess.BluemixSession.Config.IAMAccessToken != "" && sess.BluemixSession.Config.BluemixAPIKey == "" {
		err := refreshToken(sess.BluemixSession)
		if err != nil {
			return nil, err
		}

	}
	userConfig, err := fetchUserDetails(sess.BluemixSession)
	if err != nil {
		session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
	}
	session.bmxUserDetails = userConfig

	if sess.SoftLayerSession != nil && sess.SoftLayerSession.IAMToken != "" {
		sess.SoftLayerSession.IAMToken = sess.BluemixSession.Config.IAMAccessToken
		sess.SoftLayerSession.IAMRefreshToken = sess.BluemixSession.Config.IAMRefreshToken
	}

	session.functionClient, session.functionConfigErr = FunctionClient(sess.BluemixSession.Config, c.FunctionNameSpace)

	BluemixRegion = sess.BluemixSession.Config.Region

	accv1API, err := accountv1.New(sess.BluemixSession)
	if err != nil {
		session.accountV1ConfigErr = fmt.Errorf("Error occured while configuring Bluemix Accountv1 Service: %q", err)
	}
	session.bmxAccountv1ServiceAPI = accv1API

	accAPI, err := accountv2.New(sess.BluemixSession)
	if err != nil {
		session.accountConfigErr = fmt.Errorf("Error occured while configuring  Account Service: %q", err)
	}
	session.bmxAccountServiceAPI = accAPI

	cfAPI, err := mccpv2.New(sess.BluemixSession)
	if err != nil {
		session.cfConfigErr = fmt.Errorf("Error occured while configuring MCCP service: %q", err)
	}
	session.cfServiceAPI = cfAPI

	clusterAPI, err := containerv1.New(sess.BluemixSession)
	if err != nil {
		session.csConfigErr = fmt.Errorf("Error occured while configuring Container Service for K8s cluster: %q", err)
	}
	session.csServiceAPI = clusterAPI

	v2clusterAPI, err := containerv2.New(sess.BluemixSession)
	if err != nil {
		session.csv2ConfigErr = fmt.Errorf("Error occured while configuring vpc Container Service for K8s cluster: %q", err)
	}
	session.csv2ServiceAPI = v2clusterAPI

	schematicService, err := schematics.New(sess.BluemixSession)
	if err != nil {
		session.stxConfigErr = fmt.Errorf("Error occured while fetching schematics Configuration: %q", err)
	}
	session.stxServiceAPI = schematicService

	cisAPI, err := cisv1.New(sess.BluemixSession)
	if err != nil {
		session.cisConfigErr = fmt.Errorf("Error occured while configuring Cloud Internet Services: %q", err)
	}
	session.cisServiceAPI = cisAPI

	globalSearchAPI, err := globalsearchv2.New(sess.BluemixSession)
	if err != nil {
		session.globalSearchConfigErr = fmt.Errorf("Error occured while configuring Global Search: %q", err)
	}
	session.globalSearchServiceAPI = globalSearchAPI

	globalTaggingAPI, err := globaltaggingv3.New(sess.BluemixSession)
	if err != nil {
		session.globalTaggingConfigErr = fmt.Errorf("Error occured while configuring Global Tagging: %q", err)
	}
	session.globalTaggingServiceAPI = globalTaggingAPI

	iampap, err := iampapv1.New(sess.BluemixSession)
	if err != nil {
		session.iamPAPConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAMPAP Service: %q", err)
	}
	session.iamPAPServiceAPI = iampap

	iam, err := iamv1.New(sess.BluemixSession)
	if err != nil {
		session.iamConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAM Service: %q", err)
	}
	session.iamServiceAPI = iam

	iamuum, err := iamuumv1.New(sess.BluemixSession)
	if err != nil {
		session.iamUUMConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAMUUM Service: %q", err)
	}
	session.iamUUMServiceAPI = iamuum

	issession, err := issession.New(sess.BluemixSession.Config.IAMAccessToken, c.Region, c.Generation, true, c.BluemixTimeout)
	if err != nil {
		session.isConfigErr = err
		return nil, err
	}
	session.isSession = issession

	icdAPI, err := icdv4.New(sess.BluemixSession)
	if err != nil {
		session.icdConfigErr = fmt.Errorf("Error occured while configuring IBM Cloud Database Services: %q", err)
	}
	session.icdServiceAPI = icdAPI

	resourceCatalogAPI, err := catalog.New(sess.BluemixSession)
	if err != nil {
		session.resourceCatalogConfigErr = fmt.Errorf("Error occured while configuring Resource Catalog service: %q", err)
	}
	session.resourceCatalogServiceAPI = resourceCatalogAPI

	resourceManagementAPI, err := management.New(sess.BluemixSession)
	if err != nil {
		session.resourceManagementConfigErr = fmt.Errorf("Error occured while configuring Resource Management service: %q", err)
	}
	session.resourceManagementServiceAPI = resourceManagementAPI

	resourceControllerAPI, err := controller.New(sess.BluemixSession)
	if err != nil {
		session.resourceControllerConfigErr = fmt.Errorf("Error occured while configuring Resource Controller service: %q", err)
	}
	session.resourceControllerServiceAPI = resourceControllerAPI

	userManagementAPI, err := usermanagementv2.New(sess.BluemixSession)
	if err != nil {
		session.userManagementErr = fmt.Errorf("Error occured while configuring user management service: %q", err)
	}
	session.userManagementAPI = userManagementAPI
	certManagementAPI, err := certificatemanager.New(sess.BluemixSession)
	if err != nil {
		session.certManagementErr = fmt.Errorf("Error occured while configuring Certificate manager service: %q", err)
	}
	session.certManagementAPI = certManagementAPI

	ibmpisession, err := ibmpisession.New(sess.BluemixSession.Config.IAMAccessToken, c.Region, true, c.BluemixTimeout, session.bmxUserDetails.userAccount)
	if err != nil {
		session.ibmpiConfigErr = err
		return nil, err
	}

	session.ibmpiSession = ibmpisession

	return session, nil
}

func newSession(c *Config) (*Session, error) {
	ibmSession := &Session{}

	softlayerSession := &slsession.Session{
		Endpoint:  c.SoftLayerEndpointURL,
		Timeout:   c.SoftLayerTimeout,
		UserName:  c.SoftLayerUserName,
		APIKey:    c.SoftLayerAPIKey,
		Debug:     os.Getenv("TF_LOG") != "",
		Retries:   c.RetryCount,
		RetryWait: c.RetryDelay,
	}

	if c.IAMToken != "" {
		log.Println("Configuring SoftLayer Session with token")
		softlayerSession.IAMToken = c.IAMToken
		softlayerSession.IAMRefreshToken = c.IAMRefreshToken
	}
	if c.SoftLayerAPIKey != "" && c.SoftLayerUserName != "" {
		log.Println("Configuring SoftLayer Session with API key")
		softlayerSession.APIKey = c.SoftLayerAPIKey
		softlayerSession.UserName = c.SoftLayerUserName
	}
	softlayerSession.AppendUserAgent(fmt.Sprintf("terraform-provider-ibm/%s", version.Version))
	ibmSession.SoftLayerSession = softlayerSession

	if (c.IAMToken != "" && c.IAMRefreshToken == "") || (c.IAMToken == "" && c.IAMRefreshToken != "") {
		return nil, fmt.Errorf("iam_token and iam_refresh_token must be provided")
	}

	if c.IAMToken != "" && c.IAMRefreshToken != "" {
		log.Println("Configuring IBM Cloud Session with token")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			IAMAccessToken:  c.IAMToken,
			IAMRefreshToken: c.IAMRefreshToken,
			//Comment out debug mode for v0.12
			//Debug:           os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
		}
		sess, err := bxsession.New(bmxConfig)
		if err != nil {
			return nil, err
		}
		ibmSession.BluemixSession = sess
	}

	if c.BluemixAPIKey != "" {
		log.Println("Configuring IBM Cloud Session with API key")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			BluemixAPIKey: c.BluemixAPIKey,
			//Comment out debug mode for v0.12
			//Debug:         os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
			//PowerServiceInstance: c.PowerServiceInstance,
		}
		sess, err := bxsession.New(bmxConfig)
		if err != nil {
			return nil, err
		}
		ibmSession.BluemixSession = sess
	}

	return ibmSession, nil
}

func authenticateAPIKey(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func authenticateCF(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewUAARepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func fetchUserDetails(sess *bxsession.Session) (*UserConfig, error) {
	config := sess.Config
	user := UserConfig{}

	var bluemixToken string

	if strings.HasPrefix(config.IAMAccessToken, "Bearer") {
		bluemixToken = config.IAMAccessToken[7:len(config.IAMAccessToken)]
	} else {
		bluemixToken = config.IAMAccessToken
	}

	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return &user, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.userEmail = email.(string)
	}
	user.userID = claims["id"].(string)
	user.userAccount = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.cloudName = "bluemix"
	} else {
		user.cloudName = "staging"
	}
	user.cloudType = "public"
	return &user, nil
}

func refreshToken(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	_, err = tokenRefresher.RefreshToken()
	return err
}
