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

// 菜单设置
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"

type Menu struct {
	*revel.Controller
}

//首页
func (c Menu) Index(menu *models.Menu) revel.Result {
	title := "菜单管理--GoCMS管理系统"

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		menus := menu.GetMenuHtml(admin_info)

		c.Render(title, menus)
	} else {
		c.Render(title)
	}

	return c.RenderTemplate("Setting/Menu/Index.html")
}

//添加菜单
func (c Menu) Add(menu *models.Menu) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加菜单--GoCMS管理系统"

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
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

				//返回菜单Option的HTML
				menus := menu.GetMenuOptionHtml(Id, admin_info)

				c.Render(title, menus, Id)
			} else {
				c.Render(title, Id)
			}

		} else {

			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				//返回菜单Option的HTML
				menus := menu.GetMenuOptionHtml(0, admin_info)
				c.Render(title, menus)
			} else {
				c.Render(title)
			}
		}

		return c.RenderTemplate("Setting/Menu/Add.html")
	} else {

		var pid string = c.Params.Get("pid")
		if len(pid) > 0 {
			Pid, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			menu.Pid = Pid
		} else {
			c.Flash.Error("请选择父菜单!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var name string = c.Params.Get("name")
		if len(name) > 0 {
			menu.Name = name
		} else {
			c.Flash.Error("请输入中文语言名称!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var enname string = c.Params.Get("enname")
		if len(enname) > 0 {
			menu.Enname = enname
		} else {
			c.Flash.Error("请输入英文语言名称!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var url string = c.Params.Get("url")
		if len(url) > 0 {
			menu.Url = url
		} else {
			c.Flash.Error("请输入菜单地址!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var order string = c.Params.Get("order")
		if len(order) > 0 {
			Order, err := strconv.ParseInt(order, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			menu.Order = Order
		} else {
			c.Flash.Error("请输入排序!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		var data string = c.Params.Get("data")
		menu.Data = data

		var display string = c.Params.Get("display")
		if len(display) > 0 {
			Display, err := strconv.ParseInt(display, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			menu.Display = Display
		} else {
			c.Flash.Error("请选择是否显示菜单!")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}

		if menu.Save() {

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
				desc := "添加菜单:" + name + "|^|菜单管理"
				logs.Save(admin_info, c.Controller, desc)
			}

			//*****************************************

			c.Flash.Success("添加菜单成功")
			c.Flash.Out["url"] = "/Menu/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加菜单失败")
			c.Flash.Out["url"] = "/Menu/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑菜单
func (c Menu) Edit(menu *models.Menu) revel.Result {
	if c.Request.Method == "GET" {
		title := "编辑菜单--GoCMS管理系统"

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//获取菜单信息
			menu_info := menu.GetById(Id)

			if UserID, ok := c.Session["UserID"]; ok {

				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				//返回菜单Option的HTML
				menus := menu.GetMenuOptionHtml(menu_info.Pid, admin_info)

				c.Render(title, menus, menu_info)
			} else {
				c.Render(title, menu_info)
			}
		} else {

			if UserID, ok := c.Session["UserID"]; ok {

				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				//返回菜单Option的HTML
				menus := menu.GetMenuOptionHtml(0, admin_info)

				c.Render(title, menus)
			} else {
				c.Render(title)
			}
		}
		return c.RenderTemplate("Setting/Menu/Edit.html")
	} else {

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var pid string = c.Params.Get("pid")
			if len(pid) > 0 {
				Pid, err := strconv.ParseInt(pid, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				menu.Pid = Pid
			} else {
				c.Flash.Error("请选择父菜单!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var name string = c.Params.Get("name")
			if len(name) > 0 {
				menu.Name = name
			} else {
				c.Flash.Error("请输入中文语言名称!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var enname string = c.Params.Get("enname")
			if len(enname) > 0 {
				menu.Enname = enname
			} else {
				c.Flash.Error("请输入英文语言名称!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var url string = c.Params.Get("url")
			if len(url) > 0 {
				menu.Url = url
			} else {
				c.Flash.Error("请输入菜单地址!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var order string = c.Params.Get("order")
			if len(order) > 0 {
				Order, err := strconv.ParseInt(order, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				menu.Order = Order
			} else {
				c.Flash.Error("请输入排序!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var data string = c.Params.Get("data")
			menu.Data = data

			var display string = c.Params.Get("display")
			if len(display) > 0 {
				Display, err := strconv.ParseInt(display, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				menu.Display = Display
			} else {
				c.Flash.Error("请选择是否显示菜单!")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if menu.Edit(Id) {

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
					desc := "编辑菜单:" + name + "|^|菜单管理|^|ID:" + id
					logs.Save(admin_info, c.Controller, desc)
				}

				//*****************************************

				c.Flash.Success("编辑菜单成功")
				c.Flash.Out["url"] = "/Menu/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑菜单失败")
				c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑菜单失败")
			c.Flash.Out["url"] = "/Menu/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}

//删除
func (c Menu) Delete(menu *models.Menu) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) <= 0 {
		data["status"] = "0"
		data["message"] = "参数错误!"
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	if menu.DelByID(Id) {

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
			desc := "删除菜单|^|ID:" + id
			logs.Save(admin_info, c.Controller, desc)
		}
		//*****************************************

		data["status"] = "1"
		data["message"] = "删除成功!"
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
	}

	return c.RenderJson(data)
}
