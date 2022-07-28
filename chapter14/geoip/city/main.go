package main

import (
	"encoding/json"
	"fmt"
	"github.com/oschwald/maxminddb-golang"
	"log"
	"net"
)

type GeoCityRecord struct {
	Continent struct {
		Code      string                 `json:"code"`
		GeonameId int                    `json:"geoname_id"`
		Names     map[string]interface{} `json:"names"`
	} `json:"continent"`
	Country struct {
		GeonameId int                    `json:"geoname_id"`
		IsoCode   string                 `json:"iso_code"`
		Names     map[string]interface{} `json:"names"`
	} `json:"country"`
	Location struct {
		AccuracyRadius int     `json:"accuracy_radius"`
		Latitude       float32 `json:"latitude"`
		Longitude      float32 `json:"longitude"`
		TimeZone       string  `json:"time_zone"`
	} `json:"location"`
	RegisteredCountry struct {
		GeoNameID int                    `json:"geoname_id"`
		IsoCode   string                 `json:"iso_code"`
		Names     map[string]interface{} `json:"names"`
	} `json:"registered_country"`
}

func main() {
	db, err := maxminddb.Open("/home/nanik/GolandProjects/cloudprogramminggo/chapter14/geoip/city/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, network, err := net.ParseCIDR("2.0.0.0/8")
	if err != nil {
		log.Panic(err)
	}

	networks := db.NetworksWithin(network, maxminddb.SkipAliasedNetworks)
	for networks.Next() {
		var rec interface{}
		r := GeoCityRecord{}

		ip, err := networks.Network(&rec)

		if err != nil {
			log.Panic(err)
		}
		j, _ := json.Marshal(rec)

		err = json.Unmarshal([]byte(j), &r)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("IP : %s, Long : %f, Lat : %f, Country : %s, Continent: %s\n", ip.String(), r.Location.Longitude, r.Location.Latitude,
			r.Country.IsoCode, r.Continent.Code)
	}
}
