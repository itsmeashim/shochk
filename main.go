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

// Response represents the valid response structure from Shodan.
type Response struct {
	ScanCredits  int `json:"scan_credits"`
	QueryCredits int `json:"query_credits"`
}

func main() {
	// Command-line flags
	tokenFlag := flag.String("token", "", "Check a single Shodan API key.")
	fileFlag := flag.String("file", "", "Check multiple Shodan API keys from a file.")
	flag.Parse()

	// Check Shodan API key from a single token input
	if *tokenFlag != "" {
		checkToken(*tokenFlag)
		return
	}

	// Check Shodan API keys from a file input
	if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			checkToken(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
		}
		return
	}

	// Check Shodan API keys from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkToken(scanner.Text())
	}
}

// checkToken sends a request to the Shodan API to check the validity of a token.
func checkToken(token string) {
	resp, err := http.Get("https://api.shodan.io/api-info?key=" + token)
	if err != nil {
		fmt.Printf("%s - Error: %v\n", token, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		var response Response
		json.Unmarshal(body, &response)
		fmt.Printf("%s - Valid (Scan Credits: %d, Query Credits: %d)\n", token, response.ScanCredits, response.QueryCredits)
	} else {
		fmt.Printf("%s - Invalid\n", token)
	}
}
