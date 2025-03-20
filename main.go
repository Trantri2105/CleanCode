package main

import "context"

type userService struct {
	userRepo repository.UserRepository
	jwt      util.JwtUtil
	logger   *zap.Logger
}

func (u *userService) Register(ctx context.Context, user model.User, userRepo repository.UserRepository) error {
	// Process register business logic
}
