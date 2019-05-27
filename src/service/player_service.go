package service

import (
	"context"

	"github.com/ferromera/go-playground/src/domain"
	"github.com/ferromera/go-playground/src/dto"
)

type PlayerService struct {
	DAO domain.PlayerService
}

func (ts *PlayerService) GetPlayer(ctx context.Context, id string) (*domain.Player, error) {
	return ts.DAO.GetPlayer(ctx, id)
}

func (ts *PlayerService) Load(ctx context.Context) (*dto.ItemCount, error) {
	return ts.DAO.Load(ctx)
}

func (ts *PlayerService) GetPlayers(ctx context.Context, minOverrall int, maxOverall int) ([]*domain.Player, error) {
	return ts.DAO.GetPlayers(ctx, minOverrall, maxOverall)
}
