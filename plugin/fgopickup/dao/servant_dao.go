package dao

import (
	"fmt"
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/gormengine"
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/model"
)

type ServantDao struct {
	DbEngine *gormengine.Orm
}

func (sd *ServantDao) ListServantIdsByPickup(pickupId int) *[]int {
	servantIds := make([]int, 0)
	err := sd.DbEngine.Where("pickup_id = ? ", pickupId).Find(&servantIds).Error
	if err != nil {
		fmt.Println(err)
	}
	return &servantIds
}

func (sd *ServantDao) SelectServantsByIds(ids []int) *[]model.Servant {
	servant := make([]model.Servant, 0)
	sd.DbEngine.Where(ids).Find(&servant)
	return &servant
}
