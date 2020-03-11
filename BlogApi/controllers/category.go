package controllers

import (
	"FinalProject/BlogApi/models"

	"github.com/astaxie/beego"
)

// Operations about Category
type CategoryController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
// func (u *CategoryController) Post() {
// 	var user models.User
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 	uid := models.AddUser(user)
// 	u.Data["json"] = map[string]string{"uid": uid}
// 	u.ServeJSON()
// }

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Article
// @router / [get]
func (c *CategoryController) GetAll() {
	category := models.GetCategoryAll()
	c.Data["json"] = category
	c.ServeJSON()
}
