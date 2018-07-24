// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/getlunaform/lunaform/models"
)

// UndeployStackReader is a Reader for the UndeployStack structure.
type UndeployStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UndeployStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUndeployStackNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewUndeployStackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUndeployStackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUndeployStackNoContent creates a UndeployStackNoContent with default headers values
func NewUndeployStackNoContent() *UndeployStackNoContent {
	return &UndeployStackNoContent{}
}

/*UndeployStackNoContent handles this case with default header values.

No Content
*/
type UndeployStackNoContent struct {
}

func (o *UndeployStackNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tf/stack/{id}][%d] undeployStackNoContent ", 204)
}

func (o *UndeployStackNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUndeployStackUnprocessableEntity creates a UndeployStackUnprocessableEntity with default headers values
func NewUndeployStackUnprocessableEntity() *UndeployStackUnprocessableEntity {
	return &UndeployStackUnprocessableEntity{}
}

/*UndeployStackUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type UndeployStackUnprocessableEntity struct {
	Payload *models.ServerError
}

func (o *UndeployStackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /tf/stack/{id}][%d] undeployStackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UndeployStackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUndeployStackInternalServerError creates a UndeployStackInternalServerError with default headers values
func NewUndeployStackInternalServerError() *UndeployStackInternalServerError {
	return &UndeployStackInternalServerError{}
}

/*UndeployStackInternalServerError handles this case with default header values.

Internal Server Error
*/
type UndeployStackInternalServerError struct {
	Payload *models.ServerError
}

func (o *UndeployStackInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /tf/stack/{id}][%d] undeployStackInternalServerError  %+v", 500, o.Payload)
}

func (o *UndeployStackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
