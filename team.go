package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TeamMember struct {
	Name string `json:"name"`
	Id   int64  `json:"userid"`
}

type Team struct {
	Members []TeamMember `json:"members"`
}

func GetTeam() Team {
	// Datei einlesen und parsen

	// Datei einlesen...
	jsonFile, err := os.Open("team.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// ... und parsen
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var team Team

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &team)

	return team
}
