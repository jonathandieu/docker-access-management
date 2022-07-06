package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

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

func (c *Client) GetOrganizations(ctx echo.Context) error {
	orgs := Orgs{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/users/%s/orgs/?page_size=%d", ctx.QueryParam("username"), ctx.QueryParam("max_results")), nil, &orgs)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, orgs)
}

func (c *Client) GetOrganization(ctx echo.Context) error {
	org := Org{}
	err := c.sendRequest(ctx.Request().Context(), "GET", fmt.Sprintf("/orgs/%s/", ctx.QueryParam("org_name"),), nil, &org)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, org)

}

// Given an orgname and company name, create a new Organization
func (c *Client) CreateOrganization(ctx echo.Context) error {
	org := Org{
		OrgName: ctx.QueryParam("org_name"),
		Company: ctx.QueryParam("company"),
	}
	orgJson, err := json.Marshal(org)
	if err != nil {
		return err
	}

	err = c.sendRequest(ctx.Request().Context(), "POST", fmt.Sprintf("/orgs/"), orgJson, &org)
	if err != nil {
		return err
	}
	
	return ctx.JSON(http.StatusOK, org)
}
