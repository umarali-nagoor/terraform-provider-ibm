// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

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

// NewGetVpnGatewaysIDParams creates a new GetVpnGatewaysIDParams object
// with the default values initialized.
func NewGetVpnGatewaysIDParams() *GetVpnGatewaysIDParams {
	var ()
	return &GetVpnGatewaysIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpnGatewaysIDParamsWithTimeout creates a new GetVpnGatewaysIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpnGatewaysIDParamsWithTimeout(timeout time.Duration) *GetVpnGatewaysIDParams {
	var ()
	return &GetVpnGatewaysIDParams{

		timeout: timeout,
	}
}

// NewGetVpnGatewaysIDParamsWithContext creates a new GetVpnGatewaysIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpnGatewaysIDParamsWithContext(ctx context.Context) *GetVpnGatewaysIDParams {
	var ()
	return &GetVpnGatewaysIDParams{

		Context: ctx,
	}
}

// NewGetVpnGatewaysIDParamsWithHTTPClient creates a new GetVpnGatewaysIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpnGatewaysIDParamsWithHTTPClient(client *http.Client) *GetVpnGatewaysIDParams {
	var ()
	return &GetVpnGatewaysIDParams{
		HTTPClient: client,
	}
}

/*GetVpnGatewaysIDParams contains all the parameters to send to the API endpoint
for the get vpn gateways ID operation typically these are written to a http.Request
*/
type GetVpnGatewaysIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The VPN gateway idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) WithTimeout(timeout time.Duration) *GetVpnGatewaysIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) WithContext(ctx context.Context) *GetVpnGatewaysIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) WithHTTPClient(client *http.Client) *GetVpnGatewaysIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) WithGeneration(generation int64) *GetVpnGatewaysIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) WithID(id string) *GetVpnGatewaysIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) WithVersion(version string) *GetVpnGatewaysIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpn gateways ID params
func (o *GetVpnGatewaysIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpnGatewaysIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
