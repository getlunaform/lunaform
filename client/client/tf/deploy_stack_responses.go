// Code generated by go-swagger; DO NOT EDIT.

package tf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/drewsonne/terraform-server/server/models"
)

// DeployStackReader is a Reader for the DeployStack structure.
type DeployStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeployStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeployStackAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeployStackAccepted creates a DeployStackAccepted with default headers values
func NewDeployStackAccepted() *DeployStackAccepted {
	return &DeployStackAccepted{}
}

/*DeployStackAccepted handles this case with default header values.

Accepted
*/
type DeployStackAccepted struct {
	Payload *models.ResourceTfStack
}

func (o *DeployStackAccepted) Error() string {
	return fmt.Sprintf("[POST /tf/stacks][%d] deployStackAccepted  %+v", 202, o.Payload)
}

func (o *DeployStackAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceTfStack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
