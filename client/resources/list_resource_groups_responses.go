// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/drewsonne/terraform-server/models"
)

// ListResourceGroupsReader is a Reader for the ListResourceGroups structure.
type ListResourceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListResourceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListResourceGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListResourceGroupsOK creates a ListResourceGroupsOK with default headers values
func NewListResourceGroupsOK() *ListResourceGroupsOK {
	return &ListResourceGroupsOK{}
}

/*ListResourceGroupsOK handles this case with default header values.

OK
*/
type ListResourceGroupsOK struct {
	Payload *models.ResponseListResources
}

func (o *ListResourceGroupsOK) Error() string {
	return fmt.Sprintf("[GET /][%d] listResourceGroupsOK  %+v", 200, o.Payload)
}

func (o *ListResourceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseListResources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourceGroupsNotFound creates a ListResourceGroupsNotFound with default headers values
func NewListResourceGroupsNotFound() *ListResourceGroupsNotFound {
	return &ListResourceGroupsNotFound{}
}

/*ListResourceGroupsNotFound handles this case with default header values.

Not Found
*/
type ListResourceGroupsNotFound struct {
}

func (o *ListResourceGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /][%d] listResourceGroupsNotFound ", 404)
}

func (o *ListResourceGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
