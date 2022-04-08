package main

import (
	"fmt"
	"log"
	pb "monolithicapp/pb"
	"net/http"
	"text/template"
)

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

// helloRunHandler responds to requests by rendering an HTML page.
func helloRunHandler(w http.ResponseWriter, r *http.Request) {

	// Setting userUID because there is not Firebase Auth at this point
	remainders, err := ListRemainders("no-user")
	if err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("Getting firestore data: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)

	}

	if err := tmpl.Execute(w, remainders); err != nil {
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

	remainderID := r.Form.Get("remainderID")
	err = DeleteRemainder(&pb.Remainder{
		RemainderID: remainderID,
		UserUID:     "no-user",
	})
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
