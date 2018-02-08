package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Channel struct {
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	BroadcasterType              string      `json:"broadcaster_type"`
	CreatedAt                    time.Time   `json:"created_at"`
	Description                  string      `json:"description"`
	DisplayName                  string      `json:"display_name"`
	Email                        string      `json:"email"`
	Followers                    int         `json:"followers"`
	Game                         string      `json:"game"`
	ID                           json.Number `json:"_id"`
	Language                     string      `json:"language"`
	Logo                         string      `json:"logo"`
	Mature                       bool        `json:"mature"`
	Name                         string      `json:"name"`
	ProfileBanner                string      `json:"profile_banner"`
	ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
	Partner                      bool        `json:"partner"`
	Status                       string      `json:"status"`
	StreamKey                    string      `json:"stream_key"`
	Url                          string      `json:"url"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	VideoBanner                  string      `json:"video_banner"`
	Views                        int         `json:"views"`
}

type ChannelUpdateParams struct {
	ChannelFeedEnabled bool   `json:"channel_feed_enabled,omitempty"`
	Game               string `json:"game,omitempty"`
	Delay              int    `json:"delay,omitempty"`
	Status             string `json:"status,omitempty"`
}

type ChannelFollowers struct {
	Cursor  string            `json:"_cursor"`
	Follows []ChannelFollower `json:"follows"`
	Total   int               `json:"_total"`
}

type ChannelFollower struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	User          User      `json:"user"`
}

type ChannelSubscribersParams struct {
	Direction int    `json:"direction,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Offset    string `json:"offset,omitempty"`
}

