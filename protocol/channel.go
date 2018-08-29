package protocol

type Channel struct {
	Mature                       bool   `json:"mature"`
	Status                       string `json:"status"`
	BroadcasterLanguage          string `json:"broadcaster_language"`
	DisplayName                  string `json:"display_name"`
	Game                         string `json:"game"`
	Language                     string `json:"language"`
	ID                           int    `json:"_id"`
	Name                         string `json:"name"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	Partner                      bool   `json:"partner"`
	Logo                         string `json:"logo"`
	VideoBanner                  string `json:"video_banner"`
	ProfileBanner                string `json:"profile_banner"`
	ProfileBannerBackgroundColor string `json:"profile_banner_background_color"`
	URL                          string `json:"url"`
	Views                        int    `json:"views"`
	Followers                    int    `json:"followers"`
	BroadcasterType              string `json:"broadcaster_type"`
	StreamKey                    string `json:"stream_key"`
	Email                        string `json:"email"`
}

type ChannelUpdate struct {
	Status             *string `json:"status,omitempty"`
	Game               *string `json:"game,omitempty"`
	Delay              *string `json:"string,omitempty"`
	ChannelFeedEnabled *bool   `json:"channel_feed_enabled,omitempty"`
}

type Update struct {
	Channel *ChannelUpdate `json:"channel,omitempty"`
}
