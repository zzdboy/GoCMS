// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package controllers

//会员管理 登陆
import "fmt"
import "time"
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"
import "admin/utils"
import "github.com/dchest/captcha"

type User struct {
	*revel.Controller
}

//首页
func (c *User) Index(user *models.User) revel.Result {
	title := "会员首页--GoCMS管理系统"

	var page string = c.Params.Get("page")
	var search string = c.Params.Get("search")

	var Page int64 = 1

	if len(page) > 0 {
		Page, _ = strconv.ParseInt(page, 10, 64)
	}

	user_list, pages, where := user.GetUserList(search, Page, 10)

	//会员组
	user_group := new(models.User_Group)
	group_list := user_group.GetGroupList()

	c.Render(title, user_list, group_list, pages, where)
	return c.RenderTemplate("User/Manage/Index.html")
}

//添加会员
func (c *User) Add(user *models.User) revel.Result {

	if c.Request.Method == "GET" {

		title := "添加会员--GoCMS管理系统"

		//会员组
		user_group := new(models.User_Group)
		group_list := user_group.GetGroupList()

		c.Render(title, group_list)
		return c.RenderTemplate("User/Manage/Add.html")
	} else {

		data := make(map[string]string)

		var username string = c.Params.Get("username")
		if len(username) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入用户名!"
			return c.RenderJson(data)
		} else {
			user.Username = username
		}

		if user.HasName() {
			data["status"] = "0"
			data["message"] = "用户名‘" + username + "’已有!"
			return c.RenderJson(data)
		}

		var email string = c.Params.Get("email")
		if len(email) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入邮箱!"
			return c.RenderJson(data)
		} else {
			user.Email = email
		}

		if user.HasEmail() {
			data["status"] = "0"
			data["message"] = "邮箱‘" + email + "’已有!"
			return c.RenderJson(data)
		}

		var password string = c.Params.Get("password")
		if len(password) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入密码!"
			return c.RenderJson(data)
		}

		var pwdconfirm string = c.Params.Get("pwdconfirm")
		if len(pwdconfirm) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入确认密码!"
			return c.RenderJson(data)
		} else {
			if password != pwdconfirm {
				data["status"] = "0"
				data["message"] = "两次输入密码不一致!"
				return c.RenderJson(data)
			} else {
				user.Password = password
			}
		}

		var nickname string = c.Params.Get("nickname")
		if len(nickname) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入昵称!"
			return c.RenderJson(data)
		} else {
			user.Nickname = nickname
		}

		var mobile string = c.Params.Get("mobile")
		if len(mobile) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入手机号码!"
			return c.RenderJson(data)
		} else {
			user.Mobile = mobile
		}

		var groupid string = c.Params.Get("groupid")
		if len(groupid) <= 0 {
			data["status"] = "0"
			data["message"] = "请选择会员组!"
			return c.RenderJson(data)
		} else {
			Groupid, err := strconv.ParseInt(groupid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user.Groupid = Groupid
		}

		var islock string = c.Params.Get("islock")
		if len(islock) <= 0 {
			data["status"] = "0"
			data["message"] = "请选择是否定锁!"
			return c.RenderJson(data)
		} else {
			Islock, err := strconv.ParseInt(islock, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user.Islock = Islock
		}

		var point string = c.Params.Get("point")
		if len(point) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入积分点数!"
			return c.RenderJson(data)
		} else {
			Point, err := strconv.ParseInt(point, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user.Point = Point
		}

		var birthday string = c.Params.Get("birthday")
		if len(birthday) <= 0 {
			user.Birthday = "0000-00-00"
		} else {
			user.Birthday = birthday
		}

		var vip string = c.Params.Get("vip")
		if vip == "1" {
			user.Vip = 1

			var overduedate string = c.Params.Get("overduedate")

			if len(overduedate) <= 0 {
				data["status"] = "0"
				data["message"] = "请选择过期时间!"
				return c.RenderJson(data)
			} else {
				user.Overduedate = overduedate
			}

		} else {
			user.Vip = 0
			user.Overduedate = time.Now().Format("2006-01-02 15:04:04")
		}

		if user.Save() {
			data["status"] = "1"
			data["message"] = "添加会员成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "添加会员失败!"
			return c.RenderJson(data)
		}
	}
}

