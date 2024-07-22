package exercises

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type APIInfo struct {
	Added     string                `json:"added"`
	Preferred string                `json:"preferred"`
	Versions  map[string]APIVersion `json:"versions"`
}

type APIVersion struct {
	Added          string `json:"added"`
	Info           Info   `json:"info"`
	Updated        string `json:"updated"`
	SwaggerURL     string `json:"swaggerUrl"`
	SwaggerYamlURL string `json:"swaggerYamlUrl"`
	OpenAPIVer     string `json:"openapiVer"`
	Link           string `json:"link"`
}

type Info struct {
	Contact      Contact  `json:"contact"`
	Description  string   `json:"description"`
	Title        string   `json:"title"`
	Version      string   `json:"version"`
	Categories   []string `json:"x-apisguru-categories"`
	Logo         Logo     `json:"x-logo"`
	Origins      []Origin `json:"x-origin"`
	ProviderName string   `json:"x-providerName"`
	ServiceName  string   `json:"x-serviceName"`
}

type Contact struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Logo struct {
	BackgroundColor string `json:"backgroundColor"`
	URL             string `json:"url"`
}

type Origin struct {
	Format  string `json:"format"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

func fetchAPIData(url string) (map[string]APIInfo, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]APIInfo
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func Exercise5() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := "https://api.apis.guru/v2/list.json"
		data, err := fetchAPIData(url)
		if err != nil {
			http.Error(w, "Error fetching data", http.StatusInternalServerError)
			log.Printf("Error fetching data: %v", err)
			return
		}

		// Only handle the first item from the data
		for id, apiInfo := range data {
			fmt.Fprintf(w, "API ID: %s\n", id)
			for version, versionInfo := range apiInfo.Versions {
				fmt.Fprintf(w, "  Version: %s\n", version)
				fmt.Fprintf(w, "  Title: %s\n", versionInfo.Info.Title)
				fmt.Fprintf(w, "  Description: %s\n", versionInfo.Info.Description)
				fmt.Fprintf(w, "  Provider: %s\n", versionInfo.Info.ProviderName)
				if versionInfo.Info.Contact.Email != "" {
					fmt.Fprintf(w, "  Contact Email: %s\n", versionInfo.Info.Contact.Email)
				}
				fmt.Fprintf(w, "\n")
				break // Break after handling the first version
			}
			break // Break after handling the first API
		}
	})

	fmt.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
