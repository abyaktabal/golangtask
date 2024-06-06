package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dboperation "server.go/Dboperation"
	structure "server.go/Structure"
	validation "server.go/Validation"
)

func CreateEmployee(ctx *gin.Context) {
	var err error
	conn, err := dboperation.Connection() //CHECK DB FOR CONNECTION
	if err == nil {
		insertstruct := structure.Employee{} //STRUCTURE ASSIGN
		resposestruct := structure.Respose{}
		err = ctx.ShouldBindJSON(&insertstruct) //JSON BINDING
		if err != nil {
			ctx.JSON(400, gin.H{

				"StatusDesc": "ERROR IN BINDING DATA PLEASE PROVIDE THE CORRECT REQUEST FORMAT" + err.Error(),
			})

		} else {
			validationres := validation.Checkvalidationforcreateemployee(insertstruct) //VALIDATION CHECK
			if validationres != nil {
				ctx.JSON(201, gin.H{
					"Message": "validation Error" + validationres.Error(),
				})
			} else {

				resposestruct, err = dboperation.Insertemployee(insertstruct, conn) //CALL FOR FUNCTION QUERRY PACKAGE iNSERTBENE

				if err != nil {
					ctx.JSON(200, gin.H{
						"StatusCode": resposestruct.StatusCode,
						"StatusDesc": resposestruct.StatusDesc,
					})
				} else {

					ctx.JSON(201, gin.H{
						"StatusCode": resposestruct.StatusCode,
						"StatusDesc": resposestruct.StatusDesc,
					})
				}
			}

		}

	} else {
		ctx.JSON(201, gin.H{
			"StatusCode": -1,
			"StatusDesc": "DATABASE CONNECTION ERROR",
		})

	}
}
func SelectEmployee(ctx *gin.Context) {
	var err error
	conn, err := dboperation.Connection() //CHECK DB FOR CONNECTION
	if err == nil {
		selectstruct := structure.EmployeeSelect{} //STRUCTURE ASSIGN
		resposestruct := structure.Respose{}
		err = ctx.ShouldBindJSON(&selectstruct) //JSON BINDING
		if err != nil {
			ctx.JSON(400, gin.H{

				"StatusDesc": "ERROR IN BINDING DATA PLEASE PROVIDE THE CORRECT REQUEST FORMAT" + err.Error(),
			})

		} else {
			validationres := validation.CheckValidationForFetchEmployee(selectstruct) //VALIDATION CHECK
			if validationres != nil {
				ctx.JSON(201, gin.H{
					"Message": "validation Error" + validationres.Error(),
				})
			} else {

				resposestruct, err = dboperation.FetchEmployee(selectstruct, conn) //CALL FOR FUNCTION QUERRY PACKAGE iNSERTBENE

				if err != nil {
					ctx.JSON(200, gin.H{
						"StatusCode": resposestruct.StatusCode,
						"StatusDesc": resposestruct.StatusDesc,
					})
				} else {

					ctx.JSON(201, gin.H{
						"StastusCode":  resposestruct.StatusCode,
						"StatusDesc":   resposestruct.StatusDesc,
						"Emp-Name":     resposestruct.Name,
						"Emp-Id":       resposestruct.Id,
						"Emp-Position": resposestruct.Position,
						"Emp-Salary":   resposestruct.Salary,
					})
				}
			}

		}

	} else {
		ctx.JSON(201, gin.H{
			"StatusCode": -1,
			"StatusDesc": "DATABASE CONNECTION ERROR",
		})

	}
}
func UpdateEmployee(ctx *gin.Context) {
	var err error
	conn, err := dboperation.Connection() //CHECK DB FOR CONNECTION
	if err == nil {
		updatestruct := structure.EmployeeUpdatedetails{} //STRUCTURE ASSIGN
		selectstruct := structure.EmployeeSelect{}        //STRUCTURE ASSIGN FOR SELECTING EMPLOYEE DETAILS
		resposestruct := structure.Respose{}              //STRUCTURE ASSIGN FOR PROVIDING RESPONSE
		err = ctx.ShouldBindJSON(&updatestruct)           //JSON BINDING
		if err != nil {
			ctx.JSON(400, gin.H{

				"StatusDesc": "ERROR IN BINDING DATA PLEASE PROVIDE THE CORRECT REQUEST FORMAT" + err.Error(),
			})

		} else {
			validationres := validation.CheckValidationForUpdateEmployee(updatestruct) //VALIDATION CHECK
			if validationres != nil {
				ctx.JSON(201, gin.H{
					"Message": "validation Error" + validationres.Error(),
				})
			} else {
				selectstruct.Id = updatestruct.Id
				resposestruct, err = dboperation.FetchEmployee(selectstruct, conn) //CALL FOR FUNCTION QUERRY PACKAGE iNSERTBENE

				if err != nil {
					ctx.JSON(201, gin.H{
						"StatusCode": resposestruct.StatusCode,
						"StatusDesc": resposestruct.StatusDesc,
					})
				} else {
					fmt.Println("Now comming for update")
					resposestruct, err = dboperation.UpdateEmployeeDetais(updatestruct, conn)

					if err != nil {
						ctx.JSON(201, gin.H{
							"StatusCode": resposestruct.StatusCode,
							"StatusDesc": resposestruct.StatusDesc,
						})
					} else {
						ctx.JSON(200, gin.H{
							"StatusCode": resposestruct.StatusCode,
							"StatusDesc": resposestruct.StatusDesc,
						})
					}

				}
			}

		}

	} else {
		ctx.JSON(201, gin.H{
			"StatusCode": -1,
			"StatusDesc": "DATABASE CONNECTION ERROR",
		})

	}
}
func DeleteEmployeeDetials(ctx *gin.Context) {
	var err error
	conn, err := dboperation.Connection() //CHECK DB FOR CONNECTION
	if err == nil {
		selectstruct := structure.EmployeeSelect{} //STRUCTURE ASSIGN FOR SELECTING EMPLOYEE DETAILS
		resposestruct := structure.Respose{}       //STRUCTURE ASSIGN FOR PROVIDING RESPONSE
		fmt.Println("deleteEmployee details")
		err = ctx.ShouldBindJSON(&selectstruct) //JSON BINDING
		if err != nil {
			ctx.JSON(400, gin.H{

				"StatusDesc": "ERROR IN BINDING DATA PLEASE PROVIDE THE CORRECT REQUEST FORMAT" + err.Error(),
			})

		} else {
			validationres := validation.CheckValidationForFetchEmployee(selectstruct) //VALIDATION CHECK
			if validationres != nil {
				ctx.JSON(201, gin.H{
					"Message": "validation Error" + validationres.Error(),
				})
			} else {
				resposestruct, err = dboperation.FetchEmployee(selectstruct, conn) //CALL FOR FUNCTION QUERRY PACKAGE iNSERTBENE

				if err != nil {
					ctx.JSON(201, gin.H{
						"StatusCode": resposestruct.StatusCode,
						"StatusDesc": resposestruct.StatusDesc,
					})
				} else {
					fmt.Println("Now comming for delete")
					resposestruct, err = dboperation.DeleteEmployeeDetais(resposestruct, conn) //

					if err != nil {
						ctx.JSON(201, gin.H{
							"StatusCode": resposestruct.StatusCode,
							"StatusDesc": resposestruct.StatusDesc,
						})
					} else {
						ctx.JSON(200, gin.H{
							"StatusCode": resposestruct.StatusCode,
							"StatusDesc": resposestruct.StatusDesc,
						})
					}

				}
			}

		}

	} else {
		ctx.JSON(201, gin.H{
			"StatusCode": -1,
			"StatusDesc": "DATABASE CONNECTION ERROR",
		})

	}
}
