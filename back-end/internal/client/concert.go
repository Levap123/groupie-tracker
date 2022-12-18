package client

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/back-end/internal/entity"
	"groupie-tracker/back-end/pkg/errors"
	"net/http"
)

type ConcertClient struct {
	client *http.Client
}

func NewConcertClient(client *http.Client) *ConcertClient {
	return &ConcertClient{
		client: client,
	}
}

func (c *ConcertClient) GetByGroupId(id int) (entity.Concert, error) {
	var concert entity.Concert
	resp, err := c.client.Get(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id))
	if err != nil {
		return entity.Concert{}, errors.Fail(err, "cleint-GetByGroupId")
	}
	if err := json.NewDecoder(resp.Body).Decode(&concert); err != nil {
		return entity.Concert{}, errors.Fail(err, "cleint-GetByGroupId")
	}
	return concert, nil
}
