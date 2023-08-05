package service

import (
	"context"
	"elasticsearch/model"
	"elasticsearch/repo"
)

type EsService struct {
	repo repo.IRepo
}

func NewEsService(repo repo.IRepo) *EsService {
	return &EsService{
		repo: repo,
	}
}

type IEsService interface {
	Test(ctx context.Context) (interface{}, error)
	Insert(ctx context.Context, ESRequest model.ESRequest) error
}

func (s *EsService) Test(ctx context.Context) (interface{}, error) {
	return s.repo.Test(ctx)
}

func (s *EsService) Insert(ctx context.Context, ESRequest model.ESRequest) error {
	return s.repo.Insert(ctx, ESRequest)
}
