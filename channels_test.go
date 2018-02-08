package twitch

import (
	"testing"
)

func TestGetChannel(t *testing.T) {
	_, err := client.Channels.GetChannel(config.AccessToken)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetChannelByID(t *testing.T) {
	_, err := client.Channels.GetChannelByID(AmouranthID)

	if err != nil {
		t.Errorf("%s", AmouranthID, err)
	}
}

func TestUpdateChannel(t *testing.T) {
	var options ChannelUpdateParams
	options.Game = "Overwatch"

	_, err := client.Channels.UpdateChannel(config.AccessToken, TserkovID, &options)

	if err != nil {
		t.Errorf("%s", err)

		return
	}
}

func TestGetChannelEditors(t *testing.T) {
	_, err := client.Channels.GetChannelEditors(config.AccessToken, TserkovID)

	if err != nil {
		t.Errorf("%s", TserkovID, err)
	}
}

func TestGetChannelFollowers(t *testing.T) {
	_, err := client.Channels.GetChannelFollowers(AmouranthID)

	if err != nil {
		t.Errorf("%s", AmouranthID, err)
	}
}

func TestGetChannelTeams(t *testing.T) {
	_, err := client.Channels.GetChannelTeams(TexcubsfID)

	if err != nil {
		t.Errorf("%s", TexcubsfID, err)
	}
}

func TestGetChannelSubscribers(t *testing.T) {
	_, err := client.Channels.GetChannelSubscribers(
		config.AccessToken,
		TserkovID,
		nil,
	)

	if err == nil {
		t.Error("Expected error about not having a subscription program; did not get...")
	}
}

func TestCheckChannelSubscriptionByUser(t *testing.T) {
	_, err := client.Channels.CheckChannelSubscriptionByUser(
		config.AccessToken,
		TserkovID,
		TserkovID,
	)

	if err == nil {
		t.Error("Expected error; did not get one")
	}
}

func TestGetChannelVideos(t *testing.T) {
	_, err := client.Channels.GetChannelVideos(
		TserkovID,
	)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestStartChannelCommercial(t *testing.T) {
	_, err := client.Channels.StartChannelCommercial(
		config.AccessToken,
		TserkovID,
		60,
	)

	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestResetChannelStreamKey(t *testing.T) {
	_, err := client.Channels.ResetChannelStreamKey(
		config.AccessToken,
		TserkovID,
	)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetChannelCommunities(t *testing.T) {
	_, err := client.Channels.GetChannelCommunities(
		TserkovID,
	)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestSetChannelCommunities(t *testing.T) {
	err := client.Channels.SetChannelCommunities(
		config.AccessToken,
		TserkovID,
		[]string{"fd0eab99-832a-4d7e-8cc0-04d73deb2e54"},
	)

	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeleteChannelFromCommunities(t *testing.T) {
	err := client.Channels.DeleteChannelFromCommunities(
		config.AccessToken,
		TserkovID,
	)

	if err != nil {
		t.Errorf("%s", err)
	}
}
