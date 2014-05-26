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

//日志管理
import "strconv"
import "admin/app/models"
import "github.com/revel/revel"

type Logs struct {
	*revel.Controller
}

//日志列表
func (c Logs) Index(logs *models.Logs) revel.Result {

	title := "日志管理--GoCMS管理系统"

	var page string = c.Params.Get("page")
	var search string = c.Params.Get("search")

	var Page int64 = 1

	if len(page) > 0 {
		Page, _ = strconv.ParseInt(page, 10, 64)
	}

	logs_list, pages, where := logs.GetByAll(search, Page, 10)

	c.Render(title, logs_list, where, pages)

	return c.RenderTemplate("Setting/Logs/Index.html")
}

//清空日志
func (c Logs) DelAll(logs *models.Logs) revel.Result {

	data := make(map[string]string)

	IsDel := logs.DelAll()

	if IsDel {

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
			desc := "清空日志|^|日志管理"
			logs.Save(admin_info, c.Controller, desc)
		}

		//*****************************************

		data["status"] = "1"
		data["url"] = "/Message/"
		data["message"] = "清空日志完成!"
		return c.RenderJson(data)
	} else {
		data["status"] = "0"
		data["url"] = "/Message/"
		data["message"] = "清空日志失败!"
		return c.RenderJson(data)
	}
}
