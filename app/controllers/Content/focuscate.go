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

//焦点图管理--焦点图分类
import "strconv"
import "admin/app/models"
import "github.com/revel/revel"

type FocusCate struct {
	*revel.Controller
}

func (c FocusCate) Index(focusCate *models.FocusCate) revel.Result {

	title := "焦点图分类--GoCMS管理系统"

	var page string = c.Params.Get("page")

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		focusCate_list, pages := focusCate.GetByAll(Page, 10)

		c.Render(title, focusCate_list, pages)
	} else {
		focusCate_list, pages := focusCate.GetByAll(1, 10)

		c.Render(title, focusCate_list, pages)
	}

	return c.RenderTemplate("Content/FocusCate/Index.html")
}

//添加分类
func (c FocusCate) Add(focusCate *models.FocusCate) revel.Result {

	if c.Request.Method == "GET" {

		title := "添加分类--GoCMS管理系统"

		c.Render(title)
		return c.RenderTemplate("Content/FocusCate/Add.html")

	} else {

		var name string = c.Params.Get("name")
		if len(name) > 0 {
			focusCate.Name = name
		} else {
			c.Flash.Error("请输入分类名称!")
			c.Flash.Out["url"] = "/focusCate/add/"
			return c.Redirect("/Message/")
		}

		var width string = c.Params.Get("width")
		if len(width) > 0 {
			Width, err := strconv.ParseInt(width, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			focusCate.Width = Width
		} else {
			c.Flash.Error("请输宽度!")
			c.Flash.Out["url"] = "/focusCate/add/"
			return c.Redirect("/Message/")
		}

		var height string = c.Params.Get("height")
		if len(height) > 0 {
			Height, err := strconv.ParseInt(height, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			focusCate.Height = Height
		} else {
			c.Flash.Error("请输入高度!")
			c.Flash.Out["url"] = "/focusCate/add/"
			return c.Redirect("/Message/")
		}

		if focusCate.Save() {
			c.Flash.Success("添加分类成功")
			c.Flash.Out["url"] = "/FocusCate/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加分类失败!")
			c.Flash.Out["url"] = "/FocusCate/add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑分类
func (c FocusCate) Edit(focusCate *models.FocusCate) revel.Result {

	if c.Request.Method == "GET" {

		title := "编辑分类--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			focusCate_info := focusCate.GetById(Id)

			c.Render(title, focusCate_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Content/FocusCate/Edit.html")

	} else {

		var id string = c.Params.Get("id")

		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		var name string = c.Params.Get("name")
		if len(name) > 0 {
			focusCate.Name = name
		} else {
			c.Flash.Error("请输入分类名称!")
			c.Flash.Out["url"] = "/focusCate/edit/" + id + "/"
			return c.Redirect("/Message/")
		}

		var width string = c.Params.Get("width")
		if len(width) > 0 {
			Width, err := strconv.ParseInt(width, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			focusCate.Width = Width
		} else {
			c.Flash.Error("请输宽度!")
			c.Flash.Out["url"] = "/focusCate/edit/" + id + "/"
			return c.Redirect("/Message/")
		}

		var height string = c.Params.Get("height")
		if len(height) > 0 {
			Height, err := strconv.ParseInt(height, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			focusCate.Height = Height
		} else {
			c.Flash.Error("请输入高度!")
			c.Flash.Out["url"] = "/focusCate/edit/" + id + "/"
			return c.Redirect("/Message/")
		}

		if focusCate.Edit(Id) {
			c.Flash.Success("编辑分类成功")
			c.Flash.Out["url"] = "/FocusCate/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("编辑分类失败!")
			c.Flash.Out["url"] = "/focusCate/edit/" + id + "/"
			return c.Redirect("/Message/")
		}

	}
}

//删除分类
func (c FocusCate) Delete(focusCate *models.FocusCate) revel.Result {

	var id string = c.Params.Get("id")

	data := make(map[string]string)

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	if focusCate.DelByID(Id) {
		data["status"] = "1"
		data["message"] = "删除成功!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
		return c.RenderJson(data)
	}

}
