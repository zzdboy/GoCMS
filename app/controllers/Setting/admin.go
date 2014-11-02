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

// 管理员设置
import "strings"
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"

type Admin struct {
	*revel.Controller
}

//首页
func (c Admin) Index(admin *models.Admin) revel.Result {
	title := "管理员管理--GoCMS管理系统"

	var page string = c.Params.Get("page")

	where := make(map[string]string)

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_list, pages := admin.GetByAll(0, where, Page, 10)

		c.Render(title, admin_list, pages)
	} else {
		admin_list, pages := admin.GetByAll(0, where, 1, 10)

		c.Render(title, admin_list, pages)
	}

	return c.RenderTemplate("Setting/Admin/Index.html")
}

//添加管理员
func (c Admin) Add(admin *models.Admin) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加管理员--GoCMS管理系统"

		role := new(models.Role)
		role_list := role.GetRoleList()

		c.Render(title, role_list)
		return c.RenderTemplate("Setting/Admin/Add.html")
	} else {

		var username string = c.Params.Get("username")
		if len(username) > 0 {
			admin.Username = username
		} else {
			c.Flash.Error("请输入用户名!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		if admin.HasName() {
			c.Flash.Error("用户名“" + username + "”已存在！")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var password string = c.Params.Get("password")
		if len(password) > 0 {
			admin.Password = password
		} else {
			c.Flash.Error("请输入密码!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var pwdconfirm string = c.Params.Get("pwdconfirm")
		if len(pwdconfirm) > 0 {
			if password != pwdconfirm {
				c.Flash.Error("两次输入密码不一致!")
				c.Flash.Out["url"] = "/Admin/Add/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("请输入确认密码!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var email string = c.Params.Get("email")
		if len(email) > 0 {
			admin.Email = email
		} else {
			c.Flash.Error("请输入E-mail!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		if admin.HasEmail() {
			c.Flash.Error("E-mail已存在！")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var realname string = c.Params.Get("realname")
		if len(realname) > 0 {
			admin.Realname = realname
		} else {
			c.Flash.Error("请输入真实姓名!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var lang string = c.Params.Get("lang")
		if len(lang) > 0 {
			admin.Lang = lang
		} else {
			c.Flash.Error("请选择语言!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var roleid string = c.Params.Get("roleid")
		if len(roleid) > 0 {

			Roleid, err := strconv.ParseInt(roleid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin.Roleid = Roleid
		} else {
			c.Flash.Error("请选择所属角色!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		var status string = c.Params.Get("status")
		if len(status) > 0 {
			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			admin.Status = Status
		} else {
			c.Flash.Error("请选择状态!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}

		if ip := c.Request.Header.Get("X-Forwarded-For"); ip != "" {
			ips := strings.Split(ip, ",")
			if len(ips) > 0 && ips[0] != "" {
				rip := strings.Split(ips[0], ":")
				admin.Lastloginip = rip[0]
			}
		} else {
			ip := strings.Split(c.Request.RemoteAddr, ":")
			if len(ip) > 0 {
				if ip[0] != "[" {
					admin.Lastloginip = ip[0]
				}
			}
		}

		if admin.Save() {

			//******************************************
			//管理员日志
			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				desc := "添加管理员:" + username + "|^|管理员管理"
				logs.Save(admin_info, c.Controller, desc)
			}

			//*****************************************

			c.Flash.Success("添加管理员成功!")
			c.Flash.Out["url"] = "/Admin/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加管理员失败!")
			c.Flash.Out["url"] = "/Admin/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑管理员
func (c Admin) Edit(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "编辑管理员--GoCMS管理系统"

		role := new(models.Role)
		role_list := role.GetRoleList()

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(Id)

			c.Render(title, admin_info, role_list)
		} else {
			c.Render(title, role_list)
		}

		return c.RenderTemplate("Setting/Admin/Edit.html")
	} else {

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var username string = c.Params.Get("username")
			if len(username) > 0 {
				admin.Username = username
			} else {
				c.Flash.Error("请输入用户名!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var password string = c.Params.Get("password")
			if len(password) > 0 {
				admin.Password = password
			}

			var pwdconfirm string = c.Params.Get("pwdconfirm")
			if len(pwdconfirm) > 0 {
				if password != pwdconfirm {
					c.Flash.Error("两次输入密码不一致!")
					c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}
			}

			var email string = c.Params.Get("email")
			if len(email) > 0 {
				admin.Email = email
			} else {
				c.Flash.Error("请输入E-mail!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var realname string = c.Params.Get("realname")
			if len(realname) > 0 {
				admin.Realname = realname
			} else {
				c.Flash.Error("请输入真实姓名!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var lang string = c.Params.Get("lang")
			if len(lang) > 0 {
				admin.Lang = lang
			} else {
				c.Flash.Error("请选择语言!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var roleid string = c.Params.Get("roleid")
			if len(roleid) > 0 {

				Roleid, err := strconv.ParseInt(roleid, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin.Roleid = Roleid
			} else {
				c.Flash.Error("请选择所属角色!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var status string = c.Params.Get("status")
			if len(status) > 0 {
				Status, err := strconv.ParseInt(status, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				admin.Status = Status
			} else {
				c.Flash.Error("请选择是否启用!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if admin.Edit(Id) {

				//******************************************
				//管理员日志
				if UserID, ok := c.Session["UserID"]; ok {
					UserID, err := strconv.ParseInt(UserID, 10, 64)
					if err != nil {
						revel.WARN.Println(err)
					}

					admin_info := admin.GetById(UserID)

					logs := new(models.Logs)
					desc := "编辑管理员:" + username + "|^|管理员管理"
					logs.Save(admin_info, c.Controller, desc)
				}
				//*****************************************

				c.Flash.Success("编辑管理员成功!")
				c.Flash.Out["url"] = "/Admin/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑管理员失败!")
				c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑管理员失败!")
			c.Flash.Out["url"] = "/Admin/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}

	}
}

//删除管理员
func (c Admin) Delete(admin *models.Admin) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if admin.DelByID(Id) {

			//******************************************
			//管理员日志
			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				desc := "删除管理员|^|ID:" + id
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			data["status"] = "1"
			data["message"] = "删除成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "删除失败!"
			return c.RenderJson(data)
		}
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
		return c.RenderJson(data)
	}
}
