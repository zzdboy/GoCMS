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

//来源管理
import "time"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type Copyfrom struct {
	Id         int64  `xorm:"pk autoincr"`
	Sitename   string `xorm:"varchar:(30)"`
	Siteurl    string `xorm:"varchar(100)"`
	Thumb      string `xorm:"varchar(100)"`
	Createtime string `xorm:"DateTime"`
}

//获取来源列表
func (c *Copyfrom) GetByAll(Page int64, Perpage int64) ([]*Copyfrom, template.HTML) {
	copyfrom_list := []*Copyfrom{}

	//查询总数
	copyfrom := new(Copyfrom)
	Total, err := DB_Read.Count(copyfrom)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Copyfrom/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&copyfrom_list)
	return copyfrom_list, pages
}

//获取来源列表
func (c *Copyfrom) GetRoleList() []*Copyfrom {
	copyfrom_list := []*Copyfrom{}
	DB_Read.Find(&copyfrom_list)
	return copyfrom_list
}

//根据Id获取来源信息
func (c *Copyfrom) GetById(Id int64) *Copyfrom {

	copyfrom := new(Copyfrom)
	has, err := DB_Read.Id(Id).Get(copyfrom)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return copyfrom
}

//添加来源
func (c *Copyfrom) Save() bool {

	copyfrom := new(Copyfrom)
	copyfrom.Sitename = c.Sitename
	copyfrom.Siteurl = c.Siteurl
	copyfrom.Thumb = c.Thumb
	copyfrom.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(copyfrom)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑来源
func (c *Copyfrom) Edit(Id int64) bool {
	copyfrom := new(Copyfrom)

	if len(c.Sitename) > 0 {
		copyfrom.Sitename = c.Sitename
	}

	if len(c.Siteurl) > 0 {
		copyfrom.Siteurl = c.Siteurl
	}

	if len(c.Thumb) > 0 {
		copyfrom.Thumb = c.Thumb
		has, err := DB_Write.Id(Id).Cols("sitename", "siteurl", "thumb").Update(copyfrom)

		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	} else {
		has, err := DB_Write.Id(Id).Cols("sitename", "siteurl").Update(copyfrom)

		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	}

	return true
}
