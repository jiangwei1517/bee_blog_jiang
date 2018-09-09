package controllers

import (
	"bee_blog_jiang/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (p *HomeController) Get() {
	beego.Debug("enter home page")
	p.Data["IsHome"] = true
	p.Data["IsCategory"] = false
	p.Data["IsTopic"] = false
	p.TplName = "home.html"
	p.Data["IsLogin"] = checkIsLogin(p.Ctx)

	categorys, err := models.GetAllCategory()
	if err != nil {
		beego.Error("GetAllCategory failed", err)
		return
	}
	p.Data["Categories"] = categorys
	cate := p.Input().Get("cate")
	var topics = make([]*models.Topic, 0)
	if len(cate) != 0 {
		topics, err = models.GetAllTopics(cate)
	} else {
		topics, err = models.GetAllTopics("")
	}
	if err != nil {
		beego.Error("GetAllTopics failed", err)
		return
	}
	p.Data["Topics"] = topics
}
