package flightService

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/tomasarras/GoLang/entity"
)

// Service of flight
type ServiceFlight interface {
	Save(entity.Flight) (entity.Flight, error)
	FindByID(int) (entity.Flight, error)
	FindAll() []entity.Flight
	FindAllByAgency(int) []entity.Flight
	Remove(int) error
	Update(entity.Flight) (entity.Flight, error)
}

type service struct {
	db *sql.DB
	//conf *config.Config
}

// New instance of service
func New(db *sql.DB) (ServiceFlight, error) {
	return service{db}, nil
}

func (s service) Save(f entity.Flight) (entity.Flight, error) {
	query := "INSERT INTO flight(name,start,end,aircraft,id_agency) VALUES (?,?,?,?,?)"
	prepare, err := s.db.Prepare(query)
	if err != nil {
		panic(err.Error())
		//return entity.Flight{}, err
	}

	row, err := prepare.Exec(f.Name, f.Start, f.End, f.Aircraft, f.IdAgency)

	if err != nil {
		panic(err.Error())
		//return entity.Flight{}, err
	}

	f.ID, err = row.LastInsertId()

	return f, err
}

func (s service) FindByID(ID int) (entity.Flight, error) {
	rows, err := s.db.Query("SELECT * FROM flight WHERE id_flight = ?", ID)
	if err != nil {
		return entity.Flight{}, err
	}

	rows.Next()

	var flight entity.Flight
	var id int64
	var name string
	var start string
	var end string
	var aircraft string
	var idAgency int64
	err2 := rows.Scan(&id, &name, &start, &end, &aircraft, &idAgency)

	if err2 != nil {
		return flight, err
	} else {
		flight = entity.Flight{ID: id, Name: name, Start: start, End: end, Aircraft: aircraft, IdAgency: idAgency}
		//agency = entity.Flight{id, name}
	}

	return flight, nil
}

func (s service) FindAll() []entity.Flight {
	rows, err := s.db.Query("SELECT * FROM flight")
	if err != nil {
		return nil
	} else {
		flights := []entity.Flight{}
		for rows.Next() {
			var id int64
			var name string
			var start string
			var end string
			var aircraft string
			var idAgency int64
			err2 := rows.Scan(&id, &name, &start, &end, &aircraft, &idAgency)

			if err2 != nil {
				return nil
			} else {
				flight := entity.Flight{ID: id, Name: name, Start: start, End: end, Aircraft: aircraft, IdAgency: idAgency}
				flights = append(flights, flight)
			}
		}
		return flights
	}
}

func (s service) FindAllByAgency(idAgency int) []entity.Flight {
	rows, err := s.db.Query("SELECT flight.* FROM flight JOIN agency ON(flight.id_agency = agency.id_agency) WHERE agency.id_agency = ?", idAgency)
	if err != nil {
		return nil
	} else {
		flights := []entity.Flight{}
		for rows.Next() {
			var id int64
			var name string
			var start string
			var end string
			var aircraft string
			var idAgency int64
			err2 := rows.Scan(&id, &name, &start, &end, &aircraft, &idAgency)

			if err2 != nil {
				panic(err2.Error())
			} else {
				flight := entity.Flight{ID: id, Name: name, Start: start, End: end, Aircraft: aircraft, IdAgency: idAgency}
				flights = append(flights, flight)
			}
		}
		return flights
	}
}

func (s service) Remove(ID int) error {
	result, err := s.db.Exec("DELETE FROM flight WHERE id_flight = ?", ID)
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
		return errors.New("flight with id=" + strconv.Itoa(ID) + " was not eliminated")
	}
}

func (s service) Update(f entity.Flight) (entity.Flight, error) {
	result, err := s.db.Exec("UPDATE flight SET name = ?, start = ?, end = ?, aircraft = ?, id_agency = ? WHERE id_flight = ?", f.Name, f.Start, f.End, f.Aircraft, f.IdAgency, f.ID)
	if err != nil {
		return entity.Flight{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return entity.Flight{}, err
	}

	if rows > 0 {
		return f, nil
	} else {
		return entity.Flight{}, err
	}
}
