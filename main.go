package main

import (
	"encoding/json"
	"fmt"
	"ghactivitycli/types"
	"io"
	"net/http"
	"os"
)

func main() {

	if os.Args[1] == "github-activity" {
		username := os.Args[2]

		req, err := http.NewRequest("GET", "https://api.github.com/users/"+username+"/events", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		}

		githubEvents := []types.GitHubEvent{}
		err = json.Unmarshal(data, &githubEvents)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}
		fmt.Printf("%+v\n", githubEvents)

	}

}
