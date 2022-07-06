package main

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	var socketPath string
	flag.StringVar(&socketPath, "socket", "/run/guest/volumes-service.sock", "Unix domain socket to listen on")
	flag.Parse()

	os.RemoveAll(socketPath)

	logrus.New().Infof("Starting listening on %s\n", socketPath)
	router := echo.New()
	router.HideBanner = true

	startURL := ""

	ln, err := listen(socketPath)
	if err != nil {
		log.Fatal(err)
	}
	router.Listener = ln
	c := NewClient("https://hub-stage.docker.com/v2", "ryanhristovski", "Hackathon2022")

	router.GET("/repositories/", c.GetRepositories)   // landing page for repositories
	router.GET("/organizations/", c.GetOrganizations) // landing page for organizations

	log.Fatal(router.Start(startURL))
}

func listen(path string) (net.Listener, error) {
	return net.Listen("unix", path)
}

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

func (c *Client) GetRepositories(ctx context.Context, namespace string, maxResults int) (Repositories, error) {
	repositories := Repositories{
		User:       namespace,
		MaxResults: maxResults,
	}
	err := c.sendRequest(ctx, "GET", fmt.Sprintf("/repositories/%s/?page_size=%d", namespace, maxResults), nil, &repositories)
	return repositories, err
}

// Given the repository's namespace and the repository's name, get the repository
func (c *Client) GetRepository(ctx echo.Context, namespace string, name string) error {
	repository := Repository{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/repositories/%s/%s/", namespace, name), nil, &repository)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, repository)
}
func (c *Client) CreateRepository(ctx context.Context, name string, namespace string) error {
	repo := Repository{
		Name:      name,
		Namespace: namespace,
	}
	repoJson, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	err = c.sendRequest(ctx, "POST", fmt.Sprintf("/repositories/"), repoJson, &repo)
	if err != nil {
		return err
	}
	return nil
}

// Given context and an id (repository name), returns an error that is nil on success
func (c *Client) DeleteRepository(ctx context.Context, id string) error {
	return c.sendRequest(ctx, "DELETE", fmt.Sprintf("/repositories/%s/", id), nil, nil)
}

type Org struct {
	Id         string `json:"id,omitempty"`
	OrgName    string `json:"orgname"`
	FullName   string `json:"full_name"`
	Location   string `json:"location"`
	Company    string `json:"company"`
	DateJoined string `json:"date_joined"`
}

type Orgs struct {
	User       string `json:"user,omitempty"`
	Orgs       []Org  `json:"results"`
	MaxResults int
}

func (c *Client) GetOrganizations(ctx context.Context, username string, maxresults int) (Orgs, error) {
	orgs := Orgs{
		User:       username,
		MaxResults: maxresults,
	}
	err := c.sendRequest(ctx, "GET", fmt.Sprintf("/users/%s/orgs/?page_size=%d", username, maxresults), nil, &orgs)
	return orgs, err
}

func (c *Client) GetOrganization(ctx context.Context, orgname string) (Org, error) {
	org := Org{}
	err := c.sendRequest(ctx, "GET", fmt.Sprintf("/orgs/%s/", orgname), nil, &org)

	return org, err
}

// Given an orgname and company name, create a new Organization
func (c *Client) CreateOrganization(ctx context.Context, orgname string, company string) error {
	org := Org{
		OrgName: orgname,
		Company: company,
	}
	orgJson, err := json.Marshal(org)
	if err != nil {
		return err
	}

	err = c.sendRequest(ctx, "POST", fmt.Sprintf("/orgs/"), orgJson, &org)
	if err != nil {
		return err
	}

	fmt.Println(orgJson, org)

	return nil
}

type HTTPMessageBody struct {
	Message string
}
