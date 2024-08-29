package employees

import (
	"errors"
	"fmt"
	"net/http"
	employees "sgcu-backend-employee/domain/employees"
	auth "sgcu-backend-employee/router/http/auth"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	Service employees.EmployeeService
}

// GetAllEmployee  godoc
// @security ApiKeyAuth
// @Summary  get all employee
// @Description Getting all employee from database
// @Accept   json
// @Produce  json
// @Param    firstName   query     string  false  "search by firstName"
// @Param    lastName    query     string  false  "search by lastName"
// @Param    role        query     string  false  "search by role"
// @Response  200   {object}  ListResponse "Ok"
// @Response  404   string  "NotFound"
// @Response  401   string  "UnAuthorize"
// @Response  500   string  "InternalServerError"
// @Router   /api/employees [get]
func (handler *EmployeeHandler) GetAllEmployee(c *gin.Context) {
	claim, err := auth.Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if claim.Role != "HR" {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}
	query := c.Request.URL.Query()

	firstName := query.Get("firstName")
	lastName := query.Get("lastName")
	role := query.Get("role")
	results, code, err := handler.Service.GetAllEmployee(firstName, lastName, role)
	if results == nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	var responseItems = make([]EmployeeResponse, len(results))

	for i, element := range results {
		responseItems[i] = *toResponseModel(&element)
	}

	response := &ListResponse{
		Data: responseItems,
	}

	c.JSON(code, response)
}

// CreateEmployee  godoc
// @security ApiKeyAuth
// @Summary  create employee
// @Description Adding new employee to database
// @Accept   json
// @Produce  json
// @Param   employee body  EmployeeRequest true "Employee Data"
// @Response  201   {object} ListResponse "Created"
// @Response  400   string                "BadRequest"
// @Response  401   string   			  "UnAuthorize"
// @Response  404   string   			  "NotFound"
// @Response  500   string   			  "InternalServerError"
// @Router   /api/employees [post]
func (handler *EmployeeHandler) CreateEmployee(c *gin.Context) {
	claim, err := auth.Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if claim.Role != "HR" {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}
	employee, err := Bind(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newEmployee, code, err := handler.Service.CreateEmployee(employee)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, *toResponseModel(newEmployee))
}

// GetByIdEmployee  godoc
//
// @Summary get employee by id
// @security ApiKeyAuth
// @Description getting employee by id from database
// @Accept json
// @Product json
// @Param id path string true "id"
// @Response  200   {object} ListResponse "Ok"
// @Response  400   string                "BadRequest"
// @Response  401   string                "UnAuthorize"
// @Response  404   string 				  "NotFound"
// @Response  500   string 				  "InternalServerError"
// @Router /api/employees/{id} [get]
func (handler *EmployeeHandler) GetByIdEmployee(c *gin.Context) {

	claim, err := auth.Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if claim.Role != "HR" {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}
	id := c.Param("employeeId")
	result, code, err := handler.Service.GetByIdEmployee(id)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, *toResponseModel(result))
}

// UpdateEmployee  godoc
//
// @Summary update employee
// @security ApiKeyAuth
// @Description update employee to database
// @Accept json
// @Product json
// @Param   employee body  EmployeeRequest true "Employee Data"
// @Param id path string true "id"
// @Response  200   {object} ListResponse "Ok"
// @Response  400   string                "BadRequest"
// @Response  401   string 				  "UnAuthorize"
// @Response  404   string 				  "NotFound"
// @Response  500   string 				  "InternalServerError"
// @Router /api/employees/{id} [put]
func (handler *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	claim, err := auth.Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if claim.Role != "HR" {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}
	id := c.Param("employeeId")
	employee, err := Bind(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedEmployee, code, err := handler.Service.UpdateEmployee(employee, id)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}
	fmt.Println(updatedEmployee)
	c.JSON(code, toResponseModel(updatedEmployee))
}

// DeleteEmployee  godoc
//
// @Summary delete employee
// @security ApiKeyAuth
// @Description delete employee from database
// @Accept json
// @Product json
// @Param id path string true "id"
// @Response  200   string 				  "Success"
// @Response  400   string                "BadRequest"
// @Response  401   string 				  "UnAuthorize"
// @Response  404   string 				  "NotFound"
// @Response  500   string 				  "InternalServerError"
// @Router /api/employees/{id} [delete]
func (handler *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	claim, err := auth.Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if claim.Role != "HR" {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}
	id := c.Param("employeeId")
	code, err := handler.Service.DeleteEmployee(id)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}

	c.JSON(code, "Success")
}

func NewRoutesFactory(group *gin.RouterGroup) func(service employees.EmployeeService) {
	customerRoutesFactory := func(service employees.EmployeeService) {
		handler := EmployeeHandler{Service: service}
		group.GET("/", handler.GetAllEmployee)
		group.POST("/", handler.CreateEmployee)
		group.GET("/:employeeId", handler.GetByIdEmployee)
		group.PUT("/:employeeId", handler.UpdateEmployee)
		group.DELETE("/:employeeId", handler.DeleteEmployee)
	}

	return customerRoutesFactory
}
