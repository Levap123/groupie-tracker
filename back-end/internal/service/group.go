package service

import (
	"groupie-tracker/back-end/internal/client"
	"groupie-tracker/back-end/internal/entity"
)

type GroupService struct {
	client client.Group
}

func NewGroupService(groupClient client.Group) *GroupService {
	return &GroupService{
		client: groupClient,
	}
}

func (gs *GroupService) GetAll() ([]entity.GroupModel, error) {
	return gs.client.GetAll()
}

func (gs *GroupService) GetById(id int) (entity.Group, error) {
	return gs.client.GetById(id)
}
