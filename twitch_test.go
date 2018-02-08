package twitch

import (
	"encoding/json"
	"fmt"
	"os"
)

const TserkovID = "36038523"
const AmouranthID = "125387632"
const TexcubsfID = "127025184"

type Configuration struct {
	ClientId      string
	ClientSecret  string
	OauthRedirect string
	AccessToken   string
}

var client *TwitchClient
var config Configuration

func init() {
	file, _ := os.Open("config_test.json")

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)

	if err != nil {
		fmt.Println("Failed to load config: %s\n\n", err)

		os.Exit(-1)
	}

	client = NewClient(config.ClientId, config.ClientSecret, config.OauthRedirect)
}
