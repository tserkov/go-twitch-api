package twitch

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Configuration struct {
	ClientId      string
	ClientSecret  string
	OauthRedirect string
}

var config Configuration

func TestMain(m *testing.M) {
	file, _ := os.Open("config_test.json")
	decoder := json.NewDecoder(file)
	config  = Configuration{}
	err     := decoder.Decode(&config)

	if err != nil {
		fmt.Println("Failed to load config: %s\n\n", err)

		os.Exit(-1)
	} else {
		os.Exit(m.Run())
	}
}

func TestAccessToken(t *testing.T) {
	twitch := NewClient(config.ClientId, config.ClientSecret, config.OauthRedirect)

	accessToken, err := twitch.Users.GetAccessToken("fdsa")

	if err != nil {
		t.Errorf("Failed to get access token: %s", err)
	} else {
		t.Logf("access token is %s", accessToken)
	}
}