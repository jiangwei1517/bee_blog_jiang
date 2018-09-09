package main

import (
	"bee_blog_jiang/models"
	_ "bee_blog_jiang/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func main() {
	err := models.RegisterDb()
	if err != nil {
		beego.Error("register db failed", err)
		panic(err)
	}
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	dirExist, err := PathExists("attachment")
	if err != nil {
		beego.Error("judge dir exist error", err)
		return
	}
	if !dirExist {
		os.Mkdir("attachment", os.ModePerm)
	}
	beego.Run()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
