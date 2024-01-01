package tools

import "github.com/sirupsen/logrus"

type db struct {}

var employees map[uint]Employee = map[uint]Employee {
	1 : {
		Id: 1,
		Username: "Aayush Vyas",
		Age: 24,
		Craft: "SDE",
	},
	2 : {
		Id: 2,
		Username: "Someone from",
	},
}
var auth map[uint]string = map[uint]string {
	1 : "pat_token",
	2 : "pat",
}

func (db *db) GetEmployeeById(id uint) (*Employee){
	var employee , found = employees[id]
	if(!found){
		return nil
	}
	return &employee
}

func (db *db) CheckAuth(id uint, token string) bool {
	var authToken , present  = auth[id]
	if present {
		logrus.Info(token)
		return authToken == token
	}
	return present
}

func (db *db) GetEmployees() map[uint]Employee {
	return employees
}