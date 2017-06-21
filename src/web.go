package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"
)

func ipHandler(writer http.ResponseWriter, req *http.Request, clientConfig *webConfig) {
	fmt.Fprintf(writer, "%s", clientConfig.getIp())
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
	// If it is a curl request, we just give the ip
	if clientConfig.isCurl() == true {
		ipHandler(writer, req, clientConfig)
	} else {
		// Otherwise display a page with all the information
		renderTemplate(writer, clientConfig)
	}
}

/* Cache and render the main template when needed */
var templates = template.Must(template.ParseFiles("templates/main.html"))

func renderTemplate(writer http.ResponseWriter, clientConfig *webConfig) {
	err := templates.ExecuteTemplate(writer, "main.html", clientConfig)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

/* Parse the web request */
func formatClientInformation(req *http.Request) *webConfig {
	// Get the cient ip
	ip, port, _ := net.SplitHostPort(req.RemoteAddr)

	config := new(webConfig)
	config.Ip = ip
	config.Host = req.Host
	config.Port = port
	config.Method = req.Method
	config.Language = req.Header.Get("Accept-Language")
	config.UserAgent = req.Header.Get("User-Agent")
	config.XForwardedFor = req.Header.Get("X-Forwarded-For")

	return config
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *webConfig)) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		config := formatClientInformation(req)
		fn(writer, req, config)
	}
}
