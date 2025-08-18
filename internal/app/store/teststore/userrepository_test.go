package teststore_test

import (
	"testing"

	"github.com/shikidy/golang-rest/internal/app/model"
	"github.com/shikidy/golang-rest/internal/app/store"
	"github.com/shikidy/golang-rest/internal/app/store/teststore"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	assert.NoError(t, s.User().Create(model.TestUser(t)))
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	ut := model.TestUser(t)
	email := ut.Email

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
