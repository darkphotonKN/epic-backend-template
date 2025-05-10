package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/darkphotonKN/epic-backend-template/internal/auth"
	"github.com/darkphotonKN/epic-backend-template/internal/constants"
	"github.com/darkphotonKN/epic-backend-template/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	Repo Repository
}

type Repository interface {
	Create(ctx context.Context, user models.User) error
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetAll(ctx context.Context) ([]*Response, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewService(repo Repository) Service {
	return &service{
		Repo: repo,
	}
}

func (s *service) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.Repo.GetById(ctx, id)
}

func (s *service) Create(ctx context.Context, user models.User) error {
	hashedPw, err := s.HashPassword(user.Password)

	if err != nil {
		return fmt.Errorf("Error when attempting to hash password.")
	}

	// update user's password with hashed password.
	user.Password = hashedPw

	return s.Repo.Create(ctx, user)
}

// HashPassword hashes the given password using bcrypt.
func (s *service) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *service) GetAll(ctx context.Context) ([]*Response, error) {
	return s.Repo.GetAll(ctx)
}

func (s *service) Login(ctx context.Context, loginReq LoginRequest) (*LoginResponse, error) {
	user, err := s.Repo.GetUserByEmail(ctx, loginReq.Email)

	if err != nil {
		return nil, errors.New("Could not get user with provided email.")
	}

	// extract password, and compare hashes
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return nil, errors.New("The credentials provided was incorrect.")
	}

	// construct response with both user info and auth credentials
	accessExpiryTime := time.Minute * 60
	accessToken, err := auth.GenerateJWT(*user, constants.Access, accessExpiryTime)
	refreshExpiryTime := time.Hour * 24 * 7
	refreshToken, err := auth.GenerateJWT(*user, constants.Refresh, refreshExpiryTime)

	user.Password = ""

	res := &LoginResponse{
		AccessToken:      accessToken,
		AccessExpiresIn:  int(accessExpiryTime),
		RefreshToken:     refreshToken,
		RefreshExpiresIn: int(refreshExpiryTime),
		UserInfo:         user,
	}

	return res, nil
}
