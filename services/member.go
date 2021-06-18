package services

import (
	"errors"
	"sort"
	"time"

	"members-club/models"
)

func New() *service {
	return &service{
		members: make(map[string]models.Member),
	}
}

type service struct {
	members map[string]models.Member
}

func (s *service) Add(email, name string) error {
	member := models.Member{
		ID:               len(s.members) + 1,
		Email:            email,
		Name:             name,
		RegistrationDate: time.Now().Format("02.01.2006"),
	}

	if _, ok := s.members[member.Email]; ok {
		return errors.New("member already exists")
	}

	s.members[member.Email] = member

	return nil
}

type ListResponse struct {
	Members []models.Member
}

func (s *service) List() ListResponse {
	var members []models.Member
	for _, member := range s.members {
		members = append(members, member)
	}

	sort.SliceStable(members, func(i, j int) bool {
		return members[i].ID < members[j].ID
	})

	return ListResponse{
		Members: members,
	}
}
