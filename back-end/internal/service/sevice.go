package service

import (
	"groupie-tracker/back-end/internal/client"
	"groupie-tracker/back-end/internal/entity"
)

type Service struct {
	Group
	Concert
}

func NewService(client *client.Client) *Service {
	return &Service{
		Group:   NewGroupService(client.Group),
		Concert: NewConcertService(client.Concert),
	}
}

type Group interface {
	GetAll() ([]entity.GroupModel, error)
	GetById(id int) (entity.Group, error)
}

type Concert interface {
	GetByGroupId(id int) (entity.Concert, error)
}
