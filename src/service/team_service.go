package service

import (
	"context"

	"github.com/ferromera/go-playground/src/domain"
)

type TeamService struct {
	DAO domain.TeamResolver
}

func (ts *TeamService) GetTeam(ctx context.Context, id string) (*domain.Team, error) {
	return ts.DAO.GetTeam(ctx, id)
}
