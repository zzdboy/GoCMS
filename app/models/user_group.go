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

//会员组
import "time"
import "github.com/revel/revel"

type User_Group struct {
	Id               int64   `xorm:"pk"`
	Name             string  `xorm:"char(15)"`
	Issystem         int64   `xorm:"bool"`
	Usernum          int64   `xorm:"- <- ->"`
	Star             int64   `xorm:"tinyint(2)"`
	Point            int64   `xorm:"smallint(6)"`
	Allowmessage     int64   `xorm:"tinyint(5)"`
	Allowvisit       int64   `xorm:"bool"`
	Allowpost        int64   `xorm:"bool"`
	Allowpostverify  int64   `xorm:"bool"`
	Allowsearch      int64   `xorm:"bool"`
	Allowupgrade     int64   `xorm:"bool"`
	Allowsendmessage int64   `xorm:"bool"`
	Allowpostnum     int64   `xorm:"bool"`
	Allowattachment  int64   `xorm:"bool"`
	Priceyear        float64 `xorm:"float(8,2)"`
	Pricemonth       float64 `xorm:"float(8,2)"`
	Priceday         float64 `xorm:"float(8,2)"`
	Icon             string  `xorm:"char(100)"`
	Usernamecolor    string  `xorm:"char(7)"`
	Desc             string  `xorm:"char(100)"`
	Status           int64   `xorm:"bool"`
	Createtime       string  `xorm:"DateTime"`
}

//根据Id获取信息
func (u *User_Group) GetById(Id int64) *User_Group {

	user_group := new(User_Group)
	has, err := DB_Read.Table("user_group").Id(Id).Get(user_group)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return user_group
}

//获取会员组列表
func (u *User_Group) GetGroupList() []*User_Group {
	//初始化菜单
	group_list := []*User_Group{}

	err := DB_Read.Table("user_group").Find(&group_list)

	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	if len(group_list) > 0 {
		for i, v := range group_list {
			group_list[i].Usernum = v.Id
		}
	}

	return group_list
}

//添加会员组
func (u *User_Group) Save() bool {
	user_group := new(User_Group)

	user_group.Name = u.Name
	user_group.Issystem = u.Issystem
	user_group.Star = u.Star
	user_group.Point = u.Point
	user_group.Allowmessage = u.Allowmessage
	user_group.Allowvisit = u.Allowvisit
	user_group.Allowpost = u.Allowpost
	user_group.Allowpostverify = u.Allowpostverify
	user_group.Allowsearch = u.Allowsearch
	user_group.Allowupgrade = u.Allowupgrade
	user_group.Allowsendmessage = u.Allowsendmessage
	user_group.Allowpostnum = u.Allowpostnum
	user_group.Allowattachment = u.Allowattachment
	user_group.Priceyear = u.Priceyear
	user_group.Pricemonth = u.Pricemonth
	user_group.Priceday = u.Priceday
	user_group.Icon = u.Icon
	user_group.Usernamecolor = u.Usernamecolor
	user_group.Desc = u.Desc
	user_group.Status = u.Status
	user_group.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Table("user_group").Insert(user_group)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑会员组
func (u *User_Group) Edit(Id int64) bool {
	user_group := new(User_Group)

	user_group.Name = u.Name
	user_group.Star = u.Star
	user_group.Point = u.Point
	user_group.Allowmessage = u.Allowmessage
	user_group.Allowvisit = u.Allowvisit
	user_group.Allowpost = u.Allowpost
	user_group.Allowpostverify = u.Allowpostverify
	user_group.Allowsearch = u.Allowsearch
	user_group.Allowupgrade = u.Allowupgrade
	user_group.Allowsendmessage = u.Allowsendmessage
	user_group.Allowpostnum = u.Allowpostnum
	user_group.Allowattachment = u.Allowattachment
	user_group.Priceyear = u.Priceyear
	user_group.Pricemonth = u.Pricemonth
	user_group.Priceday = u.Priceday
	user_group.Icon = u.Icon
	user_group.Usernamecolor = u.Usernamecolor
	user_group.Desc = u.Desc
	user_group.Status = u.Status

	has, err := DB_Write.Table("user_group").Id(Id).Cols("name", "star", "point", "allowmessage", "allowvisit", "allowpost", "allowpostverify", "allowsearch", "allowupgrade", "allowsendmessage", "allowpostnum", "allowattachment", "priceyear", "pricemonth", "priceday", "icon", "usernamecolor", "desc", "status").Update(user_group)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
