package repository

import (
	"errors"
	"go-hexagonal/internal/core/domain"
	"sync"
)

type InMemoryRepo struct {
	mu sync.RWMutex
	users map[string]domain.User
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		users: make(map[string]domain.User),
	}
}

// Save implementa el puerto UserRepository
func (r *InMemoryRepo) Save(user domain.User) error {
    r.mu.Lock()         // Bloqueamos para escritura (Nadie más puede leer ni escribir)
    defer r.mu.Unlock() // Aseguramos el desbloqueo al salir de la función

    // Validación simple de unicidad (simulando una constraint de DB)
    if _, exists := r.users[user.Nickname]; exists {
        return errors.New("user with this nickname already exists")
    }

    // Usamos el Nickname como llave primaria para búsqueda rápida O(1)
    // En una DB real, probablemente indexarías por ID y tendrías un índice único en Nickname.
    r.users[user.Nickname] = user

    return nil
}

// GetByNickname implementa el puerto UserRepository
func (r *InMemoryRepo) GetByNickname(nickname string) (domain.User, error) {
    r.mu.RLock()         // Bloqueo de LECTURA (otros pueden leer, nadie puede escribir)
    defer r.mu.RUnlock()

    user, exists := r.users[nickname]
    if !exists {
        // Devolvemos un objeto vacío y el error estándar
        return domain.User{}, errors.New("user not found")
    }

    return user, nil
}