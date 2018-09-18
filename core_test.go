package twitchapi

import (
	"errors"
	"math"
	"testing"

	"github.com/konkers/mocktwitch"
)

func TestHttpBadUrlBase(t *testing.T) {

	api := NewConnection("", "")
	api.UrlBase = "gar\nbage://GARBAGE"

	_, err := api.GetChannel()
	if err == nil {
		t.Errorf("No error returned with UrlBase: %s.", api.UrlBase)
	}

	err = api.SetChannelGame("test", "test")
	if err == nil {
		t.Errorf("No error returned with UrlBase: %s.", api.UrlBase)
	}

	api.UrlBase = "GARBAGE"

	_, err = api.GetChannel()
	if err == nil {
		t.Errorf("No error returned with UrlBase: %s.", api.UrlBase)
	}

	err = api.SetChannelGame("test", "test")
	if err == nil {
		t.Errorf("No error returned with UrlBase: %s.", api.UrlBase)
	}
}
func TestHttpErrorFromServer(t *testing.T) {
	mock, err := mocktwitch.NewTwitch()
	if err != nil {
		t.Fatalf("Can't create mock twitch: %v.", err)
	}
	defer mock.Close()

	mock.ForceErrors = true
	api := NewConnection("", "")
	api.UrlBase = mock.ApiUrlBase
	api.VerboseLogging = true

	_, err = api.GetChannel()
	if err == nil {
		t.Errorf("No error returned with ForceErrors.")
	}

	err = api.SetChannelGame("test", "test")
	if err == nil {
		t.Errorf("No error returned with ForceErrors.")
	}
}

func TestHttpBadJSON(t *testing.T) {
	mock, err := mocktwitch.NewTwitch()
	if err != nil {
		t.Fatalf("Can't create mock twitch: %v.", err)
	}
	defer mock.Close()

	mock.ForceErrors = true
	api := NewConnection("", "")
	api.UrlBase = mock.ApiUrlBase

	err = api.put("channels", math.Inf(1))
	if err == nil {
		t.Errorf("No error returned when tring to put math.Inf(1).")
	}
}

type errReader struct{}

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestBodyReadError(t *testing.T) {
	c := &Connection{
		VerboseLogging: true,
	}
	err := c.decodeResponse(errReader{}, nil)
	if err == nil {
		t.Errorf("Read error not propagated in decodeResponse().")
	}
}
