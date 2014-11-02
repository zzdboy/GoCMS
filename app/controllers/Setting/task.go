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

// 管理员设置
import "github.com/robfig/revel/modules/jobs/app/jobs"
import "github.com/revel/revel"

type Task struct {
	*revel.Controller
}

//首页
func (c Task) Index() revel.Result {
	title := "计划任务管理--GoCMS管理系统"

	entries := jobs.MainCron.Entries()

	c.Render(title, entries)

	return c.RenderTemplate("Setting/Task/Index.html")
}
