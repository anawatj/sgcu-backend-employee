package auth

import (
	domain "sgcu-backend-employee/domain/auth"
	employees "sgcu-backend-employee/domain/employees"
)

func toResponseModel(entity *domain.ClaimData) *AuthResponse {
	return &AuthResponse{
		EmployeeId: entity.EmployeeId,
		Role:       entity.Role,
		Token:      entity.Token,
	}

}

func toResponseModel2(entity *domain.Auth) *AuthResponse2 {
	return &AuthResponse2{
		EmployeeId: entity.EmployeeId,
		Password:   entity.Password,
	}

}
func toResponseModel3(entity *employees.Employee) *EmployeeResponse {
	return &EmployeeResponse{
		EmployeeId: entity.EmployeeId,
		Password:   entity.Password,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		Salary:     entity.Salary,
		Role:       entity.Role,
	}

}
