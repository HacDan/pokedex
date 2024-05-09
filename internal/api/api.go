package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	cache "github.com/hacdan/pokedex/internal/cache"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count       int        `json:"count"`
	NextUrl     string     `json:"next"`
	PreviousUrl string     `json:"previous"`
	Results     []Location `json:"results"`
}

var MCache cache.Cache

func GetLocations(url string) LocationResponse {
	locationResponse := LocationResponse{} //TODO: Change this to better practice. "New" Keyboard is bad.

	MCache = cache.NewCache(5 * time.Minute)

	if cache.Get(url)
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

	return locationResponse
}
