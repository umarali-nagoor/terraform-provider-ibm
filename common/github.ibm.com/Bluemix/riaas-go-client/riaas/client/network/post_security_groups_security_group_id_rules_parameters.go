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

// NewPostSecurityGroupsSecurityGroupIDRulesParams creates a new PostSecurityGroupsSecurityGroupIDRulesParams object
// with the default values initialized.
func NewPostSecurityGroupsSecurityGroupIDRulesParams() *PostSecurityGroupsSecurityGroupIDRulesParams {
	var ()
	return &PostSecurityGroupsSecurityGroupIDRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSecurityGroupsSecurityGroupIDRulesParamsWithTimeout creates a new PostSecurityGroupsSecurityGroupIDRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSecurityGroupsSecurityGroupIDRulesParamsWithTimeout(timeout time.Duration) *PostSecurityGroupsSecurityGroupIDRulesParams {
	var ()
	return &PostSecurityGroupsSecurityGroupIDRulesParams{

		timeout: timeout,
	}
}

// NewPostSecurityGroupsSecurityGroupIDRulesParamsWithContext creates a new PostSecurityGroupsSecurityGroupIDRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSecurityGroupsSecurityGroupIDRulesParamsWithContext(ctx context.Context) *PostSecurityGroupsSecurityGroupIDRulesParams {
	var ()
	return &PostSecurityGroupsSecurityGroupIDRulesParams{

		Context: ctx,
	}
}

// NewPostSecurityGroupsSecurityGroupIDRulesParamsWithHTTPClient creates a new PostSecurityGroupsSecurityGroupIDRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSecurityGroupsSecurityGroupIDRulesParamsWithHTTPClient(client *http.Client) *PostSecurityGroupsSecurityGroupIDRulesParams {
	var ()
	return &PostSecurityGroupsSecurityGroupIDRulesParams{
		HTTPClient: client,
	}
}

/*PostSecurityGroupsSecurityGroupIDRulesParams contains all the parameters to send to the API endpoint
for the post security groups security group ID rules operation typically these are written to a http.Request
*/
type PostSecurityGroupsSecurityGroupIDRulesParams struct {

	/*Body*/
	Body *models.PostSecurityGroupsSecurityGroupIDRulesParamsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*SecurityGroupID
	  The security group identifier

	*/
	SecurityGroupID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithTimeout(timeout time.Duration) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithContext(ctx context.Context) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithHTTPClient(client *http.Client) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithBody(body *models.PostSecurityGroupsSecurityGroupIDRulesParamsBody) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetBody(body *models.PostSecurityGroupsSecurityGroupIDRulesParamsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithGeneration(generation int64) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithSecurityGroupID adds the securityGroupID to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithSecurityGroupID(securityGroupID string) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetSecurityGroupID(securityGroupID)
	return o
}

// SetSecurityGroupID adds the securityGroupId to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetSecurityGroupID(securityGroupID string) {
	o.SecurityGroupID = securityGroupID
}

// WithVersion adds the version to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WithVersion(version string) *PostSecurityGroupsSecurityGroupIDRulesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post security groups security group ID rules params
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostSecurityGroupsSecurityGroupIDRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param security_group_id
	if err := r.SetPathParam("security_group_id", o.SecurityGroupID); err != nil {
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
