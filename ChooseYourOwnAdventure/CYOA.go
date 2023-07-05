package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type StoryOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Story struct {
	Title   string         `json:"title"`
	Story   []string       `json:"story"`
	Options []StoryOptions `json:"options"`
}

func loadStoriesFile() map[string]Story {
	result := make(map[string]Story)

	data, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("Error loading json file", err)
	}

	json.Unmarshal(data, &result)

	return result
}

func main() {
	loadStoriesFile()

	mux := http.NewServeMux()

	storiesHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		io.WriteString(w, "Hello")
	}

	mux.HandleFunc("/", storiesHandler)

	// log.Fatal(http.ListenAndServe(":4000", mux))

}
