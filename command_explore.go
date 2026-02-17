package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <locations>")
	}
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(args[0])
	if err != nil {
		return err
	}
	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
