package slack_test

import (
	"github.com/kwmt/slack"
)

func ExampleChatPostMessage() {

	c := slack.NewClient("<TOKEN>")
	c.SetChannel("#general").SetText("こんにちは").ChatPostMessage()
}
