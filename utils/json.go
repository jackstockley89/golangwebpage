package utils

import (
	"encoding/json"
	"os"
)

type Rides struct {
	Rides []Ride `json:"rides"`
}

type Ride struct {
	ID      string  `json:"id"`
	Details Details `json:"details"`
}

type Details struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Distance string `json:"distance"`
	Time     string `json:"time"`
	AvgSpeed string `json:"avgspeed"`
	Route    string `json:"route"`
}

func UnmarshalJSON(data []byte) (*Rides, error) {
	var rides Rides
	err := json.Unmarshal(data, &rides)
	if err != nil {
		return nil, err
	}
	return &rides, nil
}

func ReadJSON(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
