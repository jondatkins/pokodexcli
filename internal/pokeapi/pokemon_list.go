package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(areaName string) (RespShallowPokemons, error) {
	url := baseURL + areaName

	if cached, ok := c.cache.Get(url); ok {
		var resp RespShallowPokemons
		// fmt.Println("Returning Cached response for %s", url)
		err := json.Unmarshal(cached, &resp)
		return resp, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	c.cache.Add(url, dat)

	pokemonResp := RespShallowPokemons{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	return pokemonResp, nil
}
