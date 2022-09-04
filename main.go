package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Data []struct {
	Movie string `json:"movie"`
	Year int16 `json:"year"`
	ReleaseDate string `json:"release_date"`
	Director string `json:"director"`
	Character string `json:"character"`
	MovieDuration string `json:"movie_duration"`
	Timestamp string `json:"timestamp"`
	FullLine string `json:"full_line"`
	CurrentWowInMovie uint8 `json:"current_wow_in_movie"`
	TotalWowsInMovie uint8 `json:"total_wows_in_movie"`
	Poster string `json:"poster"`
	Video struct {
		Res1080p string `json:"1080p"`
		Res720p string `json:"720p"`
		Res480p string `json:"480p"`
		Res360p string `json:"360p"`
	} `json:"video"`
	Audio string `json:"audio"`
}

var wowsLoaded = false
var wows Data

func getWows() Data {
	if wowsLoaded {
		return wows
	}

	content, err := ioutil.ReadFile("./data.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &wows)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	wowsLoaded = true
	
	return wows
}

func routeHome(w http.ResponseWriter, r *http.Request) {
	wows := getWows()
	jsonData, err := json.Marshal(wows)

	if err != nil {
		log.Printf("Could not marshal JSON: %s\n", err)
		io.WriteString(w, fmt.Sprintf("Could not marshal JSON: %s\n", err));
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routeHome)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(":8080", handler)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed\n")
	} else {
		log.Fatalf("Error starting server: %s\n", err)
	}
}