type ChannelTeams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	Banner      string    `json:"banner"`
	Background  string    `json:"background"`
	CreatedAt   time.Time `json:"created_at"`
	DisplayName string    `json:"display_name"`
	ID          int       `json:"_id"`
	Info        string    `json:"info"`
	Logo        string    `json:"logo"`
	Name        string    `json:"name"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ChannelSubscriptions struct {
	Subscriptions []ChannelSubscription `json:"subscriptions"`
	Total         int                   `json:"_total"`
}

type ChannelSubscription struct {
	CreatedAt   time.Time `json:"created_at"`
	ID          string    `json:"_id"`
	SubPlan     string    `json:"sub_plan"`
	SubPlanName string    `json:"sub_plan_name"`
	User        User      `json:"user"`
}

type Videos struct {
	Total  int     `json:"_total"`
	Videos []Video `json:"videos"`
}

type Video struct {
	AnimatedPreviewUrl string           `json:"animated_preview_url"`
	BroadcastID        int              `json:"broadcast_id"`
	BroadcastType      string           `json:"broadcast_type"`
	Communities        string           `json:"communities"`
	CreatedAt          time.Time        `json:"created_at"`
	Channel            Channel          `json:"channel"`
	DescriptionHtml    string           `json:"description_html"`
	Description        string           `json:"description"`
	Fps                VideoFPS         `json:"fps"`
	Game               string           `json:"game"`
	ID                 string           `json:"_id"`
	Language           string           `json:"language"`
	Length             int              `json:"length"`
	Preview            VideoPreviews    `json:"preview"`
	PublishedAt        time.Time        `json:"published_at"`
	RecordedAt         time.Time        `json:"recorded_at"`
	Resolutions        VideoResolutions `json:"resolutions"`
	Status             string           `json:"status"`
	TagList            string           `json:"tag_list"`
	Thumbnails         VideoThumbnails  `json:"thumbnails"`
	Title              string           `json:"title"`
	Url                string           `json:"url"`
	Viewable           string           `json:"viewable"`
	Views              int              `json:"views"`
	ViewableAt         time.Time        `json:"viewable_at"`
}

type VideoFPS struct {
	Chunked float32 `json:"chunked"`
	High    float32 `json:"high"`
	Low     float32 `json:"low"`
	Medium  float32 `json:"medium"`
	Mobile  float32 `json:"mobile"`
}

type VideoPreviews struct {
	Large    string `json:"large"`
	Medium   string `json:"medium"`
	Small    string `json:"small"`
	Template string `json:"template"`
}

type VideoResolutions struct {
	Chunked string `json:"chunked"`
	High    string `json:"high"`
	Low     string `json:"low"`
	Medium  string `json:"medium"`
	Mobile  string `json:"mobile"`
}

type VideoThumbnails struct {
	Large    []VideoThumbnail `json:"large"`
	Medium   []VideoThumbnail `json:"medium"`
	Small    []VideoThumbnail `json:"small"`
	Template []VideoThumbnail `json:"template"`
}

type VideoThumbnail struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type CommercialResponse struct {
	Length     int    `json:"Length"`
	Message    string `json:"Message"`
	Retryafter int    `json:"RetryAfter"`
}

type Communities struct {
	Communities []Community `json:"communities"`
}

type Community struct {
	AvatarImageUrl  string `json:"avatar_image_url"`
	CoverImageUrl   string `json:"cover_image_url"`
	Description     string `json:"description"`
	DescriptionHtml string `json:"description_html"`
	DisplayName     string `json:"display_name"`
	ID              string `json:"_id"`
	Language        string `json:"language"`
	Name            string `json:"name"`
	OwnerID         string `json:"owner_id"`
	Rules           string `json:"rules"`
	RulesHtml       string `json:"rules_html"`
	Summary         string `json:"summary"`
}

type ChannelsService struct {
	client *TwitchClient
}

// Gets a channel object based on a specified OAuth token
// Returns more data than GetChannelByID because this is privileged
// Required scopes: channel_read
func (s *ChannelsService) GetChannel(accessToken string) (Channel, error) {
	var channel Channel

	s.client.setAccessToken(accessToken)

	err := s.client.request(
		"GET",
		"channel",
		nil,
		"",
		&channel,
	)

	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}

// Gets a specified channel object
// Required scopes: none
func (s *ChannelsService) GetChannelByID(channelID string) (Channel, error) {
	var channel Channel

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s", channelID),
		nil,
		"",
		&channel,
	)

	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}

// Updates specified properties of a specified channel
// Required scopes: channel_editor
func (s *ChannelsService) UpdateChannel(accessToken string, channelID string, params *ChannelUpdateParams) (Channel, error) {
	var channel Channel

	s.client.setAccessToken(accessToken)

	var body []byte

	if params != nil {
		body, _ = json.Marshal(params)
	}

	err := s.client.request(
		"PUT",
		fmt.Sprintf("channels/%s", channelID),
		nil,
		strings.Join([]string{"{\"channel\":", string(body), "}"}, ""),
		&channel,
	)

	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}

// Gets a list of users who are editors for a specified channel
// Required scopes: channel_read
func (s *ChannelsService) GetChannelEditors(accessToken string, channelID string) (Users, error) {
	var users Users

	s.client.setAccessToken(accessToken)

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/editors", channelID),
		nil,
		"",
		&users,
	)

	if err != nil {
		return Users{}, nil
	}

	return users, nil
}

// Gets a list of users who follow a specified channel, sorted by the date when they started following the channel, up to 25 at a time
// Required scopes: none
func (s *ChannelsService) GetChannelFollowers(channelID string) (ChannelFollowers, error) {
	var channelFollowers ChannelFollowers

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/follows", channelID),
		nil,
		"",
		&channelFollowers,
	)

	if err != nil {
		return ChannelFollowers{}, err
	}

	return channelFollowers, nil
}

// Gets a list of teams to which a specified channel belongs
// Required scopes: none
func (s *ChannelsService) GetChannelTeams(channelID string) (ChannelTeams, error) {
	var channelTeams ChannelTeams

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/teams", channelID),
		nil,
		"",
		&channelTeams,
	)

	if err != nil {
		return ChannelTeams{}, err
	}

	return channelTeams, nil
}

// Gets a list of users subscribed to a specific channel, sortedb y the date when the subscribed, up to 25 at a time
// Required scopes: channel_subscriptions
func (s *ChannelsService) GetChannelSubscribers(accessToken string, channelID string, params *ChannelSubscribersParams) (ChannelSubscriptions, error) {
	var channelSubscribers ChannelSubscriptions

	s.client.setAccessToken(accessToken)

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/subscriptions", channelID),
		nil,
		"",
		&channelSubscribers,
	)

	if err != nil {
		return ChannelSubscriptions{}, err
	}

	return channelSubscribers, nil
}

// Checks if a specified channel has a specified user subscribed to it.
// Required scope: channel_check_subscription
func (s *ChannelsService) CheckChannelSubscriptionByUser(accessToken string, channelID string, userID string) (ChannelSubscription, error) {
	var subscription ChannelSubscription

	s.client.setAccessToken(accessToken)

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/subscriptions/%s", channelID, userID),
		nil,
		"",
		&subscription,
	)

	if err != nil {
		return ChannelSubscription{}, err
	}

	return subscription, nil
}

// Gets a list of VODs (Video on Demand) from a specified channel.
// Required scope: none
func (s *ChannelsService) GetChannelVideos(channelID string) (Videos, error) {
	var videos Videos

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/videos", channelID),
		nil,
		"",
		&videos,
	)

	if err != nil {
		return Videos{}, err
	}

	return videos, nil
}

// Starts a commercial (advertisement) on a specified channel.
// This is valid only for channels that are Twitch partnes. You cannot start a commercial more often than once every 8 minutes.
// Valid length values: 0, 60, 90, 120, 150, 180
// Required scope: channel_commercial
func (s *ChannelsService) StartChannelCommercial(accessToken string, channelID string, length int) (CommercialResponse, error) {
	var commercialResponse CommercialResponse

	s.client.setAccessToken(accessToken)

	p := url.Values{}
	p.Set("length", string(length))

	err := s.client.request(
		"POST",
		fmt.Sprintf("channels/%s/commercial", channelID),
		nil,
		strings.Join([]string{"{\"length\":", strconv.Itoa(length), "}"}, ""),
		&commercialResponse,
	)

	if err != nil {
		return CommercialResponse{}, err
	}

	return commercialResponse, nil
}

// Deletes the stream key for a specified channel. Once it is deleted, the stream key is automatically reset.
// Required scope: channel_stream
func (s *ChannelsService) ResetChannelStreamKey(accessToken string, channelID string) (Channel, error) {
	var channel Channel

	s.client.setAccessToken(accessToken)

	err := s.client.request(
		"DELETE",
		fmt.Sprintf("channels/%s/stream_key", channelID),
		nil,
		"",
		&channel,
	)

	if err != nil {
		return Channel{}, err
	}

	return channel, nil
}

// Gets the communities for a specified channel.
// Required scope: none
func (s *ChannelsService) GetChannelCommunities(channelID string) (Communities, error) {
	var communities Communities

	err := s.client.request(
		"GET",
		fmt.Sprintf("channels/%s/communities", channelID),
		nil,
		"",
		&communities,
	)

	if err != nil {
		return Communities{}, err
	}

	return communities, nil
}

// Sets a specified channel to be in up to three specified communities. The list of community IDs is specified in the body of the request.
// Required scope: channel_editor
func (s *ChannelsService) SetChannelCommunities(accessToken string, channelID string, communities []string) error {
	s.client.setAccessToken(accessToken)

	body, _ := json.Marshal(communities)

	err := s.client.request(
		"PUT",
		fmt.Sprintf("channels/%s/communities", channelID),
		nil,
		strings.Join([]string{"{\"community_ids\":", string(body), "}"}, ""),
		nil,
	)

	if err != nil {
		if err.Error() == "EOF" {
			return nil
		}

		return err
	}

	return errors.New("Unexpected response from Twitch")
}

// Deletes a specified channel from its communities.
// Required scope: channel_editor
func (s *ChannelsService) DeleteChannelFromCommunities(accessToken string, channelID string) error {
	s.client.setAccessToken(accessToken)

	err := s.client.request(
		"DELETE",
		fmt.Sprintf("channels/%s/community", channelID),
		nil,
		"",
		nil,
	)

	if err != nil {
		if err.Error() == "EOF" {
			return nil
		}

		return err
	}

	return errors.New("Unexpected response from Twitch")
}
