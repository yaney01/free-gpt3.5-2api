package config

import (
	"free-gpt3.5-2api/common"
	"github.com/joho/godotenv"
	"net/url"
	"os"
	"strings"
)

type config struct {
	LogLevel   string
	Bind       string
	Port       string
	Proxy      *url.URL
	AuthTokens []string
}

var CONFIG *config

func init() {
	_ = godotenv.Load()
	CONFIG = &config{}
	// Bind
	CONFIG.Bind = os.Getenv("BIND")
	if CONFIG.Bind == "" {
		CONFIG.Bind = "127.0.0.1"
	}
	// PORT
	CONFIG.Port = os.Getenv("PORT")
	if CONFIG.Port == "" {
		CONFIG.Port = "3040"
	}
	// PROXY
	proxy := os.Getenv("PROXY")
	if proxy == "" {
		CONFIG.Proxy = nil
	} else {
		CONFIG.Proxy = common.ParseUrl(proxy)
	}
	// AUTH_TOKEN
	authTokens := os.Getenv("AUTH_TOKENS")
	if authTokens == "" {
		CONFIG.AuthTokens = []string{"1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	} else {
		//以,分割 AUTH_TOKEN
		CONFIG.AuthTokens = strings.Split(authTokens, ",")
	}
}
