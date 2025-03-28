package Usecases

import (
	domain "Task-Management/Domain"
	infrastructure "Task-Management/Infrastructure"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	// Check if user already exists
	existingUser, err := u.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Create user
	return u.userRepo.Create(ctx, user)
}

func (u *userUseCase) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	// Get user by email
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Verify password
	if !infrastructure.ComparePasswords(user.Password, password) {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := infrastructure.GenerateToken(user.ID.Hex(), user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *userUseCase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *userUseCase) GetUserByID(ctx context.Context, id primitive.ObjectID) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	// If password is being updated, hash it
	if user.Password != "" {
		hashedPassword, err := infrastructure.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	return u.userRepo.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return u.userRepo.Delete(ctx, id)
}
