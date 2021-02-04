package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type TestRide struct {
	ID       int
	Link     string
	Name     string
	Date     string
	Distance string
	Time     string
	AvgSpeed string
	Route    string
}

// RideHandler Page
func TestRideHandler(t *testing.T) {

	var idRef int = 99
	var link string = "Test Link"
	var name string = "Test Name"
	var date string = "01/01/2021"
	var distance string = "100"
	var time string = "05:00:00"
	var avgSpeed string = "15"
	var route string = "Test Route"

	ps := TestRide{ID: idRef, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route}
	result := TestRide{ID: 99, Link: "Test Link", Name: "Test Name", Date: "01/01/2021", Distance: "100", Time: "05:00:00", AvgSpeed: "15", Route: "Test Route"}

	psJSON, err := json.Marshal(ps)
	if err != nil {
		log.Fatalf(err.Error())
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		log.Fatalf(err.Error())
	}

	output := fmt.Sprintf("%s", string(psJSON))
	expected := fmt.Sprintf("%s", string(resultJSON))

	if output != expected {
		t.Errorf("\nIncorrect result recieved: %s\nExpecting: %s\n", output, expected)
	}
}
