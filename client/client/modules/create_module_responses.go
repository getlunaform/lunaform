// Code generated by go-swagger; DO NOT EDIT.

package modules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/drewsonne/lunaform/server/models"
)

// CreateModuleReader is a Reader for the CreateModule structure.
type CreateModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateModuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateModuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateModuleCreated creates a CreateModuleCreated with default headers values
func NewCreateModuleCreated() *CreateModuleCreated {
	return &CreateModuleCreated{}
}

/*CreateModuleCreated handles this case with default header values.

OK
*/
type CreateModuleCreated struct {
	Payload *models.ResourceTfModule
}

func (o *CreateModuleCreated) Error() string {
	return fmt.Sprintf("[POST /tf/modules][%d] createModuleCreated  %+v", 201, o.Payload)
}

func (o *CreateModuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceTfModule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModuleBadRequest creates a CreateModuleBadRequest with default headers values
func NewCreateModuleBadRequest() *CreateModuleBadRequest {
	return &CreateModuleBadRequest{}
}

/*CreateModuleBadRequest handles this case with default header values.

Bad Request
*/
type CreateModuleBadRequest struct {
}

func (o *CreateModuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /tf/modules][%d] createModuleBadRequest ", 400)
}

func (o *CreateModuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
