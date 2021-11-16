package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type Body struct {
	Name string `json:"name"`
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var body Body
	json.Unmarshal(reqBody, &body)
	fmt.Fprintf(w, "Hello %s!", body.Name)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHi)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("Server listening on port: %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
