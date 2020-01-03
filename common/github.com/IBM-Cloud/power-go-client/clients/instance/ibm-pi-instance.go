package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"

	"log"
)

type IBMPIInstanceClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewIBMPIInstanceClient ...
func NewIBMPIInstanceClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIInstanceClient {
	return &IBMPIInstanceClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

//Get information about a single pvm only
func (f *IBMPIInstanceClient) Get(id, powerinstanceid string) (*models.PVMInstance, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesGetParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Create

func (f *IBMPIInstanceClient) Create(powerdef *p_cloud_p_vm_instances.PcloudPvminstancesPostParams, powerinstanceid string) (*models.PVMInstanceList, *models.PVMInstanceList, *models.PVMInstanceList, error) {

	log.Printf("Calling the Power PVM Create Method %s", powerdef.Body)
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithBody(powerdef.Body)

	log.Printf("Printing the params to be passed %+v", params)

	_, _, resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		log.Printf("failed to process the request..")
		return nil, nil, nil, errors.ToError(err)
	}

	return &resp.Payload, nil, nil, nil
}

// PVM Instances Delete
func (f *IBMPIInstanceClient) Delete(id, powerinstanceid string) error {

	log.Printf("Calling the Power PVM Delete Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesDeleteParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)
	_, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return errors.ToError(err)
	}

	return nil
}

// PVM Instances Update
func (f *IBMPIInstanceClient) Update(id, powerinstanceid string, powerupdateparams *p_cloud_p_vm_instances.PcloudPvminstancesPutParams) (*models.PVMInstanceUpdateResponse, error) {

	log.Printf("Calling the Power PVM Update Instance Method")
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesPutParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(powerupdateparams.Body)
	resp, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesPut(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)

	}
	return resp.Payload, nil
}

// PVM Instances Operations
func (f *IBMPIInstanceClient) Action(poweractionparams *p_cloud_p_vm_instances.PcloudPvminstancesActionPostParams, id, powerinstanceid string) (models.Object, error) {

	log.Printf("Calling the Power PVM Action Method")
	log.Printf("the params are %s - powerinstance id is %s", id, powerinstanceid)
	log.Printf("printing the poweraction params %s", *poweractionparams.Body)
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesActionPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(poweractionparams.Body)

	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesActionPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return postok.Payload, nil

}

// Generate the Console URL

func (f *IBMPIInstanceClient) PostConsoleURL(id, powerinstanceid string) (models.Object, error) {
	params := p_cloud_p_vm_instances.NewPcloudPvminstancesConsolePostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id)

	postok, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesConsolePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		return nil, errors.ToError(err)
	}
	return postok.Payload, nil
}

// Capture an instance

func (f *IBMPIInstanceClient) CaptureInstanceToImageCatalog(id, powerinstanceid string, picaptureparams *p_cloud_p_vm_instances.PcloudPvminstancesCapturePostParams) (models.Object, error) {

	params := p_cloud_p_vm_instances.NewPcloudPvminstancesCapturePostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(powerinstanceid).WithPvmInstanceID(id).WithBody(picaptureparams.Body)
	postok, _, err := f.session.Power.PCloudPVMInstances.PcloudPvminstancesCapturePost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return postok.Payload, nil

}
