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

//测试
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"

type Test struct {
	*revel.Controller
}

func (c Test) Index(admin *models.Admin) revel.Result {
	title := "测试--GoCMS管理系统"

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_info := admin.GetById(UserID)

		admin_panel := new(models.Admin_Panel)
		admin_panel.IsAdd(18, admin_info)
	}

	c.Render(title)
	return c.RenderTemplate("Public/Test.html")
}
