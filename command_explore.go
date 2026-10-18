package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	//ensure only given one location
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]

	// get detailed location data
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")

	// print the pokemones
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
