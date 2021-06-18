package services_test

import (
	"github.com/stretchr/testify/assert"
	"members-club/services"
	"testing"
)

func TestAddMember(t *testing.T) {
	members := services.New()

	resp := members.List()
	assert.Equal(t, len(resp.Members), 0)

	err := members.Add("andrii@gmail.com", "Andrii Andrii")
	assert.Nil(t, err)

	err = members.Add("andrii@gmail.com", "Andrii Andrii")
	assert.Error(t, err)
}

func TestListMembers(t *testing.T) {
	members := services.New()

	resp := members.List()
	assert.Equal(t, len(resp.Members), 0)

	err := members.Add("andrii@gmail.com", "Andrii Andrii")
	assert.Nil(t, err)

	resp = members.List()
	assert.Equal(t, len(resp.Members), 1)
}
