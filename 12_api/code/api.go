package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Planet is a planet
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// Person is a person
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople is a bunch of people
type AllPeople struct {
	People []Person `json:"results"`
}

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Println("Error fetching person's home world")
	}

	if dataBytes, err := ioutil.ReadAll(res.Body); err != nil {
		log.Println("Error parsing res.body")
	} else {
		json.Unmarshal(dataBytes, &p.Homeworld)
	}

}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting people")
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	fmt.Println(res)

	dataBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	var people AllPeople

	if err := json.Unmarshal(dataBytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, person := range people.People {
		person.getHomeWorld()
		fmt.Println(person)
	}

}

func main() {
	//	create route for /people
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
