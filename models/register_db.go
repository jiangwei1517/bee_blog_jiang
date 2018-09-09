package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDb() (err error) {
	beego.SetLogFuncCall(true)
	mySqlAdmin := beego.AppConfig.String("mysql_admin")
	mySqlPwd := beego.AppConfig.String("mysql_pwd")
	mySqlAddr := beego.AppConfig.String("mysql_addr")
	mySqlPort := beego.AppConfig.String("mysql_port")
	mySqlDataBase := beego.AppConfig.String("mysql_database")
	err = orm.RegisterDriver(SQL_DRIVER, orm.DRMySQL)
	if err != nil {
		beego.Error("orm.RegisterDriver failed", err)
		return
	}
	orm.RegisterModel(new(Topic), new(Comment), new(Category))
	err = orm.RegisterDataBase("default", SQL_DRIVER, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mySqlAdmin, mySqlPwd, mySqlAddr, mySqlPort, mySqlDataBase))
	if err != nil {
		beego.Error("orm.RegisterDataBase failed", err)
		return
	}
	beego.Debug("register Db success...")
	return
}
