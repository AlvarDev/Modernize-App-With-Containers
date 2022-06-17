package main

import (
	pb "apigatewayservice/pb"
	"encoding/json"
	"net/http"
)

// listHandler.
func (fe *frontendServer) listHandler(w http.ResponseWriter, r *http.Request) {

	// Setting userUID because there is not Firebase Auth at this point
	remaindersResp, err := fe.listRemainders(r.Context(), &pb.ListRemaindersRequest{UserUID: "no-user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rms := make([]*pb.Remainder, len(remaindersResp.GetRemainders()))
	for i, r := range remaindersResp.GetRemainders() {
		rms[i] = r
	}

	responseJSON, _ := json.Marshal(rms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func (fe *frontendServer) addHandler(w http.ResponseWriter, r *http.Request) {

	// Setting userUID because there is not Firebase Auth at this point
	rmd := RmdPost{}
	err := json.NewDecoder(r.Body).Decode(&rmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	newRmd, err := fe.addRemainder(r.Context(),
		&pb.AddRemainderRequest{
			Remainder: &pb.Remainder{
				Remainder: rmd.Remainder.GetRemainder(),
				UserUID:   "no-user",
			}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseJSON, _ := json.Marshal(newRmd)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

// TODO: avoid use GET
func (fe *frontendServer) deleteHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Setting userUID because there is not Firebase Auth at this point
	remainderID := r.Form.Get("remainderID")
	delRmd, err := fe.deleteRemainder(r.Context(),
		&pb.DeleteRemainderRequest{
			Remainder: &pb.Remainder{
				RemainderID: remainderID,
				UserUID:     "no-user",
			}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseJSON, _ := json.Marshal(delRmd)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func (fe *frontendServer) updateHandler(w http.ResponseWriter, r *http.Request) {

	rmd := RmdPost{}
	err := json.NewDecoder(r.Body).Decode(&rmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Setting userUID because there is not Firebase Auth at this point
	upRmd, err := fe.updateRemainder(r.Context(),
		&pb.UpdateRemainderRequest{
			Remainder: &pb.Remainder{
				RemainderID: rmd.Remainder.GetRemainderID(),
				UserUID:     "no-user",
				Remainder:   rmd.Remainder.GetRemainder(),
			}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseJSON, _ := json.Marshal(upRmd)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

type RmdPost struct {
	Remainder *pb.Remainder `json:"remainder"`
}
