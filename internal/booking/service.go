package booking

import (
	"context"

	"github.com/darkphotonKN/epic-backend-template/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) GetById(ctx context.Context, userId uuid.UUID, id uuid.UUID) (*models.Booking, error) {
	return s.Repo.GetById(ctx, userId, id)
}

func (s *Service) Create(ctx context.Context, userId uuid.UUID, req CreateRequest) error {
	return s.Repo.Create(ctx, userId, req)
}

func (s *Service) CreateTwo(ctx context.Context, userId uuid.UUID, req CreateRequest) error {
	return s.Repo.Create(ctx, userId, req)
}
