package domain

import (
	"context"

	"github.com/ferromera/go-playground/src/dto"
)

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}
type TeamService interface {
	GetTeam(ctx context.Context, id string) (*Team, error)
	Load(ctx context.Context) (*dto.TeamsLoaded, error)
}

