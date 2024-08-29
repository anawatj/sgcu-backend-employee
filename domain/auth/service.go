package auth

import (
	"errors"
	"net/http"
	"sgcu-backend-employee/domain/employees"
	"sgcu-backend-employee/utils/hash"
	"sgcu-backend-employee/utils/jwt"
	"strings"
)

const (
	NotFound = "record not found"
)

type AuthService interface {
	Login(*Auth) (*ClaimData, int, error)
	ChangePassword(*Auth) (*Auth, int, error)
	IsEmployee(string) (bool, error)
	GetCurrentUser(string) (*employees.Employee, int, error)
}
type Service struct {
	Repository employees.EmployeeRepository
}

func (svc *Service) Login(auth *Auth) (*ClaimData, int, error) {
	user, err := svc.Repository.GetByIdEmployee(auth.EmployeeId)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Login Error")

	}

	if !hash.CheckPasswordHash(auth.Password, user.Password) {
		return nil, http.StatusBadRequest, errors.New("Login Error")

	}

	token, _ := jwt.CreateToken(user)
	return &ClaimData{
		EmployeeId: user.EmployeeId,
		Role:       user.Role,
		Token:      token,
	}, http.StatusOK, nil

}
func (svc *Service) ChangePassword(auth *Auth) (*Auth, int, error) {
	var dataErrors []string
	if len(auth.Password) == 0 {
		dataErrors = append(dataErrors, "Password is required")
	}
	if len(dataErrors) > 0 {
		return nil, http.StatusBadRequest, errors.New(strings.Join(dataErrors, ","))
	}

	employeeDb, err := svc.Repository.GetByIdEmployee(auth.EmployeeId)
	if err != nil {
		if err.Error() == NotFound {
			return nil, http.StatusNotFound, errors.New("Employee Not Found")
		} else {
			return nil, http.StatusInternalServerError, errors.New(err.Error())
		}
	}
	hashedPassword, err := hash.HashPassword(auth.Password)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("Password Field Error")
	}
	employeeDb.Password = hashedPassword
	ret, err := svc.Repository.UpdateEmployee(employeeDb, auth.EmployeeId)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New(err.Error())
	}
	return &Auth{
		EmployeeId: ret.EmployeeId,
		Password:   ret.Password,
	}, http.StatusOK, nil
}
func (svc *Service) IsEmployee(employeeId string) (bool, error) {
	employee, err := svc.Repository.GetByIdEmployee(employeeId)
	if err != nil {
		return false, errors.New("Cannot found Employee")
	}
	return employee.Role == "Employee", nil
}
func (svc *Service) GetCurrentUser(employeeId string) (*employees.Employee, int, error) {
	employee, err := svc.Repository.GetByIdEmployee(employeeId)
	if err != nil {
		if err.Error() != NotFound {
			return nil, http.StatusInternalServerError, errors.New("Database Error")
		} else {
			return nil, http.StatusNotFound, errors.New("Employee Not Found")
		}
	}
	return employee, http.StatusOK, nil
}

func NewService(repository employees.EmployeeRepository) *Service {
	return &Service{Repository: repository}
}
