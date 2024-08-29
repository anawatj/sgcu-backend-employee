package employees

import (
	domain "sgcu-backend-employee/domain/employees"
)

func toDBModel(entity *domain.Employee) *Employee {
	return &Employee{
		Id:        entity.EmployeeId,
		Password:  entity.Password,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Salary:    entity.Salary,
		Role:      entity.Role,
	}
}

func toDomainModel(entity *Employee) *domain.Employee {
	return &domain.Employee{
		EmployeeId: entity.Id,
		Password:   entity.Password,
		FirstName:  entity.FirstName,
		LastName:   entity.LastName,
		Salary:     entity.Salary,
		Role:       entity.Role,
	}
}
