package slack_test

import (
	"github.com/kwmt/slack"
)

func Example() {

	c := slack.NewClient("TOKEN")
	c.SetChannel("#general").SetText("こんにちは").ChatPostMessage()
}
