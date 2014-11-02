// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package models

//内容评论表
import "time"
import "fmt"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type Comment struct {
	Id         int64  `xorm:"pk autoincr"`
	Replyid    int64  `xorm:"index"`
	Content    string `xorm:"text"`
	Uid        int64  `xorm:"index"`
	User       *User  `xorm:"- <- ->"`
	Agree      int64  `xorm:"int(11)"`
	Against    int64  `xorm:"int(11)"`
	Ip         string `xorm:"char(11)"`
	Createtime string `xorm:"DateTime"`
}

//根据Id获取信息
func (c *Comment) GetById(Id int64) *Comment {

	comment := new(Comment)
	has, err := DB_Read.Table("comment").Id(Id).Get(comment)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		user := new(User)
		comment.User = user.GetById(comment.Uid)
	}

	return comment
}

//获取评论列表
func (c *Comment) GetCommentList(search string, Page int64, Perpage int64) (comment_arr []*Comment, html template.HTML, where map[string]interface{}) {

	comment_list := []*Comment{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(search) > 0 {

		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if where["start_time"] != "" {
			WhereStr += " `regdate` >='" + fmt.Sprintf("%s", where["start_time"]) + " 00:00:00' AND "
		}
	}

	WhereStr += " 1 "

	//查询总数
	comment := new(Comment)
	Total, err := DB_Read.Table("comment").Where(WhereStr).Count(comment)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	if len(search) > 0 {
		Pager.SubPage_link = "/Comment/" + search + "/"
	} else {
		Pager.SubPage_link = "/Comment/"
	}

	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	DB_Read.Table("comment").Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&comment_list)

	if len(comment_list) > 0 {
		user := new(User)
		for i, v := range comment_list {
			comment_list[i].User = user.GetById(v.Uid)
		}
	}

	return comment_list, pages, where
}

//添加评论
func (c *Comment) Save() bool {
	comment := new(Comment)

	comment.Replyid = c.Replyid
	comment.Content = c.Content
	comment.Uid = c.Uid
	comment.Agree = 0
	comment.Against = 0
	comment.Ip = c.Ip
	comment.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Table("comment").Insert(comment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
