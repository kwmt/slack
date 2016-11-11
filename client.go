package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient httpRunner
	token      string
	values     url.Values
}

type httpRunner interface {
	Do(*http.Request) (*http.Response, error)
}

const (
	slack_api_base_url = "https://slack.com/api"
)

type Response struct {
	Ok      bool    `json:"ok,omitempty"`
	Channel string  `json:"channel,omitempty"`
	Ts      string  `json:"ts,omitempty"`
	Message Message `json:"message,omitempty"`
	Error   string  `json:"error,omitempty"`
}

type Message struct {
	Text     string `json:"text,omitempty"`
	Username string `json:"username,omitempty"`
	BotID    string `json:"bot_id,omitempty"`
	Type     string `json:"type,omitempty"`
	Subtype  string `json:"subtype,omitempty"`
	Ts       string `json:"ts,omitempty"`
}



func NewClient() *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: time.Duration(30) * time.Second},
		values:     url.Values{},
	}
	return c
}

func (c *Client) SetToken(token string) *Client {
	c.values.Set("token", token)
	c.token = token
	return c
}

func (c *Client) SetChannel(channel string) *Client {
	c.values.Set("channel", channel)
	return c
}
func (c *Client) SetText(text string) *Client {
	c.values.Set("text", text)
	return c
}
func (c *Client) SetUserName(name string) *Client {
	c.values.Set("username", name)
	return c
}
func (c *Client) SetIconEmoji(iconEmoji string) *Client {
	c.values.Set("icon_emoji", iconEmoji)
	return c
}

func (c *Client) PostMessage() (*Response, error) {
	return c.post("/chat.postMessage")
}

func (c *Client) post(method string) (*Response, error) {
	fmt.Println(c.values.Encode())

	req, err := http.NewRequest("POST", slack_api_base_url+method+"?"+c.values.Encode(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "text/plain")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return parse(resp)
}

func parse(resp *http.Response) (*Response, error) {
	var res Response
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
