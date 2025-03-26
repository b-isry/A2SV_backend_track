package data

import (
	"context"
	"time"

	"github.com/Bisrat/task-manager/errors"
	"github.com/Bisrat/task-manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService() *UserService {
	return &UserService{
		collection: GetCollection(models.UserCollection),
	}
}

func validateUser(user models.User) error {
	if user.Name == "" {
		return &errors.ValidationError{
			Field:   "name",
			Message: "name is required",
		}
	}
	if user.Email == "" {
		return &errors.ValidationError{
			Field:   "email",
			Message: "email is required",
		}
	}
	if user.Password == "" {
		return &errors.ValidationError{
			Field:   "password",
			Message: "password is required",
		}
	}
	return nil
}

func (s *UserService) Register(user models.User) (*models.User, error) {
	if err := validateUser(user); err != nil {
		return nil, err
	}

	var existingUser models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return nil, &errors.ValidationError{
			Field:   "email",
			Message: "email already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	result, err := s.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, &errors.ValidationError{
			Field:   "email",
			Message: "invalid email or password",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, &errors.ValidationError{
			Field:   "password",
			Message: "invalid email or password",
		}
	}

	return &user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
