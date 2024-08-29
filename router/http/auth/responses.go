package auth

type AuthResponse struct {
	EmployeeId string `json:"employeeId"`

	Role  string `json:"role"`
	Token string `json:"token"`
}

type ListResponse struct {
	Data []AuthResponse `json:"data"`
}

type AuthResponse2 struct {
	EmployeeId string `json:"employeeId"`

	Password string `json:"password"`
}

type ListResponse2 struct {
	Data []AuthResponse2 `json:"data"`
}

type EmployeeResponse struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`

	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Salary    float64 `json:"salary"`
	Role      string  `json:"role"`
}

type ListResponse3 struct {
	Data []EmployeeResponse `json:"data"`
}