//编辑会员
func (c *User) Edit(user *models.User) revel.Result {

	if c.Request.Method == "GET" {

		title := "编辑会员--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//内容
			user_info := user.GetById(Id)

			//会员组
			user_group := new(models.User_Group)
			group_list := user_group.GetGroupList()

			c.Render(title, id, group_list, user_info)
		} else {
			c.Render(title, id)
		}

		return c.RenderTemplate("User/Manage/Edit.html")
	} else {
		data := make(map[string]string)

		var id string = c.Params.Get("id")
		if len(id) <= 0 {
			data["status"] = "0"
			data["message"] = "参数错误，编辑失败!"
			return c.RenderJson(data)
		}

		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		var email string = c.Params.Get("email")
		if len(email) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入邮箱!"
			return c.RenderJson(data)
		} else {
			user.Email = email
		}

		var password string = c.Params.Get("password")

		var pwdconfirm string = c.Params.Get("pwdconfirm")
		if len(pwdconfirm) > 0 {
			if password != pwdconfirm {
				data["status"] = "0"
				data["message"] = "两次输入密码不一致!"
				return c.RenderJson(data)
			} else {
				user.Password = password
			}
		}

		var nickname string = c.Params.Get("nickname")
		if len(nickname) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入昵称!"
			return c.RenderJson(data)
		} else {
			user.Nickname = nickname
		}

		var mobile string = c.Params.Get("mobile")
		if len(mobile) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入手机号码!"
			return c.RenderJson(data)
		} else {
			user.Mobile = mobile
		}

		var groupid string = c.Params.Get("groupid")
		if len(groupid) <= 0 {
			data["status"] = "0"
			data["message"] = "请选择会员组!"
			return c.RenderJson(data)
		} else {
			Groupid, err := strconv.ParseInt(groupid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user.Groupid = Groupid
		}

		var islock string = c.Params.Get("islock")
		if len(islock) <= 0 {
			data["status"] = "0"
			data["message"] = "请选择是否定锁!"
			return c.RenderJson(data)
		} else {
			Islock, err := strconv.ParseInt(islock, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user.Islock = Islock
		}

		var point string = c.Params.Get("point")
		if len(point) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入积分点数!"
			return c.RenderJson(data)
		} else {
			Point, err := strconv.ParseInt(point, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user.Point = Point
		}

		var birthday string = c.Params.Get("birthday")
		if len(birthday) <= 0 {
			user.Birthday = "0000-00-00"
		} else {
			user.Birthday = birthday
		}

		var vip string = c.Params.Get("vip")
		if vip == "1" {
			var overduedate string = c.Params.Get("overduedate")

			if len(overduedate) <= 0 {
				data["status"] = "0"
				data["message"] = "请选择过期时间!"
				return c.RenderJson(data)
			} else {
				user.Vip = 1
				user.Overduedate = overduedate
			}

		} else {
			user.Vip = 0
			user.Overduedate = "0000-00-00 00:00:00"
		}

		if user.Edit(Id) {
			data["status"] = "1"
			data["message"] = "编辑会员成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "编辑会员失败!"
			return c.RenderJson(data)
		}
	}
}

//删除
func (c *User) Delete(user *models.User) revel.Result {

	var ids string = c.Params.Get("ids")

	data := make(map[string]string)

	if len(ids) <= 0 {
		data["status"] = "0"
		data["message"] = "请至少选择一个!"
		return c.RenderJson(data)
	}

	if user.DelByIDS(ids) {
		data["status"] = "1"
		data["message"] = "删除成功!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
		return c.RenderJson(data)
	}
}

//锁定
func (c *User) Lock(user *models.User) revel.Result {

	var ids string = c.Params.Get("ids")

	data := make(map[string]string)

	if len(ids) <= 0 {
		data["status"] = "0"
		data["message"] = "请至少选择一个!"
		return c.RenderJson(data)
	}

	if user.Lock(ids) {
		data["status"] = "1"
		data["message"] = "锁定成功!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["message"] = "锁定失败!"
		return c.RenderJson(data)
	}
}

