package employees

type Employee struct {
	Id        string `gorm:"primary_key";`
	Password  string
	FirstName string
	LastName  string
	Salary    float64
	Role      string
}
