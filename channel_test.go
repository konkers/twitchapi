package twitchapi

import (
	"testing"

	"github.com/konkers/mocktwitch"
	"github.com/konkers/twitchapi/protocol"
)

func TestChannel(t *testing.T) {
	mock, err := mocktwitch.NewTwitch()
	if err != nil {
		t.Fatalf("Can't create mock twitch: %v.", err)
	}
	defer mock.Close()

	api := NewConnection("", "")
	api.UrlBase = mock.ApiUrlBase

	channel := protocol.Channel{
		Game: "Mega Man 2",
	}

	mock.SetChannelStatus(&channel)

	channel2, err := api.GetChannel()
	if err != nil {
		t.Fatalf("Can't get channel: %v.", err)
	}

	if channel.Game != channel2.Game {
		t.Errorf("Expected game == \"%s\", got \"%s\" instead.",
			channel.Game, channel2.Game)
	}

	newGame := "Mega Man 1"
	err = api.SetChannelGame("test", newGame)
	if err != nil {
		t.Fatalf("Can't set channel game: %v.", err)
	}

	channel3, err := api.GetChannel()
	if err != nil {
		t.Fatalf("Can't get channel: %v.", err)
	}

	if channel3.Game != newGame {
		t.Errorf("Expected game == \"%s\", got \"%s\" instead.",
			newGame, channel3.Game)
	}
}
