package service

import (
	"github.com/tatsuo48/go_ddd/domain/model"
)

type CircleService struct {
	circleRepositry model.ICircleRepositry
}

func NewCircleService(cr model.ICircleRepositry) CircleService {
	return CircleService{
		circleRepositry: cr,
	}
}

func (cs CircleService) Exists(c model.Circle) bool {
	circle := cs.circleRepositry.FindByName(c.CircleName)
	return circle != nil
}
