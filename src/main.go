package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"
  "os"
)

type webConfig struct {
	Ip            net.IP
	Host          string
	Port          string
	Method        string
	Language      string
	UserAgent     string
	XForwardedFor string
}

const envWebPort = "WEBCONFIG_PORT"
const envWebListenIp = "WEBCONFIG_LISTEN_IP"

const defaultWebPort = "8080"
const defaultWebListenIp = "0.0.0.0"

func formatClientInformation(req *http.Request) *webConfig {
	// Get the cient ip
	ip, port, _ := net.SplitHostPort(req.RemoteAddr)

	config := new(webConfig)
	config.Ip = net.ParseIP(ip)
	config.Host = req.Host
	config.Port = port
	config.Method = req.Method
	config.Language = req.Header.Get("Accept-Language")
	config.UserAgent = req.Header.Get("User-Agent")
	config.XForwardedFor = req.Header.Get("X-Forwarded-For")

	return config
}

func ipHandler(writer http.ResponseWriter, req *http.Request, clientConfig *webConfig) {
	fmt.Fprintf(writer, "%s", clientConfig.Ip)
}

func jsonHandler(writer http.ResponseWriter, req *http.Request, clientConfig *webConfig) {
	jsonClientConfig, err := json.Marshal(clientConfig)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "%s", jsonClientConfig)
}

func mainHandler(writer http.ResponseWriter, req *http.Request, clientConfig *webConfig) {
  template, err := template.ParseFiles("templates/main.html")
  if err != nil {
  		http.Error(writer, err.Error(), http.StatusInternalServerError)
  		return
  }
  template.Execute(writer, clientConfig)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *webConfig)) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		config := formatClientInformation(req)
		fn(writer, req, config)
	}
}

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
  // Get information from the environment variable 
  WebPort := getEnv(envWebPort, defaultWebPort)
  ListenIp := getEnv(envWebListenIp, defaultWebListenIp)

  // Declare web Handlers
	http.HandleFunc("/ip", makeHandler(ipHandler))
	http.HandleFunc("/all.json", makeHandler(jsonHandler))
	http.HandleFunc("/", makeHandler(mainHandler))
	http.ListenAndServe(ListenIp + ":" + WebPort, nil)
}