package main

import (
	"github.com/velcrogecko/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(cfg)

}
