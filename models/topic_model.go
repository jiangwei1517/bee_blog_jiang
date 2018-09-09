package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"time"
)

const (
	SQL_DRIVER = "mysql"
)

// 文章
type Topic struct {
	Id                int64
	Uid               int64
	Title             string
	Category          string
	Lables            string
	Content           string `orm:"size(5000)"`
	Attachment        string
	Created           time.Time `orm:"index"`
	Updated           time.Time `orm:"index"`
	Views             int64     `orm:"index"`
	Author            string
	ReplyTime         time.Time `orm:"index"`
	ReplyCount        int64
	ReplyLastUserName string
}

func AddTopic(title, category, lable, content, attachment string) (err error) {
	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Category:   category,
		Lables:     lable,
		Content:    content,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
	}
	_, err = o.Insert(topic)
	if err != nil {
		beego.Error("orm insert failed", err)
		return
	}
	cate := &Category{}
	o.QueryTable("category").Filter("title", category).One(cate)
	beego.Error(cate)
	cate.TopicCount++
	o.Update(cate, "topic_count")
	return
}

func GetAllTopics(cate string) (list []*Topic, err error) {
	o := orm.NewOrm()
	if len(cate) != 0 {
		_, err = o.QueryTable("topic").OrderBy("-updated").Filter("category", cate).All(&list)
	} else {
		_, err = o.QueryTable("topic").OrderBy("-updated").All(&list)
	}
	if err != nil {
		beego.Error("GetAllTopics failed", err)
		return
	}
	return
}

func DeleteTopic(id int64) (err error) {
	o := orm.NewOrm()
	topic := &Topic{
		Id: id,
	}
	o.QueryTable("topic").Filter("id", id).One(topic)
	title := topic.Category
	_, err = o.Delete(topic)
	if err != nil {
		beego.Error("delete topic failed,", err)
		return
	}
	cate := &Category{}
	o.QueryTable("category").Filter("title", title).One(cate)
	cate.TopicCount--
	o.Update(cate, "topic_count")
	return
}

func ModifyTopic(id int64, title, category, lable, content, attachment string) (err error) {
	o := orm.NewOrm()
	topic := &Topic{}
	err = o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		beego.Error("not found topic by id = %v", id)
	}
	if topic.Attachment != attachment {
		os.Remove(path.Join("attachment", topic.Attachment))
	}
	if category != topic.Category {
		cate := &Category{}
		o.QueryTable("category").Filter("title", category).One(cate)
		cate.TopicCount++
		o.Update(cate, "topic_count")

		cate = &Category{}
		o.QueryTable("category").Filter("title", topic.Category).One(cate)
		cate.TopicCount--
		o.Update(cate, "topic_count")
	}
	topic.Title = title
	topic.Category = category
	topic.Lables = lable
	topic.Content = content
	topic.Attachment = attachment
	topic.Updated = time.Now()
	_, err = o.Update(topic, "title", "category", "lables", "content", "attachment", "updated")
	if err != nil {
		beego.Error("ModifyTopic failed", id)
		return
	}
	return
}

func GetOneTopic(id int64) (topic *Topic, err error) {
	o := orm.NewOrm()
	topic = &Topic{}
	err = o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		beego.Error("GetOneTopic failed", err)
		return
	}
	topic.Views++
	o.Update(topic, "views")
	return
}
