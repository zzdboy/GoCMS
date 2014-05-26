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

//模块管理--公告管理
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"

type Announce struct {
	*revel.Controller
}

func (c Announce) Index(announce *models.Announce) revel.Result {
	title := "公告管理--GoCMS管理系统"

	var page string = c.Params.Get("page")

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		announce_list, pages := announce.GetByAll(Page, 10)

		c.Render(title, announce_list, pages)
	} else {
		announce_list, pages := announce.GetByAll(1, 10)

		c.Render(title, announce_list, pages)
	}

	return c.RenderTemplate("Module/Announce/Index.html")
}

//添加公告
func (c Announce) Add(announce *models.Announce) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加公告--GoCMS管理系统"

		c.Render(title)
		return c.RenderTemplate("Module/Announce/Add.html")
	} else {
		var title string = c.Params.Get("title")
		if len(title) > 0 {
			announce.Title = title
		} else {
			c.Flash.Error("请输入公告标题!")
			c.Flash.Out["url"] = "/Announce/Add/"
			return c.Redirect("/Message/")
		}

		var starttime string = c.Params.Get("starttime")
		if len(starttime) > 0 {
			announce.Starttime = starttime
		} else {
			c.Flash.Error("请输入起始日期!")
			c.Flash.Out["url"] = "/Announce/Add/"
			return c.Redirect("/Message/")
		}

		var endtime string = c.Params.Get("endtime")
		if len(endtime) > 0 {
			announce.Endtime = endtime
		} else {
			c.Flash.Error("请输入截止日期!")
			c.Flash.Out["url"] = "/Announce/Add/"
			return c.Redirect("/Message/")
		}

		var content string = c.Params.Get("content")
		if len(content) > 0 {
			announce.Content = content
		} else {
			c.Flash.Error("请输入公告内容!")
			c.Flash.Out["url"] = "/Announce/Add/"
			return c.Redirect("/Message/")
		}

		var status string = c.Params.Get("status")
		if len(status) > 0 {
			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			announce.Status = Status
		} else {
			c.Flash.Error("请选择是否启用!")
			c.Flash.Out["url"] = "/Announce/Add/"
			return c.Redirect("/Message/")
		}

		if announce.Save() {

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
				desc := "添加公告:" + title
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			c.Flash.Success("添加公告成功!")
			c.Flash.Out["url"] = "/Announce/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加公告失败!")
			c.Flash.Out["url"] = "/Announce/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑栏目
func (c Announce) Edit(announce *models.Announce) revel.Result {

	if c.Request.Method == "GET" {
		title := "编辑公告--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			announce_info := announce.GetById(Id)

			c.Render(title, announce_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Module/Announce/Edit.html")
	} else {

		var id string = c.Params.Get("id")

		if len(id) > 0 {

			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var title string = c.Params.Get("title")
			if len(title) > 0 {
				announce.Title = title
			} else {
				c.Flash.Error("请输入公告标题!")
				c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var starttime string = c.Params.Get("starttime")
			if len(starttime) > 0 {
				announce.Starttime = starttime
			} else {
				c.Flash.Error("请输入起始日期!")
				c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var endtime string = c.Params.Get("endtime")
			if len(endtime) > 0 {
				announce.Endtime = endtime
			} else {
				c.Flash.Error("请输入截止日期!")
				c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var content string = c.Params.Get("content")
			if len(content) > 0 {
				announce.Content = content
			} else {
				c.Flash.Error("请输入公告内容!")
				c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var status string = c.Params.Get("status")
			if len(status) > 0 {
				Status, err := strconv.ParseInt(status, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				announce.Status = Status
			} else {
				c.Flash.Error("请选择是否启用!")
				c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if announce.Edit(Id) {

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
					desc := "编辑公告|^|ID:" + id
					logs.Save(admin_info, c.Controller, desc)
				}
				//*****************************************

				c.Flash.Success("编辑公告成功!")
				c.Flash.Out["url"] = "/Announce/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑公告失败!")
				c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑公告失败!")
			c.Flash.Out["url"] = "/Announce/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}

//删除公告
func (c Announce) Delete(announce *models.Announce) revel.Result {

	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if announce.DelByID(Id) {

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
				desc := "删除公告|^|ID:" + id
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
