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

//角色表
import "time"
import "html/template"
import "admin/utils"
import "github.com/revel/revel"

type Role struct {
	Id         int64  `xorm:"pk"`
	Rolename   string `xorm:"unique varchar(255)"`
	Desc       string `xorm:"varchar:(255)"`
	Data       string `xorm:"text"`
	Status     int64  `xorm:"bool"`
	Createtime string `xorm:"DateTime"`
}

//根据Id获取角色信息
func (r *Role) GetById(Id int64) *Role {

	role := new(Role)
	has, err := DB_Read.Id(Id).Get(role)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return role
}

//获取角色列表
func (r *Role) GetByAll(Page int64, Perpage int64) ([]*Role, template.HTML) {
	role_list := []*Role{}

	//查询总数
	role := new(Role)
	Total, err := DB_Read.Count(role)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Role/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&role_list)
	return role_list, pages
}

//获取角色
func (r *Role) GetRoleList() []*Role {
	role_list := []*Role{}
	DB_Read.Find(&role_list)
	return role_list
}

//添加角色
func (r *Role) Save() bool {

	role := new(Role)
	role.Rolename = r.Rolename
	role.Desc = r.Desc
	role.Data = r.Data
	role.Status = r.Status
	role.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(role)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑角色
func (r *Role) Edit(Id int64) bool {
	role := new(Role)

	if len(r.Rolename) > 0 {
		role.Rolename = r.Rolename
	}

	if len(r.Desc) > 0 {
		role.Desc = r.Desc
	}

	if len(r.Data) > 0 {
		role.Data = r.Data
	}

	role.Status = r.Status

	has, err := DB_Write.Id(Id).Cols("rolename", "desc", "data", "status").Update(role)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//设置状态
func (r *Role) SetStatus(Id int64) bool {
	role := new(Role)

	role.Status = r.Status

	has, err := DB_Write.Id(Id).Cols("status").Update(role)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//删除角色
func (r *Role) DelByID(Id int64) bool {

	role := new(Role)

	has, err := DB_Write.Id(Id).Delete(role)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
