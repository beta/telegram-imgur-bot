package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read web server port from env var.
	port, err := readPort()
	if err != nil {
		log.Fatalf("[web] %v", err)
	}

	// Start HTTP server for requests from wakemydyno.com.
	http.HandleFunc("/wakemydyno.txt", handleWakeMyDyno)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func readPort() (string, error) {
	const envPort = "PORT"
	port, ok := os.LookupEnv(envPort)
	if !ok {
		return "", fmt.Errorf("environment variable PORT is not properly set")
	}
	return port, nil
}

func handleWakeMyDyno(w http.ResponseWriter, r *http.Request) {
	const message = `# wakemydyno #`
	fmt.Fprintf(w, message)
}
