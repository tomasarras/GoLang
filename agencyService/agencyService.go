package agencyService

import "database/sql"

// Agency entity
type Agency struct {
	Name string `json: "name"`
}

// Service of agency
type ServiceAgency interface {
	Save(Agency) error
	FindByID(int) *Agency
	FindAll() []*Agency
	Delete(int)
}

type service struct {
	db *sql.DB
	//conf *config.Config
}

// New instance of service
func New(db *sql.DB) (ServiceAgency, error) {
	return service{db}, nil
}

func (s service) Save(a Agency) error {
	return nil
}

func (s service) FindByID(ID int) *Agency {
	return nil
}

func (s service) FindAll() []*Agency {
	return nil
}

func (s service) Delete(int) {

}
