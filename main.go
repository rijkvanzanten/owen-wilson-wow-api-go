package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func main() {
	content, err := ioutil.ReadFile("./data.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var wows Data

	err = json.Unmarshal(content, &wows)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Printf("%v", wows)
}