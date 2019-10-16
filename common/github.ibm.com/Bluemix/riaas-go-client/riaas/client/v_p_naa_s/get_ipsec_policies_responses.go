// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetIpsecPoliciesReader is a Reader for the GetIpsecPolicies structure.
type GetIpsecPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpsecPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIpsecPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIpsecPoliciesOK creates a GetIpsecPoliciesOK with default headers values
func NewGetIpsecPoliciesOK() *GetIpsecPoliciesOK {
	return &GetIpsecPoliciesOK{}
}

/*GetIpsecPoliciesOK handles this case with default header values.

The IPsec policies were retrieved successfully.
*/
type GetIpsecPoliciesOK struct {
	Payload *models.IpsecPolicyCollection
}

func (o *GetIpsecPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /ipsec_policies][%d] getIpsecPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetIpsecPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpsecPolicyCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
