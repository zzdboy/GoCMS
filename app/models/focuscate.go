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

//焦点图分类表
import "time"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type FocusCate struct {
	Id         int64  `xorm:"pk autoincr"`
	Name       string `xorm:"varchar(255)"`
	Width      int64  `xorm:"int(6)"`
	Height     int64  `xorm:"int(6)"`
	Createtime string `xorm:"DateTime"`
}

//根据Id获取焦点图分类
func (c *FocusCate) GetById(Id int64) *FocusCate {

	focuscate := new(FocusCate)
	has, err := DB_Read.Id(Id).Get(focuscate)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return focuscate
}

//获取分类列表
func (c *FocusCate) GetCateList() []*FocusCate {
	cate_list := []*FocusCate{}
	DB_Read.Find(&cate_list)
	return cate_list
}

//获取分类列表
func (c *FocusCate) GetByAll(Page int64, Perpage int64) ([]*FocusCate, template.HTML) {
	focuscate_list := []*FocusCate{}

	//查询总数
	focuscate := new(FocusCate)
	Total, err := DB_Read.Count(focuscate)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/FocusCate/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&focuscate_list)
	return focuscate_list, pages
}

//添加分类
func (c *FocusCate) Save() bool {

	focuscate := new(FocusCate)
	focuscate.Name = c.Name
	focuscate.Width = c.Width
	focuscate.Height = c.Height
	focuscate.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(focuscate)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑分类
func (c *FocusCate) Edit(Id int64) bool {
	focuscate := new(FocusCate)

	if len(c.Name) > 0 {
		focuscate.Name = c.Name
	}

	if c.Width > 0 {
		focuscate.Width = c.Width
	}

	if c.Height > 0 {
		focuscate.Height = c.Height
	}

	has, err := DB_Write.Id(Id).Cols("name", "width", "height").Update(focuscate)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//删除分类
func (c *FocusCate) DelByID(Id int64) bool {
	focuscate := new(FocusCate)

	has, err := DB_Write.Id(Id).Delete(focuscate)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
