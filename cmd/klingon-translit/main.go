package main

import (
	"encoding/json"
	"fmt"
	"github.com/TriAnMan/jexiatest/usecase/klingon/translit"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	name := strings.Join(os.Args[1:], " ")

	klingon, err := translit.String(name)
	if err != nil {
		panic(err.Error())
	}

	for i, char := range klingon {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("0x%04X", char)
	}
	fmt.Print("\n")

	uri := "http://stapi.co/api/v1/rest/character/search"
	data := url.Values{"name": []string{name}}
	httpResp, err := http.PostForm(uri, data)
	if err != nil {
		panic(err.Error())
	}

	type species struct {
		Name string `json:"name"`
	}
	type character struct {
		Uid     string    `json:"uid"`
		Species []species `json:"characterSpecies"`
	}
	type searchResponse struct {
		Character []character `json:"characters"`
	}

	searchResp := searchResponse{}
	err = json.NewDecoder(httpResp.Body).Decode(&searchResp)
	if err != nil {
		panic(err.Error())
	}

	if len(searchResp.Character) == 0 {
		print("character not found")
		os.Exit(1)
	}

	haveSpecies := false
	for _, char := range searchResp.Character {
		uri = "http://stapi.co/api/v1/rest/character?" +
			url.Values{"uid": []string{char.Uid}}.Encode()
		httpResp, err = http.Get(uri)
		if err != nil {
			panic(err.Error())
		}

		type characterResponse struct {
			Character character `json:"character"`
		}

		charResp := characterResponse{}
		err = json.NewDecoder(httpResp.Body).Decode(&charResp)
		if err != nil {
			panic(err.Error())
		}

		if len(charResp.Character.Species) > 0 {
			for _, spec := range charResp.Character.Species {
				fmt.Printf("%s ", spec.Name)
			}
			haveSpecies = true
			break
		}
	}

	if !haveSpecies {
		print("species not found")
		os.Exit(1)
	}
}
