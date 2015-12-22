package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*UpdateUserBadRequest Invalid user supplied

swagger:response updateUserBadRequest
*/
type UpdateUserBadRequest struct {
}

// Create UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() UpdateUserBadRequest {
	return UpdateUserBadRequest{}
}

// WriteResponse to the client
func (o *UpdateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(400)
}

/*UpdateUserNotFound User not found

swagger:response updateUserNotFound
*/
type UpdateUserNotFound struct {
}

// Create UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() UpdateUserNotFound {
	return UpdateUserNotFound{}
}

// WriteResponse to the client
func (o *UpdateUserNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
}
