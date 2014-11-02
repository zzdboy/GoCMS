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

//后台公用处理
import "strconv"
import "github.com/revel/revel"
import "admin/app/models"

type Public struct {
	*revel.Controller
}

//后台地图
func (c *Public) Map(menu *models.Menu) revel.Result {

	title := "后台地图--GoCMS管理系统"

	if UserID, ok := c.Session["UserID"]; ok {
		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		//返回后台地图
		map_html := menu.GetMenuMap(admin_info)

		c.Render(title, map_html)
	} else {
		c.Render(title)
	}

	return c.RenderTemplate("Public/map.html")
}

//生成网站首页
func (c *Public) CreateHtml() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/createhtml.html")
}

//搜索
func (c *Public) Search() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/search.html")
}

//消息提示
func (c *Public) Message() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/message.html")
}
