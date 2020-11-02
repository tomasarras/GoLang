package main

import (
	"testing"
)

var fileTest string = "flights_test.json"

var ft []Flight = []Flight{
	Flight{
		Id:          1,
		Nombre:      "flight1",
		FechaInicio: "2020/02/11 13:38",
		FechaFin:    "2020/02/11 19:00",
	},
	Flight{
		Id:          2,
		Nombre:      "flight2",
		FechaInicio: "2020/02/11 13:38",
		FechaFin:    "2020/02/11 19:00",
	},
	Flight{
		Id:          3,
		Nombre:      "flight3",
		FechaInicio: "2020/02/11 13:38",
		FechaFin:    "2020/02/11 19:00",
	},
	Flight{
		Id:          4,
		Nombre:      "flight4",
		FechaInicio: "2020/02/11 13:38",
		FechaFin:    "2020/02/11 19:00",
	},
	Flight{
		Id:          5,
		Nombre:      "flight5",
		FechaInicio: "2020/02/11 13:38",
		FechaFin:    "2020/02/11 19:00",
	},
}

func TestBinarySearch(t *testing.T) {
	var index int
	for i, v := range ft {
		index = BinarySearch(ft, i+1)
		if v.Id != index+1 {
			t.Errorf("BinarySerch(ft,%v) failed, expected %v, got %v", i+1, i, index)
		} else {
			t.Logf("BinarySerch(ft,%v) success, expected %v, got %v\n", i+1, i, index)
		}
	}
}

func TestBinarySearchNotFoundId(t *testing.T) {
	i := BinarySearch(ft, 256)
	if i != -1 {
		t.Errorf("BinarySerch(ft,256) failed, expected -1, got %v", i)
	} else {
		t.Logf("BinarySerch(ft,256) success, expected -1, got %v", i)
	}
}

func TestLoadFlights(t *testing.T) {
	fileName = fileTest
	err := loadFlights()

	if err != nil {
		t.Errorf("loadFlights() failed, got error %v", err.Error())
	} else {
		t.Logf("loadFlights() flights loaded")
	}

	if len(s) == 0 {
		t.Errorf("loadFlights() failed, s []Flight is empty")
	} else {
		t.Logf("loadFlights() s []Flight is not empty")
	}

	for i, v := range s {
		if v.Id != i+1 {
			t.Errorf("loadFlights() failed, expected %v, got %v", i+1, v.Id)
		} else {
			t.Logf("loadFlights() success, expected %v, got %v", i+1, v.Id)
		}
	}
}

func TestGetNewId(t *testing.T) {
	lastInsertId = 5
	id := getNewId()
	if id != 6 {
		t.Errorf("getNewId() failed, expected 6, got %v", id)
	} else {
		t.Logf("getNewId() success, expected 6, got %v", id)
	}
}

func TestCreateFlight(t *testing.T) {
	s = ft
	prevLen := len(s)
	f := Flight{
		Nombre:      "flight",
		FechaInicio: "0000/00/00 00:00",
		FechaFin:    "0000/00/00 00:00",
	}
	createFlight(f)

	if prevLen+1 != len(s) {
		t.Errorf("creteFlight(flight) failed, the flight was not added, length of s []Flight before trying add flight %v, length of s after trying add flight %v", prevLen, len(s))
	} else {
		t.Logf("createFlight(flight) the flight was added into s []Flight")
	}

	fSaved := s[len(s)-1]

	if f.Id != 0 || f.Nombre != fSaved.Nombre ||
		f.FechaInicio != fSaved.FechaInicio || f.FechaFin != fSaved.FechaFin {
		t.Errorf("creteFlight(flight) failed, the flight was added but the data not match")
	} else {
		t.Logf("creteFlight(flight) success, the flight was added and the data match")
	}

	s = s[:len(s)-1]
}

func TestDeleteFlight(t *testing.T) {
	s = ft
	lastFlight := s[len(s)-1]
	prevLen := len(s)

	d := deleteFlight(lastFlight.Id)

	if !d {
		t.Errorf("deleteFlight(%v) failed, expected true, got %v", lastFlight.Id, d)
	} else {
		t.Logf("deleteFlight(%v) flight with id %v was deleted", lastFlight.Id, lastFlight.Id)
	}

	if len(s) != (prevLen - 1) {
		t.Errorf("deleteFlight(%v) failed, length before trying delete %v, length after trying delete %v", lastFlight.Id, prevLen, len(s))
	} else {
		t.Logf("deleteFlight(%v) success, length before trying delete %v, length after trying delete %v", lastFlight.Id, prevLen, len(s))
	}

	s = append(s, lastFlight)
}

