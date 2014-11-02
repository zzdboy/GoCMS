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

//模块管理--投诉建议
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"

type Complaints struct {
	*revel.Controller
}

func (c Complaints) Index(complaints *models.Complaints) revel.Result {
	title := "投诉建议--GoCMS管理系统"

	var page string = c.Params.Get("page")

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		complaints_list, pages := complaints.GetByAll(Page, 10)

		c.Render(title, complaints_list, pages)
	} else {
		complaints_list, pages := complaints.GetByAll(1, 10)
		c.Render(title, complaints_list, pages)
	}

	return c.RenderTemplate("Module/Complaints/Index.html")
}

//删除投诉建议
func (c Complaints) Delete(complaints *models.Complaints) revel.Result {

	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if complaints.DelByID(Id) {
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
