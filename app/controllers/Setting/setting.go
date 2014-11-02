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

//设置首页

import "github.com/revel/revel"

type Setting struct {
	*revel.Controller
}

func (c Setting) Index() revel.Result {
	title := "设置--GoCMS管理系统"

	c.Render(title)
	return c.RenderTemplate("Setting/Index.html")
}
