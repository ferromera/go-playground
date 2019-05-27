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

type PlayerDao struct {
	Session *mgo.Session
}

func (td *PlayerDao) GetPlayer(ctx context.Context, id string) (*domain.Player, error) {
	c := td.Session.DB("football").C("players")
	var player domain.Player
	e := c.Find(bson.M{"id": id}).One(&player)

	if e != nil && e != mgo.ErrNotFound {
		return nil, errors.New(e.Error())
	}
	if e == mgo.ErrNotFound {
		return nil, nil
	}

	return &player, nil

}

func (td *PlayerDao) GetPlayers(ctx context.Context, minOverrall int, maxOverall int) ([]*domain.Player, error) {
	c := td.Session.DB("football").C("players")
	var results []*domain.Player
	e := c.Find(bson.M{"overall": bson.M{"$gt": minOverrall, "$lt": maxOverall}}).All(&results)
	if e != nil && e != mgo.ErrNotFound {
		return nil, errors.New(e.Error())
	}
	if e == mgo.ErrNotFound {
		return nil, nil
	}

	return results, nil
}

func (td *PlayerDao) Load(ctx context.Context) (*dto.ItemCount, error) {
	c := td.Session.DB("football").C("players")
	c.RemoveAll(nil)

	csvFile, _ := os.Open("../resources/players.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	players := dto.ItemCount{Count: 0}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.New(err.Error())
		}
		overall, err := strconv.Atoi(line[3])
		if err != nil {
			return nil, errors.New(err.Error())
		}
		age, err := strconv.Atoi(line[4])
		if err != nil {
			return nil, errors.New(err.Error())
		}
		err = c.Insert(&domain.Player{
			ID:      line[0],
			Name:    line[1],
			Team:    line[2],
			Overall: overall,
			Age:     age,
		})
		if err != nil {
			return nil, errors.New(err.Error())
		}
		players.Count++
	}
	return &players, nil

}
