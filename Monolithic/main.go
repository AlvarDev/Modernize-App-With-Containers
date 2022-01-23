package main

import (
	"html/template"
	"log"
	dt "monolithicdemo/data"
	"net/http"
	"os"
)

// Variables used to generate the HTML page.
var (
	tmpl *template.Template
)

func main() {
	// Prepare template for execution.
	tmpl = template.Must(template.ParseFiles("index.html"))

	// Define HTTP server.
	http.HandleFunc("/", helloRunHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

/***********************
* Handlers
***********************/

// helloRunHandler responds to requests by rendering an HTML page.
func helloRunHandler(w http.ResponseWriter, r *http.Request) {

	data, _ := dt.GetMessages()
	if err := tmpl.Execute(w, data); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("template.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	//	addData()
	helloRunHandler(w, r)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	//	deleteData()
	helloRunHandler(w, r)

}
