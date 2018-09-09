package controllers

import (
	"bee_blog_jiang/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (p *CategoryController) Get() {
	beego.Debug("enter category page")
	p.Data["IsHome"] = false
	p.Data["IsCategory"] = true
	p.Data["IsTopic"] = false
	p.TplName = "category.html"
	p.Data["IsLogin"] = checkIsLogin(p.Ctx)
	name := p.Input().Get("name")
	op := p.Input().Get("op")
	id := p.Input().Get("id")
	switch op {
	case "add":
		err := models.AddCategory(name)
		if err != nil {
			beego.Error("AddCategory failed", err)
			return
		}
		p.Redirect("/category", http.StatusMovedPermanently)
	case "del":
		if id != "" {
			idnum, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				beego.Error("strconv.ParseInt failed", err)
				return
			}
			err = models.DelCategory(idnum)
			if err != nil {
				beego.Error("DelCategory failed", err)
			}
			p.Redirect("/category", http.StatusMovedPermanently)
		}
	}
	list, err := models.GetAllCategory()
	if err != nil && err != orm.ErrNoRows {
		beego.Error("GetAllCategory failed", err)
		return
	}
	p.Data["Categories"] = list
}
