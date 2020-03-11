package main

import (
	_ "FinalProject/ClientBlog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Default database
	orm.RegisterDataBase("default", "mysql", "root:@/ci-blog?charset=utf8")
	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxIdleConns("default", 40)

	//orm.Debug = true
	orm.RunCommand()
	beego.Run()
}
