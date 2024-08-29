package employees

import (
	"errors"
	"net/http"
	"sgcu-backend-employee/utils/hash"
	"strings"
)

const (
	NotFound = "record not found"
)

type EmployeeService interface {
	CreateEmployee(*Employee) (*Employee, int, error)
	GetAllEmployee(string, string, string) ([]Employee, int, error)
	GetByIdEmployee(string) (*Employee, int, error)
	UpdateEmployee(*Employee, string) (*Employee, int, error)
	DeleteEmployee(string) (int, error)
}
type Service struct {
	Repository EmployeeRepository
}

func (svc *Service) CreateEmployee(employee *Employee) (*Employee, int, error) {
	var dataErrors []string
	if len(employee.EmployeeId) == 0 {
		dataErrors = append(dataErrors, "Employee Id is required")
	}

	if len(employee.Password) == 0 {
		dataErrors = append(dataErrors, "Password is required")
	}

	if len(employee.FirstName) == 0 {
		dataErrors = append(dataErrors, "FirstName is required")
	}

	if len(employee.LastName) == 0 {
		dataErrors = append(dataErrors, "LastName is required")
	}

	if len(employee.Role) == 0 {
		dataErrors = append(dataErrors, "Role is required")
	}
	if employee.Salary < 0 {
		dataErrors = append(dataErrors, "Salary is required")
	}
	if len(dataErrors) > 0 {
		return nil, http.StatusBadRequest, errors.New(strings.Join(dataErrors, ","))
	}
	employeeDb, err := svc.Repository.GetByIdEmployee(employee.EmployeeId)
	if err != nil {

		if err.Error() != NotFound {

			return nil, http.StatusInternalServerError, errors.New(err.Error())
		}
	}

	if len(employeeDb.EmployeeId) > 0 {
		return nil, http.StatusBadRequest, errors.New("Record duplicate")
	}

	hashedPassword, err := hash.HashPassword(employee.Password)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Password Field Error")
	}
	employee.Password = hashedPassword

	ret, err := svc.Repository.CreateEmployee(employee)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(err.Error())
	}
	return ret, http.StatusCreated, nil
}
func (svc *Service) GetAllEmployee(firstName string, lastName string, role string) ([]Employee, int, error) {
	ret, err := svc.Repository.GetAllEmployee(firstName, lastName, role)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(err.Error())
	}
	if len(ret) == 0 {
		return nil, http.StatusNotFound, errors.New("Employee Not Found")
	}
	return ret, http.StatusOK, nil
}
func (svc *Service) GetByIdEmployee(id string) (*Employee, int, error) {
	ret, err := svc.Repository.GetByIdEmployee(id)
	if err != nil {
		if err.Error() == NotFound {
			return nil, http.StatusNotFound, errors.New("Employee Not Found")
		} else {
			return nil, http.StatusInternalServerError, errors.New(err.Error())
		}
	}

	return ret, http.StatusOK, nil
}
func (svc *Service) UpdateEmployee(employee *Employee, id string) (*Employee, int, error) {
	var dataErrors []string
	if len(employee.EmployeeId) == 0 {
		dataErrors = append(dataErrors, "Employee Id is required")
	}

	if len(employee.Password) == 0 {
		dataErrors = append(dataErrors, "Password is required")
	}

	if len(employee.FirstName) == 0 {
		dataErrors = append(dataErrors, "FirstName is required")
	}

	if len(employee.LastName) == 0 {
		dataErrors = append(dataErrors, "LastName is required")
	}

	if len(employee.Role) == 0 {
		dataErrors = append(dataErrors, "Role is required")
	}
	if employee.Salary < 0 {
		dataErrors = append(dataErrors, "Salary is required")
	}
	if len(dataErrors) > 0 {
		return nil, http.StatusBadRequest, errors.New(strings.Join(dataErrors, ","))
	}

	employeeDb, err := svc.Repository.GetByIdEmployee(id)
	if err != nil {
		if err.Error() == NotFound {
			return nil, http.StatusNotFound, errors.New("Employee Not Found")
		} else {
			return nil, http.StatusInternalServerError, errors.New(err.Error())
		}
	}

	employeeDb.FirstName = employee.FirstName
	employeeDb.LastName = employee.LastName
	employeeDb.Salary = employee.Salary
	employeeDb.Role = employee.Role
	ret, err := svc.Repository.UpdateEmployee(employeeDb, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return ret, http.StatusOK, nil
}

func (svc *Service) DeleteEmployee(id string) (int, error) {

	_, err := svc.Repository.GetByIdEmployee(id)
	if err != nil {
		if err.Error() == NotFound {
			return http.StatusNotFound, errors.New("Employee Not Found")
		} else {
			return http.StatusInternalServerError, errors.New(err.Error())
		}

	}

	err = svc.Repository.DeleteEmployee(id)
	if err != nil {
		return http.StatusInternalServerError, errors.New(err.Error())
	}

	return http.StatusOK, nil
}
func NewService(repository EmployeeRepository) *Service {
	return &Service{Repository: repository}
}
