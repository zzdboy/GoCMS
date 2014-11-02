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

//投诉建议
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type Complaints struct {
	Id          int64  `xorm:"pk"`
	Type        int64  `xorm:"tinyint(1)"`
	Question    string `xorm:"varchar(255)"`
	Companyname string `xorm:"varchar(255)"`
	Name        string `xorm:"varchar(50)"`
	Phone       string `xorm:"varchar(50)"`
	Email       string `xorm:"varchar(50)"`
	Qq          string `xorm:"varchar(50)"`
	Createtime  string `xorm:"DateTime"`
}

//根据Id获取投诉建议信息
func (c *Complaints) GetById(Id int64) *Complaints {

	complaints := new(Complaints)
	has, err := DB_Read.Id(Id).Get(complaints)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return complaints
}

//获取投诉建议列表
func (c *Complaints) GetByAll(Page int64, Perpage int64) ([]*Complaints, template.HTML) {
	complaints_list := []*Complaints{}

	//查询总数
	complaints := new(Complaints)
	Total, err := DB_Read.Count(complaints)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Complaints/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&complaints_list)
	return complaints_list, pages
}

//删除投诉建议
func (c *Complaints) DelByID(Id int64) bool {
	cmplaints := new(Complaints)

	has, err := DB_Write.Id(Id).Delete(cmplaints)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
