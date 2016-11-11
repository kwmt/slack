package slack

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPostChatMessage(t *testing.T) {
	const token = "token"

	test := `{
  "ok": true,
  "channel": "C0JNUD6T0",
  "ts": "1477724506.000004",
  "message": {
    "text": "こんにちは",
    "username": "Slack API Tester",
    "bot_id": "B19M5GGAH",
    "type": "message",
    "subtype": "bot_message",
    "ts": "1477724506.000004"
  }
}`

	tc := &testClient{
		resp: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(test))),
		},
		err: nil,
	}
	c := NewClient()
	c.httpClient = tc
	c.SetToken(token).SetChannel("#general").SetText("こんにちは")
	resp, err := c.PostMessage()

	if err != nil {
		t.Error(err)
	}
	if !resp.Ok {
		t.Errorf("ok: got %s, expect %s", resp.Ok, true)
	}

	if resp.Message.Text != "こんにちは" {
		t.Errorf("text: got %s, expect %s", resp.Message.Text, "こんにちは")
	}

}
// TODO:共通化
func TestPostChatMessageError(t *testing.T) {
	const token = "token"

	test := `{"ok":false,"error":"invalid_auth"}`

	tc := &testClient{
		resp: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(test))),
		},
		err: nil,
	}
	c := NewClient()
	c.httpClient = tc
	c.SetToken(token).SetChannel("#general").SetText("こんにちは")
	resp, err := c.PostMessage()

	if err != nil {
		t.Error(err)
	}
	if resp.Ok {
		t.Errorf("ok: got %s, expect %s", resp.Ok, false)
	}

	if resp.Error != "invalid_auth" {
		t.Errorf("text: got %s, expect %s", resp.Message.Text, "invalid_auth")
	}

}

type testClient struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (c *testClient) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return c.resp, c.err
}
