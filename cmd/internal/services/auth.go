package services

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"testTask/cmd/internal/db"
)

type AuthService struct {
	db db.IAuthDatabase
}

func NewAuthService(db db.IAuthDatabase) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) IsExistAuthKey(ctx context.Context, key string) bool {
	isExists, err := s.db.IsExistAuthKey(ctx, key)
	if err != nil {
		log.Error("Error in IsExistAuthKey: ", err)
		return false
	}
	return isExists
}
