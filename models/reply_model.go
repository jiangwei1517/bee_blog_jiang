package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// 评论
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func AddReply(id int64, nickName, content string) (err error) {
	o := orm.NewOrm()
	topic := &Topic{}
	err = o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		beego.Error("AddReply failed not found topic id = %v", id, err)
		return
	}

	comment := &Comment{
		Tid:     id,
		Name:    nickName,
		Content: content,
		Created: time.Now(),
	}
	_, err = o.Insert(comment)
	if err != nil {
		beego.Error("AddReply Insert failed", err)
		return
	}
	topic.ReplyTime = time.Now()
	topic.ReplyCount++
	topic.ReplyLastUserName = nickName
	topic.Updated = time.Now()
	o.Update(topic, "reply_time", "reply_count", "reply_last_user_name", "updated")
	return
}

func GetAllReplys(tid int64) (list []*Comment, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("comment").Filter("tid", tid).OrderBy("-created").All(&list)
	if err != nil {
		if err == orm.ErrNoRows {
			err = nil
			return
		}
		beego.Error("GetAllReplys failed", err)
		return
	}
	return
}

func DeleteReply(tid, rid int64) (err error) {
	o := orm.NewOrm()
	comment := &Comment{
		Id: rid,
	}
	_, err = o.Delete(comment)
	if err != nil {
		beego.Error("DeleteReply failed", err)
		return
	}
	comment = &Comment{}
	o.QueryTable("comment").OrderBy("-created").One(comment)

	topic := &Topic{}
	o.QueryTable("topic").Filter("id", tid).One(topic)
	topic.Updated = time.Now()
	topic.ReplyCount--
	topic.ReplyTime = comment.Created
	o.Update(topic, "updated", "reply_count", "reply_time")
	return
}
