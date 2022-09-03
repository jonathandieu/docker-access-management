package main

import (
	"flag"
	"log"
	"net"
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
	c := NewClient("https://hub-stage.docker.com/v2", "chefjon", "Hackathon2022")

	// Repos routes
	router.GET("/repositories", c.GetRepositories)
	router.GET("/repository", c.GetRepository)
	router.POST("/repository", c.CreateRepository)
	router.DELETE("/repository", c.DeleteRepository)

	// Orgs routes
	router.GET("/organizations", c.GetOrganizations)
	router.GET("/organization", c.GetOrganization)
	router.POST("/organization", c.CreateOrganization)

	log.Fatal(router.Start(startURL))
}

func listen(path string) (net.Listener, error) {
	return net.Listen("unix", path)
}
