package main

import (
	"fmt"
	"net/http"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/plugin/ochttp"
)

const (
	port = ":8080"
)

type frontendServer struct {
	apigatewaySvcAddr string
}

func main() {

	svc := new(frontendServer)
	mustMapEnv(&svc.apigatewaySvcAddr, "API_GATEWAY_SERVICE_ADDR")

	r := mux.NewRouter()
	r.HandleFunc("/", svc.rootHandler).Methods(http.MethodGet)
	//r.HandleFunc("/add", svc.addHandler).Methods(http.MethodPost)
	//r.HandleFunc("/delete", svc.deleteHandler).Methods(http.MethodGet)
	//r.HandleFunc("/update", svc.updateHandler).Methods(http.MethodPost)

	httpHandler := &ochttp.Handler{
		Propagation: &propagation.HTTPFormat{},
		Handler:     r,
	}

	log.Info().Msg("Starting frontend service")

	if err := http.ListenAndServe(port, httpHandler); err != nil {
		log.Fatal().Err(err).Msg("Can't start service")
	}
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
