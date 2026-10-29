package main

import (
	"time"

	"github.com/velcrogecko/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
	}
	startRepl(cfg)

}
