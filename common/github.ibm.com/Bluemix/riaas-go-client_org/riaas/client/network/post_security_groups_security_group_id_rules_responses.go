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

// PostSecurityGroupsSecurityGroupIDRulesReader is a Reader for the PostSecurityGroupsSecurityGroupIDRules structure.
type PostSecurityGroupsSecurityGroupIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSecurityGroupsSecurityGroupIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSecurityGroupsSecurityGroupIDRulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSecurityGroupsSecurityGroupIDRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostSecurityGroupsSecurityGroupIDRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSecurityGroupsSecurityGroupIDRulesCreated creates a PostSecurityGroupsSecurityGroupIDRulesCreated with default headers values
func NewPostSecurityGroupsSecurityGroupIDRulesCreated() *PostSecurityGroupsSecurityGroupIDRulesCreated {
	return &PostSecurityGroupsSecurityGroupIDRulesCreated{}
}

/*PostSecurityGroupsSecurityGroupIDRulesCreated handles this case with default header values.

dummy
*/
type PostSecurityGroupsSecurityGroupIDRulesCreated struct {
	Payload *models.SecurityGroupRule
}

func (o *PostSecurityGroupsSecurityGroupIDRulesCreated) Error() string {
	return fmt.Sprintf("[POST /security_groups/{security_group_id}/rules][%d] postSecurityGroupsSecurityGroupIdRulesCreated  %+v", 201, o.Payload)
}

func (o *PostSecurityGroupsSecurityGroupIDRulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSecurityGroupsSecurityGroupIDRulesBadRequest creates a PostSecurityGroupsSecurityGroupIDRulesBadRequest with default headers values
func NewPostSecurityGroupsSecurityGroupIDRulesBadRequest() *PostSecurityGroupsSecurityGroupIDRulesBadRequest {
	return &PostSecurityGroupsSecurityGroupIDRulesBadRequest{}
}

/*PostSecurityGroupsSecurityGroupIDRulesBadRequest handles this case with default header values.

error
*/
type PostSecurityGroupsSecurityGroupIDRulesBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /security_groups/{security_group_id}/rules][%d] postSecurityGroupsSecurityGroupIdRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSecurityGroupsSecurityGroupIDRulesInternalServerError creates a PostSecurityGroupsSecurityGroupIDRulesInternalServerError with default headers values
func NewPostSecurityGroupsSecurityGroupIDRulesInternalServerError() *PostSecurityGroupsSecurityGroupIDRulesInternalServerError {
	return &PostSecurityGroupsSecurityGroupIDRulesInternalServerError{}
}

/*PostSecurityGroupsSecurityGroupIDRulesInternalServerError handles this case with default header values.

error
*/
type PostSecurityGroupsSecurityGroupIDRulesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PostSecurityGroupsSecurityGroupIDRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /security_groups/{security_group_id}/rules][%d] postSecurityGroupsSecurityGroupIdRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSecurityGroupsSecurityGroupIDRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
