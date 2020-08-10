// Code generated by go-swagger; DO NOT EDIT.

package receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prymitive/karma/internal/mapper/v017/models"
)

// GetReceiversReader is a Reader for the GetReceivers structure.
type GetReceiversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReceiversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReceiversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReceiversOK creates a GetReceiversOK with default headers values
func NewGetReceiversOK() *GetReceiversOK {
	return &GetReceiversOK{}
}

/*GetReceiversOK handles this case with default header values.

Get receivers response
*/
type GetReceiversOK struct {
	Payload []*models.Receiver
}

func (o *GetReceiversOK) Error() string {
	return fmt.Sprintf("[GET /receivers][%d] getReceiversOK  %+v", 200, o.Payload)
}

func (o *GetReceiversOK) GetPayload() []*models.Receiver {
	return o.Payload
}

func (o *GetReceiversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
