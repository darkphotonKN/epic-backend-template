package item

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	// in-memory storage for testing / demo purposes
	items map[uuid.UUID]*Item
}

func NewService() *Service {
	return &Service{
		items: make(map[uuid.UUID]*Item),
	}
}

func (s *Service) GetById(ctx context.Context, id uuid.UUID) (*Item, error) {
	item, exists := s.items[id]
	if !exists {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*Item, error) {
	items := make([]*Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}
	return items, nil
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Item, error) {
	item := &Item{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if item.Status == "" {
		item.Status = "active"
	}

	s.items[item.ID] = item
	return item, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (*Item, error) {
	item, exists := s.items[id]
	if !exists {
		return nil, errors.New("item not found")
	}

	// Update fields if provided
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.Description != "" {
		item.Description = req.Description
	}
	if req.Status != "" {
		item.Status = req.Status
	}

	item.UpdatedAt = time.Now()

	return item, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if _, exists := s.items[id]; !exists {
		return fmt.Errorf("item not found")
	}

	delete(s.items, id)
	return nil
}

