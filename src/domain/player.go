package domain

import (
	"context"

	"github.com/ferromera/go-playground/src/dto"
)

type Player struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Team    string `json:"team"`
	Overall int    `json:"overall"`
	Age     int    `json:"age"`
}

type PlayerService interface {
	GetPlayer(ctx context.Context, id string) (*Player, error)
	GetPlayers(ctx context.Context, minOverrall int, maxOverall int) ([]*Player, error)
	Load(ctx context.Context) (*dto.ItemCount, error)
}
