package main

import (
	"net/http"
	"sgcu-backend-employee/config"
	db "sgcu-backend-employee/data/database"
	employeeStore "sgcu-backend-employee/data/employees"
	_ "sgcu-backend-employee/docs"
	"sgcu-backend-employee/domain/auth"
	"sgcu-backend-employee/domain/employees"
	router "sgcu-backend-employee/router/http"
)

// @title Employees API
// @version 1.0

// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.ConnectSqlite(configuration.Database)
	if err != nil {
		panic(err)
	}

	employeeRepo := employeeStore.New(db)
	employeeSvc := employees.NewService(employeeRepo)
	authSvc := auth.NewService(employeeRepo)

	httpRouter := router.NewHTTPHandler(employeeSvc, authSvc)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
