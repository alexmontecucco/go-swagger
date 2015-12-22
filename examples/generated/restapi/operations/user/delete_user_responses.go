package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*DeleteUserBadRequest Invalid username supplied

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {
}

// Create DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() DeleteUserBadRequest {
	return DeleteUserBadRequest{}
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(400)
}

/*DeleteUserNotFound User not found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct {
}

// Create DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() DeleteUserNotFound {
	return DeleteUserNotFound{}
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
}
