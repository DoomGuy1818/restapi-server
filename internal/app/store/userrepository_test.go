package store_test

import (
	"testing"

	"github.com/DoomGuy1818/restapi-server/internal/app/store"
	"github.com/DoomGuy1818/restapi-server/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T){
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "userexample@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}


func TestUserRepository_FindByEmail(t *testing.T){
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "userexample@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}
