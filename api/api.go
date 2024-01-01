package api

import (
	"encoding/json"
	"net/http"
)

type EmployeeParams struct{
	Id uint
}

type EmployeeResponse struct {
	Id uint 
	Username string
	Age uint
	EmployementType string
	Designation string
	Craft string
}

type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	var errorResponse = Error{
		Code : code,
		Message : message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(errorResponse)

}

var RequestErrorHandler = func (w http.ResponseWriter, err error)  {
	writeError(w, err.Error(), http.StatusBadRequest)
}

var InternalErrorHandler = func (w http.ResponseWriter, err error)  {
	writeError(w, "Internal Server Error", http.StatusInternalServerError)
}