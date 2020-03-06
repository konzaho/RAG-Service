package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Request struct {
	Nickname  string
	Locations []string
	Times     []string
	Factions  []string
}

type Response struct {
	Location string `json:"location"`
	Time     string `json:"time"`
	Faction  string `json:"faction"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/rag", homeLink).Methods("POST")
	router.HandleFunc("/rag/generate", generateLevel).Methods("POST").Headers("Content-Type", "application/json")

	log.Print(http.ListenAndServe(":8080", router))
}

func generateLevel(w http.ResponseWriter, r *http.Request) {
	var request Request
	var response Response
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Error on parsing body request")
	}
	json.Unmarshal(reqBody, &request)
	response.Location = request.Locations[randomIndex(len(request.Locations))]
	response.Time = request.Times[randomIndex(len(request.Times))]
	response.Faction = request.Factions[randomIndex(len(request.Factions))]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func randomIndex(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	return rand.Intn(max-min) + min
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Randomize your ArmA Escape scenario with this service! -by: KonZaho")
}
