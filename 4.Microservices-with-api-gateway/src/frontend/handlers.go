package main

import (
	"bytes"
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

func (fe *frontendServer) addHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type RmdReq struct {
		Remainder RmdResp `json:"remainder"`
	}
	rmd := RmdReq{
		Remainder: RmdResp{
			Remainder: r.PostForm.Get("remainder"),
		}}
	rmdJSON, _ := json.Marshal(rmd)

	// Setting userUID on server because there is not Firebase Auth at this point
	targetURL := fmt.Sprintf("http://%v/add", fe.apigatewaySvcAddr)
	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(rmdJSON))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (fe *frontendServer) deleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Setting userUID because there is not Firebase Auth at this point
	targetURL := fmt.Sprintf("http://%v/delete?remainderID=%v", fe.apigatewaySvcAddr, r.Form.Get("remainderID"))
	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (fe *frontendServer) updateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type RmdReq struct {
		Remainder RmdResp `json:"remainder"`
	}
	rmd := RmdReq{
		Remainder: RmdResp{
			Remainder:   r.PostForm.Get("remainder"),
			RemainderID: r.PostForm.Get("remainderID"),
		}}
	rmdJSON, _ := json.Marshal(rmd)

	// Setting userUID on server because there is not Firebase Auth at this point
	targetURL := fmt.Sprintf("http://%v/update", fe.apigatewaySvcAddr)
	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(rmdJSON))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type RmdResp struct {
	RemainderID string `json:"remainderID"`
	Remainder   string `json:"remainder"`
}
