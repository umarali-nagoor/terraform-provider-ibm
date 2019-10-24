// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicGatewaysParams creates a new GetPublicGatewaysParams object
// with the default values initialized.
func NewGetPublicGatewaysParams() *GetPublicGatewaysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetPublicGatewaysParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGatewaysParamsWithTimeout creates a new GetPublicGatewaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGatewaysParamsWithTimeout(timeout time.Duration) *GetPublicGatewaysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetPublicGatewaysParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetPublicGatewaysParamsWithContext creates a new GetPublicGatewaysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGatewaysParamsWithContext(ctx context.Context) *GetPublicGatewaysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetPublicGatewaysParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetPublicGatewaysParamsWithHTTPClient creates a new GetPublicGatewaysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGatewaysParamsWithHTTPClient(client *http.Client) *GetPublicGatewaysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetPublicGatewaysParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetPublicGatewaysParams contains all the parameters to send to the API endpoint
for the get public gateways operation typically these are written to a http.Request
*/
type GetPublicGatewaysParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*ResourceGroupID
	  Filters the collection to resources within the resource group of the specified identifier

	*/
	ResourceGroupID *string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcCrn
	  Filters the collection to resources within the VPC of the specified CRN

	*/
	VpcCrn *string
	/*VpcID
	  Filters the collection to resources within the VPC of the specified identifier

	*/
	VpcID *string
	/*VpcName
	  Filters the collection to resources within the VPC of the specified name

	*/
	VpcName *string
	/*ZoneName
	  Filters the collection to resources within the specified zone

	*/
	ZoneName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public gateways params
func (o *GetPublicGatewaysParams) WithTimeout(timeout time.Duration) *GetPublicGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public gateways params
func (o *GetPublicGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public gateways params
func (o *GetPublicGatewaysParams) WithContext(ctx context.Context) *GetPublicGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public gateways params
func (o *GetPublicGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public gateways params
func (o *GetPublicGatewaysParams) WithHTTPClient(client *http.Client) *GetPublicGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public gateways params
func (o *GetPublicGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get public gateways params
func (o *GetPublicGatewaysParams) WithGeneration(generation int64) *GetPublicGatewaysParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get public gateways params
func (o *GetPublicGatewaysParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get public gateways params
func (o *GetPublicGatewaysParams) WithLimit(limit *int32) *GetPublicGatewaysParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get public gateways params
func (o *GetPublicGatewaysParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithResourceGroupID adds the resourceGroupID to the get public gateways params
func (o *GetPublicGatewaysParams) WithResourceGroupID(resourceGroupID *string) *GetPublicGatewaysParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get public gateways params
func (o *GetPublicGatewaysParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithStart adds the start to the get public gateways params
func (o *GetPublicGatewaysParams) WithStart(start *string) *GetPublicGatewaysParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get public gateways params
func (o *GetPublicGatewaysParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get public gateways params
func (o *GetPublicGatewaysParams) WithVersion(version string) *GetPublicGatewaysParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get public gateways params
func (o *GetPublicGatewaysParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcCrn adds the vpcCrn to the get public gateways params
func (o *GetPublicGatewaysParams) WithVpcCrn(vpcCrn *string) *GetPublicGatewaysParams {
	o.SetVpcCrn(vpcCrn)
	return o
}

// SetVpcCrn adds the vpcCrn to the get public gateways params
func (o *GetPublicGatewaysParams) SetVpcCrn(vpcCrn *string) {
	o.VpcCrn = vpcCrn
}

// WithVpcID adds the vpcID to the get public gateways params
func (o *GetPublicGatewaysParams) WithVpcID(vpcID *string) *GetPublicGatewaysParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get public gateways params
func (o *GetPublicGatewaysParams) SetVpcID(vpcID *string) {
	o.VpcID = vpcID
}

// WithVpcName adds the vpcName to the get public gateways params
func (o *GetPublicGatewaysParams) WithVpcName(vpcName *string) *GetPublicGatewaysParams {
	o.SetVpcName(vpcName)
	return o
}

// SetVpcName adds the vpcName to the get public gateways params
func (o *GetPublicGatewaysParams) SetVpcName(vpcName *string) {
	o.VpcName = vpcName
}

// WithZoneName adds the zoneName to the get public gateways params
func (o *GetPublicGatewaysParams) WithZoneName(zoneName *string) *GetPublicGatewaysParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the get public gateways params
func (o *GetPublicGatewaysParams) SetZoneName(zoneName *string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if o.VpcCrn != nil {

		// query param vpc.crn
		var qrVpcCrn string
		if o.VpcCrn != nil {
			qrVpcCrn = *o.VpcCrn
		}
		qVpcCrn := qrVpcCrn
		if qVpcCrn != "" {
			if err := r.SetQueryParam("vpc.crn", qVpcCrn); err != nil {
				return err
			}
		}

	}

	if o.VpcID != nil {

		// query param vpc.id
		var qrVpcID string
		if o.VpcID != nil {
			qrVpcID = *o.VpcID
		}
		qVpcID := qrVpcID
		if qVpcID != "" {
			if err := r.SetQueryParam("vpc.id", qVpcID); err != nil {
				return err
			}
		}

	}

	if o.VpcName != nil {

		// query param vpc.name
		var qrVpcName string
		if o.VpcName != nil {
			qrVpcName = *o.VpcName
		}
		qVpcName := qrVpcName
		if qVpcName != "" {
			if err := r.SetQueryParam("vpc.name", qVpcName); err != nil {
				return err
			}
		}

	}

	if o.ZoneName != nil {

		// query param zone.name
		var qrZoneName string
		if o.ZoneName != nil {
			qrZoneName = *o.ZoneName
		}
		qZoneName := qrZoneName
		if qZoneName != "" {
			if err := r.SetQueryParam("zone.name", qZoneName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
