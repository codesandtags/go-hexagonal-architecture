package services

import (
	"errors"
	"go-hexagonal/internal/core/domain"
	"go-hexagonal/internal/core/ports"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Create(name, email, nickname string) (domain.User, error) {
	user := domain.User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Nickname: nickname,
	}

	if err := s.repo.Save(user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserServiceImpl) Get(nickname string) (domain.User, error) {
   user, error := s.repo.GetByNickname(nickname)

	if error != nil {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}