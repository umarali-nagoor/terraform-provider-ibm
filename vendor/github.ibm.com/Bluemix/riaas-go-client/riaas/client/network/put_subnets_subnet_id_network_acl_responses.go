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

// PutSubnetsSubnetIDNetworkACLReader is a Reader for the PutSubnetsSubnetIDNetworkACL structure.
type PutSubnetsSubnetIDNetworkACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSubnetsSubnetIDNetworkACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutSubnetsSubnetIDNetworkACLCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutSubnetsSubnetIDNetworkACLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutSubnetsSubnetIDNetworkACLInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSubnetsSubnetIDNetworkACLCreated creates a PutSubnetsSubnetIDNetworkACLCreated with default headers values
func NewPutSubnetsSubnetIDNetworkACLCreated() *PutSubnetsSubnetIDNetworkACLCreated {
	return &PutSubnetsSubnetIDNetworkACLCreated{}
}

/*PutSubnetsSubnetIDNetworkACLCreated handles this case with default header values.

dummy
*/
type PutSubnetsSubnetIDNetworkACLCreated struct {
	Payload *models.NetworkACL
}

func (o *PutSubnetsSubnetIDNetworkACLCreated) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnet_id}/network_acl][%d] putSubnetsSubnetIdNetworkAclCreated  %+v", 201, o.Payload)
}

func (o *PutSubnetsSubnetIDNetworkACLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSubnetsSubnetIDNetworkACLBadRequest creates a PutSubnetsSubnetIDNetworkACLBadRequest with default headers values
func NewPutSubnetsSubnetIDNetworkACLBadRequest() *PutSubnetsSubnetIDNetworkACLBadRequest {
	return &PutSubnetsSubnetIDNetworkACLBadRequest{}
}

/*PutSubnetsSubnetIDNetworkACLBadRequest handles this case with default header values.

error
*/
type PutSubnetsSubnetIDNetworkACLBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutSubnetsSubnetIDNetworkACLBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnet_id}/network_acl][%d] putSubnetsSubnetIdNetworkAclBadRequest  %+v", 400, o.Payload)
}

func (o *PutSubnetsSubnetIDNetworkACLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSubnetsSubnetIDNetworkACLInternalServerError creates a PutSubnetsSubnetIDNetworkACLInternalServerError with default headers values
func NewPutSubnetsSubnetIDNetworkACLInternalServerError() *PutSubnetsSubnetIDNetworkACLInternalServerError {
	return &PutSubnetsSubnetIDNetworkACLInternalServerError{}
}

/*PutSubnetsSubnetIDNetworkACLInternalServerError handles this case with default header values.

error
*/
type PutSubnetsSubnetIDNetworkACLInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PutSubnetsSubnetIDNetworkACLInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnet_id}/network_acl][%d] putSubnetsSubnetIdNetworkAclInternalServerError  %+v", 500, o.Payload)
}

func (o *PutSubnetsSubnetIDNetworkACLInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
