package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

// helloRunHandler responds to requests by rendering an HTML page.
func (fe *frontendServer) rootHandler(w http.ResponseWriter, r *http.Request) {

	targetURL := fmt.Sprintf("http://%v/", fe.apigatewaySvcAddr)
	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	type RmdResp struct {
		RemainderID string `json:"remainderID"`
		Remainder   string `json:"remainder"`
	}

	var rms []RmdResp
	err = json.Unmarshal([]byte(body), &rms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, rms); err != nil {
		fmt.Println(err)
		msg := http.StatusText(http.StatusInternalServerError)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

/*
func (fe *frontendServer) addHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Setting userUID because there is not Firebase Auth at this point
	remainder := r.PostForm.Get("remainder")
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

func (fe *frontendServer) updateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	// Setting userUID because there is not Firebase Auth at this point
	remainderID := r.PostForm.Get("remainderID")
	remainder := r.PostForm.Get("remainder")
	_, err = fe.updateRemainder(r.Context(),
		&pb.UpdateRemainderRequest{
			Remainder: &pb.Remainder{
				RemainderID: remainderID,
				UserUID:     "no-user",
				Remainder:   remainder,
			}})

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}*/
