package lib

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	utils "github.com/jackstockley89/golangwebpage/utils"
)

func ActivitiesHandler(w http.ResponseWriter, r *http.Request, wantJson bool) {
	t := template.Must(template.ParseFiles("templates/activities.html"))

	file, err := utils.ReadJSON("static/json/rides.json")
	if err != nil {
		http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
		return
	}
	ride, err := utils.UnmarshalJSON(file)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
		return
	}
	if wantJson {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ride); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		return
	}

	data := struct {
		Activities interface{}
	}{
		Activities: ride.Rides,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		return
	}
}

// ...existing code...

func ActivityDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the activity ID from the URL path
	path := r.URL.Path
	activityID := strings.TrimPrefix(path, "/activities/")

	if activityID == "" {
		http.NotFound(w, r)
		return
	}

	// Load the rides data
	file, err := utils.ReadJSON("static/json/rides.json")
	if err != nil {
		http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
		return
	}

	rideData, err := utils.UnmarshalJSON(file)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
		return
	}

	// Find the specific ride
	var selectedRide *utils.Ride
	for _, ride := range rideData.Rides {
		if ride.ID == activityID {
			selectedRide = &ride
			break
		}
	}

	if selectedRide == nil {
		http.NotFound(w, r)
		return
	}

	// Parse and execute template for individual activity
	t := template.Must(template.ParseFiles("templates/activity-detail.html"))

	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, selectedRide); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
