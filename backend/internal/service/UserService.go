package service

import (
	"RecommendationService/internal/model"
	"RecommendationService/internal/repository"
	"database/sql"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetOrCreateUser(username string) (*model.User, error) {
	user, err := s.repo.GetUserByUsername(username)

	if err == nil {
		return user, nil
	}

	if err != sql.ErrNoRows {
		return nil, err
	}

	userID, err := s.repo.CreateUser(username)
	if err != nil {
		return nil, err
	}

	return s.repo.GetUserByID(userID)
}
