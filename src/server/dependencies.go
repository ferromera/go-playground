package server

import(
	"github.com/ferromera/go-playground/src/controller"
	"github.com/ferromera/go-playground/src/dao"
	"github.com/ferromera/go-playground/src/service"
)

func getTeamController()(*controller.TeamController){
	return &controller.TeamController{
		Service: &service.TeamService{
			DAO: &dao.TeamDao{},
		},
	}
}