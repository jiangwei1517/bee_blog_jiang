package routers

import (
	"bee_blog_jiang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.AutoRouter(&controllers.ReplyController{})
	beego.Router("/attachment/:all", &controllers.AttachmentController{})
	beego.Router("/category", &controllers.CategoryController{})
}
