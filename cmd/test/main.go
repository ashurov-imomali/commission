package main

import (
	"encoding/json"
	"log"
	"main/models"
	"os"
)

func main() {
	//dbSettings := `host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable`
	//repos, err := db.GetConnection(dbSettings)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
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
	log.Println(testStr)
	testStr.Profile.Active = new(bool)
	*testStr.Profile.Active = true
	log.Println("ok", testStr)
	return
	//responses := handler.GetHandler(repos)
	//rout := router.GetRouter(responses)
	//
	//err = http.ListenAndServe("localhost:8080", rout)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
}
