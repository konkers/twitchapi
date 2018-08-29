package twitchapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/konkers/twitchapi/protocol"
)

type Connection struct {
	clientID string
	oauth    string

	// UrlBase of the twitch API endpoint.  Defaults to
	// https://api.twitch.tv/kraken
	UrlBase string
}

func NewConnection(clientID string, oauth string) *Connection {
	return &Connection{
		UrlBase:  "https://api.twitch.tv/kraken",
		clientID: clientID,
		oauth:    oauth,
	}
}

func (c *Connection) getClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport)
	url, _ := url.Parse(c.UrlBase)

	// localhost is used for testing with a self signed certificate.
	if url.Hostname() == "localhost" {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	return &http.Client{
		Transport: transport,
	}
}

func (c *Connection) put(urlPath string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.UrlBase+"/"+urlPath,
		bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Add("Client-ID", c.clientID)
	req.Header.Add("Authorization", "OAuth "+c.oauth)
	req.Header.Set("Content-Type", "application/json")

	client := c.getClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non OK status code %d", resp.StatusCode)
	}

	return nil
}

func (c *Connection) get(urlPath string, data interface{}) error {
	req, err := http.NewRequest("GET", c.UrlBase+"/"+urlPath, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Client-ID", c.clientID)
	req.Header.Add("Authorization", "OAuth "+c.oauth)

	client := c.getClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *Connection) GetChannel() (*protocol.Channel, error) {
	var channel protocol.Channel
	err := c.get("channel", &channel)
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (c *Connection) SetChannelGame(channel string, game string) error {
	params := &protocol.Update{
		Channel: &protocol.ChannelUpdate{
			Game: &game,
		},
	}
	return c.put(path.Join("channels", channel), params)
}
