package main

import (
	"time"

	"github.com/velcrogecko/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
	}
	startRepl(cfg)

}
