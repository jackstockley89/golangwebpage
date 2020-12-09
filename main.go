package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "cyclingblog-db"
	port     = 5432
	user     = "demouser"
	password = "P455w0rd"
	dbname   = "cyclingblog"
)

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

// DB variable used by query function to connect to the database
var (
	DB *sql.DB
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
func rideHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	//id := []string{"1", "2"}
	//sql := "SELECT * FROM rides_table WHERE id=$1 "

	//row1, err := DB.Query(sql, pq.Array([]int{1, 2}))
	row1, err := DB.Query(`SELECT * FROM rides_table WHERE id = ANY($1)`, pq.Array([]int{1, 2}))

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

	fmt.Printf(sql)
}

//
// Load env files
//

func getEnvVars() {
	err := godotenv.Load("variables.env")
	if err != nil {
		log.Fatal("Error loading .envdb file")
	}
}

// uses const above to add in the connection values for the datatbase
// and then open a connection to the database
func dbConnect() {
	//getEnvVars()
	//postgresHost := os.Getenv("HOSTNAME")
	//postgresPort := 5432
	//postgresUser := os.Getenv("POSTGRES_USER")
	//postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	//postgresName := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
	http.HandleFunc("/activities/rideOne", rideHandler)
	http.HandleFunc("/activities/rideTwo", rideHandler)
	http.HandleFunc("/activities/rideThree", rideHandler)
	http.HandleFunc("/activities/rideFour", rideHandler)
	http.HandleFunc("/activities/rideFive", rideHandler)
	http.HandleFunc("/activities/rideSix", rideHandler)
	http.HandleFunc("/activities/rideSeven", rideHandler)
	http.HandleFunc("/activities/ridEight", rideHandler)
	http.HandleFunc("/activities/rideNine", rideHandler)
	http.HandleFunc("/activities/rideTen", rideHandler)
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
