// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostVpcsReader is a Reader for the PostVpcs structure.
type PostVpcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVpcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostVpcsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostVpcsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostVpcsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostVpcsCreated creates a PostVpcsCreated with default headers values
func NewPostVpcsCreated() *PostVpcsCreated {
	return &PostVpcsCreated{}
}

/*PostVpcsCreated handles this case with default header values.

dummy
*/
type PostVpcsCreated struct {
	Payload *models.Vpc
}

func (o *PostVpcsCreated) Error() string {
	return fmt.Sprintf("[POST /vpcs][%d] postVpcsCreated  %+v", 201, o.Payload)
}

func (o *PostVpcsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vpc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVpcsBadRequest creates a PostVpcsBadRequest with default headers values
func NewPostVpcsBadRequest() *PostVpcsBadRequest {
	return &PostVpcsBadRequest{}
}

/*PostVpcsBadRequest handles this case with default header values.

error
*/
type PostVpcsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostVpcsBadRequest) Error() string {
	return fmt.Sprintf("[POST /vpcs][%d] postVpcsBadRequest  %+v", 400, o.Payload)
}

func (o *PostVpcsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVpcsInternalServerError creates a PostVpcsInternalServerError with default headers values
func NewPostVpcsInternalServerError() *PostVpcsInternalServerError {
	return &PostVpcsInternalServerError{}
}

/*PostVpcsInternalServerError handles this case with default header values.

error
*/
type PostVpcsInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PostVpcsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vpcs][%d] postVpcsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostVpcsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
