package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/person", getPerson())
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func getPerson() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		baseUrl := fmt.Sprintf("http://%s", os.Getenv("GO_GIN_SERVICE_HOST"))
		res, err := http.Post(baseUrl+"/people", "application/json", nil)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		dec := json.NewDecoder(res.Body)
		var p Person
		err = dec.Decode(&p)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write([]byte(fmt.Sprintf("person name's is %s ", p.FirstName)))
	}
}

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
