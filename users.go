package twitch

import (
	"fmt"
	"net/url"
	"time"
)

type TokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
}

type UsersService struct {
	client *TwitchClient
}

type User struct {
	ID                 string    `json:"_id"`
	Name               string    `json:"name"`
	DisplayName        string    `json:"display_name"`
	Email              string    `json:"email"`
	IsEmailVerified    bool      `json:"email_verified"`
	Logo               string    `json:"logo"`
	Bio                string    `json:"bio"`
	IsPartnered        bool      `json:"partnered"`
	IsTwitterConnected bool      `json:"twitter_connected"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Users struct {
	Total int    `json:"_total"`
	Users []User `json:"users"`
}

type Follows struct {
	CreatedAt     time.Time `json:"created_id"`
	Notifications bool      `json:"notifications"`
	Channel       Channel
}

func (u *UsersService) GetAccessToken(code string) (string, error) {
	var tokenResponse TokenResponse

	err := u.client.request(
		"POST",
		"oauth2/token",
		url.Values{
			"grant_type":    []string{"authorization_code"},
			"clientId":      []string{u.client.clientId},
			"clientSecret":  []string{u.client.clientSecret},
			"oauthRedirect": []string{u.client.oauthRedirect},
			"code":          []string{code},
		},
		"",
		&tokenResponse,
	)

	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

// Gets a user object based on the OAuth token provided.
// Required scopes: user_read
func (u *UsersService) GetUser(accessToken string) (User, error) {
	var user User

	u.client.setAccessToken(accessToken)

	err := u.client.request(
		"GET",
		"user",
		nil,
		"",
		&user,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Gets a specified user object.
// Required scopes: none
func (u *UsersService) GetUserByID(userID string) (User, error) {
	var user User

	err := u.client.request(
		"GET",
		fmt.Sprintf("users/%s", userID),
		nil,
		"",
		&user,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Gets the user objects for the specified Twitch login names (up to 100).
// If a specified user’s Twitch-registered email address is not verified, null is returned for that user.
// Required scopes: none
func (u *UsersService) GetUsers(usernames []string) (Users, error) {
	var users Users

	err := u.client.request(
		"GET",
		"users",
		url.Values{
			"login": usernames,
		},
		"",
		&users,
	)

	if err != nil {
		return Users{}, err
	}

	return users, nil
}

func (u *UsersService) GetFollowedChannelInfo(followerId string, channelId string) (Follows, error) {
	var follows Follows

	err := u.client.request(
		"GET",
		fmt.Sprintf("users/%s/follows/channels/%s", followerId, channelId),
		nil,
		"",
		&follows,
	)

	if err != nil {
		return Follows{}, err
	}

	return follows, nil
}
