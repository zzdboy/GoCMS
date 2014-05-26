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

//管理员表
import "time"
import "regexp"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"

type Admin struct {
	Id            int64  `xorm:"pk autoincr"`
	Username      string `xorm:"unique index varchar(255)"`
	Password      string `xorm:"varchar:(32)"`
	Roleid        int64  `xorm:"index"`
	Role          *Role  `xorm:"- <- ->"`
	Lastloginip   string `xorm:"varchar(32)"`
	Lastlogintime string `xorm:"varchar(32)"`
	Email         string `xorm:"varchar(32)"`
	Realname      string `xorm:"varchar(32)"`
	Lang          string `xorm:"varchar(6)"`
	Status        int64  `xorm:"bool"`
	Createtime    string `xorm:"DateTime"`
}

type Password struct {
	Password        string
	PasswordConfirm string
}

func (a *Admin) Validate(v *revel.Validation) {
	v.Required(a.Username).Message("请输入用户名!")
	valid := v.Match(a.Username, regexp.MustCompile("^\\w*$")).Message("只能使用字母、数字和下划线!")
	if valid.Ok {
		if a.HasName() {
			err := &revel.ValidationError{
				Message: "该用户名已经注册过!",
				Key:     "a.Username",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(a.Email).Message("请输入Email")
	valid = v.Email(a.Email).Message("无效的电子邮件!")
	if valid.Ok {
		if a.HasEmail() {
			err := &revel.ValidationError{
				Message: "该邮件已经注册过!",
				Key:     "a.Email",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(a.Password).Message("请输入密码!")
	v.MinSize(a.Password, 3).Message("密码最少三位!")
}

//验证密码
func (P *Password) ValidatePassword(v *revel.Validation) {
	v.Required(P.Password).Message("请输入密码!")
	v.Required(P.PasswordConfirm).Message("请输入确认密码!")

	v.MinSize(P.Password, 6).Message("密码最少六位!")
	v.Required(P.Password == P.PasswordConfirm).Message("两次密码不相同!")
}

//获取管理员列表
func (a *Admin) GetByAll(RoleId int64, where map[string]string, Page int64, Perpage int64) ([]*Admin, template.HTML) {
	admin_list := []*Admin{}

	if RoleId > 0 {

		//查询总数
		admin := new(Admin)
		Total, err := DB_Read.Where("roleid=?", RoleId).Count(admin)
		if err != nil {
			revel.WARN.Printf("错误: %v", err)
		}

		//分页
		Pager := new(utils.Page)
		Pager.SubPage_link = "/Admin/"
		Pager.Nums = Total
		Pager.Perpage = Perpage
		Pager.Current_page = Page
		Pager.SubPage_type = 2
		pages := Pager.Show()

		DB_Read.Where("roleid=?", RoleId).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&admin_list)

		if len(admin_list) > 0 {
			role := new(Role)

			for i, v := range admin_list {
				admin_list[i].Role = role.GetById(v.Roleid)
			}
		}

		return admin_list, pages
	} else {

		//查询总数
		admin := new(Admin)
		Total, err := DB_Read.Count(admin)
		if err != nil {
			revel.WARN.Printf("错误: %v", err)
		}

		//分页
		Pager := new(utils.Page)
		Pager.SubPage_link = "/Admin/"
		Pager.Nums = Total
		Pager.Perpage = Perpage
		Pager.Current_page = Page
		Pager.SubPage_type = 2
		pages := Pager.Show()

		DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&admin_list)

		if len(admin_list) > 0 {
			role := new(Role)

			for i, v := range admin_list {
				admin_list[i].Role = role.GetById(v.Roleid)
			}
		}

		return admin_list, pages
	}
}

func (a *Admin) HasName() bool {

	admin := new(Admin)
	has, err := DB_Read.Where("username=?", a.Username).Get(admin)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if admin.Id > 0 {
		return true
	}
	return false
}

func (a *Admin) HasEmail() bool {

	admin := new(Admin)
	has, err := DB_Read.Where("email=?", a.Email).Get(admin)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if admin.Id > 0 {
		return true
	}
	return false
}

//根据Id获取管理员信息
func (a *Admin) GetById(Id int64) *Admin {
	admin := new(Admin)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		role := new(Role)
		admin.Role = role.GetById(admin.Roleid)
	}

	return admin
}

//根据真实姓名获取管理员信息
func (a *Admin) GetByRealName(name string) *Admin {
	admin := new(Admin)
	has, err := DB_Read.Where("realname=?", name).Get(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		role := new(Role)
		admin.Role = role.GetById(admin.Roleid)
	}

	return admin
}

//根据用户名获取管理员信息
func (a *Admin) GetByName(name string) *Admin {
	admin := new(Admin)
	has, err := DB_Read.Where("username=?", name).Get(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		role := new(Role)
		admin.Role = role.GetById(admin.Roleid)
	}

	return admin
}

//添加管理员
func (a *Admin) Save() bool {

	admin := new(Admin)
	admin.Username = a.Username
	admin.Password = utils.Md5(a.Password)
	admin.Roleid = a.Roleid
	admin.Lastloginip = utils.GetClientIP()
	admin.Email = a.Email
	admin.Realname = a.Realname
	admin.Lang = a.Lang
	admin.Lastlogintime = "0000-00-00 00:00:00"
	admin.Status = a.Status
	admin.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

// 更新登陆时间
func (a *Admin) UpdateLoginTime(Id int64) bool {
	admin := new(Admin)

	admin.Lastloginip = utils.GetClientIP()
	admin.Lastlogintime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Id(Id).Cols("lastloginip", "lastlogintime").Update(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//修改个人信息
func (a *Admin) EditInfo(Id int64) bool {
	admin := new(Admin)

	if len(a.Email) > 0 {
		admin.Email = a.Email
	}

	if len(a.Realname) > 0 {
		admin.Realname = a.Realname
	}

	if len(a.Lang) > 0 {
		admin.Lang = a.Lang
	}

	has, err := DB_Write.Id(Id).Cols("email", "realname", "lang").Update(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//修改密码
func (a *Admin) EditPwd(Id int64) bool {
	admin := new(Admin)

	if len(a.Password) > 0 {
		admin.Password = utils.Md5(a.Password)
	}

	has, err := DB_Write.Id(Id).Cols("password").Update(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//编辑管理员
func (a *Admin) Edit(Id int64) bool {

	admin := new(Admin)

	if len(a.Username) > 0 {
		admin.Username = a.Username
	}

	if len(a.Password) > 0 {
		admin.Password = utils.Md5(a.Password)
	}

	if a.Roleid > 0 {
		admin.Roleid = a.Roleid
	}

	if len(a.Email) > 0 {
		admin.Email = a.Email
	}

	if len(a.Realname) > 0 {
		admin.Realname = a.Realname
	}

	if len(a.Lang) > 0 {
		admin.Lang = a.Lang
	}

	admin.Status = a.Status

	if len(a.Password) > 0 {
		has, err := DB_Write.Id(Id).Cols("username", "password", "email", "realname", "roleid", "lang", "status").Update(admin)
		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	} else {
		has, err := DB_Write.Id(Id).Cols("username", "email", "realname", "roleid", "lang", "status").Update(admin)
		if err != nil {
			revel.WARN.Println(has)
			revel.WARN.Printf("错误: %v", err)
			return false
		}
	}

	return true
}

//删除管理员
func (a *Admin) DelByID(Id int64) bool {

	admin := new(Admin)

	has, err := DB_Write.Id(Id).Delete(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//获取MySQL版本
func (a *Admin) GetMysqlVer() string {
	sql := "SELECT VERSION() AS version;"
	results, err := DB_Read.Query(sql)

	if err != nil {
		revel.WARN.Println(results)
		revel.WARN.Printf("错误: %v", err)
	}

	return string(results[0]["version"])
}
