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

// PostNetworkAclsNetworkACLIDRulesReader is a Reader for the PostNetworkAclsNetworkACLIDRules structure.
type PostNetworkAclsNetworkACLIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNetworkAclsNetworkACLIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostNetworkAclsNetworkACLIDRulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostNetworkAclsNetworkACLIDRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostNetworkAclsNetworkACLIDRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNetworkAclsNetworkACLIDRulesCreated creates a PostNetworkAclsNetworkACLIDRulesCreated with default headers values
func NewPostNetworkAclsNetworkACLIDRulesCreated() *PostNetworkAclsNetworkACLIDRulesCreated {
	return &PostNetworkAclsNetworkACLIDRulesCreated{}
}

/*PostNetworkAclsNetworkACLIDRulesCreated handles this case with default header values.

dummy
*/
type PostNetworkAclsNetworkACLIDRulesCreated struct {
	Payload *models.NetworkACLRule
}

func (o *PostNetworkAclsNetworkACLIDRulesCreated) Error() string {
	return fmt.Sprintf("[POST /network_acls/{network_acl_id}/rules][%d] postNetworkAclsNetworkAclIdRulesCreated  %+v", 201, o.Payload)
}

func (o *PostNetworkAclsNetworkACLIDRulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkACLRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkAclsNetworkACLIDRulesBadRequest creates a PostNetworkAclsNetworkACLIDRulesBadRequest with default headers values
func NewPostNetworkAclsNetworkACLIDRulesBadRequest() *PostNetworkAclsNetworkACLIDRulesBadRequest {
	return &PostNetworkAclsNetworkACLIDRulesBadRequest{}
}

/*PostNetworkAclsNetworkACLIDRulesBadRequest handles this case with default header values.

error
*/
type PostNetworkAclsNetworkACLIDRulesBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostNetworkAclsNetworkACLIDRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /network_acls/{network_acl_id}/rules][%d] postNetworkAclsNetworkAclIdRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostNetworkAclsNetworkACLIDRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNetworkAclsNetworkACLIDRulesInternalServerError creates a PostNetworkAclsNetworkACLIDRulesInternalServerError with default headers values
func NewPostNetworkAclsNetworkACLIDRulesInternalServerError() *PostNetworkAclsNetworkACLIDRulesInternalServerError {
	return &PostNetworkAclsNetworkACLIDRulesInternalServerError{}
}

/*PostNetworkAclsNetworkACLIDRulesInternalServerError handles this case with default header values.

error
*/
type PostNetworkAclsNetworkACLIDRulesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PostNetworkAclsNetworkACLIDRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /network_acls/{network_acl_id}/rules][%d] postNetworkAclsNetworkAclIdRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostNetworkAclsNetworkACLIDRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
