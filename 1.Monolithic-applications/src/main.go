package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

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

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
