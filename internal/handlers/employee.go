package handlers

import (
	"encoding/json"
	"go-rest-api/api"
	"go-rest-api/internal/tools"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

// Returns an Employee given Id

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	
	var params = api.EmployeeParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error 
	err = decoder.Decode(&params , r.URL.Query()) 

	if  err != nil {
		logrus.Error(err)
		api.InternalErrorHandler(w, err)
	}

	var db *tools.DB = tools.GetDB()
	var employee *tools.Employee = (*db).GetEmployeeById(params.Id)
	var response = api.EmployeeResponse{
		Age: (*employee).Age,
		EmployementType: (*employee).EmployementType,
		Designation: (*employee).Designation,
		Username: (*employee).Username,
		Id: (*employee).Id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.Error(err)
		api.InternalErrorHandler(w, err)
	}
}
// Returns a list of all Employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {

	var db *tools.DB = tools.GetDB()
	var employees map[uint]tools.Employee = (*db).GetEmployees()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var err error = json.NewEncoder(w).Encode(employees)
	if(err != nil) {
		logrus.Error("Error decoding")
		api.InternalErrorHandler(w,err)
	}
}