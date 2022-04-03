package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	pb "mymessagesapigateway/pb"

	"google.golang.org/protobuf/encoding/protojson"
)

func (fe *frontendServer) listMessagesHandler(w http.ResponseWriter, r *http.Request) {

	messages, err := fe.listMessages(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}

	jsonBytes, _ := m.Marshal(messages)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (fe *frontendServer) addMessageHandler(w http.ResponseWriter, r *http.Request) {

	req := &pb.AddMyMessageRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = protojson.Unmarshal(body, req)
	if err != nil {
		fmt.Println(err)
	}

	messageAdded, err := fe.addMessage(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := protojson.MarshalOptions{
		Indent:          "  ",
		EmitUnpopulated: true,
	}

	jsonBytes, _ := m.Marshal(messageAdded)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
