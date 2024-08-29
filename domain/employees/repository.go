package employees

type EmployeeRepository interface {
	CreateEmployee(*Employee) (*Employee, error)
	GetAllEmployee(string, string, string) ([]Employee, error)
	GetByIdEmployee(string) (*Employee, error)
	UpdateEmployee(*Employee, string) (*Employee, error)
	DeleteEmployee(string) error
}
