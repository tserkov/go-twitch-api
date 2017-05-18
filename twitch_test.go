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

var client *TwitchClient

func init() {
	config  := Configuration{}

	file, _ := os.Open("config_test.json")

	decoder := json.NewDecoder(file)
	err     := decoder.Decode(&config)

	if err != nil {
		fmt.Println("Failed to load config: %s\n\n", err)

		os.Exit(-1)
	}

	client = NewClient(config.ClientId, config.ClientSecret, config.OauthRedirect)
}

func TestAccessToken(t *testing.T) {
	_, err := client.Users.GetAccessToken("fdsa")

	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestUserFollowedChannels(t *testing.T) {
	userFollows, err := client.Users.GetFollowedChannelInfo("28456583", "52722790")

	if err != nil {
		t.Errorf("Failed to get user follow info: %s", err)
	}

	if userFollows.Channel.Name != "toptsun" {
		t.Errorf("Wrong channel returned: %s", userFollows.Channel.Name)
	}
}
