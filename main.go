package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Flight struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	FechaInicio string `json:"fechaInicio"`
	FechaFin    string `json:"fechaFin"`
}

func (f *Flight) print() {
	fmt.Printf("id=%v, nombre=%v, inicio=%v, fin=%v\n", f.Id, f.Nombre, f.FechaInicio, f.FechaFin)
}

const fileName = "flights.json"

var lastInsertId int
var s []Flight

func getNewId() int {
	lastInsertId++
	return lastInsertId
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func BinarySearch(fs []Flight, id int) int {

	start := 0
	end := len(fs) - 1

	for start <= end {
		median := (start + end) / 2

		if fs[median].Id < id {
			start = median + 1
		} else {
			end = median - 1
		}
	}

	if start == len(fs) || fs[start].Id != id {
		return -1
	} else {
		return start
	}

}

func loadFlights() error {
	flightsJson, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}
	err = json.Unmarshal(flightsJson, &s)

	return err
}

func saveFlights() {
	mar, err := json.Marshal(s)
	check(err)
	d1 := []byte(mar)
	err = ioutil.WriteFile(fileName, d1, 0644)
	check(err)
}

func createFlight(f Flight) []Flight {
	f.Id = getNewId()
	s = append(s, f)
	return s
}

func getFlightById(id int) (Flight, error) {
	i := BinarySearch(s, id)
	if i != -1 {
		return s[i], nil
	} else {
		return Flight{}, errors.New("el vuelo con el id no existe")
	}
}

func updateFlight(f Flight) error {
	i := BinarySearch(s, f.Id)
	if i != -1 {
		s[i] = f
		return nil
	} else {
		return errors.New("el vuelo con el id no existe")
	}
}

func deleteFlight(id int) bool {
	i := BinarySearch(s, id)
	if i != -1 {
		s = append(s[:i], s[i+1:]...)
		return true
	} else {
		return false
	}
}

func getKeysPressed() string {
	r := bufio.NewReader(os.Stdin)
	entry, _ := r.ReadString('\n')         // Leer hasta el separador de salto de línea
	ch := strings.TrimRight(entry, "\r\n") // Remover el salto de línea de la entrada del usuario
	return ch
}

func showCreateFlight() {
	var f Flight
	fmt.Println("Ingresa el nombre del vuelo que quieras crear")
	input := getKeysPressed()
	clear()
	f.Nombre = input
	fmt.Println("Ingresa la fecha de salida")
	input = getKeysPressed()
	clear()
	f.FechaInicio = input
	fmt.Println("Ingresa la fecha de llegada")
	input = getKeysPressed()
	clear()
	f.FechaFin = input

	createFlight(f)
	fmt.Println("El vuelo se creo.")
	enterToContinue()
}

func enterToContinue() {
	fmt.Println("")
	fmt.Println("Apreta enter para continuar.")
	getKeysPressed()
}

func showFlights() {
	for _, f := range s {
		f.print()
	}
}

func showFlight() {
	fmt.Println("Ingresa el id del vuelo que quieras ver")
	input := getIdInput()
	clear()
	if input == -1 {
		return
	}
	f, err := getFlightById(input)

	if err != nil {
		idNotFound(input)
		return
	}

	f.print()
	enterToContinue()
}

func idNotFound(id int) {
	fmt.Printf("El id %v no existe.\n", id)
	enterToContinue()
	return
}

func getIdInput() int {
	input, err := strconv.Atoi(getKeysPressed())
	clear()
	if err != nil {
		fmt.Println("Tenes que ingresar un numero.")
		enterToContinue()
		return -1
	} else {
		if input >= 0 {
			return input
		} else {
			fmt.Println("Tenes que ingresar un numero positivo.")
			enterToContinue()
			return -1
		}
	}
}

func showDeleteFlight() {
	fmt.Println("Ingresa el id del vuelo que quieras borrar.")
	input := getIdInput()
	clear()
	if input == -1 {
		return
	}

	deleted := deleteFlight(input)

	if deleted {
		fmt.Printf("Se borro el vuelo con el id=%v", input)
		enterToContinue()
	} else {
		idNotFound(input)
	}

}

func showUpdateFlight() {
	fmt.Println("Ingresa el id del vuelo que quieras modificar.")
	input := getIdInput()
	clear()
	if input == -1 {
		return
	}
	id := BinarySearch(s, input)
	if id != -1 {
		f := s[id]
		f.Id = input
		f.print()
		fmt.Println("Ingresa el nuevo nombre")
		input := getKeysPressed()
		clear()
		f.Nombre = input
		f.print()
		fmt.Println("Ingresa la nueva fecha de salida")
		input = getKeysPressed()
		clear()
		f.FechaInicio = input
		f.print()
		fmt.Println("Ingresa la nueva fecha de llegada")
		input = getKeysPressed()
		clear()
		f.FechaFin = input

		updateFlight(f)
		fmt.Println("El vuelo se modifico.")
		enterToContinue()
	} else {
		idNotFound(input)
	}
}

func showMenu() {
	menu :=
		`
1. Listar vuelos
2. Crear un vuelo
3. Ver un vuelo
4. Modificar un vuelo
5. Borrar un vuelo

Q. Guardar y salir

Ingrese una opcion:
`
	fmt.Print(menu)
}

func main() {
	loadFlights()
	if len(s) == 0 {
		lastInsertId = 0
	} else {
		lastInsertId = s[len(s)-1].Id
	}

	exit := false
	for !exit {
		clear()
		showMenu()
		key := strings.ToUpper(getKeysPressed())
		clear()

		switch key {
		case "1":
			if len(s) == 0 {
				fmt.Print("No hay vuelos")
			} else {
				showFlights()
			}
			enterToContinue()
		case "2":
			showCreateFlight()
		case "3":
			showFlight()
		case "4":
			showUpdateFlight()
		case "5":
			showDeleteFlight()
		case "Q":
			saveFlights()
			exit = true
		default:
			fmt.Println("Opcion incorrecta")
			enterToContinue()
		}
	}

}
