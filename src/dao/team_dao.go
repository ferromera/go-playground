package dao

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/ferromera/go-playground/src/domain"
	"github.com/ferromera/go-playground/src/dto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TeamDao struct {
	Session *mgo.Session
}

func (td *TeamDao) GetTeam(ctx context.Context, id string) (*domain.Team, error) {
	c := td.Session.DB("football").C("teams")
	var team domain.Team
	e := c.Find(bson.M{"id": id}).One(&team)

	if e != nil && e != mgo.ErrNotFound {
		return nil, errors.New(e.Error())
	}
	if e == mgo.ErrNotFound {
		return nil, nil
	}

	return &team, nil

}

func (td *TeamDao) Load(ctx context.Context) (*dto.ItemCount, error) {
	c := td.Session.DB("football").C("teams")
	c.RemoveAll(nil)

	csvFile, _ := os.Open("../resources/teams.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	teams := dto.ItemCount{Count: 0}
	id := 1
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.New(err.Error())
		}
		rank, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, errors.New(err.Error())
		}
		err = c.Insert(&domain.Team{
			ID:   strconv.Itoa(id),
			Name: line[0],
			Rank: rank,
		})
		if err != nil {
			return nil, errors.New(err.Error())
		}
		teams.Count++
		id++
	}
	return &teams, nil
}

func (td *TeamDao) GetAllTeams(ctx context.Context) ([]*domain.Team, error) {

	c := td.Session.DB("football").C("teams")
	var results []*domain.Team
	e := c.Find(nil).All(&results)
	if e != nil && e != mgo.ErrNotFound {
		return nil, errors.New(e.Error())
	}
	if e == mgo.ErrNotFound {
		return nil, nil
	}

	return results, nil
}

func (td *TeamDao) GetPlayers(ctx context.Context, id string) ([]*domain.Player, error) {
	c := td.Session.DB("football").C("players")
	var results []*domain.Player
	e := c.Find(bson.M{"team": id}).All(&results)
	if e != nil && e != mgo.ErrNotFound {
		return nil, errors.New(e.Error())
	}
	if e == mgo.ErrNotFound {
		return nil, nil
	}

	return results, nil
}
