package structure

type Employee struct {
	Id       int     `json:"id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Position string  `json:"position" validate:"required"`
	Salary   float64 `json:"salary" validate:"required"`
	Checkid  int
}

type Respose struct {
	StatusCode int     `json:"StatusCode"`
	StatusDesc string  `json:"StatusDesc"`
	Id         int     `json:"Id"`
	Name       string  `json:"Name"`
	Position   string  `json:"Position"`
	Salary     float64 `json:"Salary"`
}
type EmployeeSelect struct {
	Id       int `json:"id" validate:"required"`
	Name     string
	Position string
	Salary   float64
	Selectid int
}
type EmployeeUpdatedetails struct {
	Id       int     `json:"id" validate:"required"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	StatusCode int     `json:"StatusCode"`
	StatusDesc string  `json:"StatusDesc"`
	Selectid int
}
