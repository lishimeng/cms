package cms

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
)

// Website 注入网站的基本配置
func Website(ctx server.Context) {
	ws, err := GetWebsiteInfo()
	if err != nil {
		ctx.C.Next()
		return
	}
	ctx.C.ViewData(tool.WebsiteCtx, ws)
	ctx.C.Next()
}
