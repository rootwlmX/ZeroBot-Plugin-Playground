package dao

import (
	"fmt"
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/gormengine"
	"github.com/FloatTech/ZeroBot-Plugin-Playground/plugin/fgopickup/model"
	"time"
)

type PickupDao struct {
	DbEngine *gormengine.Orm
}

func (pd *PickupDao) ListPickup() *[]model.Pickup {
	pickup := make([]model.Pickup, 0)
	unixTime := time.Now().Unix()
	err := pd.DbEngine.Where("end_time >= ?", unixTime).Find(&pickup).Error
	if err != nil {
		fmt.Println(err)
	}
	return &pickup
}

func (pd *PickupDao) ListPickupIdsByServant(servantId int) *[]int {
	pickupIds := make([]int, 0)
	err := pd.DbEngine.Where("servant_id = ?", servantId).Find(&pickupIds).Error
	if err != nil {
		fmt.Println(err)
	}
	return &pickupIds
}

func (pd *PickupDao) SelectPickupsByIds(ids []int) *[]model.Pickup {
	pickup := make([]model.Pickup, 0)
	unixTime := time.Now().Unix()
	err := pd.DbEngine.Where(ids).Where("end_time >= ?", unixTime).Find(&pickup).Error
	if err != nil {
		fmt.Println(err)
	}
	return &pickup
}
