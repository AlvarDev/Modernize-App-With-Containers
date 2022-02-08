package main

import (
	"net/http"

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
