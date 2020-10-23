package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

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

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("activitiesHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "activities.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride One Page
//
func ridesOneHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 1)

	if err != nil {
		log.Panic("ridesOneHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("ridesOneHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Two Page
//
func ridesTwoHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 2)

	if err != nil {
		log.Panic("ridesTwoHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("ridesTwoHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Three Page
//
func rideThreeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 3)

	if err != nil {
		log.Panic("rideThreeHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideThreeHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Four Page
//
func rideFourHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 4)

	if err != nil {
		log.Panic("rideFourHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideFourHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Five Page
//
func rideFiveHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 5)

	if err != nil {
		log.Panic("rideFiveHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideFiveHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Six Page
//
func rideSixHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 6)

	if err != nil {
		log.Panic("rideSixHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideSixHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Seven Page
//
func rideSevenHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 7)

	if err != nil {
		log.Panic("rideSevenHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideSevenHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Eight Page
//
func rideEightHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 8)

	if err != nil {
		log.Panic("rideEightHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideEightHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Nine Page
//
func rideNineHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 9)

	if err != nil {
		log.Panic("rideNineHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideNineHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// Ride Ten Page
//
func rideTenHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/rides.html"))
	sql := `SELECT * FROM rides_table WHERE id=$1`

	row1, err := DB.Query(sql, 10)

	if err != nil {
		log.Panic("rideTenHandler: Query Error", err)
	}

	defer row1.Close()

	var id int
	var link string
	var name string
	var date string
	var distance string
	var time string
	var avgSpeed string
	var route string

	var ps []Ride
	for row1.Next() {
		err := row1.Scan(&id, &link, &name, &date, &distance, &time, &avgSpeed, &route)
		if err != nil {
			log.Panic("rideTenHandler: Scan Error", err)
		}
		ps = append(ps, Ride{ID: id, Link: link, Name: name, Date: date, Distance: distance, Time: time, AvgSpeed: avgSpeed, Route: route})
	}

	if err := t.ExecuteTemplate(w, "rides.html", ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// uses const above to add in the connection values for the datatbase
// and then open a connection to the database
func dbConnect() {
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
	http.HandleFunc("/activities/rideOne", ridesOneHandler)
	http.HandleFunc("/activities/rideTwo", ridesTwoHandler)
	http.HandleFunc("/activities/rideThree", rideThreeHandler)
	http.HandleFunc("/activities/rideFour", rideFourHandler)
	http.HandleFunc("/activities/rideFive", rideFiveHandler)
	http.HandleFunc("/activities/rideSix", rideSixHandler)
	http.HandleFunc("/activities/rideSeven", rideSevenHandler)
	http.HandleFunc("/activities/ridEight", rideEightHandler)
	http.HandleFunc("/activities/rideNine", rideNineHandler)
	http.HandleFunc("/activities/rideTen", rideTenHandler)
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