//解锁
func (c *User) Unlock(user *models.User) revel.Result {

	var ids string = c.Params.Get("ids")

	data := make(map[string]string)

	if len(ids) <= 0 {
		data["status"] = "0"
		data["message"] = "请至少选择一个!"
		return c.RenderJson(data)
	}

	if user.Unlock(ids) {
		data["status"] = "1"
		data["message"] = "解锁成功!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["message"] = "解锁失败!"
		return c.RenderJson(data)
	}
}

//移动
func (c *User) Move(user *models.User) revel.Result {

	var groupid string = c.Params.Get("groupid")
	var ids string = c.Params.Get("ids")

	data := make(map[string]string)

	if len(ids) <= 0 {
		data["status"] = "0"
		data["message"] = "请至少选择一个!"
		return c.RenderJson(data)
	}

	if len(groupid) <= 0 {
		data["status"] = "0"
		data["message"] = "请选择会员组!"
		return c.RenderJson(data)
	}

	Groupid, err := strconv.ParseInt(groupid, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	if user.Move(Groupid, ids) {
		data["status"] = "1"
		data["message"] = "移动成功!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["message"] = "移动失败!"
		return c.RenderJson(data)
	}
}

//个人信息
func (c *User) UserInfo(user *models.User) revel.Result {
	title := "个人信息--GoCMS管理系统"

	var id string = c.Params.Get("id")

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		//内容
		user_info := user.GetById(Id)

		c.Render(title, id, user_info)
	} else {
		c.Render(title, id)
	}
	return c.RenderTemplate("User/Manage/UserInfo.html")
}

//登陆
func (c *User) Login(admin *models.Admin) revel.Result {

	if c.Request.Method == "GET" {
		title := "登陆--GoCMS管理系统"

		CaptchaId := captcha.NewLen(6)

		return c.Render(title, CaptchaId)
	} else {
		var username string = c.Params.Get("username")
		var password string = c.Params.Get("password")

		var captchaId string = c.Params.Get("captchaId")
		var verify string = c.Params.Get("verify")

		data := make(map[string]string)

		if LANG, ok := c.Session["Lang"]; ok {
			//设置语言
			c.Request.Locale = LANG
		} else {
			//设置默认语言
			c.Request.Locale = "zh"
		}

		if !captcha.VerifyString(captchaId, verify) {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("verification_code")
			return c.RenderJson(data)
		}

		if len(username) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_user_name")
			return c.RenderJson(data)
		}

		if len(password) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password")
			return c.RenderJson(data)
		}

		if len(verify) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_verification_code")
			return c.RenderJson(data)
		}

		admin_info := admin.GetByName(username)

		if admin_info.Id <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_username_error")
		} else if admin_info.Status == 0 && admin_info.Id != 1 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_forbid_login")
		} else if admin_info.Role.Status == 0 && admin_info.Id != 1 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_forbid_role_login")
		} else if username == admin_info.Username && utils.Md5(password) == admin_info.Password {
			c.Session["UserID"] = fmt.Sprintf("%d", admin_info.Id)
			c.Session["Lang"] = admin_info.Lang

			c.Flash.Success(c.Message("login_success"))
			c.Flash.Out["url"] = "/"

			//更新登陆时间
			admin.UpdateLoginTime(admin_info.Id)

			//******************************************
			//管理员日志
			logs := new(models.Logs)
			desc := "登陆用户名:" + admin_info.Username + "|^|登陆系统!|^|登陆ID:" + fmt.Sprintf("%d", admin_info.Id)
			logs.Save(admin_info, c.Controller, desc)
			//*****************************************

			data["status"] = "1"
			data["url"] = "/Message/"
			data["message"] = c.Message("login_success")
		} else {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password_error")
		}

		return c.RenderJson(data)
	}
}

