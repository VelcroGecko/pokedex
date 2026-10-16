package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {

	// get next 20 locations
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// update config
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// print the locations
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {

	// make sure not on first page:
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	// get prev 20 locations
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// update config
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// print the locations
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
