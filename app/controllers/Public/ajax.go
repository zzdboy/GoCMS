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

//AJAX操作
import "strconv"
import "admin/utils"
import "github.com/revel/revel"
import "github.com/dchest/captcha"
import "admin/app/models"

type Ajax struct {
	*revel.Controller
}

//获取验证码
func (c *Ajax) GetCaptcha() revel.Result {
	CaptchaId := captcha.NewLen(6)
	return c.RenderText(CaptchaId)
}

//当前位置
func (c *Ajax) Pos(menu *models.Menu) revel.Result {
	var id string = c.Params.Get("id")

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		revel.WARN.Println(err)
	}

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		//获取登陆用户信息
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		menu_str := menu.GetPos(Id, admin_info)
		return c.RenderText(menu_str)
	} else {
		return c.RenderText("")
	}
}

//获取快捷方式
func (c *Ajax) GetPanel(admin_panel *models.Admin_Panel) revel.Result {
	if UserID, ok := c.Session["UserID"]; ok {

		var mid string = c.Params.Get("mid")

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		Mid, err := strconv.ParseInt(mid, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		//获取登陆用户信息
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		panel_info := admin_panel.GetByMid(Mid, admin_info)

		if panel_info.Id > 0 {
			Html := "<span><a target='right' href='/" + panel_info.Url + "/'>" + panel_info.Name + "</a><a class='panel-delete' href='javascript:delete_panel();'></a></span>"
			return c.RenderText(Html)
		} else {
			Html := ""
			return c.RenderText(Html)
		}

	} else {
		Html := "<span><a href='javascript:;'>未登陆</a></span>"
		return c.RenderText(Html)
	}
}

//删除快捷方式
func (c *Ajax) DelPanel(admin_panel *models.Admin_Panel) revel.Result {
	if UserID, ok := c.Session["UserID"]; ok {
		var mid string = c.Params.Get("mid")

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		Mid, err := strconv.ParseInt(mid, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		//获取登陆用户信息
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		is_True := admin_panel.DelPanel(Mid, admin_info)

		data := make(map[string]string)

		if is_True {
			data["status"] = "0"
			data["message"] = "取消成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "2"
			data["message"] = "取消失败!"
			return c.RenderJson(data)
		}
	} else {
		data := make(map[string]string)
		data["status"] = "1"
		data["message"] = "未登陆!"
		return c.RenderJson(data)
	}
}

//添加快捷方式
func (c *Ajax) AddPanel(admin_panel *models.Admin_Panel) revel.Result {

	if UserID, ok := c.Session["UserID"]; ok {

		var mid string = c.Params.Get("mid")

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		Mid, err := strconv.ParseInt(mid, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		//获取登陆用户信息
		admin := new(models.Admin)
		admin_info := admin.GetById(UserID)

		//是否已添加快捷方式
		isAdd := admin_panel.IsAdd(Mid, admin_info)

		if isAdd {
			panel_info := admin_panel.GetByMid(Mid, admin_info)

			Html := "<span><a target='right' href='/" + panel_info.Url + "/'>" + panel_info.Name + "</a><a class='panel-delete' href='javascript:delete_panel();'></a></span>"
			return c.RenderText(Html)
		} else {
			isFinish := admin_panel.AddPanel(Mid, admin_info)

			if isFinish {
				panel_info := admin_panel.GetByMid(Mid, admin_info)

				Html := "<span><a target='right' href='/" + panel_info.Url + "/'>" + panel_info.Name + "</a><a class='panel-delete' href='javascript:delete_panel();'></a></span>"
				return c.RenderText(Html)
			} else {
				Html := "<span><a href='javascript:;'>请重新添加</a></span>"
				return c.RenderText(Html)
			}
		}

	} else {
		Html := "<span><a href='javascript:;'>未登陆</a></span>"
		return c.RenderText(Html)
	}
}

//检查消息
func (c *Ajax) GetMessage() revel.Result {
	data := make(map[string]string)

	data["status"] = "0"
	data["message"] = "请填写用户名!"
	return c.RenderJson(data)
}

//锁屏
func (c *Ajax) ScreenLock() revel.Result {
	data := make(map[string]string)

	c.Session["lock_screen"] = "1"

	data["status"] = "1"
	data["message"] = "锁屏!"
	return c.RenderJson(data)
}

//解锁
func (c *Ajax) ScreenUnlock(admin *models.Admin) revel.Result {
	var lock_password string = c.Params.Get("lock_password")

	if lock_password == "" || len(lock_password) <= 0 {
		return c.RenderText("2")
	}

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_info := admin.GetById(UserID)

		if admin_info.Password != utils.Md5(lock_password) {
			return c.RenderText("3")
		} else {
			c.Session["lock_screen"] = "0"
			return c.RenderText("1")
		}
	} else {
		return c.RenderText("4")
	}
}
