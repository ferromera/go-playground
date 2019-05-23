package domain

import "context"

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}
type TeamResolver interface {
	GetTeam(ctx context.Context, id string) (*Team, error)
}
