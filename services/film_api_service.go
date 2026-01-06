package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Struktur response dari OMDb
type OMDbSearchResponse struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

type OMDbDetailResponse struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Plot     string `json:"Plot"`
	Poster   string `json:"Poster"`
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

// Cari film berdasarkan keyword
func SearchFilms(keyword string) (interface{}, error) {

	apiKey := os.Getenv("OMDB_API_KEY")
	url := fmt.Sprintf(
		"http://www.omdbapi.com/?i=tt3896198&apikey=f42d455f",
		apiKey,
		keyword,
	)

	// request ke API OMDb
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode JSON
	var result OMDbSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	// kalau film tidak ditemukan
	if result.Response == "False" {
		return nil, fmt.Errorf(result.Error)
	}

	return result.Search, nil
}

// Ambil detail film
func GetFilmDetail(imdbID string) (interface{}, error) {

	apiKey := os.Getenv("OMDB_API_KEY")
	url := fmt.Sprintf(
		"https://www.omdbapi.com/?apikey=%s&i=%s",
		apiKey,
		imdbID,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result OMDbDetailResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.Response == "False" {
		return nil, fmt.Errorf(result.Error)
	}

	return result, nil
}
