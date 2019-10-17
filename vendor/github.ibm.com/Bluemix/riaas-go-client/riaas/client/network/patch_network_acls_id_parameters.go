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

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewPatchNetworkAclsIDParams creates a new PatchNetworkAclsIDParams object
// with the default values initialized.
func NewPatchNetworkAclsIDParams() *PatchNetworkAclsIDParams {
	var ()
	return &PatchNetworkAclsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchNetworkAclsIDParamsWithTimeout creates a new PatchNetworkAclsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchNetworkAclsIDParamsWithTimeout(timeout time.Duration) *PatchNetworkAclsIDParams {
	var ()
	return &PatchNetworkAclsIDParams{

		timeout: timeout,
	}
}

// NewPatchNetworkAclsIDParamsWithContext creates a new PatchNetworkAclsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchNetworkAclsIDParamsWithContext(ctx context.Context) *PatchNetworkAclsIDParams {
	var ()
	return &PatchNetworkAclsIDParams{

		Context: ctx,
	}
}

// NewPatchNetworkAclsIDParamsWithHTTPClient creates a new PatchNetworkAclsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchNetworkAclsIDParamsWithHTTPClient(client *http.Client) *PatchNetworkAclsIDParams {
	var ()
	return &PatchNetworkAclsIDParams{
		HTTPClient: client,
	}
}

/*PatchNetworkAclsIDParams contains all the parameters to send to the API endpoint
for the patch network acls ID operation typically these are written to a http.Request
*/
type PatchNetworkAclsIDParams struct {

	/*Body*/
	Body *models.PatchNetworkAclsIDParamsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The network ACL identifier

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

// WithTimeout adds the timeout to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithTimeout(timeout time.Duration) *PatchNetworkAclsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithContext(ctx context.Context) *PatchNetworkAclsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithHTTPClient(client *http.Client) *PatchNetworkAclsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithBody(body *models.PatchNetworkAclsIDParamsBody) *PatchNetworkAclsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetBody(body *models.PatchNetworkAclsIDParamsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithGeneration(generation int64) *PatchNetworkAclsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithID(id string) *PatchNetworkAclsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) WithVersion(version string) *PatchNetworkAclsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch network acls ID params
func (o *PatchNetworkAclsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchNetworkAclsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
