package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func isPokemonCaught(baseXP int) bool {
	baseExp := float32(baseXP)
	maxExp := float32(300)

	catchProbability := 1.0 - (baseExp / maxExp)
	randNum := rand.Float32()
	// fmt.Println("random num is:", randNum)
	if randNum < catchProbability {
		// fmt.Println("Caught!")
		return true
	} else {
		// fmt.Println("Escaped!")
		return false
	}
}

var pokedex map[string]RespShallowPokemonsInfo

func CatchPokemon(pokemon RespShallowPokemonsInfo) {
	if pokedex == nil {
		pokedex = make(map[string]RespShallowPokemonsInfo)
	}
	isCaught := isPokemonCaught(pokemon.BaseExperience)
	if isCaught {
		pokedex[pokemon.Name] = pokemon
		// p.Add(pokemon.Name, pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
}

func (c *Client) GetPokemon(pokemonName string) (RespShallowPokemonsInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName
	// fmt.Println(url)
	if cached, ok := c.cache.Get(url); ok {
		var resp RespShallowPokemonsInfo
		err := json.Unmarshal(cached, &resp)
		return resp, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}
	if resp.StatusCode != 200 {
		return RespShallowPokemonsInfo{}, fmt.Errorf("non-OK HTTP status: %s for URL %s", resp.Status, url)
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}
	// c.cache.Add(url, dat)
	pokemonResp := RespShallowPokemonsInfo{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemonsInfo{}, err
	}
	CatchPokemon(pokemonResp)
	return pokemonResp, nil
}
