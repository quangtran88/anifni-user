package repositories

import (
	"github.com/quangtran88/anifni-user/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockedUserRepository struct {
	mock.Mock
}

func (m *MockedUserRepository) FindById(id domain.ID) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockedUserRepository) Create(user domain.User) (domain.ID, error) {
	args := m.Called(user)
	return args.Get(0).(domain.ID), args.Error(1)
}
