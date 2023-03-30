// Package mahjong 日麻相关插件
package mahjong

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	engine := control.Register("mahjong", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "blank",
		Help: "- 牌理\n" +
			"- blank\n" +
			"- blank\n" +
			"- blank",
		// Banner:            "https://wx2.sinaimg.cn/large/0083LFbYgy1hcfkreklmbj31e012w7i5.jpg",
		PrivateDataFolder: "mahjong",
	})

	engine.OnPrefix(`牌理`).
		SetBlock(true).
		Limit(ctxext.LimitByGroup).
		Handle(nil)
}
