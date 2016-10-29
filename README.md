# slack api
[![CircleCI](https://circleci.com/gh/kwmt/slack.svg?style=svg&circle-token=1cca846f826acd4dcca75661aec5af4a2cbd1702)](https://circleci.com/gh/kwmt/slack)

Usage
-----

```go
c := slack.NewClient("<YOUR SLACK API TOKEN>")
c.SetChannel("#general").SetText("こんにちは").ChatPostMessage()
```

You can check <YOUR SLACK API TOKEN> from [here](https://api.slack.com/docs/oauth-test-tokens).

Install
-------

```
go get github.com/kwmt/slack
```
