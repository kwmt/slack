# slack api
[![CircleCI](https://circleci.com/gh/kwmt/slack.svg?style=svg&circle-token=1cca846f826acd4dcca75661aec5af4a2cbd1702)](https://circleci.com/gh/kwmt/slack) [![GoDoc](https://godoc.org/github.com/kwmt/slack?status.svg)](http://godoc.org/github.com/kwmt/slack) 
Usage
-----

```go
attachments := slack.NewAttachments(1)
attachments[0].SetFallback("Required plain-text summary of the attachment.").SetTitle("タイトル")

json, err := slack.JSON(attachments)
if err != nil {
	log.Println(err)
}

c := slack.NewClient()
c.SetToken("YOUR TOKEN").SetChannel("#general").SetText("こんにちは").SetAttachments(json).PostMessage()
```

You can check the token for test from [here](https://api.slack.com/docs/oauth-test-tokens).

Install
-------

```
go get github.com/kwmt/slack
```
