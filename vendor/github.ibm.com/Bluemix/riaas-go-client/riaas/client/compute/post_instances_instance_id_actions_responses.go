// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostInstancesInstanceIDActionsReader is a Reader for the PostInstancesInstanceIDActions structure.
type PostInstancesInstanceIDActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInstancesInstanceIDActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostInstancesInstanceIDActionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostInstancesInstanceIDActionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInstancesInstanceIDActionsCreated creates a PostInstancesInstanceIDActionsCreated with default headers values
func NewPostInstancesInstanceIDActionsCreated() *PostInstancesInstanceIDActionsCreated {
	return &PostInstancesInstanceIDActionsCreated{}
}

/*PostInstancesInstanceIDActionsCreated handles this case with default header values.

dummy
*/
type PostInstancesInstanceIDActionsCreated struct {
	Payload *models.InstanceAction
}

func (o *PostInstancesInstanceIDActionsCreated) Error() string {
	return fmt.Sprintf("[POST /instances/{instance_id}/actions][%d] postInstancesInstanceIdActionsCreated  %+v", 201, o.Payload)
}

func (o *PostInstancesInstanceIDActionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInstancesInstanceIDActionsBadRequest creates a PostInstancesInstanceIDActionsBadRequest with default headers values
func NewPostInstancesInstanceIDActionsBadRequest() *PostInstancesInstanceIDActionsBadRequest {
	return &PostInstancesInstanceIDActionsBadRequest{}
}

/*PostInstancesInstanceIDActionsBadRequest handles this case with default header values.

error
*/
type PostInstancesInstanceIDActionsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostInstancesInstanceIDActionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{instance_id}/actions][%d] postInstancesInstanceIdActionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostInstancesInstanceIDActionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
