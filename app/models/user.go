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

//会员
import "fmt"
import "time"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type User struct {
	Id          int64       `xorm:"pk"`
	Email       string      `xorm:"char(32)"`
	Username    string      `xorm:"char(20)"`
	Password    string      `xorm:"char(32)"`
	Encrypt     string      `xorm:"char(6)"`
	Nickname    string      `xorm:"char(20)"`
	Mobile      string      `xorm:"char(11)"`
	Birthday    string      `xorm:"Date"`
	Regip       string      `xorm:"char(11)"`
	Regdate     string      `xorm:"DateTime"`
	Lastdate    string      `xorm:"DateTime"`
	Lastip      string      `xorm:"char(11)"`
	Loginnum    int64       `xorm:"int(11)"`
	Groupid     int64       `xorm:"int(11)"`
	UserGroup   *User_Group `xorm:"- <- ->"`
	Areaid      int64       `xorm:"int(11)"`
	Amount      float64     `xorm:"float(8,2)"`
	Point       int64       `xorm:"smallint(5)"`
	Ismessage   int64       `xorm:"bool"`
	Islock      int64       `xorm:"bool"`
	Vip         int64       `xorm:"tinyint(1)"`
	Overduedate string      `xorm:"DateTime"`
	Status      int64       `xorm:"bool"`
	Createtime  string      `xorm:"DateTime"`
}

//根据Id获取信息
func (u *User) GetById(Id int64) *User {

	user := new(User)
	has, err := DB_Read.Table("user").Id(Id).Get(user)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		user_group := new(User_Group)
		user.UserGroup = user_group.GetById(user.Groupid)
	}

	return user
}

//用户名是否已有
func (u *User) HasName() bool {

	user := new(User)
	has, err := DB_Read.Table("user").Where("username=?", u.Username).Get(user)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if user.Id > 0 {
		return true
	}
	return false
}

//邮箱是否已有
func (u *User) HasEmail() bool {

	user := new(User)
	has, err := DB_Read.Table("user").Where("email=?", u.Email).Get(user)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if user.Id > 0 {
		return true
	}
	return false
}

//获取会员列表
func (u *User) GetUserList(search string, Page int64, Perpage int64) (user_arr []*User, html template.HTML, where map[string]interface{}) {
	//初始化菜单
	user_list := []*User{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(search) > 0 {
		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if where["start_time"] != "" {
			WhereStr += " `regdate` >='" + fmt.Sprintf("%s", where["start_time"]) + " 00:00:00' AND "
		}

		if where["end_time"] != "" {
			WhereStr += " `regdate` <='" + fmt.Sprintf("%s", where["end_time"]) + " 23:59:59' AND "
		}

		if where["islock"] != "" && where["islock"] != "0" {
			WhereStr += " `islock` =" + fmt.Sprintf("%s", where["islock"])
		}

		if where["type"] != "" && where["keyword"] != "" {

			if where["type"] == "1" {
				//用户名
				WhereStr += " `username` ='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			} else if where["type"] == "2" {
				//用户ID
				WhereStr += " `id` =" + fmt.Sprintf("%s", where["keyword"]) + " AND "
			} else if where["type"] == "3" {
				//邮箱
				WhereStr += " `email` ='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			} else if where["type"] == "4" {
				//注册ip
				WhereStr += " `regip` ='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			} else if where["type"] == "5" {
				//昵称
				WhereStr += " `nickname` like '%" + fmt.Sprintf("%s", where["keyword"]) + "%' AND "
			}
		}
	}

	WhereStr += " 1 "

	//查询总数
	user := new(User)
	Total, err := DB_Read.Table("user").Where(WhereStr).Count(user)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	if len(search) > 0 {
		Pager.SubPage_link = "/User/" + search + "/"
	} else {
		Pager.SubPage_link = "/User/"
	}

	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	DB_Read.Table("user").Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&user_list)

	if len(user_list) > 0 {
		user_group := new(User_Group)
		for i, v := range user_list {
			user_list[i].UserGroup = user_group.GetById(v.Groupid)
		}
	}

	return user_list, pages, where
}

//添加会员
func (u *User) Save() bool {
	user := new(User)

	user.Email = u.Email
	user.Username = u.Username
	user.Password = utils.Md5(u.Password)
	user.Encrypt = utils.RandomString(6)
	user.Nickname = u.Nickname
	user.Mobile = u.Mobile
	user.Birthday = u.Birthday
	user.Regip = u.Regip
	user.Regdate = time.Now().Format("2006-01-02 15:04:04")
	user.Lastdate = time.Now().Format("2006-01-02 15:04:04")
	user.Lastip = u.Lastip
	user.Loginnum = 0
	user.Groupid = u.Groupid
	user.Areaid = u.Areaid
	user.Amount = 0
	user.Point = u.Point
	user.Ismessage = u.Ismessage
	user.Islock = u.Islock
	user.Vip = u.Vip
	user.Overduedate = u.Overduedate
	user.Status = u.Status
	user.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Table("user").Insert(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//添加会员
func (u *User) Edit(Id int64) bool {
	user := new(User)

	user.Email = u.Email
	if len(u.Password) > 0 {
		user.Password = utils.Md5(u.Password)
	}
	user.Nickname = u.Nickname
	user.Mobile = u.Mobile
	user.Birthday = u.Birthday
	user.Groupid = u.Groupid
	user.Point = u.Point
	user.Islock = u.Islock
	user.Vip = u.Vip
	user.Overduedate = u.Overduedate

	if len(u.Password) > 0 {
		has, err := DB_Write.Table("user").Id(Id).Cols("email", "password", "nickname", "mobile", "birthday", "groupid", "point", "islock", "vip", "overduedate").Update(user)
		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	} else {
		has, err := DB_Write.Table("user").Id(Id).Cols("email", "nickname", "mobile", "birthday", "groupid", "point", "islock", "vip", "overduedate").Update(user)
		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	}

	return true
}

//删除用户
func (u *User) DelByID(Id int64) bool {

	user := new(User)

	has, err := DB_Write.Id(Id).Delete(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//批量删除
func (u *User) DelByIDS(ids string) bool {
	user := new(User)

	has, err := DB_Write.Where("id in (" + ids + ")").Delete(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//锁定
func (u *User) Lock(ids string) bool {
	user := new(User)

	user.Islock = 1

	has, err := DB_Write.Where("id in (" + ids + ")").Cols("islock").Update(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//解锁
func (u *User) Unlock(ids string) bool {
	user := new(User)

	user.Islock = 2

	has, err := DB_Write.Where("id in (" + ids + ")").Cols("islock").Update(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//移动
func (u *User) Move(groupid int64, ids string) bool {
	user := new(User)

	user.Groupid = groupid

	has, err := DB_Write.Where("id in (" + ids + ")").Cols("groupid").Update(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
