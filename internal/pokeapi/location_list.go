package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {

	var loc Locations //; return struct

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check if in cache
	if val, ok := c.cache.Get(url); ok {
		// unmashal into struct
		//fmt.Println("cache hit:", url)
		err := json.Unmarshal(val, &loc)
		if err != nil {
			return Locations{}, err
		}
		return loc, nil
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

	// put new data into cache
	c.cache.Add(url, data)

	// unmashal into struct
	err = json.Unmarshal(data, &loc)
	if err != nil {
		return Locations{}, err
	}

	//return Locations struct
	return loc, nil

}
