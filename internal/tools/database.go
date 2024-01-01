package tools

type DB interface {
	GetEmployeeById(id uint) *Employee
	CheckAuth(id uint, token string) bool
	GetEmployees() map[uint]Employee
}
type Auth struct {
	id uint
	token string
}
type Employee struct {
	Id uint 
	Username string
	Age uint
	EmployementType string
	Designation string
	Craft string
}

func GetDB() *DB {
	var database DB = &db{}
	return &database
}