//退出登陆
func (c *User) Logout(admin *models.Admin) revel.Result {

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_info := admin.GetById(UserID)

		//******************************************
		//管理员日志
		logs := new(models.Logs)
		desc := "登陆用户名:" + admin_info.Username + "|^|退出系统!|^|登陆ID:" + fmt.Sprintf("%d", admin_info.Id)
		logs.Save(admin_info, c.Controller, desc)
		//*****************************************

		for k := range c.Session {
			if k != "Lang" {
				delete(c.Session, k)
			}
		}
	}

	c.Flash.Success("安全退出")
	c.Flash.Out["url"] = "/User/Login/"
	return c.Redirect("/Message/")
}

//个人信息
func (c *User) EditInfo(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "个人信息--GoCMS管理系统"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)
			c.Render(title, admin_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/EditInfo.html")
	} else {

		var realname string = c.Params.Get("realname")
		if len(realname) > 0 {
			admin.Realname = realname
		} else {
			c.Flash.Error("请输入真实姓名!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		var email string = c.Params.Get("email")
		if len(email) > 0 {
			admin.Email = email
		} else {
			c.Flash.Error("请输入电子邮件!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		var lang string = c.Params.Get("lang")
		if len(lang) > 0 {
			admin.Lang = lang
		} else {
			c.Flash.Error("请选择语言!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			if admin.EditInfo(UserID) {

				//******************************************
				//管理员日志
				if UserID, ok := c.Session["UserID"]; ok {
					UserID, err := strconv.ParseInt(UserID, 10, 64)
					if err != nil {
						revel.WARN.Println(err)
					}

					admin := new(models.Admin)
					admin_info := admin.GetById(UserID)

					c.Session["Lang"] = admin_info.Lang

					logs := new(models.Logs)
					desc := "个人设置|^|个人信息"
					logs.Save(admin_info, c.Controller, desc)
				}

				if LANG, ok := c.Session["Lang"]; ok {
					//设置语言
					c.Request.Locale = LANG
				} else {
					//设置默认语言
					c.Request.Locale = "zh"
				}

				c.Flash.Success(c.Message("operation_success"))
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error(c.Message("operation_failure"))
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error(c.Message("not_login"))
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}

//常用菜单
func (c *User) AdminPanel(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "常用菜单--GoCMS管理系统"

		c.Render(title)

		return c.RenderTemplate("User/AdminPanel.html")
	} else {
		c.Flash.Success("添加成功!")
		c.Flash.Out["url"] = "/AdminPanel/"
		return c.Redirect("/Message/")
	}
}

//修改密码
func (c *User) EditPwd(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "修改密码--GoCMS管理系统"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)

			c.Render(title, admin_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/EditPwd.html")
	} else {

		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)

			var old_password string = c.Params.Get("old_password")
			if len(old_password) > 0 {
				if admin_info.Password != utils.Md5(old_password) {
					c.Flash.Error("旧密码不正确!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				}
			} else {
				return c.Redirect("/User/EditPwd/")
			}

			var new_password string = c.Params.Get("new_password")
			if len(new_password) > 0 {

			} else {
				c.Flash.Error("新密码不能为空!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			var new_pwdconfirm string = c.Params.Get("new_pwdconfirm")
			if len(new_pwdconfirm) > 0 {
				if new_pwdconfirm != new_password {
					c.Flash.Error("两次输入密码入不一致!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				} else {
					admin.Password = new_pwdconfirm
				}
			} else {
				c.Flash.Error("请输入重复新密码!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			if admin.EditPwd(UserID) {

				//******************************************
				//管理员日志
				logs := new(models.Logs)
				desc := "个人设置|^|修改密码"
				logs.Save(admin_info, c.Controller, desc)
				//*****************************************

				c.Flash.Success("修改成功!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("修改失败!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("未登陆，请先登陆!")
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}

//左侧导航菜单
func (c *User) Left(menu *models.Menu) revel.Result {

	title := "左侧导航--GoCMS管理系统"

	var pid string = c.Params.Get("pid")

	if len(pid) > 0 {
		Pid, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)

			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(Pid, admin_info)

			c.Render(title, left_menu)
		} else {
			c.Render(title)
		}
	} else {
		//获取左侧导航菜单
		//默认获取 首页
		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)

			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(1, admin_info)

			c.Render(title, left_menu)
		} else {
			c.Render(title)
		}
	}
	return c.RenderTemplate("Public/left.html")
}
