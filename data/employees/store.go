package employees

import (
	domain "sgcu-backend-employee/domain/employees"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	NotFound = "record not found"
)

type Store struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&Employee{})

	return &Store{
		DB: db,
	}
}

func (s *Store) CreateEmployee(employee *domain.Employee) (*domain.Employee, error) {
	entity := toDBModel(employee)
	err := s.DB.Create(entity).Error
	return toDomainModel(entity), err
}
func (s *Store) GetAllEmployee(firstName string, lastName string, role string) ([]domain.Employee, error) {
	var query []string
	var args []interface{}
	if len(role) > 0 {
		query = append(query, "role=?")
		args = append(args, role)
	}
	if len(firstName) > 0 {
		firstNameWherePercent := strings.Replace(firstName, "*", "%", -1)
		firstNameWhere := strings.Replace(firstNameWherePercent, "?", "_", -1)
		query = append(query, "first_name like ?")
		args = append(args, firstNameWhere)
	}
	if len(lastName) > 0 {
		lastNameWherePercent := strings.Replace(lastName, "*", "%", -1)
		lastNameWhere := strings.Replace(lastNameWherePercent, "?", "_", -1)
		query = append(query, "last_name like ?")
		args = append(args, lastNameWhere)
	}

	var results []Employee
	err := s.DB.Where(strings.Join(query, " AND "), args...).Find(&results).Error
	var ret []domain.Employee
	if err == nil {
		var employees = make([]domain.Employee, len(results))
		for i, element := range results {
			employees[i] = *toDomainModel(&element)
		}
		ret = employees
	}
	return ret, err
}
func (s *Store) GetByIdEmployee(id string) (*domain.Employee, error) {
	result := &Employee{}
	err := s.DB.Where("id = ?", id).First(result).Error
	return toDomainModel(result), err
}
func (s *Store) UpdateEmployee(employee *domain.Employee, id string) (*domain.Employee, error) {
	entity := toDBModel(employee)
	err := s.DB.Save(entity).Error

	return toDomainModel(entity), err
}
func (s *Store) DeleteEmployee(id string) error {
	err := s.DB.Delete(&Employee{Id: id}).Error
	return err
}
