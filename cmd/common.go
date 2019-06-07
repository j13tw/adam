package cmd

import (
	"regexp"
)

const (
	serverCertFilename = "server.pem"
	serverKeyFilename  = "server-key.pem"
	defaultDatabaseURL = "./run/adam"
	jsonContentType    = "application/json"
	textContentType    = "text/plain"
)

var (
	cn          string
	certPath    string
	keyPath     string
	hosts       string
	force       bool
	databaseURL string
)

func getFriendlyCN(cn string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\\.\\-]`)
	return re.ReplaceAllString(cn, "_")
}
