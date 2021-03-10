package tipeee

import (
	"crypto/tls"
	"net/http"
	"time"
)

// ClientAPI is a structure that can be used to
// communicate with Tipeee's API
type Client struct {
	accessToken string
	httpClient  *http.Client
}

// ClientWithToken returns a Client initialized with OAuth2 access token
func ClientWithToken(accessToken string) *Client {
	clientAPI := &Client{
		accessToken: accessToken,
	}
	clientAPI.initHttpClient()
	return clientAPI
}

// utils

func (c *Client) initHttpClient() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.httpClient = &http.Client{Transport: transport, Timeout: time.Second * 30}
}
