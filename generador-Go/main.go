package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	id   int
	name string
}

type Command struct {
	games       []Game
	players     int
	rungames    int
	concurrence int
	timeout     float64
}

type Request struct {
	Gameid         int    `json:"gameid"`
	Gamename       string `json:"gamename"`
	Players        int    `json:"players"`
	Request_number int    `json:"request_number"`
}

var command Command

func processInput(input string) {
	/* Sintaxis esperada:
	rungame --gamename "1 | Game1 | 2 | Game2" --players 30 --rungames 30000 --concurrence 10 --timeout 3m
	*/
	split := strings.Split(input, " ")
	// fmt.Println(split)
	/* Variables temporales */
	var tmp string
	var game_tmp Game
	/* Variables finales */
	var games []Game
	var players int
	var rungames int
	var concurrence int
	var timeout float64
	for i := 0; i < len(split); i++ {
		tmp = split[i]
		if tmp == "--gamename" {
			aux := 0
			split[i+1] = strings.Replace(split[i+1], "\"", "", 1)
			for tmp[len(tmp)-1:] != "\"" {
				i++
				tmp = split[i]
				// fmt.Println(tmp)
				if tmp == "|" {
					continue
				}
				if aux%2 == 0 {
					game_tmp.id, _ = strconv.Atoi(tmp)
				} else {
					game_tmp.name = strings.Replace(tmp, "\"", "", -1)
					games = append(games, game_tmp)
				}
				aux++
			}
		} else if tmp == "--players" {
			i++
			tmp = split[i]
			players, _ = strconv.Atoi(tmp)
		} else if tmp == "--rungames" {
			i++
			tmp = split[i]
			rungames, _ = strconv.Atoi(tmp)
		} else if tmp == "--concurrence" {
			i++
			tmp = split[i]
			concurrence, _ = strconv.Atoi(tmp)
		} else if tmp == "--timeout" {
			i++
			tmp = split[i]
			tmp = tmp[:len(tmp)-1]
			timeout, _ = strconv.ParseFloat(tmp, 64)
		}
	}
	command.games = games
	command.players = players
	command.rungames = rungames
	command.concurrence = concurrence
	command.timeout = timeout
	// fmt.Println(command)
}

var bodies []Request

func generateJsonBodies() {
	var request Request
	var randomGame Game
	min_players := 1
	max_players := command.players
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= command.rungames; i++ {
		randomGame = command.games[rand.Intn(len(command.games))]
		randomPlayers := rand.Intn(max_players-min_players+1) + min_players
		request.Gameid = randomGame.id
		request.Gamename = randomGame.name
		request.Players = randomPlayers
		request.Request_number = i
		bodies = append(bodies, request)
		// fmt.Println(request)
	}
}

var request_number int = 0

func generateTraffic() {
	// your queue of json structs
	q := make(chan Request)
	// done channel takes the result of the request
	done := make(chan bool)
	// timeout
	timeout := time.After(time.Duration(command.timeout*60) * time.Second)

	for i := 0; i < command.concurrence; i++ {
		go doRequest(q, i, done)
	}

	// appends a struct to the queue
	for j := 0; j < command.rungames; j++ {
		// fmt.Println(j)
		go func(json Request) {
			q <- json
		}(bodies[j])
	}

	for c := 0; c < command.rungames; c++ {
		select {
		case <-done:
			request_number++
		case <-timeout:
			fmt.Println("Time limit exceeded!")
			return
		}
	}

	fmt.Println("Rungames finished!")
}

// const INGRESS = "http://193.60.11.13.nip.io"
const INGRESS = "http://localhost:10000/play"

func doRequest(queue chan Request, worknumber int, done chan bool) {
	for {
		k := <-queue
		// time.Sleep(1 * time.Second)
		req, err := json.Marshal(k)
		if err != nil {
			fmt.Println(err)
		}
		_, err = http.Post(INGRESS, "application/json", bytes.NewBuffer(req))
		if err != nil {
			fmt.Println(err)
		}
		// Read the response body
		/* body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("STATUS:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", string(body))
		fmt.Println() */
		done <- true
	}
}

func readInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Usac Squid Game >> ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	// test := `rungame --gamename "1 | MaxPlayer | 2 | MinPlayer | 3 | RandomWinner" --players 50 --rungames 25 --concurrence 10 --timeout 0.1m`
	processInput(text)
	generateJsonBodies()
	generateTraffic()
	fmt.Printf("%d requests has been sent.\n", request_number)
}

func main() {
	readInput()
}
