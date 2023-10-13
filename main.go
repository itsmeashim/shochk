package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const shodanAPIURL = "https://api.shodan.io/account/profile?key="

type ShodanProfile struct {
	Member      bool   `json:"member"`
	Credits     int    `json:"credits"`
	DisplayName string `json:"display_name"`
	Created     string `json:"created"`
}

func checkAPIKey(apiKey string) {
	resp, err := http.Get(shodanAPIURL + apiKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making the request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		fmt.Println(apiKey, "- Invalid")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		return
	}

	var profile ShodanProfile
	if err := json.Unmarshal(body, &profile); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling JSON: %v\n", err)
		return
	}

	fmt.Println(apiKey, "- Valid (Member:", profile.Member, ", Credits:", profile.Credits, ")")
}

func checkKeysFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening the file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		checkAPIKey(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the file: %v\n", err)
	}
}

func main() {
	var tokenFlag, fileFlag string
	flag.StringVar(&tokenFlag, "token", "", "Specify a single Shodan API key to check.")
	flag.StringVar(&fileFlag, "file", "", "Specify a file containing a list of Shodan API keys to check (one per line).")
	flag.Parse()

	if tokenFlag == "" && fileFlag == "" {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			checkAPIKey(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		}
		return
	}

	if tokenFlag != "" {
		checkAPIKey(tokenFlag)
	}

	if fileFlag != "" {
		checkKeysFromFile(fileFlag)
	}
}
