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

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized.
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams() *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithTimeout creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithTimeout(timeout time.Duration) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{

		timeout: timeout,
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithContext creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithContext(ctx context.Context) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{

		Context: ctx,
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithHTTPClient creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParamsWithHTTPClient(client *http.Client) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams{
		HTTPClient: client,
	}
}

/*DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams contains all the parameters to send to the API endpoint
for the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length operation typically these are written to a http.Request
*/
type DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The VPN connection identifier

	*/
	ID string
	/*PrefixAddress
	  The prefix address part of the CIDR

	*/
	PrefixAddress string
	/*PrefixLength
	  The prefix length part of the CIDR

	*/
	PrefixLength string
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

// WithTimeout adds the timeout to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithTimeout(timeout time.Duration) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithContext(ctx context.Context) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithHTTPClient(client *http.Client) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithGeneration(generation int64) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithID(id string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetID(id string) {
	o.ID = id
}

// WithPrefixAddress adds the prefixAddress to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithPrefixAddress(prefixAddress string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixAddress(prefixAddress)
	return o
}

// SetPrefixAddress adds the prefixAddress to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetPrefixAddress(prefixAddress string) {
	o.PrefixAddress = prefixAddress
}

// WithPrefixLength adds the prefixLength to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithPrefixLength(prefixLength string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixLength(prefixLength)
	return o
}

// SetPrefixLength adds the prefixLength to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetPrefixLength(prefixLength string) {
	o.PrefixLength = prefixLength
}

// WithVersion adds the version to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithVersion(version string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetVersion(version string) {
	o.Version = version
}

// WithVpnGatewayID adds the vpnGatewayID to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WithVpnGatewayID(vpnGatewayID string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams {
	o.SetVpnGatewayID(vpnGatewayID)
	return o
}

// SetVpnGatewayID adds the vpnGatewayId to the delete vpn gateways vpn gateway ID connections ID peer cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) SetVpnGatewayID(vpnGatewayID string) {
	o.VpnGatewayID = vpnGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param prefix_address
	if err := r.SetPathParam("prefix_address", o.PrefixAddress); err != nil {
		return err
	}

	// path param prefix_length
	if err := r.SetPathParam("prefix_length", o.PrefixLength); err != nil {
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
