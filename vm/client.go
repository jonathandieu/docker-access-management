package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Client struct {
	BaseURL    string
	auth       Auth
	HTTPClient *http.Client
}

// Create the API client, providing the authentication.
func NewClient(hub_host string, username string, password string) *Client {
	return &Client{
		BaseURL: hub_host,
		auth: Auth{
			Username: username,
			Password: password,
		},
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type CreatePersonalAccessToken struct {
	TokenLabel string   `json:"token_label"`
	Scopes     []string `json:"scopes"`
}

type CreatePersonalAccessTokenResponse struct {
	UUID       string   `json:"uuid"`
	Token      string   `json:"token"`
	TokenLabel string   `json:"token_label"`
	Scopes     []string `json:"scopes"`
}

func (c *Client) DeletePersonalAccessToken(ctx context.Context, uuid string) error {
	return c.sendRequest(ctx, "DELETE", fmt.Sprintf("/access-tokens/%s", uuid), nil, nil)
}

func (c *Client) sendRequest(ctx context.Context, method string, url string, body []byte, result interface{}) error {

	authJson, err := json.Marshal(c.auth)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v2/users/login/", c.BaseURL), bytes.NewBuffer(authJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyBytes))
	}
	token := Token{}
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return err
	}

	req, err = http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, url), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", token.Token))

	req = req.WithContext(ctx)

	res, err = c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyBytes))
	}

	if result != nil {
		if err = json.NewDecoder(res.Body).Decode(result); err != nil {
			return err
		}
	}

	return nil
}