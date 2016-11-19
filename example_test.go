package slack_test

import (
	"log"

	"github.com/kwmt/slack"
)

func Example() {
	c := slack.NewClient()
	c.SetToken("YOUR TOKEN").SetChannel("#general").SetText("こんにちは").PostMessage()
}

func ExampleAttachment() {
	// https://api.slack.com/docs/message-attachments
	//    [
	//        {
	//            "fallback": "Required plain-text summary of the attachment.",
	//            "color": "good",
	//            "pretext": "Optional text that appears above the attachment block",
	//            "author_name": "Bobby Tables",
	//            "author_link": "http://flickr.com/bobby/",
	//            "author_icon": "http://flickr.com/icons/bobby.jpg",
	//            "title": "Slack API Documentation",
	//            "title_link": "https://api.slack.com/",
	//            "text": "Optional text that appears within the attachment",
	//            "fields": [
	//                {
	//                    "title": "Priority",
	//                    "value": "High",
	//                    "short": false
	//                }
	//            ],
	//            "image_url": "http://my-website.com/path/to/image.jpg",
	//            "thumb_url": "http://example.com/path/to/thumb.png",
	//            "footer": "Slack API",
	//            "footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
	//            "ts": 123456789
	//        }
	//    ]

	fields := slack.NewFields(1)
	fields[0].SetTitle("Priority").SetValue("High").SetShort(false)

	attachments := slack.NewAttachments(1)
	attachments[0].SetFallback("Required plain-text summary of the attachment.").SetColor(slack.COLOR_GOOD).SetPretext(
		"Optional text that appears above the attachment block",
	).SetAuthorName("Bobby Tables").SetAuthorLink("http://flickr.com/bobby/").SetAuthorIcon("http://flickr.com/icons/bobby.jpg").SetTitle(
		"Slack API Documentatio",
	).SetTitleLink("https://api.slack.com/").SetText("Optional text that appears within the attachment").SetFields(fields).SetImageURL(
		"http://my-website.com/path/to/image.jpg",
	).SetThumbURL("http://example.com/path/to/thumb.png").SetFooterText("Slack API").SetFooterIcon(
		"https://platform.slack-edge.com/img/default_application_icon.png",
	).SetFooterTs(123456789)

	json, err := slack.JSON(attachments)
	if err != nil {
		log.Println(err)
	}

	c := slack.NewClient()
	c.SetToken("YOUR TOKEN").SetChannel("#general").SetText("こんにちは").SetAttachment(json).PostMessage()
}
