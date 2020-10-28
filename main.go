package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Flight struct {
	Id          int    `csv:"id"`
	Nombre      string `csv:"nombre"`
	FechaInicio string `csv:"fechaInicio"`
	FechaFin    string `csv:"fechaFin"`
}

func (f Flight) toString() string {
	return strconv.Itoa(f.Id) + "," + f.Nombre + "," + f.FechaInicio + "," + f.FechaFin
}

const fileName = "vuelos.csv"

var lastInsertId int

func getNewId() int {
	lastInsertId++
	return lastInsertId
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFlights() ([]Flight, *os.File, error) {
	csvFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	check(err)

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	check(err)

	var flight Flight
	var flights []Flight

	for i, record := range rawCSVdata {
		if i != 0 {
			flight.Id, err = strconv.Atoi(record[0])
			check(err)
			flight.Nombre = record[1]
			flight.FechaInicio = record[2]
			flight.FechaFin = record[3]
			flights = append(flights, flight)
		}
	}

	return flights, csvFile, nil
}

func createFlight(f Flight, file *os.File) {
	f.Id = getNewId()
	file.WriteString("\n" + f.toString())
	file.Sync()
}

func (f *Flight) print() {
	fmt.Printf("id=%v nombre=%v inicio=%v fin%v\n", f.Id, f.Nombre, f.FechaInicio, f.FechaFin)
}

func main() {
	flights, file, err := loadFlights()
	check(err)
	defer file.Close()
	lastInsertId = len(flights)
	flight := Flight{
		Id:          1,
		Nombre:      "nombre del vuelo",
		FechaInicio: "2023/08/09 18:00",
		FechaFin:    "2023/08/09 23:00",
	}
	createFlight(flight, file)
}
