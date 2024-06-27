package services

import (
	"context"
	"testTask/cmd/internal/db"
	"testTask/cmd/internal/db/models"
)

type UserProfileService struct {
	db db.IUserProfileDatabase
}

func NewUserProfileService(db db.IUserProfileDatabase) *UserProfileService {
	return &UserProfileService{db: db}
}

func (s *UserProfileService) GetAllUsersProfiles(ctx context.Context) ([]models.ProfileAPI, error) {
	return s.db.GetAllProfiles(ctx)
}

func (s *UserProfileService) GetUsersProfilesByUsername(ctx context.Context, username string) ([]models.ProfileAPI, error) {
	profile, err := s.db.GetProfileByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if profile.ID == 0 {
		return []models.ProfileAPI{}, nil
	}
	return []models.ProfileAPI{profile}, nil
}
