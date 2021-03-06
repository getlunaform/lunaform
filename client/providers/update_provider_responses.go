// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/getlunaform/lunaform/models"
)

// UpdateProviderReader is a Reader for the UpdateProvider structure.
type UpdateProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewUpdateProviderAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateProviderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProviderAccepted creates a UpdateProviderAccepted with default headers values
func NewUpdateProviderAccepted() *UpdateProviderAccepted {
	return &UpdateProviderAccepted{}
}

/*UpdateProviderAccepted handles this case with default header values.

Updated
*/
type UpdateProviderAccepted struct {
	Payload *models.ResourceTfProvider
}

func (o *UpdateProviderAccepted) Error() string {
	return fmt.Sprintf("[PUT /tf/provider/{name}][%d] updateProviderAccepted  %+v", 202, o.Payload)
}

func (o *UpdateProviderAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceTfProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProviderBadRequest creates a UpdateProviderBadRequest with default headers values
func NewUpdateProviderBadRequest() *UpdateProviderBadRequest {
	return &UpdateProviderBadRequest{}
}

/*UpdateProviderBadRequest handles this case with default header values.

Bad Request
*/
type UpdateProviderBadRequest struct {
	Payload *models.ServerError
}

func (o *UpdateProviderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tf/provider/{name}][%d] updateProviderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProviderNotFound creates a UpdateProviderNotFound with default headers values
func NewUpdateProviderNotFound() *UpdateProviderNotFound {
	return &UpdateProviderNotFound{}
}

/*UpdateProviderNotFound handles this case with default header values.

Not Found
*/
type UpdateProviderNotFound struct {
	Payload *models.ServerError
}

func (o *UpdateProviderNotFound) Error() string {
	return fmt.Sprintf("[PUT /tf/provider/{name}][%d] updateProviderNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProviderInternalServerError creates a UpdateProviderInternalServerError with default headers values
func NewUpdateProviderInternalServerError() *UpdateProviderInternalServerError {
	return &UpdateProviderInternalServerError{}
}

/*UpdateProviderInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateProviderInternalServerError struct {
	Payload *models.ServerError
}

func (o *UpdateProviderInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /tf/provider/{name}][%d] updateProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProviderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
