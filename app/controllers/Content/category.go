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

//内容管理--栏目管理
import "strconv"
import "encoding/json"
import "github.com/revel/revel"
import "admin/app/models"
import "admin/utils"

type Category struct {
	*revel.Controller
}

func (c Category) Index(category *models.Category) revel.Result {
	title := "栏目管理--GoCMS管理系统"

	UserID := utils.GetSession("UserID", c.Session)

	if len(UserID) > 0 {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		categorys := category.GetCateGoryHtml(admin_info)

		c.Render(title, categorys)
	} else {
		c.Render(title)
	}

	return c.RenderTemplate("Content/Category/Index.html")
}

//添加栏目
func (c Category) Add(category *models.Category) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加栏目--GoCMS管理系统"

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//返回菜单Option的HTML
			categorys := category.GetCateGoryOptionHtml(Id)

			c.Render(title, categorys, Id)
		} else {
			//返回菜单Option的HTML
			categorys := category.GetCateGoryOptionHtml(0)
			c.Render(title, categorys)
		}

		return c.RenderTemplate("Content/Category/Add.html")
	} else {

		//栏目类型
		var cate_type string = c.Params.Get("type")
		if len(cate_type) > 0 {
			Type, err := strconv.ParseInt(cate_type, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			category.Type = Type
		} else {
			c.Flash.Error("请选择栏目类型!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		//上级栏目
		var pid string = c.Params.Get("pid")
		if len(pid) > 0 {
			Pid, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			category.Pid = Pid
		} else {
			c.Flash.Error("请选择父栏目!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		//栏目名称
		var name string = c.Params.Get("name")
		if len(name) > 0 {
			category.Name = name
		} else {
			c.Flash.Error("请输入栏目名称!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		var enname string = c.Params.Get("enname")
		if len(enname) > 0 {
			category.Enname = enname
		} else {
			c.Flash.Error("请输入英文栏目名称!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		//栏目地址
		var url string = c.Params.Get("url")
		if len(url) > 0 {
			category.Url = url
		} else {
			c.Flash.Error("请输入栏目地址!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		//描述
		var desc string = c.Params.Get("desc")
		if len(desc) > 0 {
			category.Desc = desc
		} else {
			c.Flash.Error("请输入描述!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		//是否在导航显示
		var ismenu string = c.Params.Get("ismenu")
		if len(ismenu) > 0 {
			IsMenu, err := strconv.ParseInt(ismenu, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			category.Ismenu = IsMenu
		} else {
			c.Flash.Error("请选择是否在导航显示!")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}

		Setting := make(map[string]interface{})

		//栏目生成Html
		var ishtml string = c.Params.Get("ishtml")
		Setting["ishtml"] = ishtml

		var content_ishtml string = c.Params.Get("content_ishtml")
		Setting["content_ishtml"] = content_ishtml

		var meta_title string = c.Params.Get("meta_title")
		Setting["meta_title"] = meta_title

		var meta_keywords string = c.Params.Get("meta_keywords")
		Setting["meta_keywords"] = meta_keywords

		var meta_desc string = c.Params.Get("meta_desc")
		Setting["meta_desc"] = meta_desc

		//栏目设置
		Setting_Text, err := json.Marshal(Setting)
		if err != nil {
			revel.WARN.Println(err)
		}
		category.SettingText = Setting
		category.Setting = string(Setting_Text)

		if category.Save() {
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
				desc := "添加栏目:" + name + "|^|栏目管理"
				logs.Save(admin_info, c.Controller, desc)
			}

			//*****************************************

			c.Flash.Success("添加栏目成功")
			c.Flash.Out["url"] = "/Category/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加栏目失败")
			c.Flash.Out["url"] = "/Category/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑栏目
func (c Category) Edit(category *models.Category) revel.Result {

	if c.Request.Method == "GET" {
		title := "编辑栏目--GoCMS管理系统"

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//获取菜单信息
			category_info := category.GetById(Id)

			//返回菜单Option的HTML
			categorys := category.GetCateGoryOptionHtml(category_info.Pid)

			c.Render(title, categorys, category_info)
		} else {

			//返回菜单Option的HTML
			categorys := category.GetCateGoryOptionHtml(0)

			c.Render(title, categorys)
		}

		return c.RenderTemplate("Content/Category/Edit.html")
	} else {
		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//栏目类型
			var cate_type string = c.Params.Get("type")
			if len(cate_type) > 0 {
				Type, err := strconv.ParseInt(cate_type, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				category.Type = Type
			} else {
				c.Flash.Error("请选择栏目类型!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			//上级栏目
			var pid string = c.Params.Get("pid")
			if len(pid) > 0 {
				Pid, err := strconv.ParseInt(pid, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				category.Pid = Pid
			} else {
				c.Flash.Error("请选择父栏目!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			//栏目名称
			var name string = c.Params.Get("name")
			if len(name) > 0 {
				category.Name = name
			} else {
				c.Flash.Error("请输入栏目名称!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var enname string = c.Params.Get("enname")
			if len(enname) > 0 {
				category.Enname = enname
			} else {
				c.Flash.Error("请输入英文栏目名称!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			//栏目地址
			var url string = c.Params.Get("url")
			if len(url) > 0 {
				category.Url = url
			} else {
				c.Flash.Error("请输入栏目地址!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			//描述
			var desc string = c.Params.Get("desc")
			if len(desc) > 0 {
				category.Desc = desc
			} else {
				c.Flash.Error("请输入描述!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			//是否在导航显示
			var ismenu string = c.Params.Get("ismenu")
			if len(ismenu) > 0 {
				IsMenu, err := strconv.ParseInt(ismenu, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
				category.Ismenu = IsMenu
			} else {
				c.Flash.Error("请选择是否在导航显示!")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			Setting := make(map[string]interface{})

			//栏目生成Html
			var ishtml string = c.Params.Get("ishtml")
			Setting["ishtml"] = ishtml

			var content_ishtml string = c.Params.Get("content_ishtml")
			Setting["content_ishtml"] = content_ishtml

			var meta_title string = c.Params.Get("meta_title")
			Setting["meta_title"] = meta_title

			var meta_keywords string = c.Params.Get("meta_keywords")
			Setting["meta_keywords"] = meta_keywords

			var meta_desc string = c.Params.Get("meta_desc")
			Setting["meta_desc"] = meta_desc

			//栏目设置
			Setting_Text, err := json.Marshal(Setting)
			if err != nil {
				revel.WARN.Println(err)
			}
			category.SettingText = Setting
			category.Setting = string(Setting_Text)

			if category.Edit(Id) {
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
					desc := "编辑栏目:" + name + "|^|栏目管理|^|ID:" + id
					logs.Save(admin_info, c.Controller, desc)
				}

				//*****************************************

				c.Flash.Success("编辑栏目成功")
				c.Flash.Out["url"] = "/Category/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑栏目失败")
				c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑菜单失败")
			c.Flash.Out["url"] = "/Category/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}

//删除栏目
func (c Category) Delete(category *models.Category) revel.Result {
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

	if category.DelByID(Id) {

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
			desc := "删除栏目|^|ID:" + id
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
