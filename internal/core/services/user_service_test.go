package services_test

import (
	"errors"
	"testing"

	"go-hexagonal/internal/core/domain"
	"go-hexagonal/internal/core/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user domain.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepository) GetByNickname(nickname string) (domain.User, error) {
    args := m.Called(nickname)
    // El Get devuelve un objeto User (o cero) y un error
    return args.Get(0).(domain.User), args.Error(1)
}

// 2. Test: Crear Usuario Exitosamente
func TestCreateUser_Success(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepository)
    service := services.NewUserService(mockRepo)

    // Simulamos que el repositorio siempre guarda bien (retorna nil error)
    // Usamos mock.Anything porque el ID se genera aleatoriamente dentro del servicio
    mockRepo.On("Save", mock.AnythingOfType("domain.User")).Return(nil)

    // Act
    createdUser, err := service.Create("Edwin", "edwin@codesandtags.io", "edwin123")

    // Assert
    assert.NoError(t, err)
    assert.NotEmpty(t, createdUser.ID) // Verificamos que se generó el UUID
    assert.Equal(t, "Edwin", createdUser.Name)

    // Verificamos que el repositorio fue llamado exactamente una vez
    mockRepo.AssertExpectations(t)
}

// 3. Test: Obtener Usuario por Nickname
func TestGetUser_Success(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepository)
    service := services.NewUserService(mockRepo)
    expectedUser := domain.User{ID: "123", Name: "Edwin", Nickname: "edwin123"}

    mockRepo.On("GetByNickname", "edwin123").Return(expectedUser, nil)

    // Act
    user, err := service.Get("edwin123")

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "Edwin", user.Name)
}

// 4. Test: Usuario No Encontrado
func TestGetUser_NotFound(t *testing.T) {
    mockRepo := new(MockUserRepository)
    service := services.NewUserService(mockRepo)

    // Simulamos error desde el repo
    mockRepo.On("GetByNickname", "unknown").Return(domain.User{}, errors.New("db error"))

    _, err := service.Get("unknown")

    assert.Error(t, err)
    assert.Equal(t, "user not found", err.Error())
}