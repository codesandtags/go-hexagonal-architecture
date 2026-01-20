package ports

import "go-hexagonal/internal/core/domain"

type UserRepository interface {
	Save(user domain.User) error
	GetByNickname(nickname string) (domain.User, error)
}

type UserService interface {
	Create(name, email, nickname string) (domain.User, error)
	Get(nickname string) (domain.User, error)
}