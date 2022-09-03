package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Repository struct {
	User            string `json:"user,omitempty"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	Description     string `json:"description"`
	Private         bool   `json:"is_private"`
	PullCount       int    `json:"pull_count"`
	FullDescription string `json:"full_description,omitempty"`
}
type Repositories struct {
	User         string       `json:"user,omitempty"`
	Repositories []Repository `json:"results"`
	MaxResults   int
}

// , namespace string, maxResults int
func (c *Client) GetRepositories(ctx echo.Context) error {
	repositories := Repositories{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/repositories/%s/?page_size=%d", ctx.QueryParam("namespace"), ctx.QueryParam("max_results")), nil, &repositories)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, repositories)
}

// Given the repository's namespace and the repository's name, get the repository
func (c *Client) GetRepository(ctx echo.Context) error {
	repository := Repository{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/repositories/%s/%s/", ctx.QueryParam("namespace"), ctx.QueryParam("name")), nil, &repository)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, repository)
}

func (c *Client) CreateRepository(ctx echo.Context) error {
	repo := Repository{
		Name:      ctx.QueryParam("name"),
		Namespace: ctx.QueryParam("namespace"),
	}
	repoJson, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	err = c.sendRequest(ctx.Request().Context(), "POST", fmt.Sprintf("/repositories/"), repoJson, &repo)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, repo)
}

// Given context and an id (repository name), returns an error that is nil on success
func (c *Client) DeleteRepository(ctx echo.Context) error {
	err := c.sendRequest(ctx.Request().Context(), "DELETE", fmt.Sprintf("/repositories/%s/%s", ctx.QueryParam("namespace"), ctx.QueryParam("name")), nil, nil)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, "Deleted")
}