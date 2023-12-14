package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	schemaURL = "https://gitlab.com/api/graphql"
	query     = `{
		metadata {
			version
		}
	}`
)

func GetGitlabVersion() string {
	client := &http.Client{}

	token := os.Getenv("BEARER_TOKEN")

	// Construct the JSON payload for the POST request
	data := map[string]interface{}{
		"query": query,
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", schemaURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read the response body as JSON
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		panic(err)
	}

	metadata := result["data"].(map[string]interface{})["metadata"]
	version := metadata.(map[string]interface{})["version"].(string)

	fmt.Println("Metadata version:", version)

	return version
}
