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

// NewPutSubnetsSubnetIDPublicGatewayParams creates a new PutSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized.
func NewPutSubnetsSubnetIDPublicGatewayParams() *PutSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &PutSubnetsSubnetIDPublicGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSubnetsSubnetIDPublicGatewayParamsWithTimeout creates a new PutSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSubnetsSubnetIDPublicGatewayParamsWithTimeout(timeout time.Duration) *PutSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &PutSubnetsSubnetIDPublicGatewayParams{

		timeout: timeout,
	}
}

// NewPutSubnetsSubnetIDPublicGatewayParamsWithContext creates a new PutSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSubnetsSubnetIDPublicGatewayParamsWithContext(ctx context.Context) *PutSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &PutSubnetsSubnetIDPublicGatewayParams{

		Context: ctx,
	}
}

// NewPutSubnetsSubnetIDPublicGatewayParamsWithHTTPClient creates a new PutSubnetsSubnetIDPublicGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSubnetsSubnetIDPublicGatewayParamsWithHTTPClient(client *http.Client) *PutSubnetsSubnetIDPublicGatewayParams {
	var ()
	return &PutSubnetsSubnetIDPublicGatewayParams{
		HTTPClient: client,
	}
}

/*PutSubnetsSubnetIDPublicGatewayParams contains all the parameters to send to the API endpoint
for the put subnets subnet ID public gateway operation typically these are written to a http.Request
*/
type PutSubnetsSubnetIDPublicGatewayParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*RequestBody*/
	RequestBody PutSubnetsSubnetIDPublicGatewayBody
	/*SubnetID
	  The subnet identifier

	*/
	SubnetID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithTimeout(timeout time.Duration) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithContext(ctx context.Context) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithHTTPClient(client *http.Client) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithGeneration(generation int64) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithRequestBody adds the requestBody to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithRequestBody(requestBody PutSubnetsSubnetIDPublicGatewayBody) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetRequestBody(requestBody PutSubnetsSubnetIDPublicGatewayBody) {
	o.RequestBody = requestBody
}

// WithSubnetID adds the subnetID to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithSubnetID(subnetID string) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetSubnetID(subnetID)
	return o
}

// SetSubnetID adds the subnetId to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetSubnetID(subnetID string) {
	o.SubnetID = subnetID
}

// WithVersion adds the version to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) WithVersion(version string) *PutSubnetsSubnetIDPublicGatewayParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the put subnets subnet ID public gateway params
func (o *PutSubnetsSubnetIDPublicGatewayParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PutSubnetsSubnetIDPublicGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.RequestBody); err != nil {
		return err
	}

	// path param subnet_id
	if err := r.SetPathParam("subnet_id", o.SubnetID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
