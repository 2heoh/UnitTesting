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
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Src        string  `json:"source"`
	Group_type string  `json:"group_type"`
	Type       string  `json:"type"`
	Alt_names  string  `json:"alt_names"`
	Lat        float32 `json:"lat"`
	Lon        float32 `json:"lon"`
	Info       string  `json:"info"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id,name,src,group_type,type,alt_names,lat,lon,info FROM address WHERE type='country' ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}

	countries := []Country{}

	for rows.Next() {
		country := Country{}
		err = rows.Scan(
			&country.Id,
			&country.Name,
			&country.Src,
			&country.Group_type,
			&country.Type,
			&country.Alt_names,
			&country.Lat,
			&country.Lon,
			&country.Info,
		)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
	}

	jsonCountries, err := json.Marshal(countries)
	if err != nil {
		log.Fatal(err)
	}

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
