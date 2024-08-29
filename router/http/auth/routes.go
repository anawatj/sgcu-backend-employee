package auth

import (
	"errors"
	"net/http"
	"sgcu-backend-employee/domain/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service auth.AuthService
}

// Login  godoc
//
// @Summary  Login
// @Description Login to System
// @Accept json
// @Product json
// @Param   login body  AuthRequest true "Login Data"
// @Response  200   {object} AuthResponse "Ok"
// @Response  400   string                "BadRequest"
// @Response  401   string   			  "UnAuthorize"
// @Response  404   string   			  "NotFound"
// @Response  500   string   			  "InternalServerError"
// @Router /api/auth/login [post]
func (handler *AuthHandler) Login(c *gin.Context) {
	auth, err := Bind(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	claimData, code, err := handler.Service.Login(auth)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(code, *toResponseModel(claimData))
}

// ChangePassword  godoc
//
// @Summary  ChangePassword
// @Description ChangePassword
// @Accept json
// @Product json
// @Param   changePassword body  AuthRequest true "ChangePassword Data"
// @Response  200   {object} AuthResponse2 "Ok"
// @Response  400   string                "BadRequest"
// @Response  401   string   			  "UnAuthorize"
// @Response  404   string   			  "NotFound"
// @Response  500   string   			  "InternalServerError"
// @Router /api/auth/changePassword [put]
func (handler *AuthHandler) ChangePassword(c *gin.Context) {

	claim, err := Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(claim.EmployeeId) == 0 {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}

	auth, err := Bind(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if claim.Role == "Employee" {
		if claim.EmployeeId != auth.EmployeeId {
			c.Status(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, errors.New("Cannot change password anothor employee").Error())
			return
		}
	} else if claim.Role == "HR" {

		if claim.EmployeeId != auth.EmployeeId {
			isEmPloyee, err := handler.Service.IsEmployee(auth.EmployeeId)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				c.JSON(http.StatusInternalServerError, errors.New("database error").Error())
				return
			}
			if !isEmPloyee {
				c.Status(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, errors.New("Cannot change password anothor employee").Error())
				return
			}
		}

	}
	authData, code, err := handler.Service.ChangePassword(auth)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(code, *toResponseModel2(authData))
}

// GetCurrentUser godoc
// @security ApiKeyAuth
// @Summary  Get CurrentUser
// @Description GetCurrentUser
// @Accept json
// @Product json
// @Response  200   {object} EmployeeResponse "Ok"
// @Response  400   string                "BadRequest"
// @Response  401   string   			  "UnAuthorize"
// @Response  404   string   			  "NotFound"
// @Response  500   string   			  "InternalServerError"
// @Router /api/auth/currentUser [get]
func (handler *AuthHandler) GetCurrentUser(c *gin.Context) {
	claim, err := Header(c)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if len(claim.EmployeeId) == 0 {
		c.Status(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, errors.New("Unauthorize").Error())
		return
	}
	employee, code, err := handler.Service.GetCurrentUser(claim.EmployeeId)
	if err != nil {
		c.Status(code)
		c.JSON(code, err.Error())
		return
	}
	c.JSON(code, *toResponseModel3(employee))
}

func NewRoutesFactory(group *gin.RouterGroup) func(service auth.AuthService) {
	customerRoutesFactory := func(service auth.AuthService) {
		handler := AuthHandler{Service: service}
		group.POST("/login", handler.Login)
		group.PUT("/changePassword", handler.ChangePassword)
		group.GET("/currentUser", handler.GetCurrentUser)

	}

	return customerRoutesFactory
}
