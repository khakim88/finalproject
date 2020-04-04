package controllers

import (
	"FinalProject/BlogApi/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type ArticleController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ArticleController) Post() {
	var posts models.Posts
	json.Unmarshal(u.Ctx.Input.RequestBody, &posts)
	uid := models.AddPost(posts)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Article
// @router / [get]
func (u *ArticleController) GetAll() {
	article := models.GetPostAll()
	u.Data["json"] = article
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
// func (u *ArticleController) Get() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		user, err := models.GetUser(uid)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = user
// 		}
// 	}
// 	u.ServeJSON()
// }
