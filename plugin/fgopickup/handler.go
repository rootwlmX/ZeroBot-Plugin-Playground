package fgopickup

import (
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
	"time"
)

// 列出所有的卡池
func listPickups(ctx *zero.Ctx) {
	service := service{}
	pickups := service.getPickups()

	msg := make(message.Message, len(*pickups))
	for i, pickup := range *pickups {
		msg[i] = getMsgOfSinglePickup(ctx, pickup)
	}
	ctx.Send(msg)
}

// 以卡池id获取某一个卡池的up从者们
func pickupDetail(ctx *zero.Ctx) {
	pickupID, err := strconv.Atoi(ctx.State["args"].(string))
	if err != nil {
		ctx.Send("参数错误！")
		return
	}
	service := service{}
	detail := service.getPickupDetail(pickupID)
	servants := detail.Servants

	msg := make(message.Message, len(servants)+1)
	msg[0] = getMsgOfSinglePickup(ctx, detail.Pickup)

	for i, servant := range servants {
		avatar := message.Image(servant.Avatar)
		name := message.Text(servant.Name)
		msg[i+1] = ctxext.FakeSenderForwardNode(ctx, avatar, name)
	}
	ctx.Send(msg)
}

func getMsgOfSinglePickup(ctx *zero.Ctx, pickup pickup) message.MessageSegment {
	banner := message.Image(pickup.Banner)
	id := message.Text("\n", strconv.Itoa(pickup.ID), ". ")
	name := message.Text(pickup.Name)
	date := message.Text("\n",
		parseTime(pickup.StartTime), "~", parseTime(pickup.EndTime))
	return ctxext.FakeSenderForwardNode(ctx, banner, id, name, date)
}

func parseTime(timeInSeconds int64) string {
	return time.Unix(timeInSeconds, 0).Format("2006-01-02")
}
