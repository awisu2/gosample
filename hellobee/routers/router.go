package routers

import (
	"github.com/awisu2/gosample/hellobee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
