package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPINetworkClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerImageClient ...
func NewIBMPINetworkClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPINetworkClient {
	return &IBMPINetworkClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

func (f *IBMPINetworkClient) Get(id, powerinstanceid string) (*models.Network, error) {

	params := p_cloud_networks.NewPcloudNetworksGetParams().WithCloudInstanceID(powerinstanceid).WithNetworkID(id)
	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *IBMPINetworkClient) Create(name string, networktype string, cidr string, dnsservers []string, gateway string, startip string, endip string, powerinstanceid string) (*models.Network, *models.Network, error) {

	var ipbody = []*models.IPAddressRange{
		{&endip, &startip},
	}

	var body = models.NetworkCreate{}
	body.Cidr = cidr
	body.Gateway = gateway
	body.IPAddressRanges = ipbody
	body.Name = name
	body.DNSServers = dnsservers
	body.Type = &networktype

	log.Printf("Printing the body %+v", body)
	params := p_cloud_networks.NewPcloudNetworksPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	_, resp, err := f.session.Power.PCloudNetworks.PcloudNetworksPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	//log.Printf("The error is %d ",resp.Payload.VlanID)
	if err != nil {
		return nil, nil, errors.ToError(err)
	}

	if resp != nil {
		log.Printf("Failed to create the network ")
	}
	//if postok != nil {
	//	log.Print("Request failed ")
	//}

	return resp.Payload, nil, nil
}

func (f *IBMPINetworkClient) GetPublic(cloud_instance_id string) (*models.Networks, error) {

	filterQuery := "type=\"pub-vlan\""
	params := p_cloud_networks.NewPcloudNetworksGetallParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(cloud_instance_id).WithFilter(&filterQuery)

	resp, err := f.session.Power.PCloudNetworks.PcloudNetworksGetall(params, ibmpisession.NewAuth(f.session, cloud_instance_id))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
