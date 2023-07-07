package main

import (
	"encoding/json"
	"html/template"
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
	Title     string         `json:"title"`
	Paragraph []string       `json:"story"`
	Options   []StoryOptions `json:"options"`
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
	// load stories
	stories := loadStoriesFile()

	// create a mux.
	mux := http.NewServeMux()

	file, err := os.ReadFile("./story.tmpl.html")

	if err != nil {
		log.Fatal("Could not open html template file ", err)
	}

	// handler function which matches url path and displays correct stories.
	storiesHandler := func(w http.ResponseWriter, r *http.Request) {

		url := r.URL
		var path string

		// our home page is /intro
		if len(url.Path) <= 1 {
			io.WriteString(w, "home page is /intro")
			return
		}

		// basic path processing. removing / from start and end.
		if url.Path[len(url.Path)-1] == '/' {
			path = url.Path[1 : len(url.Path)-1]
		} else {
			path = url.Path[1:len(url.Path)]
		}

		// get correct story
		s, prs := stories[path]

		// send the correct title to user.
		if prs {
			t, err := template.New("base").Parse(string(file))
			if err != nil {
				log.Fatal("Error parsing the template file", err)
			}
			t.ExecuteTemplate(w, "base", s)
		} else {
			// wrong story requested.
			io.WriteString(w, "Cannot find a story at this url")
		}
	}

	mux.HandleFunc("/", storiesHandler)

	log.Fatal(http.ListenAndServe(":4000", mux))

}
