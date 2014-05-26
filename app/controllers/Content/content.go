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

//内容首页
import "os"
import "strconv"
import "strings"
import "admin/utils"
import "github.com/revel/revel"
import "github.com/revel/config"
import "admin/app/models"

type Content struct {
	*revel.Controller
}

func (c Content) Index(article *models.Article) revel.Result {
	title := "内容--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Content/Index.html")
}

//左侧菜单
func (c Content) Left(article *models.Article) revel.Result {
	title := "左侧菜单--GoCMS管理系统"

	category := new(models.Category)
	categorys := category.GetLeftTree()

	c.Render(title, categorys)
	return c.RenderTemplate("Content/Left.html")
}

//内容管理列表
func (c Content) List(article *models.Article) revel.Result {

	title := "内容--GoCMS管理内容"

	var cid string = c.Params.Get("cid")
	var page string = c.Params.Get("page")
	var search string = c.Params.Get("search")

	Cid, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	category := new(models.Category)
	categorys := category.GetCateGoryOptionHtml(0)

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		article_list, pages, where := article.GetByAll(search, Cid, Page, 10)

		c.Render(title, cid, categorys, article_list, where, pages)
	} else {
		article_list, pages, where := article.GetByAll(search, Cid, 1, 10)

		c.Render(title, cid, categorys, article_list, where, pages)
	}

	return c.RenderTemplate("Content/Manage/Index.html")
}

//获取关键词
func (c Content) Keywords(article *models.Article) revel.Result {

	var title string = c.Params.Get("title")

	if len(title) > 0 {

	}

	data := ""
	return c.RenderText(data)
}

//删除
func (c Content) Delete(article *models.Article) revel.Result {

	var cid string = c.Params.Get("cid")

	var ids []int64
	c.Params.Bind(&ids, "ids")

	if len(ids) <= 0 {
		c.Flash.Error("请至少选择一个!")
		c.Flash.Out["url"] = "/Content/list/" + cid + "/"
		return c.Redirect("/Message/")
	}

	for _, Id := range ids {
		article.DelByID(Id)
	}

	c.Flash.Success("删除成功!")
	c.Flash.Out["url"] = "/Content/list/" + cid + "/"
	return c.Redirect("/Message/")
}

//快速进入 搜索
func (c Content) CateNameSearch(category *models.Category) revel.Result {
	var catename string = c.Params.Get("catename")

	if len(catename) < 0 {
		return c.RenderHtml("")
	}

	data := category.GetCateNameHtml(catename)

	return c.RenderHtml(data)
}

//推送
func (c Content) Push(article *models.Article) revel.Result {
	title := "内容--推送"

	c.Render(title)
	return c.RenderTemplate("Content/Manage/Push.html")
}

//批量移动
func (c Content) Remove(article *models.Article) revel.Result {

	if c.Request.Method == "GET" {
		title := "内容--批量移动"

		c.Render(title)
		return c.RenderTemplate("Content/Manage/Remove.html")
	} else {
		var ids string = c.Params.Get("ids")
		var cid string = c.Params.Get("cid")

		data := make(map[string]string)

		if len(ids) <= 0 {
			data["status"] = "1"
			data["url"] = "/Message/"
			data["message"] = "请选择至少选择一个!"
		} else if len(cid) <= 0 {
			data["status"] = "1"
			data["url"] = "/Message/"
			data["message"] = "请选择要移动的栏目!"
		} else {
			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			is_remove := article.Remove(Cid, ids)

			if is_remove {
				data["status"] = "1"
				data["url"] = "/Message/"
				data["message"] = "移动成功!"
			} else {
				data["status"] = "1"
				data["url"] = "/Message/"
				data["message"] = "移动失败!"
			}
		}

		return c.RenderJson(data)
	}
}

//评论管理
func (c Content) Comment(article *models.Article) revel.Result {
	title := "内容--查看评论"

	c.Render(title)
	return c.RenderTemplate("Content/Manage/Comment.html")
}

