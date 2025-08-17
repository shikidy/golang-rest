package store_test

import (
	"testing"

	"github.com/shikidy/golang-rest/internal/app/model"
	"github.com/shikidy/golang-rest/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "someemail@gac.rom",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	ut := model.TestUser(t)
	email := ut.Email

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
