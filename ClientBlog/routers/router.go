package routers

import (
	"FinalProject/ClientBlog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/admin", &controllers.AdminController{})
}
