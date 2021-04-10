// Pakage icndb provides access to the Internet Chuck Norris Database (ICNDb) API.
package icndb

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"strconv"
)

const apiURL = "https://api.icndb.com"

// Joke is a Chuck Norris joke. It is the basic unit of information of the ICNDb API.
type Joke struct {
	// ID is the joke's unique numeric reference.
	ID int `json:"id"`
	// Text is the actual joke.
	Text string `json:"joke"`
	// Categories is the list of categories the joke belongs to.
	Categories []string `json:"categories"`
}

// Rand gets a random joke.
func Rand() (Joke, error) {
	resp, err := http.Get(apiURL + "/jokes/random")
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	raw, err := decodeResp(resp.Body)
	if err != nil {
		return Joke{}, err
	}
	var joke Joke
	if err := json.Unmarshal(raw, &joke); err != nil {
		return Joke{}, fmt.Errorf("failed decoding joke: %v", err)
	}
	return joke, nil
}

// Randn gets n random jokes.
func Randn(n int) ([]Joke, error) {
	resp, err := http.Get(apiURL + "/jokes/random/" + strconv.Itoa(n))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := decodeResp(resp.Body)
	if err != nil {
		return nil, err
	}
	var jokes []Joke
	if err := json.Unmarshal(raw, &jokes); err != nil {
		return nil, fmt.Errorf("failed decoding jokes: %v", err)
	}
	return jokes, nil
}

// Fetch gets a specific joke with the given id.
func Fetch(id int) (Joke, error) {
	resp, err := http.Get(apiURL + "/jokes/" + strconv.Itoa(id))
	if err != nil {
		return Joke{}, err
	}
	defer resp.Body.Close()

	raw, err := decodeResp(resp.Body)
	if err != nil {
		return Joke{}, err
	}
	var joke Joke
	if err := json.Unmarshal(raw, &joke); err != nil {
		return Joke{}, fmt.Errorf("failed decoding joke: %v", err)
	}
	joke.Text = html.UnescapeString(joke.Text)
	return joke, nil
}

// Count gets the number of available jokes.
func Count() (int, error) {
	resp, err := http.Get(apiURL + "/jokes/count")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	raw, err := decodeResp(resp.Body)
	if err != nil {
		return 0, err
	}
	var count int
	if err := json.Unmarshal(raw, &count); err != nil {
		return 0, fmt.Errorf("failed decoding number: %v", err)
	}
	return count, nil
}

type response struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func decodeResp(in io.Reader) (json.RawMessage, error) {
	var resp response
	if err := json.NewDecoder(in).Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed decoding response: %v", err)
	}
	if resp.Type != "success" {
		var msg string
		if err := json.Unmarshal(resp.Value, &msg); err != nil {
			return nil, fmt.Errorf("failed decoding json: %v", err)
		}
		return nil, errors.New(msg)
	}
	return resp.Value, nil
}
