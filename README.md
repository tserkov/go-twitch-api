# Twitch API (Golang)

[![Build Status](https://travis-ci.org/tserkov/go-twitch-api.svg?branch=master)](https://travis-ci.com/tserkov/go-twitch-api)

This is a Go wrapper for the [Twitch.tv Kraken API](https://dev.twitch.tv/docs), initially created for an internal project, but am making public as I build it out to our requirements.

### Usage
```go
package main

import (
	"fmt"

	"github.com/tserkov/go-twitch-api"
)

func main() {
	t := twitch.NewClient("CLIENT_ID", "CLIENT_SECRET", "CLIENT_REDIRECT")

	u, err := t.Users.GetByLogin("tserkov")

	if err != nil {
		fmt.Printf("Failed to get users: %s", err)

		return
	}

	if u.Total == 0 {
		fmt.Printf("No users found with login/username tserkov")

		return
	}

	tserkov := u.Users[0]

	fmt.Printf("User ID for %s is %s", tserkov.DisplayName, tserkov.Id)
}
```

### Users
* GetAccessToken (authorizationCode _string_) _string_
	- Exchanges an authorization code (which you somehow obtained from Twitch's OAuth flow) for an access token.
* GetAuthenticated (accessToken _string_) _twitch.User_, _error_
	- Gets the user associated with the provided access token.
* GetByLogin (username _string_) _twitch.Users_, _error_
	- Returns users matching the provided username/login. Useful for converting username/login to user id.
* GetFollowedChannelInfo (followerId _string_, channelId _string_) _twitch.Follows_, _error_
	- If the provided user id is following the provided channel id, returns the related follow info.

### Todo
* Channels endpoints
* Chat endpoints
* Games endpoints
* Ingests endpoints
* Search endpoints
* Streams endpoints
* Teams endpoints
* Users endpoints
* Videos endpoints

Yea... all the endpoints.