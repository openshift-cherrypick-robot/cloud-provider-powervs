// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerAuthRegistrationReader is a Reader for the ServiceBrokerAuthRegistration structure.
type ServiceBrokerAuthRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewServiceBrokerAuthRegistrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerAuthRegistrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerAuthRegistrationOK creates a ServiceBrokerAuthRegistrationOK with default headers values
func NewServiceBrokerAuthRegistrationOK() *ServiceBrokerAuthRegistrationOK {
	return &ServiceBrokerAuthRegistrationOK{}
}

/*
ServiceBrokerAuthRegistrationOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthRegistrationOK struct {
	Payload *models.AccessToken
}

// IsSuccess returns true when this service broker auth registration o k response has a 2xx status code
func (o *ServiceBrokerAuthRegistrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker auth registration o k response has a 3xx status code
func (o *ServiceBrokerAuthRegistrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth registration o k response has a 4xx status code
func (o *ServiceBrokerAuthRegistrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth registration o k response has a 5xx status code
func (o *ServiceBrokerAuthRegistrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth registration o k response a status code equal to that given
func (o *ServiceBrokerAuthRegistrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceBrokerAuthRegistrationOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/registration][%d] serviceBrokerAuthRegistrationOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationOK) String() string {
	return fmt.Sprintf("[GET /auth/v1/registration][%d] serviceBrokerAuthRegistrationOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationOK) GetPayload() *models.AccessToken {
	return o.Payload
}

func (o *ServiceBrokerAuthRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthRegistrationUnauthorized creates a ServiceBrokerAuthRegistrationUnauthorized with default headers values
func NewServiceBrokerAuthRegistrationUnauthorized() *ServiceBrokerAuthRegistrationUnauthorized {
	return &ServiceBrokerAuthRegistrationUnauthorized{}
}

/*
ServiceBrokerAuthRegistrationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerAuthRegistrationUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth registration unauthorized response has a 2xx status code
func (o *ServiceBrokerAuthRegistrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth registration unauthorized response has a 3xx status code
func (o *ServiceBrokerAuthRegistrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth registration unauthorized response has a 4xx status code
func (o *ServiceBrokerAuthRegistrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth registration unauthorized response has a 5xx status code
func (o *ServiceBrokerAuthRegistrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth registration unauthorized response a status code equal to that given
func (o *ServiceBrokerAuthRegistrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ServiceBrokerAuthRegistrationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/v1/registration][%d] serviceBrokerAuthRegistrationUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationUnauthorized) String() string {
	return fmt.Sprintf("[GET /auth/v1/registration][%d] serviceBrokerAuthRegistrationUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthRegistrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthRegistrationInternalServerError creates a ServiceBrokerAuthRegistrationInternalServerError with default headers values
func NewServiceBrokerAuthRegistrationInternalServerError() *ServiceBrokerAuthRegistrationInternalServerError {
	return &ServiceBrokerAuthRegistrationInternalServerError{}
}

/*
ServiceBrokerAuthRegistrationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthRegistrationInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth registration internal server error response has a 2xx status code
func (o *ServiceBrokerAuthRegistrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth registration internal server error response has a 3xx status code
func (o *ServiceBrokerAuthRegistrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth registration internal server error response has a 4xx status code
func (o *ServiceBrokerAuthRegistrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth registration internal server error response has a 5xx status code
func (o *ServiceBrokerAuthRegistrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker auth registration internal server error response a status code equal to that given
func (o *ServiceBrokerAuthRegistrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServiceBrokerAuthRegistrationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/registration][%d] serviceBrokerAuthRegistrationInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationInternalServerError) String() string {
	return fmt.Sprintf("[GET /auth/v1/registration][%d] serviceBrokerAuthRegistrationInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthRegistrationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthRegistrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
