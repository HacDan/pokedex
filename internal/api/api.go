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

type LocationResponse struct {
	Count       int        `json:"count"`
	NextUrl     string     `json:"next"`
	PreviousUrl string     `json:"previous"`
	Results     []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
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
