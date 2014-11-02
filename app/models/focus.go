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

//焦点图列表
import "os"
import "fmt"
import "time"
import "strings"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"
import "github.com/revel/config"

type Focus struct {
	Id         int64      `xorm:"pk autoincr"`
	Cid        int64      `xorm:"index"`
	Focuscate  *FocusCate `xorm:"- <- ->"`
	Aid        int64      `xorm:"index"`
	Admin      *Admin     `xorm:"- <- ->"`
	Title      string     `xorm:"varchar(255)"`
	Url        string     `xorm:"varchar(255)"`
	Img        string     `xorm:"varchar(255)"`
	Content    string     `xorm:"text"`
	Clicks     int64      `xorm:"int(10)"`
	Order      int64      `xorm:"tynyint(5)"`
	Status     int64      `xorm:"tynyint(1)"`
	Createtime string     `xorm:"DateTime"`
	Updatetime string     `xorm:"DateTime"`
}

//根据Id获取内容信息
func (c *Focus) GetById(Id int64) *Focus {
	focus := new(Focus)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(focus)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		admin := new(Admin)
		focus.Admin = admin.GetById(focus.Aid)

		//所属分类
		focuscate := new(FocusCate)
		focus.Focuscate = focuscate.GetById(focus.Cid)

		if focus.Img != "" {
			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//前台网站地址
			sitedomain, _ := config_conf.String("website", "website.sitedomain")
			focus.Img = sitedomain + focus.Img
		}
	}

	return focus
}

//获取焦点图列表
func (c *Focus) GetByAll(search string, Page int64, Perpage int64) (focus_arr []*Focus, html template.HTML, where map[string]interface{}) {
	focus_list := []*Focus{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(search) > 0 {

		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if where["cid"] != "" {
			WhereStr += " `cid`='" + fmt.Sprintf("%d", where["cid"]) + "' AND "
		}

		if where["keyword"] != "" {
			//关键字
			keyword := fmt.Sprintf("%s", where["keyword"])
			WhereStr += " `title` like '%" + keyword + "%' AND "
		}
	}

	WhereStr += " 1 "

	//查询总数
	focus := new(Focus)
	Total, err := DB_Read.Where(WhereStr).Count(focus)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Focus/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&focus_list)

	if len(focus_list) > 0 {
		admin := new(Admin)
		focuscate := new(FocusCate)

		for i, v := range focus_list {
			focus_list[i].Admin = admin.GetById(v.Aid)

			//所属栏目
			focus_list[i].Focuscate = focuscate.GetById(v.Cid)

			if v.Img != "" {
				//判断是否是系统的分隔符
				separator := "/"
				if os.IsPathSeparator('\\') {
					separator = "\\"
				} else {
					separator = "/"
				}

				config_file := (revel.BasePath + "/conf/config.conf")
				config_file = strings.Replace(config_file, "/", separator, -1)
				config_conf, _ := config.ReadDefault(config_file)

				//前台网站地址
				sitedomain, _ := config_conf.String("website", "website.sitedomain")
				v.Img = sitedomain + v.Img
			}
		}
	}

	return focus_list, pages, where
}

//添加焦点图
func (c *Focus) Save() bool {

	focus := new(Focus)
	focus.Cid = c.Cid
	focus.Aid = c.Aid
	focus.Title = c.Title
	focus.Url = c.Url
	focus.Img = c.Img
	focus.Content = c.Content
	focus.Clicks = 0
	focus.Order = c.Order
	focus.Status = 1
	focus.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(focus)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑焦点图
func (c *Focus) Edit(Id int64) bool {
	focus := new(Focus)
	focus.Cid = c.Cid
	focus.Aid = c.Aid
	focus.Title = c.Title
	focus.Url = c.Url
	focus.Img = c.Img
	focus.Content = c.Content
	focus.Order = c.Order
	focus.Status = 1
	focus.Updatetime = time.Now().Format("2006-01-02 15:04:04")

	if len(c.Img) > 0 {
		has, err := DB_Write.Id(Id).Cols("cid", "title", "url", "img", "content", "order", "updatetime").Update(focus)

		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	} else {
		has, err := DB_Write.Id(Id).Cols("cid", "title", "url", "content", "order", "updatetime").Update(focus)
		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	}

	return true
}
