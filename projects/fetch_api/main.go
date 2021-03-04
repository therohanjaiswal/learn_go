package main

import (
	"io"
	"log"
	"os"
	"time"

	"./nhlApi"
)

func main() {

	now := time.Now()

	rosterFile, err := os.OpenFile("roster.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error while opening file roster.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	for _, team := range teams {
		log.Println("---------------------------")
		log.Printf("Name: %s\n", team.Name)
		log.Println("---------------------------")
	}

	log.Println("took %v", time.Now().Sub(now).String())
}
