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

// DeleteVpcsVpcIDAddressPrefixesIDReader is a Reader for the DeleteVpcsVpcIDAddressPrefixesID structure.
type DeleteVpcsVpcIDAddressPrefixesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVpcsVpcIDAddressPrefixesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVpcsVpcIDAddressPrefixesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVpcsVpcIDAddressPrefixesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVpcsVpcIDAddressPrefixesIDNoContent creates a DeleteVpcsVpcIDAddressPrefixesIDNoContent with default headers values
func NewDeleteVpcsVpcIDAddressPrefixesIDNoContent() *DeleteVpcsVpcIDAddressPrefixesIDNoContent {
	return &DeleteVpcsVpcIDAddressPrefixesIDNoContent{}
}

/*DeleteVpcsVpcIDAddressPrefixesIDNoContent handles this case with default header values.

The prefix was deleted successfully.
*/
type DeleteVpcsVpcIDAddressPrefixesIDNoContent struct {
}

func (o *DeleteVpcsVpcIDAddressPrefixesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{vpc_id}/address_prefixes/{id}][%d] deleteVpcsVpcIdAddressPrefixesIdNoContent ", 204)
}

func (o *DeleteVpcsVpcIDAddressPrefixesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVpcsVpcIDAddressPrefixesIDNotFound creates a DeleteVpcsVpcIDAddressPrefixesIDNotFound with default headers values
func NewDeleteVpcsVpcIDAddressPrefixesIDNotFound() *DeleteVpcsVpcIDAddressPrefixesIDNotFound {
	return &DeleteVpcsVpcIDAddressPrefixesIDNotFound{}
}

/*DeleteVpcsVpcIDAddressPrefixesIDNotFound handles this case with default header values.

error
*/
type DeleteVpcsVpcIDAddressPrefixesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteVpcsVpcIDAddressPrefixesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /vpcs/{vpc_id}/address_prefixes/{id}][%d] deleteVpcsVpcIdAddressPrefixesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVpcsVpcIDAddressPrefixesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
