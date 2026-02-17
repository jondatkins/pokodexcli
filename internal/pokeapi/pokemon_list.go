package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(areaName string) (RespShallowPokemons, error) {
	url := baseURL + "/location-area/" + areaName
	if cached, ok := c.cache.Get(url); ok {
		var resp RespShallowPokemons
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
	if resp.StatusCode != 200 {
		return RespShallowPokemons{}, fmt.Errorf("non-OK HTTP status: %s for URL %s", resp.Status, url)
	}
	defer resp.Body.Close()
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
