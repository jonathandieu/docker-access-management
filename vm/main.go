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

	router.GET("/repositories/", c.GetRepository)

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
	User         string                 `json:"user,omitempty"`
	Repositories map[string]interface{} `json:"results"`
	MaxResults   int
}

func (c *Client) GetRepository(ctx echo.Context) error {
	repository := Repository{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/repositories/ryanhristovski/hackathon22/"), nil, &repository)
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

func (c *Client) CreateOrganization(ctx context.Context, orgname string, location string, company string) error {
	org := Org{
		OrgName:  orgname,
		Company:  company,
		Location: location,
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
