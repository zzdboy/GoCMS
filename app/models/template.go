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

//模板风格
import "time"
import "github.com/revel/revel"

type Template struct {
	Id         int64  `xorm:"pk autoincr"`
	Identity   string `xorm:"varchar(50)"`
	Name       string `xorm:"varchar:(50)"`
	Author     string `xorm:"varchar(50)"`
	Version    string `xorm:"varchar(20)"`
	Status     int64  `xorm:"bool"`
	Createtime string `xorm:"DateTime"`
}

//根据Id获取信息
func (t *Template) GetById(Id int64) *Template {

	template := new(Template)
	has, err := DB_Read.Table("template").Id(Id).Get(template)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return template
}

//获取模板风格列表
func (t *Template) GetTemplateList() []*Template {
	//初始化菜单
	template_list := []*Template{}

	err := DB_Read.Table("template").Find(&template_list)

	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	return template_list
}

//添加模板风格
func (t *Template) Save() bool {
	template := new(Template)

	template.Identity = t.Identity
	template.Name = t.Name
	template.Author = t.Author
	template.Version = t.Version
	template.Status = t.Status
	template.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Table("template").Insert(template)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//设置状态
func (t *Template) Setstatus(Id int64) bool {
	template := new(Template)

	template.Status = t.Status

	has, err := DB_Write.Table("template").Id(Id).Cols("status").Update(template)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑模板风格
func (t *Template) Edit(Id int64) bool {
	template := new(Template)

	template.Identity = t.Identity
	template.Name = t.Name
	template.Author = t.Author
	template.Version = t.Version
	template.Status = t.Status

	has, err := DB_Write.Table("template").Id(Id).Cols("identity", "name", "author", "version", "status").Update(template)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//删除模板风格
func (t *Template) DelByID(Id int64) bool {

	template := new(Template)

	has, err := DB_Write.Id(Id).Delete(template)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
