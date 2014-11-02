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

//Kindeditor编辑器
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

type Kindeditor struct {
	*revel.Controller
}

//编辑器管理
func (c *Kindeditor) Manager(upload *models.Upload) revel.Result {
	file := make(map[string]interface{})

	//判断是否是系统的分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	basepath, _ := filepath.Abs("")
	config_file := (revel.BasePath + "/conf/config.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	config_conf, _ := config.ReadDefault(config_file)

	//上传文件目录
	upload_dir, _ := config_conf.String("upload", "upload.dir")
	//允许上传的后缀名
	filesuffix, _ := config_conf.String("upload", "upload.filesuffix")

	revel.WARN.Println(filesuffix)

	//前台网站地址
	sitedomain, _ := config_conf.String("website", "website.sitedomain")

	//根目录路径，可以指定绝对路径，比如 /var/www/attached/
	root_path := fmt.Sprintf("%s/www/%s/", basepath, upload_dir)

	//根目录URL，可以指定绝对路径，比如 http://www.yoursite.com/attached/
	root_url := sitedomain + upload_dir

	//目录名
	dir_name := c.Params.Get("dir")

	if dir_name != "" {
		root_path += dir_name + "/"
		root_url += dir_name + "/"
	}

	//相对于根目录的上一级目录
	file["moveup_dir_path"] = ""

	//相对于根目录的当前目录
	file["current_dir_path"] = ""

	//当前目录的URL
	file["current_url"] = ""

	//文件数
	file["total_count"] = 10

	//文件列表数组
	file["file_list"] = ""

	return c.RenderJson(file)
}

//内容发布标题缩略图
func (c *Kindeditor) TitleImage(upload *models.Upload) revel.Result {
	data := make(map[string]interface{})

	//判断是否是系统的分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	//接收上传文件
	imgFile, header, err := c.Request.FormFile("imgFile")
	if err != nil {
		data["error"] = 1
		data["message"] = err.Error()

		return c.RenderJson(data)
	}
	defer imgFile.Close()

	//读取文件数据
	fileData, _ := ioutil.ReadAll(imgFile)

	if len(fileData) >= 1024*1024*2 {
		data["error"] = 1
		data["message"] = "你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!"

		return c.RenderJson(data)
	}

	basepath, _ := filepath.Abs("")
	config_file := (revel.BasePath + "/conf/config.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	config_conf, _ := config.ReadDefault(config_file)

	//上传文件目录
	upload_dir, _ := config_conf.String("upload", "upload.dir")
	//允许上传的后缀名
	titlesuffix, _ := config_conf.String("upload", "upload.titlesuffix")

	//前台网站地址
	sitedomain, _ := config_conf.String("website", "website.sitedomain")

	//文件类型检测
	if !strings.Contains(titlesuffix, path.Ext(header.Filename)) {
		data["error"] = 1
		data["message"] = "文件只支持Office文件，图片和rar存档!"

		return c.RenderJson(data)
	}

	//前台网站调用目录
	web_save_path := fmt.Sprintf("/%s/article/%s", upload_dir, time.Now().Format("2006/01/02/"))

	//文件保存目录
	save_path := fmt.Sprintf("%s/www/%s/article/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
	//字符串替换 /替换为系统分隔符
	save_path = strings.Replace(save_path, "/", separator, -1)

	//创建目录
	err = os.MkdirAll(save_path, os.ModePerm)
	if err != nil {
		revel.WARN.Println(err)
		data["error"] = 1
		data["message"] = "创建目录失败!"

		return c.RenderJson(data)
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

		data["error"] = 1
		data["message"] = "保存图片失败!"
		return c.RenderJson(data)
	}

	//*******************图片处理****************

	//缩略图400
	thumb_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_100" + path.Ext(header.Filename)

	//内容显示图片
	web_url := web_save_path + thumb_name

	//400宽度图片生成
	thumb_135 := save_path + thumb_name
	utils.Resize(old_img, thumb_135, "100x70", "center", "white")

	//*******************图片处理****************

	data["error"] = 0
	data["url"] = sitedomain + web_url

	return c.RenderJson(data)
}

//编辑器上传文件
func (c *Kindeditor) Upload(upload *models.Upload) revel.Result {

	data := make(map[string]interface{})

	//判断是否是系统的分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	//接收上传文件
	imgFile, header, err := c.Request.FormFile("imgFile")
	if err != nil {
		data["error"] = 1
		data["message"] = err.Error()

		return c.RenderJson(data)
	}
	defer imgFile.Close()

	//读取文件数据
	fileData, _ := ioutil.ReadAll(imgFile)

	if len(fileData) >= 1024*1024*2 {
		data["error"] = 1
		data["message"] = "你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!"

		return c.RenderJson(data)
	}

	basepath, _ := filepath.Abs("")
	config_file := (revel.BasePath + "/conf/config.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	config_conf, _ := config.ReadDefault(config_file)

	//上传文件目录
	upload_dir, _ := config_conf.String("upload", "upload.dir")
	//允许上传的后缀名
	filesuffix, _ := config_conf.String("upload", "upload.filesuffix")

	//前台网站地址
	sitedomain, _ := config_conf.String("website", "website.sitedomain")

	//文件类型检测
	if !strings.Contains(filesuffix, path.Ext(header.Filename)) {
		data["error"] = 1
		data["message"] = "文件只支持Office文件，图片和rar存档!"

		return c.RenderJson(data)
	}

	//前台网站调用目录
	web_save_path := fmt.Sprintf("/%s/article/%s", upload_dir, time.Now().Format("2006/01/02/"))

	//文件保存目录
	save_path := fmt.Sprintf("%s/www/%s/article/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
	//字符串替换 /替换为系统分隔符
	save_path = strings.Replace(save_path, "/", separator, -1)

	//创建目录
	err = os.MkdirAll(save_path, os.ModePerm)
	if err != nil {
		revel.WARN.Println(err)
		data["error"] = 1
		data["message"] = "创建目录失败!"

		return c.RenderJson(data)
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

		data["error"] = 1
		data["message"] = "保存图片失败!"
		return c.RenderJson(data)
	}

	//*******************图片处理****************

	//缩略图400
	new_img_thumb_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_400" + path.Ext(header.Filename)

	//内容显示图片
	web_url := web_save_path + new_img_thumb_name

	//400宽度图片生成
	new_img_400 := save_path + new_img_thumb_name
	utils.Resize(old_img, new_img_400, "400", "center", "white")

	//圆角图片生成
	new_img_thumb_0x4 := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_0x4" + path.Ext(header.Filename)
	new_img_0x4 := save_path + new_img_thumb_0x4
	utils.Vignette(new_img_400, new_img_0x4, "0x4")

	//图片反色处理
	new_img_thumb_negate := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_negate" + path.Ext(header.Filename)
	new_img_negate := save_path + new_img_thumb_negate
	utils.Negate(new_img_400, new_img_negate)

	//图片加水印之文字水印处理
	new_img_watermark := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_watermark" + path.Ext(header.Filename)
	new_img_watermark_file := save_path + new_img_watermark
	utils.WatermarkText(new_img_400, new_img_watermark_file)

	web_url = web_save_path + new_img_watermark
	//*******************图片处理****************

	data["error"] = 0
	data["url"] = sitedomain + web_url

	return c.RenderJson(data)
}

//内容发布标题缩略图
func (c *Kindeditor) AnnounceImage(upload *models.Upload) revel.Result {
	data := make(map[string]interface{})

	//判断是否是系统的分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	//接收上传文件
	imgFile, header, err := c.Request.FormFile("imgFile")
	if err != nil {
		data["error"] = 1
		data["message"] = err.Error()

		return c.RenderJson(data)
	}
	defer imgFile.Close()

	//读取文件数据
	fileData, _ := ioutil.ReadAll(imgFile)

	if len(fileData) >= 1024*1024*2 {
		data["error"] = 1
		data["message"] = "你上传大小为" + utils.FileSize(len(fileData)) + ",文件应小于2M!"

		return c.RenderJson(data)
	}

	basepath, _ := filepath.Abs("")
	config_file := (revel.BasePath + "/conf/config.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	config_conf, _ := config.ReadDefault(config_file)

	//上传文件目录
	upload_dir, _ := config_conf.String("upload", "upload.dir")
	//允许上传的后缀名
	filesuffix, _ := config_conf.String("upload", "upload.filesuffix")

	//前台网站地址
	sitedomain, _ := config_conf.String("website", "website.sitedomain")

	//文件类型检测
	if !strings.Contains(filesuffix, path.Ext(header.Filename)) {
		data["error"] = 1
		data["message"] = "文件只支持Office文件，图片和rar存档!"

		return c.RenderJson(data)
	}

	//前台网站调用目录
	web_save_path := fmt.Sprintf("/%s/announce/%s", upload_dir, time.Now().Format("2006/01/02/"))

	//文件保存目录
	save_path := fmt.Sprintf("%s/www/%s/announce/%s", basepath, upload_dir, time.Now().Format("2006/01/02/"))
	//字符串替换 /替换为系统分隔符
	save_path = strings.Replace(save_path, "/", separator, -1)

	//创建目录
	err = os.MkdirAll(save_path, os.ModePerm)
	if err != nil {
		revel.WARN.Println(err)
		data["error"] = 1
		data["message"] = "创建目录失败!"

		return c.RenderJson(data)
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

		data["error"] = 1
		data["message"] = "保存失败!"
		return c.RenderJson(data)
	}

	//*******************图片处理****************

	//缩略图400
	thumb_name := time.Now().Format("20060102150404") + strconv.Itoa(rand_num) + "_400" + path.Ext(header.Filename)

	//内容显示图片
	web_url := web_save_path + thumb_name

	//400宽度图片生成
	thumb_135 := save_path + thumb_name
	utils.Resize(old_img, thumb_135, "400", "center", "white")

	//*******************图片处理****************

	data["error"] = 0
	data["url"] = sitedomain + web_url

	return c.RenderJson(data)
}
