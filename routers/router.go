package routers

import (
	"BeeGoo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.HelloControllers{})
	//beego.Router("/simple",&SimpleController{},"get:GetFunc; post:PostFunc")
    //beego.Router("/taoba", &SimpleController{},"get:GetFunc; post:PostFunc")
}
