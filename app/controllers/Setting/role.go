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

// 角色设置
import "strconv"
import "admin/app/models"
import "github.com/revel/revel"

type Role struct {
	*revel.Controller
}

//首页
func (c Role) Index(role *models.Role) revel.Result {
	title := "角色管理--GoCMS管理系统"

	var page string = c.Params.Get("page")

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		role_list, pages := role.GetByAll(Page, 10)

		c.Render(title, role_list, pages)
	} else {
		role_list, pages := role.GetByAll(1, 10)

		c.Render(title, role_list, pages)
	}

	return c.RenderTemplate("Setting/Role/Index.html")
}

//成员管理
func (c Role) Member(role *models.Role) revel.Result {
	title := "成员管理--GoCMS管理系统"

	var id string = c.Params.Get("id")
	var page string = c.Params.Get("page")

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		where := make(map[string]string)

		admin := new(models.Admin)

		if len(page) > 0 {
			Page, err := strconv.ParseInt(page, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_list, pages := admin.GetByAll(Id, where, Page, 10)

			c.Render(title, admin_list, pages)
		} else {
			admin_list, pages := admin.GetByAll(Id, where, 1, 10)

			c.Render(title, admin_list, pages)
		}

	} else {
		c.Render(title)
	}

	return c.RenderTemplate("Setting/Role/Member.html")
}

//添加角色
func (c Role) Add(role *models.Role) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加角色--GoCMS管理系统"

		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)

			menu := new(models.Menu)
			tree := menu.GetMenuTree("", admin_info)

			c.Render(title, tree)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Setting/Role/Add.html")
	} else {

		var rolename string = c.Params.Get("rolename")
		if len(rolename) > 0 {
			role.Rolename = rolename
		} else {
			c.Flash.Error("请输入角色名称!")
			c.Flash.Out["url"] = "/Role/Add/"
			return c.Redirect("/Message/")
		}

		var desc string = c.Params.Get("desc")
		if len(desc) > 0 {
			role.Desc = desc
		} else {
			c.Flash.Error("请输入角色描述!")
			c.Flash.Out["url"] = "/Role/Add/"
			return c.Redirect("/Message/")
		}

		var data string = c.Params.Get("data")
		if len(data) > 0 {
			role.Data = data
		} else {
			c.Flash.Error("请选择所属权限!")
			c.Flash.Out["url"] = "/Role/Add/"
			return c.Redirect("/Message/")
		}

		var status string = c.Params.Get("status")
		if len(status) > 0 {
			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			role.Status = Status
		} else {
			c.Flash.Error("请选择是否启用!")
			c.Flash.Out["url"] = "/Role/Add/"
			return c.Redirect("/Message/")
		}

		if role.Save() {

			//******************************************
			//管理员日志
			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				desc := "添加角色:" + rolename + "|^|角色管理"
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			c.Flash.Success("添加角色成功")
			c.Flash.Out["url"] = "/Role/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加角色失败")
			c.Flash.Out["url"] = "/Role/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑角色
func (c Role) Edit(role *models.Role) revel.Result {
	if c.Request.Method == "GET" {
		title := "编辑角色--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			role_info := role.GetById(Id)

			if UserID, ok := c.Session["UserID"]; ok {

				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				menu := new(models.Menu)
				tree := menu.GetMenuTree(role_info.Data, admin_info)

				c.Render(title, role_info, tree, Id)
			} else {
				c.Render(title, role_info, Id)
			}

		} else {

			if UserID, ok := c.Session["UserID"]; ok {

				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				menu := new(models.Menu)
				tree := menu.GetMenuTree("", admin_info)

				c.Render(title, tree)
			} else {
				c.Render(title)
			}
		}

		return c.RenderTemplate("Setting/Role/Edit.html")
	} else {
		var id string = c.Params.Get("id")

		if len(id) > 0 {

			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var rolename string = c.Params.Get("rolename")
			if len(rolename) > 0 {
				role.Rolename = rolename
			} else {
				c.Flash.Error("请输入角色名称!")
				c.Flash.Out["url"] = "/Role/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var desc string = c.Params.Get("desc")
			if len(desc) > 0 {
				role.Desc = desc
			} else {
				c.Flash.Error("请输入角色描述!")
				c.Flash.Out["url"] = "/Role/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var data string = c.Params.Get("data")
			if len(data) > 0 {
				role.Data = data
			} else {
				c.Flash.Error("请选择所属权限!")
				c.Flash.Out["url"] = "/Role/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var status string = c.Params.Get("status")
			if len(status) > 0 {
				Status, err := strconv.ParseInt(status, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				role.Status = Status
			} else {
				c.Flash.Error("请选择是否启用!")
				c.Flash.Out["url"] = "/Role/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if role.Edit(Id) {

				//******************************************
				//管理员日志
				if UserID, ok := c.Session["UserID"]; ok {
					UserID, err := strconv.ParseInt(UserID, 10, 64)
					if err != nil {
						revel.WARN.Println(err)
					}

					admin := new(models.Admin)
					admin_info := admin.GetById(UserID)

					logs := new(models.Logs)
					desc := "编辑角色:" + rolename + "|^|角色管理|^|ID:" + id
					logs.Save(admin_info, c.Controller, desc)
				}
				//*****************************************

				c.Flash.Success("编辑角色成功")
				c.Flash.Out["url"] = "/Role/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑角色失败")
				c.Flash.Out["url"] = "/Role/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑角色失败")
			c.Flash.Out["url"] = "/Role/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}

//设置状态
func (c Role) SetStatus(role *models.Role) revel.Result {
	var id string = c.Params.Get("id")
	var status string = c.Params.Get("status")

	data := make(map[string]string)

	if len(id) > 0 && len(status) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		Status, err := strconv.ParseInt(status, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		role.Status = Status

		if role.SetStatus(Id) {

			//******************************************
			//管理员日志
			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				if Status == 1 {
					desc := "角色管理|^|设置状态|^|状态:启用"
					logs.Save(admin_info, c.Controller, desc)
				} else {
					desc := "角色管理|^|设置状态|^|状态:锁定"
					logs.Save(admin_info, c.Controller, desc)
				}
			}
			//*****************************************

			data["status"] = "1"
			data["message"] = "设置成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "设置失败!"
			return c.RenderJson(data)
		}
	} else {
		data["status"] = "0"
		data["message"] = "设置失败!"
		return c.RenderJson(data)
	}
}

//删除角色
func (c Role) Delete(role *models.Role) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if role.DelByID(Id) {
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
