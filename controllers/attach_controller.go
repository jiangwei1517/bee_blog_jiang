package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachmentController struct {
	beego.Controller
}

func (p *AttachmentController) Get() {
	url, err := url.QueryUnescape(p.Ctx.Request.RequestURI)
	if err != nil {
		beego.Error("url.QueryUnescape failed", err)
		return
	}
	url = url[1:]
	download(p, url)
}

func download(p *AttachmentController, url string) (err error) {
	beego.Debug(url)
	file, err := os.Open(url)
	if err != nil {
		beego.Error("os.Open failed", err)
		return
	}
	_, err = io.Copy(p.Ctx.ResponseWriter, file)
	if err != nil {
		beego.Error("io.Copy failed", err)
	}
	return
}
