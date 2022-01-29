package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	apiKey = os.Getenv("API_KEY")
)

type BuildStacks struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Owners struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

type Organizations struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Teams struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Regions struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Spaces struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Shield bool   `json:"shield"`
}

type Stacks struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type HerokuApp struct {
	Acm                          bool          `json:"acm"`
	ArchivedAt                   string        `json:"archived_at"`
	BuildPackProvidedDescription string        `json:"buildpack_provided_description"`
	BuildStack                   BuildStacks   `json:"build_stack"`
	CreatedAt                    string        `json:"created_at"`
	GitUrl                       string        `json:"git_url"`
	Id                           string        `json:"id"`
	InternalRouting              bool          `json:"internal_routing"`
	Maintenance                  bool          `json:"maintenance"`
	Name                         string        `json:"name"`
	Owner                        Owners        `json:"owner"`
	Org                          Organizations `json:"organization"`
	Team                         Teams         `json:"team"`
	Region                       Regions       `json:"region"`
	ReleasedAt                   string        `json:"released_at"`
	RepoSize                     int           `json:"repo_size"`
	SlugSize                     int           `json:"slug_size"`
	Space                        Spaces        `json:"space"`
	Stack                        Stacks        `json:"stack"`
	UpdatedAt                    string        `json:"updated_at"`
	WebUrl                       string        `json:"web_url"`
}

func loadDotEnvVars() error {
	err := godotenv.Load()

	if err != nil {
		return errors.New("error loading .env")
	}

	return nil
}

func main() {
	args := os.Args

	if len(args) == 0 {
		fmt.Println("Provide app name")
		return
	}

	// Load our .env vars
	if err := loadDotEnvVars(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	client := http.Client{}

	reqString := fmt.Sprintf("https://api.heroku.com/apps/%s", args[1])

	req, err := http.NewRequest("GET", reqString, nil)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	req.Header = http.Header{
		"Accept":        []string{"application/vnd.heroku+json; version=3"},
		"Authorization": []string{"Bearer", apiKey},
	}

	resp, err := client.Do(req)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var herokuApp HerokuApp

	if err := json.NewDecoder(resp.Body).Decode(&herokuApp); err != nil {
		log.Fatalf("Failed to decode json: %v\n", err)
	}

}