func TestDeleteFlightIdNotFound(t *testing.T) {
	s = ft

	d := deleteFlight(512)

	if d {
		t.Errorf("deleteFlight(512) failed, expected false, got %v", d)
	} else {
		t.Logf("deleteFlight(512) sucess, expected false, got %v", d)
	}
}

func TestUpdateFlight(t *testing.T) {
	s = ft
	lastFlight := s[len(s)-1]
	var originalFlight Flight
	originalFlight.Id = lastFlight.Id
	originalFlight.Nombre = lastFlight.Nombre
	originalFlight.FechaInicio = lastFlight.FechaInicio
	originalFlight.FechaFin = lastFlight.FechaFin

	var updatedFlight Flight
	updatedFlight.Id = lastFlight.Id
	updatedFlight.Nombre = "updated"
	updatedFlight.FechaInicio = "1234/12/12 00:00"
	updatedFlight.FechaFin = "1234/12/12 00:00"

	updateFlight(updatedFlight)

	last := s[len(s)-1]
	if last.Id != updatedFlight.Id || last.Nombre != updatedFlight.Nombre ||
		last.FechaInicio != updatedFlight.FechaInicio || last.FechaFin != updatedFlight.FechaFin {
		t.Error("updateFlight(flight) failed, the data not match")
	} else {
		t.Log("updateFlight(flight) success, the data match")
	}

	s = s[:len(s)-1]
	s = append(s, originalFlight)
}

func TestUpdateFlightIdNotFound(t *testing.T) {
	s = ft
	eMessage := "el vuelo con el id no existe"
	updatedFlight := Flight{
		Id: 80,
	}

	err := updateFlight(updatedFlight)
	if err != nil {

		if err.Error() != eMessage {
			t.Errorf("updateFlight(flight) failed, expected %v, got %v", eMessage, err.Error())
		} else {
			t.Log("updateFlight(flight) success, the id not found")
		}
	} else {
		t.Error("updateFlight(flight) failed, expected error, got nil")
	}
}

func TestGetFlightById(t *testing.T) {
	s = ft
	f, err := getFlightById(1)

	if err == nil {
		if f.Id != 1 {
			t.Errorf("getFligthById(1) failed, expected id %v, got id %v", 1, f.Id)
		} else {
			t.Logf("getFligthById(1) success, expected id %v, got id %v", 1, f.Id)
		}
	} else {
		t.Errorf("getFligthById(1) failed, expected flight, got error %v", err.Error())
	}
}

func TestGetFlightByIdIdNotFound(t *testing.T) {
	eMessage := "el vuelo con el id no existe"
	s = ft
	f, err := getFlightById(90)

	if err != nil {
		if err.Error() != eMessage {
			t.Errorf("getFligthById(90) failed, expected error %v, got error %v", eMessage, err.Error())
		} else {
			t.Logf("getFligthById(90) success, expected error %v, got error %v", eMessage, err.Error())
		}
	} else {
		t.Errorf("getFligthById(90) failed, expected error %v, got flight %v", err.Error(), f)
	}
}

func TestSaveFlights(t *testing.T) {
	fileName = fileTest
	s = []Flight{
		Flight{
			Id:          20,
			Nombre:      "test-saveFlights",
			FechaInicio: "0000/00/00 00:00",
			FechaFin:    "0000/00/00 00:00",
		},
	}
	t.Logf("saveFlights() trying to save flights")
	saveFlights()
	t.Logf("saveFlights() the flights were saved")
	s = make([]Flight, 0)
	t.Logf("saveFlights() trying open saved flights")
	err := loadFlights()

	if err != nil {
		t.Errorf("saveFlights() failed, flights saved but not accessible error %v", err.Error())
	} else {
		t.Logf("saveFlights() the flights were loaded")
	}

	if len(s) != 1 {
		t.Errorf("saveFlights() failed, expected lenght of s 1, got %v", len(s))
	} else {
		t.Logf("saveFlights() success, 1 flight was saved")
	}

	s = ft
	saveFlights()
}
