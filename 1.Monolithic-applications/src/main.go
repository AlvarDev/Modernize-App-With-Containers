package main

import (
	"fmt"
	"log"
	pb "monolithicapp/pb"
	"net/http"
	"os"
	"text/template"
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

	port := ""
	mustMapEnv(&port, "PORT")

	if port == "" {
		port = "8080"
	}

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

	// Setting userUID because there is not Firebase Auth at this point
	data, err := ListRemainders("no-user")
	if err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("Getting firestore data: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)

	}

	if err := tmpl.Execute(w, data); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("template.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Setting userUID because there is not Firebase Auth at this point
	remainder := r.Form.Get("remainder")
	_, _ = AddRemainder(&pb.Remainder{
		Remainder: remainder,
		UserUID:   "no-user",
	})

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	remainderId := r.Form.Get("remainderId")
	err = DeleteRemainder(&pb.Remainder{
		RemainderId: remainderId,
		UserUID:     "no-user",
	})
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
