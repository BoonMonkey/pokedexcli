package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeLocation struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     string
	Previous *string
}

func commandMap(cfg *Config) error {
	var fullUrl string

	if cfg.Next == "" {
		fullUrl = "https://pokeapi.co/api/v2/location-area?limit=20"
	} else {
		fullUrl = cfg.Next
	}

	fmt.Printf("Making request to: %s\n", fullUrl) // debug

	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Response status: %s\n", res.Status) // debug

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Printf("Response body length: %d\n", len(body)) // debug

	jsonData := PokeLocation{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Number of results: %d\n", len(jsonData.Results))

	for _, result := range jsonData.Results {
		fmt.Println(result.Name)
	}

	cfg.Next = jsonData.Next
	cfg.Previous = jsonData.Previous

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	fullUrl := *cfg.Previous

	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	jsonData := PokeLocation{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range jsonData.Results {
		fmt.Println(result.Name)
	}

	cfg.Next = jsonData.Next
	cfg.Previous = jsonData.Previous

	return nil
}
