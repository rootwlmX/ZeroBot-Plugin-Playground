package fgopickup

import (
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/dao"
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/gormengine"
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/model"
)

type service struct {
}

func (s *service) getPickups() []model.Pickup {
	pd := dao.PickupDao{DbEngine: gormengine.GetOrmEngine()}
	list := pd.ListPickup()
	return *list
}

func (s *service) getPickupDetail(pickupId int) []model.Servant {
	sd := dao.ServantDao{DbEngine: gormengine.GetOrmEngine()}
	ids := sd.ListServantIdsByPickup(pickupId)
	return *sd.SelectServantsByIds(*ids)
}
