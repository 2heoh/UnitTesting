package main

import (
	"fmt"
	"log"

	"encoding/json"

	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

type Country struct {
	id         int     //`json:id`
	name       string  //`json:"name"`
	src        string  //`json:"source"`
	group_type string  //`json:"group_type"`
	_type      string  //`json:"type"`
	alt_names  string  //`json:"alt_names"`
	lat        float32 //`json:"lat"`
	lon        float32 //`json:"lon"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id,name,src,group_type,type,alt_names,lat,lon FROM address WHERE type='country' ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}

	countries := make([]Country, 0, 300)

	for rows.Next() {
		country := Country{}
		err = rows.Scan(
			&country.id,
			&country.name,
			&country.src,
			&country.group_type,
			&country._type,
			&country.alt_names,
			&country.lat,
			&country.lon,
		)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
	}

	log.Printf("c: %v", countries[0].id)

	jsonCountries, err := json.Marshal(countries[0])

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("json: %v", jsonCountries)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonCountries))
}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://reader:reader@osm-db-dev.srv.pv.km/address?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
