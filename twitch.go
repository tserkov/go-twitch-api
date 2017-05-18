package twitch

import (
	"errors"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"

	//"io/ioutil"
)

const (
	twitchHost   = "api.twitch.tv"
	twitchPath   = "kraken"
	acceptHeader = "application/vnd.twitchtv.v5+json"
)

type TwitchClient struct {
	httpClient    *http.Client

	baseUrl       *url.URL

	clientId      string
	clientSecret  string
	oauthRedirect string

	accessToken   string

//	Channels   *ChannelsService
//	Chat       *ChatService
//	Games      *GamesService
//	Ingests    *IngestsService
//	Search     *SearchService
//	Streams    *StreamsService
//	Teams      *TeamsService
	Users      *UsersService
//	Videos     *VideosService
}

type TwitchError struct {
	Error string `json:"error"`
	Status int `json:"status"`
	Message string `json:"message"`
}

func NewClient(clientId, clientSecret, oauthRedirect string) (*TwitchClient) {
	baseUrl, _ := url.Parse("https://api.twitch.tv/kraken/")

	c := &TwitchClient{
		httpClient:    http.DefaultClient,
		baseUrl:       baseUrl,
		clientId:      clientId,
		clientSecret:  clientSecret,
		oauthRedirect: oauthRedirect,
	}

//	c.Channels = &ChannelsService{client: c}
//	c.Chat = &ChatService{client: c}
//	c.Games = &GamesService{client: c}
//	c.Ingests = &IngestsService{client: c}
//	c.Search = &SearchService{client: c}
//	c.Streams = &StreamsService{client: c}
//	c.Teams = &TeamsService{client: c}
	c.Users = &UsersService{client: c}
//	c.Videos = &VideosService{client: c}

	return c
}

func (c *TwitchClient) setAccessToken (code string) {
	c.accessToken = fmt.Sprintf("OAuth %s", code)
}

func (c *TwitchClient) createRequest(method, endpoint string, params url.Values) *http.Request {
	url := fmt.Sprintf("https://%s/%s/%s", twitchHost, twitchPath, endpoint)

	r, _ := http.NewRequest(method, url, nil)

	r.URL.RawQuery = params.Encode()
	r.Header.Set("Accept", acceptHeader)
	r.Header.Set("Client-ID", c.clientId)

	if c.accessToken != "" {
		r.Header.Set("Authorization", c.accessToken)
	}

	return r
}

func (c *TwitchClient) request(method, endpoint string, params url.Values, v interface{}) error {
	// Create the request
	req := c.createRequest(method, endpoint, params)

	// Perform the request
	res, err := c.httpClient.Do(req)

	if err != nil {
		return errors.New("twitch_unavailable")
	}

	defer res.Body.Close()

	// if endpoint == "user" {
	// 	robots, _ := ioutil.ReadAll(res.Body)
	// 	res.Body.Close()
	// 	fmt.Printf("%s", robots)
	// }

	decoder := json.NewDecoder(res.Body)

	// Check if request was bad
	if res.StatusCode >= 300 {
		var twitchError TwitchError

		decoder.Decode(&twitchError)

		return errors.New(twitchError.Message)
	}

	err = decoder.Decode(v)

	if err != nil {
		fmt.Println("%s", err)
		return errors.New("invalid_response")
	}

	return nil
}
