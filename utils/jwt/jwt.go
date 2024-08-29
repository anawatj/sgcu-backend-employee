package jwt

import (
	"errors"
	"sgcu-backend-employee/domain/employees"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(req *employees.Employee) (string, error) {
	var secretKey = []byte("secret-key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"EmployeeId": req.EmployeeId,
			"Role":       req.Role,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func DecodeToken(tokenString string) (*employees.Employee, error) {
	var secretKey = []byte("secret-key")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("Invalid Token")
	}

	return &employees.Employee{
		EmployeeId: claims["EmployeeId"].(string),
		Role:       claims["Role"].(string),
	}, nil
}
