package auth

import (
	"errors"
	auth "sgcu-backend-employee/domain/auth"
	"sgcu-backend-employee/utils/jwt"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`
}

func Header(c *gin.Context) (*auth.ClaimData, error) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		return nil, errors.New("Invalid token")
	}
	claim, err := jwt.DecodeToken(token)
	if err != nil {
		return nil, errors.New("Invalid token")
	}
	return &auth.ClaimData{
		EmployeeId: claim.EmployeeId,
		Role:       claim.Role,
	}, nil
}

func Bind(c *gin.Context) (*auth.Auth, error) {
	var json AuthRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	auth := &auth.Auth{
		EmployeeId: json.EmployeeId,
		Password:   json.Password,
	}

	return auth, nil
}
