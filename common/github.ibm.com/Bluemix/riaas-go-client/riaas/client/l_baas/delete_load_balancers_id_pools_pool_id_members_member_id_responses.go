// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDReader is a Reader for the DeleteLoadBalancersIDPoolsPoolIDMembersMemberID structure.
type DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent creates a DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent with default headers values
func NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent() *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent {
	return &DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent{}
}

/*DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent handles this case with default header values.

The member deletion request was accepted.
*/
type DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent struct {
}

func (o *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /load_balancers/{id}/pools/{pool_id}/members/{member_id}][%d] deleteLoadBalancersIdPoolsPoolIdMembersMemberIdNoContent ", 204)
}

func (o *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound creates a DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound with default headers values
func NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound() *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound {
	return &DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound{}
}

/*DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound handles this case with default header values.

A load balancer, pool or member with the specified identifier could not be found.
*/
type DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /load_balancers/{id}/pools/{pool_id}/members/{member_id}][%d] deleteLoadBalancersIdPoolsPoolIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLoadBalancersIDPoolsPoolIDMembersMemberIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
