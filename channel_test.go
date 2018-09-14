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

	mock.ChannelFollows = protocol.ChannelFollows{
		Cursor: "1",
		Total:  2,
		Follows: []*protocol.ChannelFollower{
			&protocol.ChannelFollower{
				User: protocol.User{
					ID: 1,
				},
			},
			&protocol.ChannelFollower{
				User: protocol.User{
					ID: 2,
				},
			},
		},
	}

	mock.ForceErrors = true
	_, err = api.GetChannelFollows("test")
	mock.ForceErrors = false
	if err == nil {
		t.Error("No error when one expected.")
	}

	follows, err := api.GetChannelFollows("test")
	if err != nil {
		t.Fatalf("Can't get channel follows: %v.", err)
	}

	if follows.Total != 2 {
		t.Errorf("channel follows.Total %d != expected 2.", follows.Total)
	}
	if len(follows.Follows) != 2 {
		t.Fatalf("channel len(follows.Follows) %d != expected 2.", len(follows.Follows))
	}
	if follows.Follows[0].User.ID != 1 {
		t.Errorf("first follower ID (%d) != 1", follows.Follows[0].User.ID)
	}
	if follows.Follows[1].User.ID != 2 {
		t.Errorf("second follower ID (%d) != 2", follows.Follows[0].User.ID)
	}
}
