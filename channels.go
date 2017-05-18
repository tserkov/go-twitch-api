package twitch

type Channel struct {
	Id int `json:"_id"`
	BroadcasterLanguage string `json:"broadcaster_language"`
	CreatedAt string `json:"created_at"`
	Displayname string `json:"display_name"`
	FollowerCount int `json:"followers"`
	Game string `json:"game"`
	Language string `json:"language"`
	Logo string `json:"logo"`
	IsMature bool `json:"mature"`
	Name string `json:"name"`
	IsPartner string `json:"partner"`
	ProfileBanner string `json:"profile_banner"`
	ProfileBannerBackgroundColor string `json:"profile_banner_background_color"`
	Status string `json:"status"`
	Url string `json:"url"`
	VideoBanner string `json:"url"`
	ViewCount int `json:"views"`
}

