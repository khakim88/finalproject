package controllers

import (
	"FinalProject/ClientBlog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	education := models.GetEducationAll()

	c.Data["About"] = "I am experienced in leveraging agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition."
	c.Data["Email"] = "agus.khakim@bluebirdgroup.com"
	c.Data["Eduction"] = education
	//var a string
	c.TplName = "index.tpl"
}
