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

//初始化入口文件
import "runtime"
import "strconv"
import "path/filepath"
import "github.com/revel/revel"
import "admin/app/models"

var BasePath, _ = filepath.Abs("")

//定义项目根目录
var ROOT_DIR string = BasePath

//定义项目上传文件目录
var UPLOAD_DIR string = BasePath + "/www/upload/"

func init() {
	revel.OnAppStart(Bootstrap)

	//检测是否登陆
	revel.InterceptFunc(CheckLogin, revel.BEFORE, revel.ALL_CONTROLLERS)
}

//系统初始化变量
func Bootstrap() {
	//多核运行
	np := runtime.NumCPU()
	if np >= 2 {
		runtime.GOMAXPROCS(np - 1)
	}

	if runtime.GOOS == "windows" {
		UPLOAD_DIR = BasePath + "\\www\\upload\\"
	} else {
		UPLOAD_DIR = BasePath + "/www/upload/"
	}
}

//检测登陆
func CheckLogin(c *revel.Controller) revel.Result {

	//登陆页面，CSS, JS, Ajax, 验证码页面 都不进行登陆验证
	if c.Name == "User" && c.MethodName == "Login" || c.Name == "Ajax" || c.Name == "Static" || c.Name == "Captcha" || c.Name == "Kindeditor" {

		if LANG, ok := c.Session["Lang"]; ok {
			//设置语言
			c.RenderArgs["currentLocale"] = LANG
		} else {
			//设置默认语言
			c.RenderArgs["currentLocale"] = "zh"
		}

		return nil
	} else {
		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
				return c.Redirect("/Login/")
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)
			if admin_info.Id <= 0 {
				return c.Redirect("/Login/")
			}

			//设置语言
			c.RenderArgs["currentLocale"] = admin_info.Lang
		} else {
			return c.Redirect("/Login/")
		}
	}

	return nil
}
