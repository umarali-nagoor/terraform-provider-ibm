// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewPostKeysParams creates a new PostKeysParams object
// with the default values initialized.
func NewPostKeysParams() *PostKeysParams {
	var ()
	return &PostKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKeysParamsWithTimeout creates a new PostKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKeysParamsWithTimeout(timeout time.Duration) *PostKeysParams {
	var ()
	return &PostKeysParams{

		timeout: timeout,
	}
}

// NewPostKeysParamsWithContext creates a new PostKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKeysParamsWithContext(ctx context.Context) *PostKeysParams {
	var ()
	return &PostKeysParams{

		Context: ctx,
	}
}

// NewPostKeysParamsWithHTTPClient creates a new PostKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKeysParamsWithHTTPClient(client *http.Client) *PostKeysParams {
	var ()
	return &PostKeysParams{
		HTTPClient: client,
	}
}

/*PostKeysParams contains all the parameters to send to the API endpoint
for the post keys operation typically these are written to a http.Request
*/
type PostKeysParams struct {

	/*Body*/
	Body PostKeysBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post keys params
func (o *PostKeysParams) WithTimeout(timeout time.Duration) *PostKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post keys params
func (o *PostKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post keys params
func (o *PostKeysParams) WithContext(ctx context.Context) *PostKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post keys params
func (o *PostKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post keys params
func (o *PostKeysParams) WithHTTPClient(client *http.Client) *PostKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post keys params
func (o *PostKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post keys params
func (o *PostKeysParams) WithBody(body PostKeysBody) *PostKeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post keys params
func (o *PostKeysParams) SetBody(body PostKeysBody) {
	o.Body = body
}

// WithGeneration adds the generation to the post keys params
func (o *PostKeysParams) WithGeneration(generation int64) *PostKeysParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post keys params
func (o *PostKeysParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the post keys params
func (o *PostKeysParams) WithVersion(version string) *PostKeysParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post keys params
func (o *PostKeysParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
