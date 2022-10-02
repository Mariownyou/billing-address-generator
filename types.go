package main

import (
	"encoding/json"
	"fmt"
)


type Place struct {
	Center	  []float32 `json:"center"`
	PlaceName string    `json:"place_name"`
	Relevance float32   `json:"relevance"`
	Address   string    `json:"address"`
}


type Response struct {
	Features []Place `json:"features"`
}


type CityCoordinates struct {
	Lat		float32
	Lng		float32
	Name	string
	Address string
}


func (c *CityCoordinates) UnmarshalJSON(data []byte) error {	
	var vals Response
    if err := json.Unmarshal(data, &vals); err != nil {
        return err
    }

	if len(vals.Features) > 0 {
		place := vals.Features[0]
		c.Name = place.PlaceName
		c.Lng = place.Center[0]
		c.Lat = place.Center[1]
		c.Address = place.Address
	}
	
	return nil
}


func (c *CityCoordinates) String() string {
	if c.Name == "" {
		return "Not Found, Try different city"
	}
	return fmt.Sprintf("%s", c.Name)
}

