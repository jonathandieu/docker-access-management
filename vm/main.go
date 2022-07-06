package main

import (
	// "encoding/json"
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

	router.GET("/repo", repo)

	log.Fatal(router.Start(startURL))
}

func listen(path string) (net.Listener, error) {
	return net.Listen("unix", path)
}

// type ProxyData struct {
// 	Status int    `json:"status"`
// 	Data   string `json:"data"`
// }

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

// func proxy(ctx echo.Context) error {
// 	url := ctx.QueryParam("url")
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("Search-Version", "v3")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return ctx.JSON(http.StatusOK, ProxyData{
// 			Status: resp.StatusCode,
// 			Data:   err.Error(),
// 		})
// 	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return ctx.JSON(http.StatusOK, ProxyData{
// 			Status: 500,
// 			Data:   err.Error(),
// 		})
// 	}
// 	sb := string(body)
// 	return ctx.JSON(http.StatusOK, ProxyData{
// 		Status: resp.StatusCode,
// 		Data:   sb,
// 	})
// }

func repo (ctx echo.Context)  error {
	c := NewClient("https://hub-stage.docker.com", "ryanhristovski", "Hackathon2022")
	repositories := Repositories{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/u/%s/?page_size=%d", "ryanhristovski", 10), nil, &repositories)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, repositories)
}

type HTTPMessageBody struct {
	Message string
}
