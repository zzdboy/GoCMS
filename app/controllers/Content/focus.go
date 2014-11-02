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

//焦点图管理--焦点图列表
import "fmt"
import "strings"
import "path"
import "time"
import "os"
import "strconv"
import "math/rand"
import "path/filepath"
import "io/ioutil"
import "admin/utils"
import "github.com/revel/revel"
import "github.com/revel/config"
import "admin/app/models"

type Focus struct {
	*revel.Controller
}

func (c Focus) Index(focus *models.Focus) revel.Result {

	title := "焦点图列表--GoCMS管理系统"

	var page string = c.Params.Get("page")
	var search string = c.Params.Get("search")

	FocusCate := new(models.FocusCate)
	Cate_list := FocusCate.GetCateList()

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		focus_list, pages, where := focus.GetByAll(search, Page, 10)

		c.Render(title, focus_list, Cate_list, where, pages)
	} else {
		focus_list, pages, where := focus.GetByAll(search, 1, 10)

		c.Render(title, focus_list, Cate_list, where, pages)
	}

	return c.RenderTemplate("Content/Focus/Index.html")
}

//添加焦点图
func (c Focus) Add(focus *models.Focus) revel.Result {

	if c.Request.Method == "GET" {

		title := "添加焦点图--GoCMS管理系统"

		FocusCate := new(models.FocusCate)
		Cate_list := FocusCate.GetCateList()

		c.Render(title, Cate_list)
		return c.RenderTemplate("Content/Focus/Add.html")

	} else {

		var cid string = c.Params.Get("cid")
		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			focus.Cid = Cid

			focusCate := new(models.FocusCate)
			focusCate_info := focusCate.GetById(Cid)

			if focusCate_info.Id <= 0 {
				c.Flash.Error("焦点图分类错误!")
				c.Flash.Out["url"] = "/Focus/Add/"
				return c.Redirect("/Message/")
			}

			//****************************************
			//图片

			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			//接收上传文件
			imgFile, header, err := c.Request.FormFile("img")
			if err != nil {
				c.Flash.Error("图片上传错误!")
				c.Flash.Out["url"] = "/Focus/Add/"
				return c.Redirect("/Message/")
			}
			defer imgFile.Close()

			//读取文件数据
			fileData, _ := ioutil.ReadAll(imgFile)

			if len(fileData) >= 1024*1024*2 {
				c.Flash.Error("你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!")
				c.Flash.Out["url"] = "/Focus/Add/"
				return c.Redirect("/Message/")
			}

			basepath, _ := filepath.Abs("")
			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//上传文件目录
			upload_dir, _ := config_conf.String("upload", "upload.dir")
			//允许上传的后缀名
			titlesuffix, _ := config_conf.String("upload", "upload.titlesuffix")

			//文件类型检测
			if !strings.Contains(titlesuffix, path.Ext(header.Filename)) {
				c.Flash.Error("文件只支持图片!")
				c.Flash.Out["url"] = "/Focus/Add/"
				return c.Redirect("/Message/")
			}

			//前台网站调用目录
			web_save_path := fmt.Sprintf("/%s/focus/%s", upload_dir, time.Now().Format("2006/01/02/"))

			//文件保存目录
			save_path := fmt.Sprintf("%s/www/%s/focus/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
			//字符串替换 /替换为系统分隔符
			save_path = strings.Replace(save_path, "/", separator, -1)

			//创建目录
			err = os.MkdirAll(save_path, os.ModePerm)
			if err != nil {
				c.Flash.Error("创建目录失败!")
				c.Flash.Out["url"] = "/Focus/Add/"
				return c.Redirect("/Message/")
			}

			//新文件名
			rand.Seed(time.Now().UnixNano())
			rand_num := rand.Intn(99999)

			//原图
			new_file_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + path.Ext(header.Filename)

			//保存文件
			old_img := save_path + new_file_name

			err = ioutil.WriteFile(old_img, fileData, os.ModePerm)
			if err != nil {
				revel.WARN.Println(err)
				c.Flash.Error("保存图片失败!")
				c.Flash.Out["url"] = "/Focus/Add/"
				return c.Redirect("/Message/")
			}

			thumb_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_thumb" + path.Ext(header.Filename)

			//内容显示图片
			web_url_img := web_save_path + thumb_name

			//400宽度图片生成
			thumb := save_path + thumb_name
			thumb_width := strconv.Itoa(int(focusCate_info.Width)) + "x" + strconv.Itoa(int(focusCate_info.Height))
			utils.Resize(old_img, thumb, thumb_width, "center", "white")

			focus.Img = web_url_img

			//***************************************

		} else {
			c.Flash.Error("请选择所属分类!")
			c.Flash.Out["url"] = "/Focus/Add/"
			return c.Redirect("/Message/")
		}

		UserID := utils.GetSession("UserID", c.Session)

		if len(UserID) > 0 {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			focus.Aid = UserID
		}

		var title string = c.Params.Get("title")
		if len(title) > 0 {
			focus.Title = title
		} else {
			c.Flash.Error("请输入标题!")
			c.Flash.Out["url"] = "/Focus/Add/"
			return c.Redirect("/Message/")
		}

		var url string = c.Params.Get("url")
		if len(url) > 0 {
			focus.Url = url
		} else {
			c.Flash.Error("请输入地址!")
			c.Flash.Out["url"] = "/Focus/Add/"
			return c.Redirect("/Message/")
		}

		var content string = c.Params.Get("content")
		if len(content) > 0 {
			focus.Content = content
		} else {
			c.Flash.Error("请输入摘要!")
			c.Flash.Out["url"] = "/Focus/Add/"
			return c.Redirect("/Message/")
		}

		var order string = c.Params.Get("order")
		if len(order) > 0 {

			Order, err := strconv.ParseInt(order, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			focus.Order = Order
		} else {
			focus.Order = 0
		}

		var status string = c.Params.Get("status")
		Status, err := strconv.ParseInt(status, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		focus.Status = Status

		if focus.Save() {
			c.Flash.Success("添加焦点图成功")
			c.Flash.Out["url"] = "/Focus/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加焦点图失败")
			c.Flash.Out["url"] = "/Focus/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑焦点图
func (c Focus) Edit(focus *models.Focus) revel.Result {

	if c.Request.Method == "GET" {

		title := "编辑焦点图--GoCMS管理系统"

		FocusCate := new(models.FocusCate)
		Cate_list := FocusCate.GetCateList()

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			focus_info := focus.GetById(Id)

			c.Render(title, Cate_list, focus_info)
		} else {
			c.Render(title, Cate_list)
		}

		return c.RenderTemplate("Content/Focus/Edit.html")

	} else {
		var cid string = c.Params.Get("cid")

		var id string = c.Params.Get("id")
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if len(cid) > 0 {

			Cid, err := strconv.ParseInt(cid, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			focus.Cid = Cid

			focusCate := new(models.FocusCate)
			focusCate_info := focusCate.GetById(Cid)

			if focusCate_info.Id <= 0 {
				c.Flash.Error("焦点图分类错误!")
				c.Flash.Out["url"] = "/Focus/edit/" + id
				return c.Redirect("/Message/")
			}

			//****************************************
			//图片

			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			//接收上传文件
			imgFile, header, err := c.Request.FormFile("img")
			if err != nil {
				defer imgFile.Close()
			} else {
				//读取文件数据
				fileData, _ := ioutil.ReadAll(imgFile)

				if len(fileData) >= 1024*1024*2 {
					c.Flash.Error("你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!")
					c.Flash.Out["url"] = "/Focus/edit/" + id
					return c.Redirect("/Message/")
				}

				basepath, _ := filepath.Abs("")
				config_file := (revel.BasePath + "/conf/config.conf")
				config_file = strings.Replace(config_file, "/", separator, -1)
				config_conf, _ := config.ReadDefault(config_file)

				//上传文件目录
				upload_dir, _ := config_conf.String("upload", "upload.dir")
				//允许上传的后缀名
				titlesuffix, _ := config_conf.String("upload", "upload.titlesuffix")

				//文件类型检测
				if !strings.Contains(titlesuffix, path.Ext(header.Filename)) {
					c.Flash.Error("文件只支持图片!")
					c.Flash.Out["url"] = "/Focus/edit/" + id
					return c.Redirect("/Message/")
				}

				//前台网站调用目录
				web_save_path := fmt.Sprintf("/%s/focus/%s", upload_dir, time.Now().Format("2006/01/02/"))

				//文件保存目录
				save_path := fmt.Sprintf("%s/www/%s/focus/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
				//字符串替换 /替换为系统分隔符
				save_path = strings.Replace(save_path, "/", separator, -1)

				//创建目录
				err = os.MkdirAll(save_path, os.ModePerm)
				if err != nil {
					c.Flash.Error("创建目录失败!")
					c.Flash.Out["url"] = "/Focus/edit/" + id
					return c.Redirect("/Message/")
				}

				//新文件名
				rand.Seed(time.Now().UnixNano())
				rand_num := rand.Intn(99999)

				//原图
				new_file_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + path.Ext(header.Filename)

				//保存文件
				old_img := save_path + new_file_name

				err = ioutil.WriteFile(old_img, fileData, os.ModePerm)
				if err != nil {
					revel.WARN.Println(err)
					c.Flash.Error("保存图片失败!")
					c.Flash.Out["url"] = "/Focus/edit/" + id
					return c.Redirect("/Message/")
				}

				thumb_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_thumb" + path.Ext(header.Filename)

				//内容显示图片
				web_url_img := web_save_path + thumb_name

				//400宽度图片生成
				thumb := save_path + thumb_name
				thumb_width := strconv.Itoa(int(focusCate_info.Width)) + "x" + strconv.Itoa(int(focusCate_info.Height))
				utils.Resize(old_img, thumb, thumb_width, "center", "white")

				focus.Img = web_url_img

				//***************************************

				defer imgFile.Close()
			}

		} else {
			c.Flash.Error("请选择所属分类!")
			c.Flash.Out["url"] = "/Focus/edit/" + id
			return c.Redirect("/Message/")
		}

		UserID := utils.GetSession("UserID", c.Session)

		if len(UserID) > 0 {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
			focus.Aid = UserID
		}

		var title string = c.Params.Get("title")
		if len(title) > 0 {
			focus.Title = title
		} else {
			c.Flash.Error("请输入标题!")
			c.Flash.Out["url"] = "/Focus/edit/" + id
			return c.Redirect("/Message/")
		}

		var url string = c.Params.Get("url")
		if len(url) > 0 {
			focus.Url = url
		} else {
			c.Flash.Error("请输入地址!")
			c.Flash.Out["url"] = "/Focus/edit/" + id
			return c.Redirect("/Message/")
		}

		var content string = c.Params.Get("content")
		if len(content) > 0 {
			focus.Content = content
		} else {
			c.Flash.Error("请输入摘要!")
			c.Flash.Out["url"] = "/Focus/edit/" + id
			return c.Redirect("/Message/")
		}

		var order string = c.Params.Get("order")
		if len(order) > 0 {

			Order, err := strconv.ParseInt(order, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			focus.Order = Order
		} else {
			focus.Order = 0
		}

		var status string = c.Params.Get("status")
		Status, err := strconv.ParseInt(status, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		focus.Status = Status

		if focus.Edit(Id) {
			c.Flash.Success("编辑焦点图成功")
			c.Flash.Out["url"] = "/Focus/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("编辑焦点图失败")
			c.Flash.Out["url"] = "/Focus/edit/" + id
			return c.Redirect("/Message/")
		}
	}
}
