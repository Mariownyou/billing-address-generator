package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const mapboxUrl = "https://api.mapbox.com/geocoding/v5/mapbox.places/"

var token string
var city string

func init() {
	flag.StringVar(&city, "c", "Moscow", "City")
	flag.Parse()
}

func main() {
	// get city coordinates
	cityText, err := getMapboxResponse(city, map[string]string{"proximity": "ip", "types": "place"})

	if err != nil {
		panic(err)
	}

	var cityCoordinates CityCoordinates
	if err := json.Unmarshal([]byte(cityText), &cityCoordinates); err != nil {
		panic(err)
	}

	// get nearest address
	coordinatesString := fmt.Sprintf("%f,%f", cityCoordinates.Lng, cityCoordinates.Lat)
	addressText, err := getMapboxResponse(coordinatesString, map[string]string{"types": "address"})
	if err != nil {
		panic(err)
	}

	var address CityCoordinates
	if err := json.Unmarshal([]byte(addressText), &address); err != nil {
		panic(err)
	}

	fmt.Println(address.String())
}

func getMapboxResponse(query string, params map[string]string) (string, error) {
	baseUrl, err := url.Parse(mapboxUrl)
	if err != nil {
		return "", err
	}

	// create url
	baseUrl.Path += fmt.Sprintf("%s.json", query)

	// add params
	p := url.Values{}
	p.Add("access_token", token)
	for key, element := range params {
		p.Add(key, element)
	}

	baseUrl.RawQuery = p.Encode()

	// create request
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	// parse response
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(bodyText), nil
}
