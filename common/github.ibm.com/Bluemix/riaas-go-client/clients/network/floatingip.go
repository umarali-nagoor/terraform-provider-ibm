package network

import (
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/network"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.com/go-openapi/strfmt"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// FloatingIPClient ...
type FloatingIPClient struct {
	session *session.Session
}

// NewFloatingIPClient ...
func NewFloatingIPClient(sess *session.Session) *FloatingIPClient {
	return &FloatingIPClient{
		sess,
	}
}

// List ...
func (f *FloatingIPClient) List(start string) ([]*models.FloatingIP, string, error) {
	return f.ListWithFilter("", "", start)
}

// ListWithFilter ...
func (f *FloatingIPClient) ListWithFilter(zoneName, resourcegroupID, start string) ([]*models.FloatingIP, string, error) {
	params := network.NewGetFloatingIpsParamsWithTimeout(f.session.Timeout)
	if zoneName != "" {
		params = params.WithZoneName(&zoneName)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetFloatingIps(params, session.Auth(f.session))

	if err != nil {
		return nil, "", riaaserrors.ToError(err)
	}

	return resp.Payload.FloatingIps, utils.GetNext(resp.Payload.Next), nil
}

// Get ...
func (f *FloatingIPClient) Get(id string) (*models.FloatingIP, error) {
	params := network.NewGetFloatingIpsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.GetFloatingIpsID(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *FloatingIPClient) Create(name, zoneName, resourcegroupID, targetID string) (*models.FloatingIP, error) {

	var body = network.PostFloatingIpsBody{
		Name: name,
	}

	if zoneName != "" {
		var zone = network.PostFloatingIpsParamsBodyZone{
			Name: zoneName,
		}
		body.Zone = &zone
	}

	if targetID != "" {
		targetUUID := strfmt.UUID(targetID)
		var target = network.PostFloatingIpsParamsBodyTarget{
			ID: targetUUID,
		}
		body.Target = &target
	}

	if resourcegroupID != "" {
		resourcegroupuuid := strfmt.UUID(resourcegroupID)
		var resourcegroup = network.PostFloatingIpsParamsBodyResourceGroup{
			ID: resourcegroupuuid,
		}
		body.ResourceGroup = &resourcegroup
	}

	params := network.NewPostFloatingIpsParamsWithTimeout(f.session.Timeout).WithBody(body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PostFloatingIps(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *FloatingIPClient) Delete(id string) error {
	params := network.NewDeleteFloatingIpsIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Network.DeleteFloatingIpsID(params, session.Auth(f.session))
	return riaaserrors.ToError(err)
}

// Update ...
func (f *FloatingIPClient) Update(id, name, targetID string) (*models.FloatingIP, error) {
	var body = network.PatchFloatingIpsIDBody{}

	if name != "" {
		body.Name = name
	}

	if targetID != "" {
		targetUUID := strfmt.UUID(targetID)
		var target = network.PatchFloatingIpsIDParamsBodyTarget{
			ID: targetUUID,
		}
		body.Target = &target
	}

	params := network.NewPatchFloatingIpsIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Network.PatchFloatingIpsID(params, session.Auth(f.session))
	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload, nil
}
