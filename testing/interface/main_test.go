package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	user := User{
		ID:   2,
		Name: "Rama",
	}

	md := MockDatastore{
		Users: map[int]User{
			2: user,
		},
	}

	s := Service{
		ds: md,
	}

	u, err := s.GetUser(2)
	assert.NoError(t, err)
	assert.Equal(t, u, user)
}
