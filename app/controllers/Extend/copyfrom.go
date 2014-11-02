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

//来源管理
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

type Copyfrom struct {
	*revel.Controller
}

func (c Copyfrom) Index(copyfrom *models.Copyfrom) revel.Result {
	title := "来源管理--GoCMS管理系统"

	var page string = c.Params.Get("page")

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

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		copyfrom_list, pages := copyfrom.GetByAll(Page, 10)

		c.Render(title, copyfrom_list, sitedomain, pages)
	} else {
		copyfrom_list, pages := copyfrom.GetByAll(1, 10)

		c.Render(title, copyfrom_list, sitedomain, pages)
	}

	return c.RenderTemplate("Extend/Copyfrom/Index.html")
}

//添加来源
func (c Copyfrom) Add(copyfrom *models.Copyfrom) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加来源--GoCMS管理系统"

		c.Render(title)
		return c.RenderTemplate("Extend/Copyfrom/Add.html")
	} else {

		//接收上传文件
		thumb, header, err := c.Request.FormFile("thumb")
		if err != nil {
			//来源logo
			copyfrom.Thumb = ""
		} else {
			defer thumb.Close()

			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			fileData, _ := ioutil.ReadAll(thumb)

			if len(fileData) >= 1024*1024*2 {
				c.Flash.Error("你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!")
				c.Flash.Out["url"] = "/Copyfrom/Add/"
				return c.Redirect("/Message/")
			}

			basepath, _ := filepath.Abs("")
			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//上传文件目录
			upload_dir, _ := config_conf.String("upload", "upload.dir")
			//允许上传的后缀名
			filesuffix, _ := config_conf.String("upload", "upload.filesuffix")

			//文件类型检测
			if !strings.Contains(filesuffix, path.Ext(header.Filename)) {
				c.Flash.Error("文件只支持图片!")
				c.Flash.Out["url"] = "/Copyfrom/Add/"
				return c.Redirect("/Message/")
			}

			//文件保存目录
			web_save_path := fmt.Sprintf("/%s/copyfrom/%s", upload_dir, time.Now().Format("2006/01/02/"))
			save_path := fmt.Sprintf("%s/www/%s/copyfrom/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
			//字符串替换 /替换为系统分隔符
			save_path = strings.Replace(save_path, "/", separator, -1)

			//新文件名
			rand.Seed(time.Now().UnixNano())
			rand_num := rand.Intn(99999)
			new_file_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + path.Ext(header.Filename)

			//创建目录
			error := os.MkdirAll(save_path, os.ModePerm)
			if error != nil {
				c.Flash.Error(error.Error())
				c.Flash.Out["url"] = "/Copyfrom/Add/"
				return c.Redirect("/Message/")
			}

			//保存文件
			file_dir := save_path + new_file_name
			save_url := web_save_path + new_file_name

			e := ioutil.WriteFile(file_dir, fileData, os.ModePerm)
			if e != nil {
				c.Flash.Error(e.Error())
				c.Flash.Out["url"] = "/Copyfrom/Add/"
				return c.Redirect("/Message/")
			}

			//来源logo
			copyfrom.Thumb = save_url
		}

		//来源名称
		var sitename string = c.Params.Get("sitename")
		if len(sitename) > 0 {
			copyfrom.Sitename = sitename
		} else {
			c.Flash.Error("请输入来源名称!")
			c.Flash.Out["url"] = "/Copyfrom/Add/"
			return c.Redirect("/Message/")
		}

		//来源链接
		var siteurl string = c.Params.Get("siteurl")
		if len(siteurl) > 0 {
			copyfrom.Siteurl = siteurl
		} else {
			c.Flash.Error("请输入来源链接!")
			c.Flash.Out["url"] = "/Copyfrom/Add/"
			return c.Redirect("/Message/")
		}

		if copyfrom.Save() {

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
				desc := "添加来源:" + sitename + "|^|来源管理"
				logs.Save(admin_info, c.Controller, desc)
			}

			//*****************************************

			c.Flash.Success("添加来源成功!")
			c.Flash.Out["url"] = "/Copyfrom/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加来源失败")
			c.Flash.Out["url"] = "/Copyfrom/Add/"
			return c.Redirect("/Message/")
		}

	}
}

