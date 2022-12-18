package service

import (
	"groupie-tracker/back-end/internal/client"
	"groupie-tracker/back-end/internal/entity"
)

type ConcertService struct {
	client client.Concert
}

func NewConcertService(client client.Concert) *ConcertService {
	return &ConcertService{
		client: client,
	}
}

func (cs *ConcertService) GetByGroupId(id int) (entity.Concert, error) {
	return cs.client.GetByGroupId(id)
}
