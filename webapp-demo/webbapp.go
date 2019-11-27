package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// make struct to show ingo
type information struct {
	Name string
	Time string
}

// set up our main function
func main() {
	// set up a default value that will appear if nothing is passed in
	webInfo := information{"Person", time.Now().Format(time.Stamp)}

	//point to html and wrap it in a call
	templates := template.Must(template.ParseFiles("templates/index.html"))

	// Tell go to look in the static DIR and allow it to search for CSS
	http.Handle("/static/",
		http.StripPrefix("/static/",
			// This is formal path to DIR
			http.FileServer(http.Dir("static"))))

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// This allows URL query ?name=<NAME> to replace name
		if name := r.FormValue("name"); name != "" {
			webInfo.Name = name
		}
		// pass webInfor var and pass
		if err := templates.ExecuteTemplate(w, "index.html", webInfo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// spin up the sever for localhost:8123 and print errors if any
	fmt.Println(http.ListenAndServe(":8123", nil))
}
