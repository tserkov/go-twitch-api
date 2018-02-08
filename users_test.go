package twitch

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	_, err := client.Users.GetAccessToken("asdf")

	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}

func TestGetUser(t *testing.T) {
	_, err := client.Users.GetUser(config.AccessToken)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetUserByID(t *testing.T) {
	_, err := client.Users.GetUserByID(TserkovID)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetUsers(t *testing.T) {
	_, err := client.Users.GetUsers([]string{"tserkov"})

	if err != nil {
		t.Errorf("%s", err)

		return
	}
}

func TestGetFollowedChannelInfo(t *testing.T) {
	_, err := client.Users.GetFollowedChannelInfo("28456583", "52722790")

	if err != nil {
		t.Errorf("%s", err)
	}
}
