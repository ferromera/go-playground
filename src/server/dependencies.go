package server

import (
	"log"

	"github.com/ferromera/go-playground/src/controller"
	"github.com/ferromera/go-playground/src/dao"
	"github.com/ferromera/go-playground/src/service"
	"gopkg.in/mgo.v2"
)

func getTeamController() *controller.TeamController {
	return &controller.TeamController{
		Service: &service.TeamService{
			DAO: &dao.TeamDao{
				Session: getSession(),
			},
		},
	}
}
func getPlayerController() *controller.PlayerController {
	return &controller.PlayerController{
		Service: &service.PlayerService{
			DAO: &dao.PlayerDao{
				Session: getSession(),
			},
		},
	}
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Unable to create session")
		panic(err)
	}
	return session
}
