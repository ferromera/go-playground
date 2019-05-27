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
	GetAllTeams(ctx context.Context) ([]*Team, error)
	GetPlayers(ctx context.Context, id string) ([]*Player, error)
	Load(ctx context.Context) (*dto.ItemCount, error)
}

