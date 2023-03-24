package fgopickup

import (
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"time"
)

func listPickups(ctx *zero.Ctx) {
	service := service{}
	pickups := service.getPickups()

	msg := message.Message{}
	for _, pickup := range pickups {
		banner := message.Image(pickup.Banner)
		name := message.Text("\n" + pickup.Name)
		date := message.Text("\n" +
			parseTime(pickup.StartTime) + "~" + parseTime(pickup.EndTime))
		msg = append(msg, ctxext.FakeSenderForwardNode(ctx, banner, name, date))
	}
	ctx.Send(msg)
}

func parseTime(timeInSeconds int64) string {
	return time.Unix(timeInSeconds, 0).Format("2006-01-02")
}
