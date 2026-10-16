package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// make request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	// use client to make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}

	// dont close response until end of function or go routine
	defer resp.Body.Close()

	// read stream into bytes array
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	// unmashal into struct
	var loc Locations
	err = json.Unmarshal(data, &loc)
	if err != nil {
		return Locations{}, err
	}

	//return Locations struct
	return loc, nil

}
