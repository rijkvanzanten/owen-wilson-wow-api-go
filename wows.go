package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var wowsLoaded = false
var wows Data

func GetWows() Data {
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