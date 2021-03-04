package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ENV struct values used for Database Connection Handler function
type ENV struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// Ride table query list used by Activites and Ride Handler functions
type Ride struct {
	ID       int
	Link     string
	Name     string
	Date     string
	Distance string
	Time     string
	AvgSpeed string
	Route    string
}

// Welcome sets home page data types
type Welcome struct {
	Name string
	Time string
}

// IDHandler set sql query for RideHandler function
type IDHandler struct {
	ID int
}

// DB variable set in the main function to set the parameters used by query in RideHandler function to return correct data
var (
	DB      *sql.DB
	idone   = &IDHandler{ID: 1}
	idtwo   = &IDHandler{ID: 2}
	idthree = &IDHandler{ID: 3}
	idfour  = &IDHandler{ID: 4}
	idfive  = &IDHandler{ID: 5}
	idsix   = &IDHandler{ID: 6}
	idseven = &IDHandler{ID: 7}
	ideight = &IDHandler{ID: 8}
	idnine  = &IDHandler{ID: 9}
	idten   = &IDHandler{ID: 10}
)

// HomeHandler Page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"to Cycling Blog", time.Now().Format(time.Stamp)}
	t := template.Must(template.ParseFiles("templates/home.html"))

	if err := t.ExecuteTemplate(w, "home.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ActivitiesHandler Page
func ActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/activities.html"))
	sql := `SELECT * FROM rides_table`

	row1, err := DB.Query(sql)

	if err != nil {
		log.Panic("activitiesHandler: Query Error", err)
	}

	defer row1.Close()

	var idRef int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&idRef, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("activitiesHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: idRef, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "activities.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RideHandler Page
func (p *IDHandler) RideHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := "SELECT * FROM rides_table WHERE id=$1 "
	id := fmt.Sprintf("%d", p.ID)
	fmt.Println(id)
	row1, err := DB.Query(sql, id)

	if err != nil {
		log.Panic("rideHandler: Query Error", err)
	}

	defer row1.Close()

	var idRef int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&idRef, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: idRef, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	fmt.Println(ps)

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DbConnect uses variable to load values from .env file and then pass them in connection parameters via a struct
func DbConnect() {
	err := godotenv.Load(".env_app")
	if err != nil {
		log.Fatal("Error loading .envdb file")
	}
	postgresHost := os.Getenv("POSTGRES_HOSTNAME")
	postgresPort := 5432
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresName := os.Getenv("POSTGRES_DB")

	env := ENV{Host: postgresHost, Port: postgresPort, User: postgresUser, Password: postgresPassword, Dbname: postgresName}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.Host, env.Port, env.User, env.Password, env.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic("DbConnect: unable to connect to database", err)
	}
	DB = db
	fmt.Println("Successfully connected!")
}

// Main
func main() {
	DbConnect()
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/activities", ActivitiesHandler)
	http.HandleFunc("/activities/rideOne", idone.RideHandler)
	http.HandleFunc("/activities/rideTwo", idtwo.RideHandler)
	http.HandleFunc("/activities/rideThree", idthree.RideHandler)
	http.HandleFunc("/activities/rideFour", idfour.RideHandler)
	http.HandleFunc("/activities/rideFive", idfive.RideHandler)
	http.HandleFunc("/activities/rideSix", idsix.RideHandler)
	http.HandleFunc("/activities/rideSeven", idseven.RideHandler)
	http.HandleFunc("/activities/rideEight", ideight.RideHandler)
	http.HandleFunc("/activities/rideNine", idnine.RideHandler)
	http.HandleFunc("/activities/rideTen", idten.RideHandler)
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
