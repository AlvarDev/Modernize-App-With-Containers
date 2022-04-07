package main

import (
	"fmt"
	pb "frontendservice/pb"
	"log"
	"net/http"
	"text/template"
)

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

// helloRunHandler responds to requests by rendering an HTML page.
func (fe *frontendServer) rootHandler(w http.ResponseWriter, r *http.Request) {

	// Setting userUID because there is not Firebase Auth at this point
	remaindersResp, err := fe.listRemainders(r.Context(), &pb.ListRemaindersRequest{UserUID: "no-user"})
	if err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("Getting firestore data: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)

	}

	if err := tmpl.Execute(w, remaindersResp); err != nil {
		msg := http.StatusText(http.StatusInternalServerError)
		log.Printf("template.Execute: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

func (fe *frontendServer) addHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Setting userUID because there is not Firebase Auth at this point
	remainder := r.Form.Get("remainder")
	_, err = fe.addRemainder(r.Context(),
		&pb.AddRemainderRequest{
			Remainder: &pb.Remainder{
				Remainder: remainder,
				UserUID:   "no-user",
			}})

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (fe *frontendServer) deleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Setting userUID because there is not Firebase Auth at this point
	remainderID := r.Form.Get("remainderID")
	_, err = fe.deleteRemainder(r.Context(),
		&pb.DeleteRemainderRequest{
			Remainder: &pb.Remainder{
				RemainderID: remainderID,
				UserUID:     "no-user",
			}})
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
