package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon name>")
	}
	name := args[0]
	if _, ok := cfg.caughtPokemon[name]; ok {
		fmt.Println("Name: ", cfg.caughtPokemon[name].Name)
		fmt.Println("Height: ", cfg.caughtPokemon[name].Height)
		fmt.Println("Weight: ", cfg.caughtPokemon[name].Weight)

		fmt.Println("Stats:")
		for _, stat := range cfg.caughtPokemon[name].Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range cfg.caughtPokemon[name].Types {
			fmt.Printf("  -%s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
