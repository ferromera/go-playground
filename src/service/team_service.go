package service

import (
	"context"

	"github.com/ferromera/go-playground/src/domain"
	"github.com/ferromera/go-playground/src/dto"
)

type TeamService struct {
	DAO domain.TeamService
}

func (ts *TeamService) GetTeam(ctx context.Context, id string) (*domain.Team, error) {
	return ts.DAO.GetTeam(ctx, id)
}

func (ts *TeamService) Load(ctx context.Context) (*dto.TeamsLoaded, error) {
	return ts.DAO.Load(ctx)
}
