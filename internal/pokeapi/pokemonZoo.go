package pokeapi

import (
	"sync"
)

type Pokedex struct {
	mu   sync.Mutex
	data map[string]RespShallowPokemonsInfo
}

func (p *Pokedex) Add(key string, pokemon RespShallowPokemonsInfo) {
	p.mu.Lock()
	p.data[key] = pokemon
	p.mu.Unlock()
}

func (p *Pokedex) Get(key string) (RespShallowPokemonsInfo, bool) {
	p.mu.Lock()
	if value, isMapContainsKey := p.data[key]; isMapContainsKey {
		p.mu.Unlock()
		return value, true
	} else {
		p.mu.Unlock()
		return RespShallowPokemonsInfo{}, false
	}
}

func NewPokedex() *Pokedex {
	newpokedex := Pokedex{
		data: make(map[string]RespShallowPokemonsInfo),
	}
	return &newpokedex
}
