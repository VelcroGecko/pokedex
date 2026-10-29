package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {

	url := baseURL + "/pokemon/" + pokemonName
	var pokemonResp Pokemon //; return struct

	//check if in cache and if so return
	// check if in cache
	if val, ok := c.cache.Get(url); ok {
		// unmashal into struct
		//fmt.Println("cache hit:", url)
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// make request obj
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// use client to make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	// dont close response until end of function or go routine
	defer resp.Body.Close()

	// read stream into bytes array
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// probs add to cache here
	c.cache.Add(url, data)

	// unmashal into struct
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil

}
