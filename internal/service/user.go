package service

import (
	"synapse/internal/model"
	"synapse/internal/repository"
	"synapse/internal/utils"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *UserService) Register(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) Login(username, password string) (*model.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(password, user.Password) {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (s *UserService) GetUserProfile(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
