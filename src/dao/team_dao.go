package dao

import(
	"github.com/ferromera/go-playground/src/domain"
	"context"
	"errors"
)

type TeamDao struct{

}


func (td * TeamDao) GetTeam(ctx context.Context, id string) (*domain.Team, error) {
	Barcelona:= &domain.Team{
		ID: "1",
		Name: "Barcelona",
		Rank: 1,
	}
	RealMadrid:= &domain.Team{
		ID: "2",
		Name: "Real Madrid",
		Rank: 4,
	}
	Liverpool:= &domain.Team{
		ID: "3",
		Name: "Liverpool",
		Rank: 3,
	}
	ManchesterCity:= &domain.Team{
		ID: "4",
		Name: "Manchester City",
		Rank: 2,
	}
	switch id {
	case "1": return Barcelona, nil
	case "2": return RealMadrid, nil
	case "3": return Liverpool, nil
	case "4": return ManchesterCity, nil
	}
	return nil, errors.New("Team not found")

}