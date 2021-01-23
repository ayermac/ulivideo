package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	UserId      int
	Stamp       int
	Status      int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init()  {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error)  {
	o := orm.NewOrm()
	var comments []Comment
	num, _ := o.Raw("SELECT id FROM comment WHERE status =1 AND episodes_id = ?", episodesId).QueryRows(&comments)
	_, err := o.Raw("SELECT id,content,add_time,user_id,stamp,praise_count,episodes_id FROM comment WHERE status=1 AND episodes_id=? ORDER BY add_time DESC LIMIT ?,?", episodesId, offset, limit).QueryRows(&comments)
	return num, comments, err
}

func SaveComment(content string, uid int, episodesId int, videoId int) error  {
	o := orm.NewOrm()
	var comment Comment
	comment.Content = content
	comment.UserId = uid
	comment.EpisodesId = episodesId
	comment.VideoId = videoId
	comment.AddTime = time.Now().Unix()
	_, err := o.Insert(&comment)
	if err == nil {
		// 修改视频总评论数
		o.Raw("UPDATE video SET comment = comment + 1 WHERE id = ?", videoId).Exec()
		// 修改视频剧集评论数
		o.Raw("UPDATE video_episodes SET comment = comment + 1 WHERE id = ?", episodesId).Exec()
	}
	return err
}