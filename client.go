package main

import (
	"io/ioutil"
	"net/http"
)

// Client represents a Meetup API v3 client.
type Client struct {
	APIKey string
	APIURL string

	client http.Client
}

// FindGroups searches for groups in Meetup.
func (c *Client) FindGroups(location, text string) (string, error) {
	r, err := http.NewRequest("GET", c.APIURL+"/find/groups", nil)
	if err != nil {
		return "", err
	}

	q := r.URL.Query()
	q.Add("key", c.APIKey)
	q.Add("location", location)
	q.Add("text", text)
	r.URL.RawQuery = q.Encode()

	response, err := c.client.Do(r)
	if err != nil {
		return "", err
	}

	groups, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(groups), nil
}

// NewClient creates a new Client with an API key.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		APIURL: "https://api.meetup.com",
	}
}
