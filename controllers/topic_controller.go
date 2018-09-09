package controllers

import (
	"bee_blog_jiang/models"
	"github.com/astaxie/beego"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (p *TopicController) Get() {
	p.TplName = "topic.html"
	p.Data["IsTopic"] = true
	p.Data["IsCategory"] = false
	p.Data["IsHome"] = false
	p.Data["IsLogin"] = checkIsLogin(p.Ctx)
	list, err := models.GetAllTopics("")
	if err != nil {
		beego.Error("getAllTopics failed", err)
		return
	}
	p.Data["Topics"] = list
}

func (p *TopicController) Add() {
	p.TplName = "topic_add.html"
	p.Data["IsTopic"] = true
	p.Data["IsCategory"] = false
	p.Data["IsHome"] = false
	p.Data["IsLogin"] = checkIsLogin(p.Ctx)
}

func (p *TopicController) Post() {
	if !checkIsLogin(p.Ctx) {
		p.Redirect("/login", http.StatusMovedPermanently)
	}
	title := p.Input().Get("title")
	category := p.Input().Get("category")
	lable := p.Input().Get("lable")
	lable = "$" + strings.Join(strings.Split(lable, " "), "#$") + "#"
	content := p.Input().Get("content")
	_, fd, err := p.GetFile("attachment")
	if err != nil {
		beego.Error("attachment upload error", err)
		return
	}
	fileName := fd.Filename
	toFile := path.Join("attachment", fileName)
	p.SaveToFile("attachment", toFile)

	tid := p.Input().Get("tid")
	if tid == "" {
		err = models.AddTopic(title, category, lable, content, fileName)
		if err != nil {
			beego.Error("AddTopic failed", err)
			return
		}
	} else {
		id, err := strconv.ParseInt(tid, 10, 64)
		if err != nil {
			beego.Error("strconv.ParseInt tid failed", err)
			return
		}
		err = models.ModifyTopic(id, title, category, lable, content, fileName)
		if err != nil {
			beego.Error("ModifyTopic failed", err)
			return
		}
	}
	p.Redirect("/topic", http.StatusMovedPermanently)
}

func (p *TopicController) Delete() {
	tid := p.Input().Get("tid")
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("strconv.ParseInt failed", err)
		return
	}
	err = models.DeleteTopic(id)
	if err != nil {
		beego.Error("DeleteTopic failed", err)
		return
	}
	p.Redirect("/topic", http.StatusMovedPermanently)
}

func (p *TopicController) Modify() {
	p.TplName = "topic_modify.html"
	p.Data["IsTopic"] = true
	p.Data["IsCategory"] = false
	p.Data["IsHome"] = false
	p.Data["IsLogin"] = checkIsLogin(p.Ctx)
	tid := p.Input().Get("tid")
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("strconv.ParseInt failed", err)
		return
	}
	p.Data["Tid"] = id
	topic, err := models.GetOneTopic(id)
	if err != nil {
		beego.Error("GetOneTopic failed", err)
		return
	}
	lables := topic.Lables
	lables = strings.Replace(strings.Replace(lables, "#", "", -1), "$", " ", -1)
	topic.Lables = lables
	p.Data["Topic"] = topic
}

func (p *TopicController) View() {
	p.TplName = "topic_view.html"
	p.Data["IsTopic"] = true
	p.Data["IsCategory"] = false
	p.Data["IsHome"] = false
	p.Data["IsLogin"] = checkIsLogin(p.Ctx)
	tid := p.Ctx.Input.Params()["0"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("strconv.ParseInt failed", err)
		return
	}
	topic, err := models.GetOneTopic(id)
	if err != nil {
		beego.Error("")
	}
	lables := topic.Lables
	lables = strings.Replace(strings.Replace(lables, "#", "", -1), "$", " ", -1)
	lablesArr := strings.Split(lables, " ")
	p.Data["Lables"] = lablesArr
	p.Data["Topic"] = topic

	list, err := models.GetAllReplys(id)
	if err != nil {
		beego.Error("GetAllReplys failed", err)
		return
	}
	p.Data["Replies"] = list
}
