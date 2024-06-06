package router

import (
	"github.com/gin-gonic/gin"
	api "server.go/Api"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/HbtChk", check)
	router.POST("/CreateEmployee", api.CreateEmployee)
	router.POST("/GetEmployeeByID", api.SelectEmployee)
	router.POST("/UpdateEmployee", api.UpdateEmployee)
	router.POST("/DeleteEmployee", api.DeleteEmployeeDetials)

	router.Run(":8080")
	return router
}
func check(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Message": "Employee onboard Application running",
	})

}
