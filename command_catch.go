package main

import "fmt"

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon name>")
	}
	fmt.Printf("- 'Throwing a Pokeball at %s...'\n", args[0])
	_, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	// cfg.pokeapiClient.CatchPokemon(pokemonResp)
	// fmt.Printf("Base Exp %d\n", pokemonResp.BaseExperience)
	// for _, encounter := range pokemonResp.PokemonEncounters {
	// 	fmt.Printf("- %s\n", encounter.Pokemon.Name)
	// }
	// cfg.pokeapiPokedex.CatchPokemon(pokemonResp)
	return nil
}
