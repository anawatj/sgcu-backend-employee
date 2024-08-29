package employees

import (
	employees "sgcu-backend-employee/domain/employees"

	"github.com/gin-gonic/gin"
)

type EmployeeRequest struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`

	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Salary    float64 `json:"salary"`
	Role      string  `json:"role"`
}

func Bind(c *gin.Context) (*employees.Employee, error) {
	var json EmployeeRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	employee := &employees.Employee{
		EmployeeId: json.EmployeeId,
		Password:   json.Password,
		FirstName:  json.FirstName,
		LastName:   json.LastName,
		Salary:     json.Salary,
		Role:       json.Role,
	}

	return employee, nil
}
