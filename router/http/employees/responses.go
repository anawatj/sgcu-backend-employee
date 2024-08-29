package employees

type EmployeeResponse struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`

	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Salary    float64 `json:"salary"`
	Role      string  `json:"role"`
}

type ListResponse struct {
	Data []EmployeeResponse `json:"data"`
}
