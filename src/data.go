package main

import (
	"strings"
)

type webConfig struct {
	Ip            string
	Host          string
	Port          string
	Method        string
	Language      string
	UserAgent     string
	XForwardedFor string
}

/* Struct helpers */
func (config *webConfig) isCurl() bool {
	return strings.HasPrefix(config.UserAgent, "curl/")
}

func (config *webConfig) getIp() string {
	if len(config.XForwardedFor) == 0 {
		return config.Ip
	}
	return config.XForwardedFor
}
