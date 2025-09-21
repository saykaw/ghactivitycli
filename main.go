package main

import (
	"fmt"
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
		fmt.Printf("Fetched GitHub activity for user: %s\n", username)
		fmt.Print(res.Body)
		for {
			data := make([]byte, 1024)
			n, err := res.Body.Read(data)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				break
			}
			fmt.Printf("Read %d bytes from response body\n", n)
			fmt.Print(string(data[:n]))
		}

	}

}
