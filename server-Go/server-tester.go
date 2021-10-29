package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	Gameid         int    `json:"gameid"`
	Gamename       string `json:"gamename"`
	Players        int    `json:"players"`
	Request_number int    `json:"request_number"`
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// var b bytes.Buffer
	// b.Write(reqBody)
	// fmt.Println(b.String())
	var game Request
	json.Unmarshal(reqBody, &game)

	fmt.Println("Request:")
	fmt.Println("Gameid:", game.Gameid)
	fmt.Println("Gamename:", game.Gamename)
	fmt.Println("Players:", game.Players)
	fmt.Println("Request_number:", game.Request_number)
	fmt.Println()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { fmt.Println("Hi from GO server!") })
	myRouter.HandleFunc("/play", getRequest).Methods("POST", "GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Go server listening on port: 10000")
	handleRequests()
}
