package router

import (
	"net/http"
	"sgcu-backend-employee/domain/auth"
	"sgcu-backend-employee/domain/employees"
	employeeRoutes "sgcu-backend-employee/router/http/employees"

	authRoutes "sgcu-backend-employee/router/http/auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPHandler(employeeSvc employees.EmployeeService, authSvc auth.AuthService) http.Handler {
	router := gin.Default()

	api := router.Group("/api")

	employeeGroup := api.Group("/employees")
	authGroup := api.Group("/auth")
	authRoutes.NewRoutesFactory(authGroup)(authSvc)
	employeeRoutes.NewRoutesFactory(employeeGroup)(employeeSvc)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