//添加相关文章
func (c Content) Relationlist(article *models.Article) revel.Result {
	title := "内容--添加相关文章"

	var page string = c.Params.Get("page")
	var cid string = c.Params.Get("cid")
	var search string = c.Params.Get("search")

	//栏目信息
	category := new(models.Category)
	category_info := category.GetCateGoryOptionHtml(0)

	var Page int64 = 1

	if len(page) > 0 {
		Page, _ = strconv.ParseInt(page, 10, 64)
	}

	Cid, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	article_list, pages, where := article.GetByList(Cid, search, Page, 10)

	c.Render(title, cid, category_info, article_list, pages, where)
	return c.RenderTemplate("Content/Manage/Relationlist.html")
}

//单页面 添加内容
func (c Content) AddContent(article *models.Article) revel.Result {
	if c.Request.Method == "GET" {
		title := "内容--GoCMS添加内容"

		var cid string = c.Params.Get("cid")

		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//栏目信息
			category := new(models.Category)
			category_info := category.GetById(Cid)

			//内容
			article_info := article.GetByCid(Cid)

			c.Render(title, cid, category_info, article_info)
			return c.RenderTemplate("Content/Manage/AddContent.html")
		} else {
			c.Render(title, cid)
			return c.RenderTemplate("Content/Manage/AddContent.html")
		}
	} else {

		var cid string = c.Params.Get("cid")

		if len(cid) < 0 {
			c.Flash.Error("请选择栏目!")
			c.Flash.Out["url"] = "/Content/addContent/" + cid + "/"
			return c.Redirect("/Message/")
		}

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			article.Aid = UserID
		}

		var title string = c.Params.Get("title")
		if len(title) > 0 {
			article.Title = title
		} else {
			c.Flash.Error("请输入标题!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		var style_color string = c.Params.Get("style_color")
		article.Color = style_color
		var style_font_weight string = c.Params.Get("style_font_weight")
		article.Font = style_font_weight

		var keywords string = c.Params.Get("keywords")
		if len(keywords) > 0 {
			article.Keywords = keywords
		} else {
			article.Keywords = ""
		}

		var content string = c.Params.Get("content")
		if len(content) > 0 {
			article.Content = content
		} else {
			c.Flash.Error("请输入内容!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Cid = Cid

			//内容
			article_info := article.GetByCid(Cid)

			if article_info.Id > 0 {
				if article.Edit(article_info.Id) {
					c.Flash.Success("编辑内容成功!")
					c.Flash.Out["url"] = "/Content/addContent/" + cid + "/"
					return c.Redirect("/Message/")
				} else {
					c.Flash.Success("编辑内容成功!")
					c.Flash.Out["url"] = "/Content/addContent/" + cid + "/"
					return c.Redirect("/Message/")
				}
			} else {
				if article.Save() {
					c.Flash.Success("添加内容成功!")
					c.Flash.Out["url"] = "/Content/addContent/" + cid + "/"
					return c.Redirect("/Message/")
				} else {
					c.Flash.Error("添加内容失败")
					c.Flash.Out["url"] = "/Content/addContent/" + cid + "/"
					return c.Redirect("/Message/")
				}
			}
		} else {
			c.Flash.Error("请选择栏目!")
			c.Flash.Out["url"] = "/Content/addContent/" + cid + "/"
			return c.Redirect("/Message/")
		}
	}
}

//添加内容
func (c Content) Add(article *models.Article) revel.Result {

	if c.Request.Method == "GET" {
		title := "内容--GoCMS添加内容"

		var cid string = c.Params.Get("cid")

		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//栏目信息
			category := new(models.Category)
			category_info := category.GetById(Cid)

			//来源
			copyfrom := new(models.Copyfrom)
			copyfrom_list := copyfrom.GetRoleList()

			c.Render(title, cid, category_info, copyfrom_list)
			return c.RenderTemplate("Content/Manage/Add.html")
		} else {
			c.Render(title, cid)
			return c.RenderTemplate("Content/Manage/Add.html")
		}

	} else {

		var dosubmit string = c.Params.Get("dosubmit")
		var dosubmit_continue string = c.Params.Get("dosubmit_continue")

		var cid string = c.Params.Get("cid")
		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Cid = Cid
		} else {
			c.Flash.Error("请选择栏目!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			article.Aid = UserID
		}

		var title string = c.Params.Get("title")
		if len(title) > 0 {
			article.Title = title
		} else {
			c.Flash.Error("请输入标题!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		var style_color string = c.Params.Get("style_color")
		article.Color = style_color
		var style_font_weight string = c.Params.Get("style_font_weight")
		article.Font = style_font_weight

		var thumb string = c.Params.Get("thumb")
		if len(thumb) > 0 {

			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//前台网站地址
			sitedomain, _ := config_conf.String("website", "website.sitedomain")
			thumb = strings.Replace(thumb, sitedomain, "", -1)

			article.Thumb = thumb
		} else {
			article.Thumb = ""
		}

		var content string = c.Params.Get("content")
		if len(content) > 0 {
			article.Content = content
		} else {
			c.Flash.Error("请输入内容!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		var copyfrom string = c.Params.Get("copyfrom")
		if len(copyfrom) > 0 {
			article.Copyfrom = copyfrom
		} else {
			article.Copyfrom = ""
		}

		var keywords string = c.Params.Get("keywords")
		if len(keywords) > 0 {
			article.Keywords = keywords
		} else {
			article.Keywords = ""
		}

		var description string = c.Params.Get("description")
		if len(description) > 0 {
			article.Description = description
		} else {

			//是否截取内容
			var add_introduce string = c.Params.Get("add_introduce")
			if add_introduce == "1" {
				var introcude_length string = c.Params.Get("introcude_length")

				Introcude_length, err := strconv.ParseInt(introcude_length, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				description = utils.Html2str(content)

				article.Description = utils.Substr(description, Introcude_length)
			} else {
				article.Description = ""
			}

		}

		var relation string = c.Params.Get("relation")
		if len(relation) > 0 {
			article.Relation = relation
		} else {
			article.Relation = ""
		}

		var pagetype string = c.Params.Get("pagetype")
		if len(pagetype) > 0 {

			Pagetype, err := strconv.ParseInt(pagetype, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Pagetype = Pagetype
		} else {
			article.Pagetype = 0
		}

		var maxcharperpage string = c.Params.Get("maxcharperpage")
		if len(maxcharperpage) > 0 {
			Maxcharperpage, err := strconv.ParseInt(maxcharperpage, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Maxcharperpage = Maxcharperpage
		} else {
			article.Maxcharperpage = 10000
		}

		var istop string = c.Params.Get("istop")
		if len(istop) > 0 {
			Istop, err := strconv.ParseInt(istop, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Istop = Istop
		} else {
			article.Istop = 0
		}

		var status string = c.Params.Get("status")
		if len(status) > 0 {
			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Status = Status
		} else {
			article.Status = 1
		}

		var iscomment string = c.Params.Get("iscomment")
		if len(iscomment) > 0 {
			Iscomment, err := strconv.ParseInt(iscomment, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Iscomment = Iscomment
		} else {
			article.Status = 1
		}

		if article.Save() {

			if len(dosubmit) > 0 {
				c.Flash.Success("添加内容成功!")
				c.Flash.Out["url"] = "/Content/list/" + cid + "/"
				return c.Redirect("/Message/")
			} else if len(dosubmit_continue) > 0 {
				c.Flash.Success("添加内容成功!")
				c.Flash.Out["url"] = "/Content/add/" + cid + "/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Success("添加内容成功!")
				c.Flash.Out["url"] = "/Content/list/" + cid + "/"
				return c.Redirect("/Message/")
			}

		} else {
			c.Flash.Error("添加内容失败")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑内容
func (c Content) Edit(article *models.Article) revel.Result {

	if c.Request.Method == "GET" {
		title := "内容--GoCMS编辑内容"

		var cid string = c.Params.Get("cid")
		var id string = c.Params.Get("id")

		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//内容
			article_info := article.GetById(Id)

			//栏目信息
			category := new(models.Category)
			category_info := category.GetById(Cid)

			//来源
			copyfrom := new(models.Copyfrom)
			copyfrom_list := copyfrom.GetRoleList()

			c.Render(title, cid, category_info, article_info, copyfrom_list)
			return c.RenderTemplate("Content/Manage/Edit.html")
		} else {
			c.Render(title, cid)
			return c.RenderTemplate("Content/Manage/Edit.html")
		}

	} else {

		var dosubmit string = c.Params.Get("dosubmit")
		var dosubmit_continue string = c.Params.Get("dosubmit_continue")

		var cid string = c.Params.Get("cid")
		var id string = c.Params.Get("id")

		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		var title string = c.Params.Get("title")
		if len(title) > 0 {
			article.Title = title
		} else {
			c.Flash.Error("请输入标题!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		var style_color string = c.Params.Get("style_color")
		article.Color = style_color
		var style_font_weight string = c.Params.Get("style_font_weight")
		article.Font = style_font_weight

		var thumb string = c.Params.Get("thumb")
		if len(thumb) > 0 {

			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//前台网站地址
			sitedomain, _ := config_conf.String("website", "website.sitedomain")
			thumb = strings.Replace(thumb, sitedomain, "", -1)

			article.Thumb = thumb
		} else {
			article.Thumb = ""
		}

		var content string = c.Params.Get("content")
		if len(content) > 0 {
			article.Content = content
		} else {
			c.Flash.Error("请输入内容!")
			c.Flash.Out["url"] = "/Content/Add/" + cid + "/"
			return c.Redirect("/Message/")
		}

		var copyfrom string = c.Params.Get("copyfrom")
		if len(copyfrom) > 0 {
			article.Copyfrom = copyfrom
		} else {
			article.Copyfrom = ""
		}

		var keywords string = c.Params.Get("keywords")
		if len(keywords) > 0 {
			article.Keywords = keywords
		} else {
			article.Keywords = ""
		}

		var description string = c.Params.Get("description")
		if len(description) > 0 {
			article.Description = description
		} else {

			//是否截取内容
			var add_introduce string = c.Params.Get("add_introduce")
			if add_introduce == "1" {
				var introcude_length string = c.Params.Get("introcude_length")

				Introcude_length, err := strconv.ParseInt(introcude_length, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				description = utils.Html2str(content)

				article.Description = utils.Substr(description, Introcude_length)
			} else {
				article.Description = ""
			}

		}

		var relation string = c.Params.Get("relation")
		if len(relation) > 0 {
			article.Relation = relation
		} else {
			article.Relation = ""
		}

		var pagetype string = c.Params.Get("pagetype")
		if len(pagetype) > 0 {

			Pagetype, err := strconv.ParseInt(pagetype, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Pagetype = Pagetype
		} else {
			article.Pagetype = 0
		}

		var maxcharperpage string = c.Params.Get("maxcharperpage")
		if len(maxcharperpage) > 0 {
			Maxcharperpage, err := strconv.ParseInt(maxcharperpage, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Maxcharperpage = Maxcharperpage
		} else {
			article.Maxcharperpage = 10000
		}

		var istop string = c.Params.Get("istop")
		if len(istop) > 0 {
			Istop, err := strconv.ParseInt(istop, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Istop = Istop
		} else {
			article.Istop = 0
		}

		var status string = c.Params.Get("status")
		if len(status) > 0 {
			Status, err := strconv.ParseInt(status, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Status = Status
		} else {
			article.Status = 1
		}

		var iscomment string = c.Params.Get("iscomment")
		if len(iscomment) > 0 {
			Iscomment, err := strconv.ParseInt(iscomment, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			article.Iscomment = Iscomment
		} else {
			article.Status = 1
		}

		if article.Edit(Id) {

			if len(dosubmit) > 0 {
				c.Flash.Success("编辑内容成功!")
				c.Flash.Out["url"] = "/Content/list/" + cid + "/"
				return c.Redirect("/Message/")
			} else if len(dosubmit_continue) > 0 {
				c.Flash.Success("编辑内容成功!")
				c.Flash.Out["url"] = "/Content/add/" + cid + "/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Success("编辑内容成功!")
				c.Flash.Out["url"] = "/Content/list/" + cid + "/"
				return c.Redirect("/Message/")
			}

		} else {
			c.Flash.Error("编辑内容失败")
			c.Flash.Out["url"] = "/Content/edit/" + cid + "/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}
