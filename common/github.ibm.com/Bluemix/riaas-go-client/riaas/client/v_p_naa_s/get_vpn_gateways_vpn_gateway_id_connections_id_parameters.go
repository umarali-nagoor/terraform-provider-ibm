// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDParams creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDParams object
// with the default values initialized.
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDParams() *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDParamsWithTimeout creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDParamsWithTimeout(timeout time.Duration) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDParams{

		timeout: timeout,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDParamsWithContext creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDParamsWithContext(ctx context.Context) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDParams{

		Context: ctx,
	}
}

// NewGetVpnGatewaysVpnGatewayIDConnectionsIDParamsWithHTTPClient creates a new GetVpnGatewaysVpnGatewayIDConnectionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpnGatewaysVpnGatewayIDConnectionsIDParamsWithHTTPClient(client *http.Client) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	var ()
	return &GetVpnGatewaysVpnGatewayIDConnectionsIDParams{
		HTTPClient: client,
	}
}

/*GetVpnGatewaysVpnGatewayIDConnectionsIDParams contains all the parameters to send to the API endpoint
for the get vpn gateways vpn gateway ID connections ID operation typically these are written to a http.Request
*/
type GetVpnGatewaysVpnGatewayIDConnectionsIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The VPN connection idenitifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpnGatewayID
	  The VPN gateway identifier

	*/
	VpnGatewayID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithTimeout(timeout time.Duration) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithContext(ctx context.Context) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithHTTPClient(client *http.Client) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithGeneration(generation int64) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithID(id string) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithVersion(version string) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetVersion(version string) {
	o.Version = version
}

// WithVpnGatewayID adds the vpnGatewayID to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WithVpnGatewayID(vpnGatewayID string) *GetVpnGatewaysVpnGatewayIDConnectionsIDParams {
	o.SetVpnGatewayID(vpnGatewayID)
	return o
}

// SetVpnGatewayID adds the vpnGatewayId to the get vpn gateways vpn gateway ID connections ID params
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) SetVpnGatewayID(vpnGatewayID string) {
	o.VpnGatewayID = vpnGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpnGatewaysVpnGatewayIDConnectionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vpn_gateway_id
	if err := r.SetPathParam("vpn_gateway_id", o.VpnGatewayID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
