package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"serveLine/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getData", &controllers.GetToken{})
	beego.Router("/ws", &controllers.WebSocket{})
}
