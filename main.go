package main

import (
	"log"
	"net/http"
)

const envWebPort = "WEBCONFIG_PORT"
const envWebListenIp = "WEBCONFIG_LISTEN_IP"

const defaultWebPort = "8080"
const defaultWebListenIp = "0.0.0.0"

func main() {
	// Get information from the environment variable
	WebPort := Getenv(envWebPort, defaultWebPort)
	ListenIp := Getenv(envWebListenIp, defaultWebListenIp)

	// Declare web Handlers
	http.HandleFunc("/ip", makeHandler(ipHandler))
	http.HandleFunc("/all.json", makeHandler(jsonHandler))
	http.HandleFunc("/", makeHandler(mainHandler))
	// Log and Listen
	log.Printf("Listening on %s:%s", ListenIp, WebPort)
	err := http.ListenAndServe(ListenIp+":"+WebPort, nil)
	log.Fatal(err)
}
