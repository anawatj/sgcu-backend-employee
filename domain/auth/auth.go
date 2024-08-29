package auth

type Auth struct {
	EmployeeId string
	Password   string
}

type ClaimData struct {
	EmployeeId string
	Role       string
	Token      string
}
