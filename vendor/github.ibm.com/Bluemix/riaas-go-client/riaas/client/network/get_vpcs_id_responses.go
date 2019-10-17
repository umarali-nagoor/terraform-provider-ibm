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

// GetVpcsIDReader is a Reader for the GetVpcsID structure.
type GetVpcsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpcsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpcsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVpcsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetVpcsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVpcsIDOK creates a GetVpcsIDOK with default headers values
func NewGetVpcsIDOK() *GetVpcsIDOK {
	return &GetVpcsIDOK{}
}

/*GetVpcsIDOK handles this case with default header values.

dummy
*/
type GetVpcsIDOK struct {
	Payload *models.Vpc
}

func (o *GetVpcsIDOK) Error() string {
	return fmt.Sprintf("[GET /vpcs/{id}][%d] getVpcsIdOK  %+v", 200, o.Payload)
}

func (o *GetVpcsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vpc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpcsIDNotFound creates a GetVpcsIDNotFound with default headers values
func NewGetVpcsIDNotFound() *GetVpcsIDNotFound {
	return &GetVpcsIDNotFound{}
}

/*GetVpcsIDNotFound handles this case with default header values.

error
*/
type GetVpcsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetVpcsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /vpcs/{id}][%d] getVpcsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetVpcsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpcsIDInternalServerError creates a GetVpcsIDInternalServerError with default headers values
func NewGetVpcsIDInternalServerError() *GetVpcsIDInternalServerError {
	return &GetVpcsIDInternalServerError{}
}

/*GetVpcsIDInternalServerError handles this case with default header values.

error
*/
type GetVpcsIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetVpcsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vpcs/{id}][%d] getVpcsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVpcsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
