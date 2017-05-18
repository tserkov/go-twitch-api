package twitch

import (
	"fmt"
	"net/url"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope []string `json:"scope"`
}

type UsersService struct {
	client *TwitchClient
}

type User struct {
	Id string `json:"_id"`
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	Email string `json:"email"`
	IsEmailVerified bool `json:"email_verified"`
	Logo string `json:"logo"`
	IsPartnered bool `json:"partnered"`
	IsTwitterConnected bool `json:"twitter_connected"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Follows struct {
	CreatedAt string `json:"created_id"`
	Notifications bool `json:"notifications"`
	Channel Channel
}

func (u *UsersService) GetAccessToken (code string) (string, error) {
	p := url.Values{}
	p.Set("grant_type", "authorization_code")
	p.Set("client_id", u.client.clientId)
	p.Set("client_secret", u.client.clientSecret)
	p.Set("redirect_uri", u.client.oauthRedirect)
	p.Set("code", code)

	var tokenResponse TokenResponse

	err := u.client.request("POST", "oauth2/token", p, &tokenResponse)

	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func (u *UsersService) GetAuthenticated (accessToken string) (User, error) {
	var userResponse User

	u.client.setAccessToken(accessToken)

	err := u.client.request(
		"GET",
		"user",
		nil,
		&userResponse,
	)

	if err != nil {
		return User{}, err
	}

	return userResponse, nil
}

func (u *UsersService) GetByLogin (username string) (User, error) {
	var userResponse User

	err := u.client.request(
		"GET",
		"users",
		url.Values{
			"login": []string{username},
		},
		&userResponse,
	)

	if err != nil {
		return User{}, err
	}

	return userResponse, nil
}

func (u *UsersService) GetFollowedChannelInfo (followerId string, channelId string) (Follows, error) {
	var followsResponse Follows

	err := u.client.request(
		"GET",
		fmt.Sprintf("users/%s/follows/channels/%s", followerId, channelId),
		nil,
		&followsResponse,
	)

	if err != nil {
		return Follows{}, err
	}

	return followsResponse, nil
}
