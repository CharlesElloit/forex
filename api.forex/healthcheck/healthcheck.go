package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port, lookupError := os.LookupEnv("PORT")
	if !lookupError {
		port = "4000"
	}

	httpClient := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, error := httpClient.Get("http://localhost:" + port + "/ping")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Argh! Broken")
		os.Exit(1)
	}
}
