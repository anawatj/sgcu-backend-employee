package employees

import (
	domain "sgcu-backend-employee/domain/employees"
)

func toResponseModel(entity *domain.Employee) *EmployeeResponse {
	return &EmployeeResponse{
		EmployeeId: entity.EmployeeId,
		Password:   entity.Password,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		Salary:     entity.Salary,
		Role:       entity.Role,
	}

}
