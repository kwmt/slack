package slack_test

import (
	"github.com/kwmt/slack"
)

func Example() {

	c := slack.NewClient()
	c.SetToken("YOUR TOKEN").SetChannel("#general").SetText("こんにちは").PostMessage()
}
