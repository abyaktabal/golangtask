package dboperation

import (
	"database/sql"
	"fmt"
	"strconv"

	structure "server.go/Structure"
)

func Insertemployee(data structure.Employee, conn *sql.DB) (structure.Respose, error) {
	res := structure.Respose{}
	err = conn.QueryRow("SELECT id FROM employee.employee_details WHERE id=" + fmt.Sprint(data.Id)).Scan(&data.Checkid)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Employee Id before insertion", data.Id)
	if data.Checkid == 0 {
		fmt.Println("Employee ID is not present")
		result, err := db.Exec("INSERT INTO employee.employee_details (id, name, position, salary) VALUES (?, ?, ?, ?)", data.Id, data.Name, data.Position, data.Salary)
		if err != nil {
			fmt.Println("Insert Error is:", err.Error(), "Of user", result)
			res.StatusDesc = "Employee on_board failed due to some Db problem."
			res.StatusCode = -1
			return res, err
		} else {
			fmt.Println("Insert Error is:", err, "Of user", result)
			res.StatusDesc = "Emoplyee_Id :" + strconv.Itoa(data.Id) + " is successfully inserted ."
			res.StatusCode = 0
			return res, err
		}

	} else {
		res.StatusCode = -1
		res.StatusDesc = "Employee_id is:" + strconv.Itoa(data.Id) + " already exits."
		return res, err
	}

}
func FetchEmployee(data structure.EmployeeSelect, conn *sql.DB) (structure.Respose, error) {
	res := structure.Respose{}
	err = conn.QueryRow("SELECT id,name,position,salary FROM employee.employee_details WHERE id="+fmt.Sprint(data.Id)).Scan(&data.Id, &data.Name, &data.Position, &data.Salary)
	if err == sql.ErrNoRows {
		fmt.Println("Employee Id before select", data.Id)
		res.StatusCode = -1
		res.StatusDesc = "Employee_id is:" + strconv.Itoa(data.Id) + " Not On_boarded please onboard again ."
		return res, err

	} else {
		if data.Id != 0 {
			fmt.Println("Employee Details:", data)
			res.Id = data.Id
			res.Name = data.Name
			res.Position = data.Position
			res.Salary = data.Salary
			res.StatusCode = 0
			res.StatusDesc = "Employee Details Successfully fetched"
			return res, err

		}
	}
	return res, nil
}
func UpdateEmployeeDetais(data structure.EmployeeUpdatedetails, conn *sql.DB) (structure.Respose, error) {
	res := structure.Respose{}
	fmt.Println("for updating this user data", data)
	query := "UPDATE employee.employee_details SET name = ?, position = ? , salary = ? WHERE id = ?"
	result, err := db.Exec(query, data.Name, data.Position, data.Salary, data.Id)
	// result, err := db.Exec("update employee.employee_details SET name = " + fmt.Sprint(data.Name) + "position = " + fmt.Sprint(data.Position) + "salary = " + fmt.Sprint(data.Salary) + "'\n" + "WHERE id = " + fmt.Sprint(data.Id))
	if err != nil {
		fmt.Println("Insert Error is:", err.Error(), "Of user", result)
		res.StatusDesc = "Employee Details failed due to some Db problem."
		res.StatusCode = -1
		return res, err
	} else {
		res.StatusDesc = "Emoplyee_Id :" + strconv.Itoa(data.Id) + " is successfully Updated ."
		res.StatusCode = 0
		return res, err
	}
}
func DeleteEmployeeDetais(data structure.Respose, conn *sql.DB) (structure.Respose, error) {
	res := structure.Respose{}
	fmt.Println("for delete this user data", data)
	query := "DELETE FROM employee.employee_details WHERE id = ?"
	result, err := db.Exec(query, data.Id)
	if err != nil {
		fmt.Println("Insert Error is:", err.Error(), "Of user", result)
		res.StatusDesc = "Employee deleted failed due to some Db problem."
		res.StatusCode = -1
		return res, err
	} else {
		res.StatusDesc = "Emoplyee_Id :" + strconv.Itoa(data.Id) + " is successfully Deleted ."
		res.StatusCode = 0
		return res, err
	}
}
