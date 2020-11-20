package agencyService

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/tomasarras/GoLang/entity"
)

// Service of agency
type ServiceFlight interface {
	Save(entity.Agency) (entity.Agency, error)
	FindByID(int) (entity.Agency, error)
	FindAll() []entity.Agency
	Remove(int) error
	Update(entity.Agency) (entity.Agency, error)
}

type service struct {
	db *sql.DB
	//conf *config.Config
}

// New instance of service
func New(db *sql.DB) (ServiceFlight, error) {
	return service{db}, nil
}

func (s service) Save(a entity.Agency) (entity.Agency, error) {
	query := "INSERT INTO agency(name) VALUES (?)"
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return entity.Agency{}, err
	}

	row, err := prepare.Exec(a.Name)

	if err != nil {
		return entity.Agency{}, err
	}

	a.ID, err = row.LastInsertId()

	return a, err
}

func (s service) FindByID(ID int) (entity.Agency, error) {
	rows, err := s.db.Query("SELECT * FROM agency WHERE id_agency = ?", ID)
	fmt.Println(rows)
	if err != nil {
		return entity.Agency{}, err
	}

	rows.Next()

	var agency entity.Agency
	var id int64
	var name string
	err2 := rows.Scan(&id, &name)

	if err2 != nil {
		return entity.Agency{}, err
	} else {
		agency = entity.Agency{id, name}
	}

	return agency, nil
}

func (s service) FindAll() []entity.Agency {
	rows, err := s.db.Query("SELECT * FROM agency")
	if err != nil {
		return nil
	} else {
		agencies := []entity.Agency{}
		for rows.Next() {
			var id int64
			var name string
			err2 := rows.Scan(&id, &name)

			if err2 != nil {
				return nil
			} else {
				agency := entity.Agency{id, name}
				agencies = append(agencies, agency)
			}
		}
		return agencies
	}
}

func (s service) Remove(ID int) error {
	result, err := s.db.Exec("DELETE FROM agency WHERE id_agency = ?", ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows > 0 {
		return nil
	} else {
		return errors.New("agency with id=" + strconv.Itoa(ID) + " was not eliminated")
	}
}

func (s service) Update(a entity.Agency) (entity.Agency, error) {
	result, err := s.db.Exec("UPDATE agency SET name = ? WHERE id_agency = ?", a.Name, a.ID)
	//result, err := productModel.Db.Exec("UPDATE agency SET name = ?, price = ?, quantity = ?, status = ? where id = ?", product.Name, product.Price, product.Quantity, product.Status, product.Id)
	if err != nil {
		return entity.Agency{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return entity.Agency{}, err
	}

	if rows > 0 {
		return a, nil
	} else {
		return entity.Agency{}, err
	}
}
