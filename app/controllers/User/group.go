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

//会员组管理
import "strconv"
import "admin/app/models"
import "github.com/revel/revel"

type Group struct {
	*revel.Controller
}

//首页
func (c *Group) Index(user_group *models.User_Group) revel.Result {
	title := "会员组首页--GoCMS管理系统"

	group_list := user_group.GetGroupList()

	c.Render(title, group_list)
	return c.RenderTemplate("User/Group/Index.html")
}

//添加会员组
func (c *Group) Add(user_group *models.User_Group) revel.Result {

	if c.Request.Method == "GET" {

		title := "添加会员组--GoCMS管理系统"

		c.Render(title)
		return c.RenderTemplate("User/Group/Add.html")
	} else {

		data := make(map[string]string)

		var name string = c.Params.Get("name")
		if len(name) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入会员组名称!"
			return c.RenderJson(data)
		} else {
			user_group.Name = name
		}

		var point string = c.Params.Get("point")
		if len(point) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入积分!"
			return c.RenderJson(data)
		} else {
			Point, err := strconv.ParseInt(point, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Point = Point
		}

		var star string = c.Params.Get("star")
		if len(star) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入星星数!"
			return c.RenderJson(data)
		} else {
			Star, err := strconv.ParseInt(star, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Star = Star
		}

		var allowpost string = c.Params.Get("allowpost")
		Allowpost, err := strconv.ParseInt(allowpost, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowpost = Allowpost

		var allowpostverify string = c.Params.Get("allowpostverify")
		Allowpostverify, err := strconv.ParseInt(allowpostverify, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowpostverify = Allowpostverify

		var allowupgrade string = c.Params.Get("allowupgrade")
		Allowupgrade, err := strconv.ParseInt(allowupgrade, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowupgrade = Allowupgrade

		var allowsendmessage string = c.Params.Get("allowsendmessage")
		Allowsendmessage, err := strconv.ParseInt(allowsendmessage, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowsendmessage = Allowsendmessage

		var allowattachment string = c.Params.Get("allowattachment")
		Allowattachment, err := strconv.ParseInt(allowattachment, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowattachment = Allowattachment

		var allowsearch string = c.Params.Get("allowsearch")
		Allowsearch, err := strconv.ParseInt(allowsearch, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowsearch = Allowsearch

		var priceday string = c.Params.Get("priceday")
		Priceday, err := strconv.ParseFloat(priceday, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Priceday = Priceday

		var pricemonth string = c.Params.Get("pricemonth")
		Pricemonth, err := strconv.ParseFloat(pricemonth, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Pricemonth = Pricemonth

		var priceyear string = c.Params.Get("priceyear")
		Priceyear, err := strconv.ParseFloat(priceyear, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Priceyear = Priceyear

		var allowmessage string = c.Params.Get("allowmessage")
		if len(allowmessage) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入最大短消息数!"
			return c.RenderJson(data)
		} else {
			Allowmessage, err := strconv.ParseInt(allowmessage, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Allowmessage = Allowmessage
		}

		var allowpostnum string = c.Params.Get("allowpostnum")
		if len(allowpostnum) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入日最大投稿数!"
			return c.RenderJson(data)
		} else {
			Allowpostnum, err := strconv.ParseInt(allowpostnum, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Allowpostnum = Allowpostnum
		}

		var usernamecolor string = c.Params.Get("usernamecolor")
		user_group.Usernamecolor = usernamecolor

		var icon string = c.Params.Get("icon")
		user_group.Icon = icon

		var desc string = c.Params.Get("desc")
		user_group.Desc = desc

		var status string = c.Params.Get("status")
		Status, err := strconv.ParseInt(status, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Status = Status

		if user_group.Save() {
			data["status"] = "1"
			data["message"] = "添加成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "添加失败!"
			return c.RenderJson(data)
		}
	}
}

//编辑会员组
func (c *Group) Edit(user_group *models.User_Group) revel.Result {

	if c.Request.Method == "GET" {

		title := "编辑会员组--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			user_group_info := user_group.GetById(Id)

			c.Render(title, user_group_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/Group/Edit.html")
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

		var name string = c.Params.Get("name")
		if len(name) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入会员组名称!"
			return c.RenderJson(data)
		} else {
			user_group.Name = name
		}

		var point string = c.Params.Get("point")
		if len(point) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入积分!"
			return c.RenderJson(data)
		} else {
			Point, err := strconv.ParseInt(point, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Point = Point
		}

		var star string = c.Params.Get("star")
		if len(star) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入星星数!"
			return c.RenderJson(data)
		} else {
			Star, err := strconv.ParseInt(star, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Star = Star
		}

		var allowpost string = c.Params.Get("allowpost")
		Allowpost, err := strconv.ParseInt(allowpost, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowpost = Allowpost

		var allowpostverify string = c.Params.Get("allowpostverify")
		Allowpostverify, err := strconv.ParseInt(allowpostverify, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowpostverify = Allowpostverify

		var allowupgrade string = c.Params.Get("allowupgrade")
		Allowupgrade, err := strconv.ParseInt(allowupgrade, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowupgrade = Allowupgrade

		var allowsendmessage string = c.Params.Get("allowsendmessage")
		Allowsendmessage, err := strconv.ParseInt(allowsendmessage, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowsendmessage = Allowsendmessage

		var allowattachment string = c.Params.Get("allowattachment")
		Allowattachment, err := strconv.ParseInt(allowattachment, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowattachment = Allowattachment

		var allowsearch string = c.Params.Get("allowsearch")
		Allowsearch, err := strconv.ParseInt(allowsearch, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Allowsearch = Allowsearch

		var priceday string = c.Params.Get("priceday")
		Priceday, err := strconv.ParseFloat(priceday, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Priceday = Priceday

		var pricemonth string = c.Params.Get("pricemonth")
		Pricemonth, err := strconv.ParseFloat(pricemonth, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Pricemonth = Pricemonth

		var priceyear string = c.Params.Get("priceyear")
		Priceyear, err := strconv.ParseFloat(priceyear, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Priceyear = Priceyear

		var allowmessage string = c.Params.Get("allowmessage")
		if len(allowmessage) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入最大短消息数!"
			return c.RenderJson(data)
		} else {
			Allowmessage, err := strconv.ParseInt(allowmessage, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Allowmessage = Allowmessage
		}

		var allowpostnum string = c.Params.Get("allowpostnum")
		if len(allowpostnum) <= 0 {
			data["status"] = "0"
			data["message"] = "请输入日最大投稿数!"
			return c.RenderJson(data)
		} else {
			Allowpostnum, err := strconv.ParseInt(allowpostnum, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			user_group.Allowpostnum = Allowpostnum
		}

		var usernamecolor string = c.Params.Get("usernamecolor")
		user_group.Usernamecolor = usernamecolor

		var icon string = c.Params.Get("icon")
		user_group.Icon = icon

		var desc string = c.Params.Get("desc")
		user_group.Desc = desc

		var status string = c.Params.Get("status")
		Status, err := strconv.ParseInt(status, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		user_group.Status = Status

		if user_group.Edit(Id) {
			data["status"] = "1"
			data["message"] = "编辑成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "编辑失败!"
			return c.RenderJson(data)
		}
	}
}
