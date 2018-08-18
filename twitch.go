package twitchapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/kr/pretty"
)

type Connection struct {
	clientID string
	oauth    string
	urlBase  string
}

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

func NewConnection(clientID string, oauth string) *Connection {
	return &Connection{
		clientID: clientID,
		oauth:    oauth,
	}
}

func (c *Connection) put(urlPath string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", "https://api.twitch.tv/kraken/"+urlPath,
		bytes.NewBuffer(b))
	log.Printf("%#v\n", pretty.Formatter(req))
	log.Printf("%#v\n", req.URL.String())
	log.Printf("%#v\n", string(b))
	if err != nil {
		return err
	}
	req.Header.Add("Client-ID", c.clientID)
	req.Header.Add("Authorization", "OAuth "+c.oauth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return err
}

func (c *Connection) get(urlPath string, data interface{}) error {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/kraken/"+urlPath, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Client-ID", c.clientID)
	req.Header.Add("Authorization", "OAuth "+c.oauth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}

func (c *Connection) GetChannel() (*Channel, error) {
	var channel Channel
	err := c.get("channel", &channel)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (c *Connection) SetChannelGame(id int, game string) error {
	params := &Update{
		Channel: &ChannelUpdate{
			Game: &game,
		},
	}
	return c.put(path.Join("channels", "djkonkers"), params)
}
