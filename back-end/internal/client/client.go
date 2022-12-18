package client

import (
	"net/http"

	"groupie-tracker/back-end/internal/entity"
)

type Client struct {
	Group
	Concert
}

func NewClient(client *http.Client) *Client {
	return &Client{
		Group:   NewGroupClient(client),
		Concert: NewConcertClient(client),
	}
}

type Group interface {
	GetAll() ([]entity.GroupModel, error)
	GetById(id int) (entity.Group, error)
}

type Concert interface {
	GetByGroupId(id int) (entity.Concert, error)
}
