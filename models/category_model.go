package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

func AddCategory(title string) (err error) {
	o := orm.NewOrm()
	category := &Category{
		Title: title,
	}
	err = o.QueryTable("category").Filter("title", title).One(category)
	beego.Debug(category)
	if err == nil {
		err = fmt.Errorf("category = %v is existed", title)
		return
	}
	if err == orm.ErrNoRows {
		err = nil
	}
	// 获取文章数
	list := make([]*Topic, 0)
	o.QueryTable("topic").Filter("category", title).All(&list)
	category.TopicCount = int64(len(list))

	category.Title = title
	category.Created = time.Now()
	category.TopicTime = time.Now()
	_, err = o.Insert(category)
	if err != nil {
		beego.Error("AddCategory failed", err)
		return
	}
	return
}

func GetAllCategory() (list []*Category, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("category").OrderBy("-created").All(&list)
	if err != nil && err != orm.ErrNoRows {
		beego.Error("GetAllCategory failed", err)
		return
	}
	return
}

func DelCategory(id int64) (err error) {
	o := orm.NewOrm()
	category := &Category{
		Id: id,
	}
	o.QueryTable("category").Filter("id", id).One(category)
	title := category.Title
	_, err = o.Delete(category)
	if err != nil {
		beego.Error("DelCategory error", err)
		return
	}
	var list []*Topic
	o.QueryTable("topic").Filter("category", title).All(&list)
	for _, v := range list {
		v.Category = ""
		o.Update(v, "category")
	}
	return
}
