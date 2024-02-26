package main

import (
	"encoding/json"
	"log"
	"main/db"
	"main/handler"
	"main/models"
	"main/router"
	"net/http"
	"os"
)

func main() {
	dbSettings := `host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable`
	repos, err := db.GetConnection(dbSettings)
	if err != nil {
		log.Println(err)
		return
	}
	bytes, err := os.ReadFile("./requests/test.json")
	if err != nil {
		log.Println(err, "there ???")
		return
	}
	var testStr models.ProfileCreatRequest
	err = json.Unmarshal(bytes, &testStr)
	if err != nil {
		log.Println(err)
		return
	}
	testStr.Profile.Active = true
	responses := handler.GetHandler(repos)
	rout := router.GetRouter(responses)

	err = http.ListenAndServe("localhost:8080", rout)
	if err != nil {
		log.Println(err)
		return
	}
}
