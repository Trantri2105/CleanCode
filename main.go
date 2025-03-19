package main

import (
	"context"
	"errors"
)

type UserService interface {
	Register(ctx context.Context, user model.User) error
}

type userService struct {
	userRepo repository.UserRepository
	jwt      util.JwtUtil
	logger   *zap.Logger
}

func (u *userService) Register(ctx context.Context, user model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("failed to generate password", zap.Error(err))
		return err
	}
	user.Password = string(hash)
	err = u.userRepo.CreateUser(ctx, user)
	return err
}

func NewUserRepository(db *gorm.DB, logger *zap.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}
