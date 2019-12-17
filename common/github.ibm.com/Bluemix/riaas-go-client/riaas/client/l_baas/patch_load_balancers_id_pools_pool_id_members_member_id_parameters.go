// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams creates a new PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized.
func NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams() *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout creates a new PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout(timeout time.Duration) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams{

		timeout: timeout,
	}
}

// NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithContext creates a new PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithContext(ctx context.Context) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams{

		Context: ctx,
	}
}

// NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithHTTPClient creates a new PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithHTTPClient(client *http.Client) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams{
		HTTPClient: client,
	}
}

/*PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams contains all the parameters to send to the API endpoint
for the patch load balancers ID pools pool ID members member ID operation typically these are written to a http.Request
*/
type PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams struct {

	/*Body
	  The member template

	*/
	Body models.MemberTemplatePatch
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The load balancer identifier

	*/
	ID string
	/*MemberID
	  The member identifier

	*/
	MemberID string
	/*PoolID
	  The pool identifier

	*/
	PoolID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithTimeout(timeout time.Duration) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithContext(ctx context.Context) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithHTTPClient(client *http.Client) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithBody(body models.MemberTemplatePatch) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetBody(body models.MemberTemplatePatch) {
	o.Body = body
}

// WithGeneration adds the generation to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithGeneration(generation int64) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithID(id string) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetID(id string) {
	o.ID = id
}

// WithMemberID adds the memberID to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithMemberID(memberID string) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WithPoolID adds the poolID to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithPoolID(poolID string) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithVersion adds the version to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithVersion(version string) *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch load balancers ID pools pool ID members member ID params
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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

	// path param member_id
	if err := r.SetPathParam("member_id", o.MemberID); err != nil {
		return err
	}

	// path param pool_id
	if err := r.SetPathParam("pool_id", o.PoolID); err != nil {
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
