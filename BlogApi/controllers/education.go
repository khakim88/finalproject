package controllers

import (
	"FinalProject/BlogApi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type EducationController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
// func (e *EducationController) Post() {
// 	var user models.User
// 	json.Unmarshal(e.Ctx.Input.RequestBody, &user)
// 	uid := models.AddUser(user)
// 	e.Data["json"] = map[string]string{"uid": uid}
// 	e.ServeJSON()
// }

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (e *EducationController) GetAll() {
	users := models.GetEducationAll()
	e.Data["json"] = users
	e.ServeJSON()
}
