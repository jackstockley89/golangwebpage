package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

type Active struct {
	Title string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"to Cycling Blog", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/home.html"))

	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}

	if err := templates.ExecuteTemplate(w, "home.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func activitiesHandler(w http.ResponseWriter, r *http.Request) {
	p := Active{Title: "Activities"}
	t := template.Must(template.ParseFiles("templates/activities.html"))
	//t.Execute(w, p)

	if err := t.ExecuteTemplate(w, "activities.html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/activities", activitiesHandler)
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
