package controllers

import (
	"bee_blog_jiang/models"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
)

type ReplyController struct {
	beego.Controller
}

func (p *ReplyController) Add() {
	tid := p.Input().Get("tid")
	nickName := p.Input().Get("nickname")
	content := p.Input().Get("content")
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("strconv.ParseInt tid failed", err)
		return
	}
	err = models.AddReply(id, nickName, content)
	if err != nil {
		beego.Error("AddReply failed", err)
		return
	}
	p.Redirect("/topic/view/"+tid, http.StatusMovedPermanently)
}

func (p *ReplyController) Delete() {
	tid := p.Input().Get("tid")
	rid := p.Input().Get("rid")
	tidnum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error("strconv.ParseInt tid failed", err)
		return
	}
	ridnum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		beego.Error("strconv.ParseInt rid failed", err)
		return
	}
	err = models.DeleteReply(tidnum, ridnum)
	if err != nil {
		beego.Error("DeleteReply failed", err)
		return
	}
	p.Redirect("/topic/view/"+tid, http.StatusMovedPermanently)
}
