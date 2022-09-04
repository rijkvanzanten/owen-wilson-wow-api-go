package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func unexpectedError(w *http.ResponseWriter, err error) {
	writer := *w
	log.Printf("Could not marshal JSON: %s\n", err)
	io.WriteString(writer, fmt.Sprintf("Could not marshal JSON: %s\n", err));
	writer.WriteHeader(http.StatusInternalServerError)
	return
}

func RouteHome(w http.ResponseWriter, r *http.Request) {
	wows := GetWows()
	jsonData, err := Marshal(wows)

	if err != nil {
		unexpectedError(&w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func RouteMovies(w http.ResponseWriter, r *http.Request) {
	wows := GetWows()

	keys := make(map[string]bool)
	list := []string{}

	for _, wow := range *wows {
		movie := wow.Movie

		if _, value := keys[movie]; !value {
			keys[movie] = true
			list = append(list, movie)
		}
	}

	jsonData, err := Marshal(list)

	if err != nil {
		unexpectedError(&w, err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func RouteDirectors(w http.ResponseWriter, r *http.Request) {
	wows := GetWows()

	keys := make(map[string]bool)
	list := []string{}

	for _, wow := range *wows {
		director := wow.Director

		if _, value := keys[director]; !value {
			keys[director] = true
			list = append(list, director)
		}
	}

	jsonData, err := Marshal(list)

	if err != nil {
		unexpectedError(&w, err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}