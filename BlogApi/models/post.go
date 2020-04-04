package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

var (
	PostList       []*Posts
	EducationList  []*Education
	CategoryList   []*Categories
	ExperienceList []*Experience
)

var Post Posts

func init() {

	orm.RegisterModel(new(Posts), new(Education), new(Categories), new(Experience))

}

type Posts struct {
	Id         int    `orm:"pk;column(id)" json:"Id"`
	CategoryId int    `orm:"column(category_id)" json:"CategoryId"`
	User_id    int    `orm:"column(user_id)" json:"UserId"`
	Title      string `orm:"column(title)" json:"Title"`
	Slug       string `orm:"column(slug)" json:"Slug"`
	Body       string `orm:"column(body)" json:"Body"`
	Postimage  string `orm:"column(post_image)" json:"Image"`
	// Createdat  string `orm:"column(created_at)" json:"Created"`
}

type Education struct {
	Id         int    `orm:"pk;column(Id)" json:"id"`
	SchoolName string `orm:"column(SchoolName)" json:"school"`
	Title      string `orm:"column(Title)" json:"title"`
	Start      string `orm:"column(Start)" json:"start"`
	End        string `orm:"column(End)" json:"end"`
	Gpa        string `orm:"column(Gpa)" json:"gpa"`
	// Createdat  string `orm:"column(created_at)" json:"Created"`
}

type Categories struct {
	Id     int    `orm:"pk;column(id)" json:"id"`
	IdUser string `orm:"column(user_id)" json:"IdUser"`
	Name   string `orm:"column(name)" json:"name"`
}

type Experience struct {
	Id       int    `orm:"pk;column(Id)" json:"id"`
	JobTitle string `orm:"column(JobTitle)" json:"jobtitle"`
	Company  string `orm:"column(Company)" json:"company"`
	Start    string `orm:"column(StartDate)" json:"start"`
	End      string `orm:"column(EndDate)" json:"end"`
	JobDesc  string `orm:"column(JobDesc)" json:"jobdesc"`
}

// func (p *Post) TableName() string {
// 	return "posts"
// }

func GetPostAll() []*Posts {
	o := orm.NewOrm()
	// o.Raw("SELECT id, category_id,user_id,title,slug,body,post_image FROM Posts",).QueryRows(&PostList)
	o.QueryTable("posts").All((&PostList))
	return PostList
}

func GetEducationAll() []*Education {
	o := orm.NewOrm()
	// o.Raw("SELECT id, category_id,user_id,title,slug,body,post_image FROM Posts",).QueryRows(&PostList)
	o.QueryTable("education").All((&EducationList))
	return EducationList
}

func GetCategoryAll() []*Categories {
	o := orm.NewOrm()
	// o.Raw("SELECT id, category_id,user_id,title,slug,body,post_image FROM Posts",).QueryRows(&PostList)
	o.QueryTable("categories").All((&CategoryList))
	return CategoryList
}

func GetExperienceAll() []*Experience {
	o := orm.NewOrm()
	// o.Raw("SELECT id, category_id,user_id,title,slug,body,post_image FROM Posts",).QueryRows(&PostList)
	o.QueryTable("experience").All((&ExperienceList))
	return ExperienceList

}

func AddPost(u Posts) string {
	o := orm.NewOrm()
	var post Posts

	id, err := o.Insert(&post)
	if err == nil {
		fmt.Println(id)
	}

	return "post"
}
