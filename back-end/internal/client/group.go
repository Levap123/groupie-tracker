package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/back-end/internal/entity"
	"groupie-tracker/back-end/pkg/errors"
)

type GroupClient struct {
	client *http.Client
}

func NewGroupClient(client *http.Client) *GroupClient {
	return &GroupClient{
		client: client,
	}
}

func (gc *GroupClient) GetAll() ([]entity.GroupModel, error) {
	var groups []entity.GroupModel
	resp, err := gc.client.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, errors.Fail(err, "groups - Getall")
	}
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, errors.Fail(err, "groups - Getall")
	}
	return groups, nil
}

func (gc *GroupClient) GetById(id int) (entity.Group, error) {
	var group entity.Group
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id), nil)
	if err != nil {
		return entity.Group{}, errors.Fail(err, "groups - Getbyid")
	}
	resp, err := gc.client.Do(req)
	if err != nil {
		return entity.Group{}, errors.Fail(err, "groups - Getbyid")
	}
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return entity.Group{}, errors.Fail(err, "groups - Getbyid")
	}
	return group, nil
}