//编辑来源
func (c Copyfrom) Edit(copyfrom *models.Copyfrom) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加来源--GoCMS管理系统"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			copyfrom_info := copyfrom.GetById(Id)

			c.Render(title, copyfrom_info, Id)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Extend/Copyfrom/Edit.html")
	} else {

		var id string = c.Params.Get("id")

		if len(id) > 0 {

			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			//接收上传文件
			thumb, header, err := c.Request.FormFile("thumb")
			if err != nil {
				//来源logo
				copyfrom.Thumb = ""
			} else {
				defer thumb.Close()

				//判断是否是系统的分隔符
				separator := "/"
				if os.IsPathSeparator('\\') {
					separator = "\\"
				} else {
					separator = "/"
				}

				fileData, _ := ioutil.ReadAll(thumb)

				if len(fileData) >= 1024*1024*2 {
					c.Flash.Error("你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!")
					c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				basepath, _ := filepath.Abs("")
				config_file := fmt.Sprintf("%s/admin/conf/config.conf", basepath)
				config_file = strings.Replace(config_file, "/", separator, -1)
				config_conf, _ := config.ReadDefault(config_file)

				//上传文件目录
				upload_dir, _ := config_conf.String("upload", "upload.dir")
				//允许上传的后缀名
				filesuffix, _ := config_conf.String("upload", "upload.filesuffix")

				//文件类型检测
				if !strings.Contains(filesuffix, path.Ext(header.Filename)) {
					c.Flash.Error("文件只支持图片!")
					c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				//文件保存目录
				web_save_path := fmt.Sprintf("/%s/copyfrom/%s", upload_dir, time.Now().Format("2006/01/02/"))
				save_path := fmt.Sprintf("%s/www/%s/copyfrom/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
				//字符串替换 /替换为系统分隔符
				save_path = strings.Replace(save_path, "/", separator, -1)

				//新文件名
				rand.Seed(time.Now().UnixNano())
				rand_num := rand.Intn(99999)
				new_file_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + path.Ext(header.Filename)

				//创建目录
				error := os.MkdirAll(save_path, os.ModePerm)
				if error != nil {
					c.Flash.Error(error.Error())
					c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				//保存文件
				file_dir := save_path + new_file_name
				save_url := web_save_path + new_file_name

				e := ioutil.WriteFile(file_dir, fileData, os.ModePerm)
				if e != nil {
					c.Flash.Error(e.Error())
					c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
					return c.Redirect("/Message/")
				}

				//来源logo
				copyfrom.Thumb = save_url
			}

			//来源名称
			var sitename string = c.Params.Get("sitename")
			if len(sitename) > 0 {
				copyfrom.Sitename = sitename
			} else {
				c.Flash.Error("请输入来源名称!")
				c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			//来源链接
			var siteurl string = c.Params.Get("siteurl")
			if len(siteurl) > 0 {
				copyfrom.Siteurl = siteurl
			} else {
				c.Flash.Error("请输入来源链接!")
				c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			if copyfrom.Edit(Id) {

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
					desc := "编辑来源:" + sitename + "|^|来源管理"
					logs.Save(admin_info, c.Controller, desc)
				}

				//*****************************************

				c.Flash.Success("编辑来源成功!")
				c.Flash.Out["url"] = "/Copyfrom/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑来源失败")
				c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑来源失败")
			c.Flash.Out["url"] = "/Copyfrom/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}
	}
}

//删除来源
func (c Copyfrom) Delete(copyfrom *models.Copyfrom) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {

		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if copyfrom.DelByID(Id) {
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
