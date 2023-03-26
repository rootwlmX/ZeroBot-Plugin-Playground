package fgopickup

import (
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"strconv"
	"time"
)

func listPickups(ctx *zero.Ctx) {
	service := service{}
	pickups := service.getPickups()

	msg := message.Message{}
	for _, pickup := range pickups {
		id := message.Text("id:" + strconv.Itoa(pickup.Id) + "\n")
		banner := message.Image(pickup.Banner)
		name := message.Text("\n" + pickup.Name)
		date := message.Text("\n" +
			parseTime(pickup.StartTime) + "~" + parseTime(pickup.EndTime))
		msg = append(msg, ctxext.FakeSenderForwardNode(ctx, id, banner, name, date))
	}
	ctx.Send(msg)
}

func pickupDetail(ctx *zero.Ctx) {
	pickupId := ctx.State["args"].(string)
	service := service{}
	id, err := strconv.Atoi(pickupId)
	if err != nil {
		ctx.Send("错误的参数")
	}
	servants := service.getPickupDetail(id)

	msg := message.Message{}
	msg = append(msg, ctxext.FakeSenderForwardNode(ctx, message.Text("卡池包含以下UP从者")))
	for _, servant := range servants {
		avatar := message.Image(servant.Avatar)
		name := message.Text(servant.Name)
		msg = append(msg, ctxext.FakeSenderForwardNode(ctx, avatar, name))
	}
	ctx.Send(msg)
}

func parseTime(timeInSeconds int64) string {
	return time.Unix(timeInSeconds, 0).Format("2006-01-02")
}
