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
	infoRaw    string
	Info       Info `json:"info"`
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
			&country.infoRaw,
		)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, country)

	}
	return countries
}

func prettyJson(countries []Country) (string, error) {

	for i, _ := range countries {
		err := json.Unmarshal([]byte(countries[i].infoRaw), &countries[i].Info)
		if err != nil {
			return "", fxmt.Errorf("can't parse \"info\"("+countries[i].infoRaw+") ", err)
		}
	}

	jsonCountries, err := json.Marshal(countries)

	if err != nil {
		return "", err
	}

	return string(jsonCountries), nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Println("GET /")

	json, err := prettyJson(getCountries(db))

	if err != nil {
		log.Fatal("ERORR: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, json)
}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://reader:reader@osm-db-dev/address?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	log.Println("Listen on 8080")
	http.ListenAndServe(":8080", nil)
}
