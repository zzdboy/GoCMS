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

//常用菜单
import "time"
import "github.com/revel/revel"

type Admin_Panel struct {
	Id         int64  `xorm:"pk"`
	Mid        int64  `xorm:"int(11)"`
	Menu       *Menu  `xorm:"- <- ->"`
	Aid        int64  `xorm:"int(11)"`
	Name       string `xorm:"varchar(40)"`
	Url        string `xorm:"char(100)"`
	Createtime string `xorm:"DateTime"`
}

//根据Id获取信息
func (a *Admin_Panel) GetById(Id int64) *Admin_Panel {

	admin_panel := new(Admin_Panel)
	has, err := DB_Read.Table("admin_panel").Id(Id).Get(admin_panel)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return admin_panel
}

//获取快捷方式列表
func (a *Admin_Panel) GetPanelList(Admin_Info *Admin) []*Admin_Panel {
	//初始化菜单
	admin_panel := []*Admin_Panel{}

	err := DB_Read.Table("admin_panel").Where("aid=?", Admin_Info.Id).Find(&admin_panel)

	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	} else {
		menu := new(Menu)

		for i, v := range admin_panel {
			admin_panel[i].Menu = menu.GetById(v.Mid)
		}
	}

	return admin_panel
}

//根据mid获取快捷方式
func (a *Admin_Panel) GetByMid(Mid int64, Admin_Info *Admin) *Admin_Panel {
	admin_panel := new(Admin_Panel)
	has, err := DB_Read.Table("admin_panel").Where("mid=? and aid=?", Mid, Admin_Info.Id).Get(admin_panel)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return admin_panel
}

//是否已添加快捷方式
func (a *Admin_Panel) IsAdd(Mid int64, Admin_Info *Admin) bool {
	admin_panel := new(Admin_Panel)

	has, err := DB_Read.Table("admin_panel").Where("mid=? AND aid=?", Mid, Admin_Info.Id).Get(admin_panel)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	if admin_panel.Id > 0 {
		return true
	} else {
		return false
	}
}

//删除快捷方式
func (a *Admin_Panel) DelPanel(Mid int64, Admin_Info *Admin) bool {
	admin_panel := new(Admin_Panel)

	has, err := DB_Write.Table("admin_panel").Where("mid=? AND aid=?", Mid, Admin_Info.Id).Delete(admin_panel)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//添加快捷方式
func (a *Admin_Panel) AddPanel(Mid int64, Admin_Info *Admin) bool {
	admin_panel := new(Admin_Panel)

	admin_panel.Mid = Mid
	admin_panel.Aid = Admin_Info.Id

	menu := new(Menu)

	//获取菜单信息
	menu_info := menu.GetById(Mid)

	admin_panel.Name = menu_info.Name
	admin_panel.Url = menu_info.Url
	admin_panel.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Table("admin_panel").Insert(admin_panel)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
