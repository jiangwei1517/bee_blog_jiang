package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

func (p *LoginController) Get() {
	beego.Debug("enter login page")
	p.TplName = "login.html"
	//  退出登陆
	if p.Input().Get("exit") == "true" {
		p.Ctx.SetCookie("uname", "", -1, "/")
		p.Ctx.SetCookie("pwd", "", -1, "/")
	}
}

// 提交登陆密码表单
func (p *LoginController) Post() {
	uname := p.Input().Get("uname")
	if len(uname) == 0 {
		beego.Debug("uname is empty")
		return
	}
	pwd := p.Input().Get("pwd")
	if len(pwd) == 0 {
		beego.Debug("pwd is empty")
		return
	}
	isAutoOn := p.Input().Get("autoLogin")
	// 保存到cookie
	if uname == beego.AppConfig.String("admin") && pwd == beego.AppConfig.String("pwd") {
		maxAge := 0
		if isAutoOn == "on" {
			maxAge = 1<<31 - 1
		}
		p.Ctx.SetCookie("uname", uname, maxAge, "/")
		p.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	} else {
		beego.Debug("admin or pwd is error")
		p.Redirect("/login", http.StatusMovedPermanently)
		return
	}
	p.Redirect("/", http.StatusMovedPermanently)
}

func checkIsLogin(ctx *context.Context) (b bool) {
	uname := ctx.GetCookie("uname")
	pwd := ctx.GetCookie("pwd")
	if uname == beego.AppConfig.String("admin") && pwd == beego.AppConfig.String("pwd") {
		return true
	}
	return
}
