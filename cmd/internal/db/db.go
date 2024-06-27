package db

import (
	"context"
	"testTask/cmd/internal/db/models"
)

type IUserProfileDatabase interface {
	GetAllProfiles(ctx context.Context) ([]models.ProfileAPI, error)
	GetProfileByUsername(ctx context.Context, username string) (models.ProfileAPI, error)
}

type IAuthDatabase interface {
	IsExistAuthKey(ctx context.Context, key string) (bool, error)
}
