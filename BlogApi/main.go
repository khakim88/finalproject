package main

import (
	_ "FinalProject/BlogApi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		//	beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Default database
	orm.RegisterDataBase("default", "mysql", "root:@/ci-blog?charset=utf8")
	orm.SetMaxOpenConns("default", 30)
	orm.SetMaxIdleConns("default", 40)

	//orm.Debug = true
	orm.RunCommand()
	beego.Run()
}
