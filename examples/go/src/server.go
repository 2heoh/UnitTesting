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

type Info struct {
	Continent string `json:"continent"`
	Zoom      int    `json:"zoom"`
	Iso2      string `json:"iso2"`
	Iso3      string `json:"iso3"`
}

type Country struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Path       string  `json:"path"`
	Src        string  `json:"source"`
	Group_type string  `json:"group_type"`
	Type       string  `json:"type"`
	Lat        float32 `json:"lat"`
	Lon        float32 `json:"lon"`
	Info       string  `json:"info"`
}

func getCountries(db *sql.DB) []Country {
	rows, err := db.Query("SELECT id,name,path,src,group_type,type,lat,lon,info FROM address WHERE type='country' ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}

	countries := []Country{}

	for rows.Next() {
		country := Country{}
		err = rows.Scan(
			&country.Id,
			&country.Name,
			&country.Path,
			&country.Src,
			&country.Group_type,
			&country.Type,
			&country.Lat,
			&country.Lon,
			&country.Info,
		)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)
	}
	return countries
}

func handler(w http.ResponseWriter, r *http.Request) {

	countries := getCountries(db)

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
