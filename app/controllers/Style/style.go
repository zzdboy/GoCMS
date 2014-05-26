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

//界面首页
import "strconv"
import "admin/app/models"
import "github.com/revel/revel"

type Style struct {
	*revel.Controller
}

func (c Style) Index(template *models.Template) revel.Result {
	title := "模板风格--GoCMS管理系统"

	template_list := template.GetTemplateList()

	c.Render(title, template_list)
	return c.RenderTemplate("Style/Template/Index.html")
}

func (c *Style) File(template *models.Template) revel.Result {
	title := "模板风格--GoCMS管理系统"

	var id string = c.Params.Get("id")

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		template_info := template.GetById(Id)

		c.Render(title, template_info)
	} else {
		c.Render(title)
	}

	return c.RenderTemplate("Style/Template/File.html")
}

//导入风格
func (c *Style) Import(template *models.Template) revel.Result {

	if c.Request.Method == "GET" {

		title := "导入风格--GoCMS管理系统"

		c.Render(title)
		return c.RenderTemplate("Style/Template/Import.html")
	} else {

		data := make(map[string]string)

		var identity string = c.Params.Get("identity")
		if len(identity) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格标识!"
			return c.RenderJson(data)
		} else {
			template.Identity = identity
		}

		var name string = c.Params.Get("name")
		if len(name) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格中文名!"
			return c.RenderJson(data)
		} else {
			template.Name = name
		}

		var author string = c.Params.Get("author")
		if len(author) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格作者!"
			return c.RenderJson(data)
		} else {
			template.Author = author
		}

		var version string = c.Params.Get("version")
		if len(version) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格版本!"
			return c.RenderJson(data)
		} else {
			template.Version = version
		}

		var status string = c.Params.Get("status")

		if len(status) <= 0 {
			data["status"] = "0"
			data["message"] = "请选择状态!"
			return c.RenderJson(data)
		} else {

			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			template.Status = Status
		}

		if template.Save() {
			data["status"] = "1"
			data["message"] = "添加风格成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "添加风格失败!"
			return c.RenderJson(data)
		}
	}
}

//设置状态
func (c *Style) Setstatus(template *models.Template) revel.Result {

	data := make(map[string]string)

	var id string = c.Params.Get("id")

	if len(id) <= 0 {
		data["status"] = "0"
		data["message"] = "参数错误!"
		return c.RenderJson(data)
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	var status string = c.Params.Get("status")

	if len(status) <= 0 {
		data["status"] = "0"
		data["message"] = "请选择状态!"
		return c.RenderJson(data)
	} else {

		Status, err := strconv.ParseInt(status, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		template.Status = Status
	}

	if template.Setstatus(Id) {
		data["status"] = "1"
		data["message"] = "设置成功!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["message"] = "设置失败!"
		return c.RenderJson(data)
	}
}

//编辑风格
func (c *Style) Edit(template *models.Template) revel.Result {

	if c.Request.Method == "GET" {

		title := "编辑风格--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			template_info := template.GetById(Id)

			c.Render(title, template_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Style/Template/Edit.html")
	} else {

		data := make(map[string]string)

		var id string = c.Params.Get("id")

		if len(id) <= 0 {
			data["status"] = "0"
			data["message"] = "参数错误!"
			return c.RenderJson(data)
		}

		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		var identity string = c.Params.Get("identity")
		if len(identity) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格标识!"
			return c.RenderJson(data)
		} else {
			template.Identity = identity
		}

		var name string = c.Params.Get("name")
		if len(name) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格中文名!"
			return c.RenderJson(data)
		} else {
			template.Name = name
		}

		var author string = c.Params.Get("author")
		if len(author) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格作者!"
			return c.RenderJson(data)
		} else {
			template.Author = author
		}

		var version string = c.Params.Get("version")
		if len(version) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入风格版本!"
			return c.RenderJson(data)
		} else {
			template.Version = version
		}

		var status string = c.Params.Get("status")

		if len(status) <= 0 {
			data["status"] = "0"
			data["message"] = "请选择状态!"
			return c.RenderJson(data)
		} else {

			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			template.Status = Status
		}

		if template.Edit(Id) {
			data["status"] = "1"
			data["message"] = "编辑风格成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "编辑风格失败!"
			return c.RenderJson(data)
		}
	}
}
