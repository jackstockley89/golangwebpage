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
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ENV struct values
type ENV struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// Ride table query list
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

// Active sets activites page data types
type Active struct {
	Title string
}

// IDHandler set sql query for ride pages
type IDHandler struct {
	ID int
}

// DB variable used by query function to connect to the database
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

//
// Home Page
//
func homeHandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"to Cycling Blog", time.Now().Format(time.Stamp)}
	t := template.Must(template.ParseFiles("templates/home.html"))

	if err := t.ExecuteTemplate(w, "home.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Activities Page
//
func activitiesHandler(w http.ResponseWriter, r *http.Request) {
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

//
// Ride Page
//
func (p *IDHandler) rideHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// database connect
// uses variable to load values from .env file and then pass them in connection parameters via a struct
//
func dbConnect() {
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

func main() {
	dbConnect()
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/activities", activitiesHandler)
	http.HandleFunc("/activities/rideOne", idone.rideHandler)
	http.HandleFunc("/activities/rideTwo", idtwo.rideHandler)
	http.HandleFunc("/activities/rideThree", idthree.rideHandler)
	http.HandleFunc("/activities/rideFour", idfour.rideHandler)
	http.HandleFunc("/activities/rideFive", idfive.rideHandler)
	http.HandleFunc("/activities/rideSix", idsix.rideHandler)
	http.HandleFunc("/activities/rideSeven", idseven.rideHandler)
	http.HandleFunc("/activities/ridEight", ideight.rideHandler)
	http.HandleFunc("/activities/rideNine", idnine.rideHandler)
	http.HandleFunc("/activities/rideTen", idten.rideHandler)
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
