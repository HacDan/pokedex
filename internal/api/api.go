package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	cache "github.com/hacdan/pokedex/internal/cache"
)

type Client struct {
	cache  cache.Cache
	client http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocations(url string) LocationResponse {
	locationResponse := LocationResponse{} //TODO: Change this to better practice. "New" Keyboard is bad.
	val, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			panic(err)
		}
		return locationResponse
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		panic(err)
	}

	c.cache.Add(url, body)
	return locationResponse
}

func (c *Client) GetPokemonFromLocation(url string) LocationAreaDetailsResponse {
	locationAreaDetailsResponse := LocationAreaDetailsResponse{}
	val, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(val, &locationAreaDetailsResponse)
		if err != nil {
			panic(err)
		}
		return locationAreaDetailsResponse
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &locationAreaDetailsResponse)
	if err != nil {
		panic(err)
	}

	c.cache.Add(url, body)
	return locationAreaDetailsResponse
}

func (c *Client) GetPokemon(url string) PokemonResponse {
	pokemonResponse := PokemonResponse{}
	val, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			panic(err)
		}
		return pokemonResponse
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		panic(err)
	}

	c.cache.Add(url, body)
	return pokemonResponse
}
