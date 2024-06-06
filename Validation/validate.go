package validation

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
	structure "server.go/Structure"
)

//VALIDATION FOR CREATE BENE API
func Checkvalidationforcreateemployee(data structure.Employee) error {
	v := validator.New()

	a := structure.Employee{
		Id:       data.Id,
		Name:     data.Name,
		Position: data.Position,
		Salary:   data.Salary,
	}
	err := v.Struct(a)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("Validation failure", e)
			return err
		}
	}

	return nil
}

//VALIDATION FOR EMPLOYEE FETCH API
func CheckValidationForFetchEmployee(data structure.EmployeeSelect) error {
	v := validator.New()
	a := structure.EmployeeSelect{
		Id: data.Id,
	}
	err := v.Struct(a)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("Validation failure", e)
			return err
		}
	}

	return nil
}
func CheckValidationForUpdateEmployee(data structure.EmployeeUpdatedetails) error {
	v := validator.New()
	a := structure.EmployeeUpdatedetails{
		Id: data.Id,
	}
	err := v.Struct(a)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("Validation failure", e)
			return err
		}
	}

	return nil
}