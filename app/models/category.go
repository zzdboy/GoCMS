// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package models

//栏目管理
import "strconv"
import "html/template"
import "encoding/json"
import "github.com/revel/revel"

type Category struct {
	Id          int64                  `xorm:"pk"`
	Pid         int64                  `xorm:"int(11)"`
	Type        int64                  `xorm:"int(11)"`
	Name        string                 `xorm:"varchar(255)"`
	Enname      string                 `xorm:"varchar(255)"`
	Desc        string                 `xorm:"text"`
	Url         string                 `xorm:"varchar(100)"`
	Hits        int64                  `xorm:"int(11)"`
	Setting     string                 `xorm:"text"`
	SettingText map[string]interface{} `xorm:"- <- ->"`
	Order       int64                  `xorm:"int(11)"`
	Ismenu      int64                  `xorm:"default 1"`
}

//根据Id获取栏目信息
func (c *Category) GetById(Id int64) *Category {

	category := new(Category)
	has, err := DB_Read.Id(Id).Get(category)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		SettingText := make(map[string]interface{})
		if len(category.Setting) > 0 {
			err = json.Unmarshal([]byte(category.Setting), &SettingText)
			if err != nil {
				revel.WARN.Printf("错误: %v", err)
			} else {
				category.SettingText = SettingText
			}
		}
	}

	return category
}

//添加栏目
func (c *Category) Save() bool {

	category := new(Category)
	category.Pid = c.Pid
	category.Type = c.Type
	category.Name = c.Name
	category.Enname = c.Enname
	category.Desc = c.Desc
	category.Url = c.Url
	category.Hits = 0
	category.Setting = c.Setting
	category.Order = c.Order
	category.Ismenu = c.Ismenu

	has, err := DB_Write.Insert(category)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑栏目
func (c *Category) Edit(Id int64) bool {
	category := new(Category)

	if c.Pid > 0 {
		category.Pid = c.Pid
	}

	if len(c.Name) > 0 {
		category.Name = c.Name
	}

	if len(c.Enname) > 0 {
		category.Enname = c.Enname
	}

	category.Type = c.Type

	if len(c.Url) > 0 {
		category.Url = c.Url
	}

	if len(c.Desc) > 0 {
		category.Desc = c.Desc
	}

	if len(c.Setting) > 0 {
		category.Setting = c.Setting
	}

	if c.Order > 0 {
		category.Order = c.Order
	}

	if c.Ismenu > 0 {
		category.Ismenu = c.Ismenu
	}

	has, err := DB_Write.Id(Id).Update(category)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//快速进入 搜索
func (c *Category) GetCateNameHtml(CateName string) string {
	categorys := make([]*Category, 0)

	Html := ""

	if len(CateName) > 0 {
		DB_Read.Where(" name like '%" + CateName + "%'").Find(&categorys)

		for _, category := range categorys {
			Html += "<li>"
			if category.Type == 0 {
				Html += "<a href=\"/Content/list/" + strconv.FormatInt(category.Id, 10) + "/\">" + category.Name + "</a>"
			} else {
				Html += "<a href=\"/Content/addContent/" + strconv.FormatInt(category.Id, 10) + "/\">" + category.Name + "</a>"
			}

			Html += "</li>"
		}
	}

	return Html
}

//获取所有菜单
//返回HTML
func (c *Category) GetCateGoryHtml(Admin_Info *Admin) template.HTML {
	categorys := make([]*Category, 0)
	DB_Read.Asc("order").Find(&categorys)

	//初始化菜单Map
	category_list := make(map[int64][]*Category)

	for _, category := range categorys {
		if _, ok := category_list[category.Pid]; !ok {
			category_list[category.Pid] = make([]*Category, 0)
		}
		category_list[category.Pid] = append(category_list[category.Pid], category)
	}

	Html := ""

	for _, category := range category_list[0] {
		Html += "<tr>"
		Html += "<td align='center'>" + strconv.FormatInt(category.Id, 10) + "</td>"

		if Admin_Info.Lang == "zh-cn" {
			Html += "<td ><b>" + category.Name + "</b>"
		} else {
			Html += "<td ><b>" + category.Enname + "</b>"
		}
		if category.Ismenu == 0 {
			Html += "&nbsp;&nbsp;<img title='不在导航显示' src='/public/img/icon/gear_disable.png'>"
		}
		Html += "</td>"

		Html += "<td >"
		if category.Type == 0 {
			if Admin_Info.Lang == "zh-cn" {
				Html += "内部栏目"
			} else {
				Html += "Column"
			}
		} else {
			if Admin_Info.Lang == "zh-cn" {
				Html += "<font color=\"blue\">单网页</font>"
			} else {
				Html += "<font color=\"blue\">Single page</font>"
			}
		}
		Html += "</td>"

		Html += "<td align='center'>" + strconv.FormatInt(category.Hits, 10) + "</td>"
		Html += "<td align='center'>" + category.Url + "</td>"

		Html += "<td align=\"center\">"
		if Admin_Info.Lang == "zh-cn" {
			Html += "<a href=\"/Category/add/" + strconv.FormatInt(category.Id, 10) + "/\">添加子菜单</a> |"
			Html += "<a href=\"/Category/edit/" + strconv.FormatInt(category.Id, 10) + "/\">修改</a> |"
			Html += "<a onclick=\"delete_cate(" + strconv.FormatInt(category.Id, 10) + ")\" href=\"javascript:;\">删除</a>"
		} else {
			Html += "<a href=\"/Category/add/" + strconv.FormatInt(category.Id, 10) + "/\">Add Submenu</a> |"
			Html += "<a href=\"/Category/edit/" + strconv.FormatInt(category.Id, 10) + "/\">Edit</a> |"
			Html += "<a onclick=\"delete_cate(" + strconv.FormatInt(category.Id, 10) + ")\" href=\"javascript:;\">Delete</a>"
		}
		Html += "</td>"

		Html += "</tr>"

		for _, category_second := range category_list[category.Id] {
			Html += "<tr>"
			Html += "<td align='center'>" + strconv.FormatInt(category_second.Id, 10) + "</td>"

			Html += "<td >&nbsp;&nbsp;&nbsp;&nbsp;&#12288;&#8866;&nbsp;&nbsp;" + category_second.Name
			if category_second.Ismenu == 0 {
				Html += "<img title='不在导航显示' src='/public/img/icon/gear_disable.png'>"
			}
			Html += "</td>"

			Html += "<td >"
			if category_second.Type == 0 {
				if Admin_Info.Lang == "zh-cn" {
					Html += "内部栏目"
				} else {
					Html += "Column"
				}
			} else {
				if Admin_Info.Lang == "zh-cn" {
					Html += "<font color=\"blue\">单网页</font>"
				} else {
					Html += "<font color=\"blue\">Single page</font>"
				}
			}
			Html += "</td>"

			Html += "<td align='center'>" + strconv.FormatInt(category_second.Hits, 10) + "</td>"
			Html += "<td align='center'>" + category_second.Url + "</td>"

			Html += "<td align=\"center\">"
			if Admin_Info.Lang == "zh-cn" {
				Html += "<a href=\"/Category/add/" + strconv.FormatInt(category_second.Id, 10) + "/\">添加子菜单</a> |"
				Html += "<a href=\"/Category/edit/" + strconv.FormatInt(category_second.Id, 10) + "/\">修改</a> |"
				Html += "<a onclick=\"delete_cate(" + strconv.FormatInt(category_second.Id, 10) + ")\" href=\"javascript:;\">删除</a>"
			} else {
				Html += "<a href=\"/Category/add/" + strconv.FormatInt(category_second.Id, 10) + "/\">Add Submenu</a> |"
				Html += "<a href=\"/Category/edit/" + strconv.FormatInt(category_second.Id, 10) + "/\">Edit</a> |"
				Html += "<a onclick=\"delete_cate(" + strconv.FormatInt(category_second.Id, 10) + ")\" href=\"javascript:;\">Delete</a>"
			}
			Html += "</td>"

			Html += "</tr>"

			for _, category_last := range category_list[category_second.Id] {
				Html += "<tr>"
				Html += "<td align='center'>" + strconv.FormatInt(category_last.Id, 10) + "</td>"

				Html += "<td >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&#12288;&#8866;&nbsp;&nbsp;" + category_last.Name
				if category_last.Ismenu == 0 {
					Html += "<img title='不在导航显示' src='/public/img/icon/gear_disable.png'>"
				}
				Html += "</td>"

				Html += "<td >"
				if category_last.Type == 0 {
					if Admin_Info.Lang == "zh-cn" {
						Html += "内部栏目"
					} else {
						Html += "Column"
					}
				} else {
					if Admin_Info.Lang == "zh-cn" {
						Html += "<font color=\"blue\">单网页</font>"
					} else {
						Html += "<font color=\"blue\">Single page</font>"
					}
				}
				Html += "</td>"

				Html += "<td align='center'>" + strconv.FormatInt(category_last.Hits, 10) + "</td>"
				Html += "<td align='center'>" + category_last.Url + "</td>"

				Html += "<td align=\"center\">"
				if Admin_Info.Lang == "zh-cn" {
					//Html += "<a href=\"/Category/add/" + strconv.FormatInt(category_last.Id, 10) + "/\">添加子菜单</a> |"
					Html += "<a href=\"/Category/edit/" + strconv.FormatInt(category_last.Id, 10) + "/\">修改</a> |"
					Html += "<a onclick=\"delete_cate(" + strconv.FormatInt(category_last.Id, 10) + ")\" href=\"javascript:;\">删除</a>"
				} else {
					//Html += "<a href=\"/Category/add/" + strconv.FormatInt(category_last.Id, 10) + "/\">Add Submenu</a> |"
					Html += "<a href=\"/Category/edit/" + strconv.FormatInt(category_last.Id, 10) + "/\">Edit</a> |"
					Html += "<a onclick=\"delete_cate(" + strconv.FormatInt(category_last.Id, 10) + ")\" href=\"javascript:;\">Delete</a>"
				}
				Html += "</td>"

				Html += "</tr>"
			}
		}
	}

	return template.HTML(Html)
}

//返回菜单树
func (c *Category) GetLeftTree() template.HTML {
	categorys := make([]*Category, 0)
	DB_Read.Asc("order").Find(&categorys)

	//初始化菜单Map
	category_list := make(map[int64][]*Category)

	for _, category := range categorys {
		if _, ok := category_list[category.Pid]; !ok {
			category_list[category.Pid] = make([]*Category, 0)
		}
		category_list[category.Pid] = append(category_list[category.Pid], category)
	}

	Html := "<SCRIPT type=\"text/javascript\">var zNodes =["

	for _, category := range category_list[0] {

		if category.Type == 0 {
			Html += "{ id:" + strconv.FormatInt(category.Id, 10) + ", pId:" + strconv.FormatInt(category.Pid, 10) + ", name:'" + category.Name + "', open:true},"
		} else {
			Html += "{ id:" + strconv.FormatInt(category.Id, 10) + ", pId:" + strconv.FormatInt(category.Pid, 10) + ", name:'" + category.Name + "', url:'/Content/addContent/" + strconv.FormatInt(category.Id, 10) + "/', target :'right', icon :'/public/img/file.gif',  open:true},"
		}

		for _, category_second := range category_list[category.Id] {
			if len(category_list[category_second.Id]) > 0 {
				Html += "{ id:" + strconv.FormatInt(category_second.Id, 10) + ", pId:" + strconv.FormatInt(category_second.Pid, 10) + ", name:'" + category_second.Name + "', open:true},"
			} else {
				if category_second.Type == 0 {
					Html += "{ id:" + strconv.FormatInt(category_second.Id, 10) + ", pId:" + strconv.FormatInt(category_second.Pid, 10) + ", name:'" + category_second.Name + "', url:'/Content/list/" + strconv.FormatInt(category_second.Id, 10) + "/', target :'right', icon :'/public/img/add_content.gif',  open:true},"
				} else {
					Html += "{ id:" + strconv.FormatInt(category_second.Id, 10) + ", pId:" + strconv.FormatInt(category_second.Pid, 10) + ", name:'" + category_second.Name + "', url:'/Content/addContent/" + strconv.FormatInt(category_second.Id, 10) + "/', target :'right', icon :'/public/img/file.gif',  open:true},"
				}
			}
			for _, category_last := range category_list[category_second.Id] {
				if category_last.Type == 0 {
					Html += "{ id:" + strconv.FormatInt(category_last.Id, 10) + ", pId:" + strconv.FormatInt(category_last.Pid, 10) + ", name:'" + category_last.Name + "', url:'/Content/list/" + strconv.FormatInt(category_last.Id, 10) + "/', target :'right', icon :'/public/img/add_content.gif',  open:true},"
				} else {
					Html += "{ id:" + strconv.FormatInt(category_last.Id, 10) + ", pId:" + strconv.FormatInt(category_last.Pid, 10) + ", name:'" + category_last.Name + "', url:'/Content/addContent/" + strconv.FormatInt(category_last.Id, 10) + "/', target :'right', icon :'/public/img/file.gif',  open:true},"
				}

			}
		}
	}

	Html += "];</SCRIPT>"

	return template.HTML(Html)
}

//返回左侧导航
func (c *Category) GetLeftHtml() template.HTML {

	categorys := make([]*Category, 0)
	DB_Read.Asc("order").Find(&categorys)

	//初始化菜单Map
	category_list := make(map[int64][]*Category)

	for _, category := range categorys {
		if _, ok := category_list[category.Pid]; !ok {
			category_list[category.Pid] = make([]*Category, 0)
		}
		category_list[category.Pid] = append(category_list[category.Pid], category)
	}

	Html := ""

	for _, category := range category_list[0] {

		if len(category_list[category.Id]) > 0 {
			Html += "<li id='" + strconv.FormatInt(category.Id, 10) + "'>"
			Html += "<div class=\"hitarea\"></div>"
			Html += "<span class='folder'>" + category.Name + "</span>"
		} else {
			Html += "<li id='" + strconv.FormatInt(category.Id, 10) + "'>"
			Html += "<div class=\"hitarea\"></div>"
			Html += "<a href=\"/Content/list/" + strconv.FormatInt(category.Id, 10) + "/\" target=\"right\"><img alt=\"添加\" src=\"/public/img/add_content.gif\"></a> "
			Html += "<a href='/Content/list/" + strconv.FormatInt(category.Id, 10) + "/' target='right' onclick='open_list(this)'>" + category.Name + "</a>"
			Html += "</li>"
		}

		if len(category_list[category.Id]) > 0 {
			for _, category_second := range category_list[category.Id] {
				if len(category_list[category_second.Id]) > 0 {
					Html += "<ul><span class=''> "
					Html += "<li id='" + strconv.FormatInt(category_second.Id, 10) + "'>"
					Html += "<div class=\"hitarea\"></div>"
					Html += "<span class='folder'>" + category_second.Name + "</span>"
				} else {
					Html += "<ul><span class=''> "
					Html += "<li id='" + strconv.FormatInt(category_second.Id, 10) + "'>"
					Html += "<a href=\"/Content/list/" + strconv.FormatInt(category_second.Id, 10) + "/\" target=\"right\"><img alt=\"添加\" src=\"/public/img/add_content.gif\"></a> "
					Html += "<a href='/Content/list/" + strconv.FormatInt(category_second.Id, 10) + "/' target='right' onclick='open_list(this)'>" + category_second.Name + "</a>"
					Html += "</li>"
				}

				for _, category_last := range category_list[category_second.Id] {
					Html += "<ul><span class=''> "
					Html += "<li id='" + strconv.FormatInt(category_last.Id, 10) + "'>"
					Html += "<a href=\"/Content/list/" + strconv.FormatInt(category_last.Id, 10) + "/\" target=\"right\"><img alt=\"添加\" src=\"/public/img/add_content.gif\"></a> "
					Html += "<a href='/Content/list/" + strconv.FormatInt(category_last.Id, 10) + "/' target='right' onclick='open_list(this)'>" + category_last.Name + "</a>"
					Html += "</span></ul>"
				}

				if len(category_list[category_second.Id]) > 0 {
					Html += "</li></span></ul>"
				} else {
					Html += "</li></span></ul>"
				}
			}
		}

		if len(category_list[category.Id]) > 0 {
			Html += "</li>"
		}
	}

	return template.HTML(Html)
}

//返回菜单Option的HTML
func (c *Category) GetCateGoryOptionHtml(Id int64) template.HTML {

	categorys := make([]*Category, 0)
	DB_Read.Asc("order").Find(&categorys)

	//初始化菜单Map
	category_list := make(map[int64][]*Category)

	for _, category := range categorys {
		if _, ok := category_list[category.Pid]; !ok {
			category_list[category.Pid] = make([]*Category, 0)
		}
		category_list[category.Pid] = append(category_list[category.Pid], category)
	}

	Html := ""
	for _, category := range category_list[0] {

		if category.Id == Id {
			Html += "<option value=" + strconv.FormatInt(category.Id, 10) + " selected ><b>" + category.Name + "</b></option>"
		} else {
			Html += "<option value=" + strconv.FormatInt(category.Id, 10) + "><b>" + category.Name + "</b></option>"
		}

		for _, category_second := range category_list[category.Id] {
			if category_second.Id == Id {
				Html += "<option value=" + strconv.FormatInt(category_second.Id, 10) + " selected >&#12288;&#8866;" + category_second.Name + "</option>"
			} else {
				Html += "<option value=" + strconv.FormatInt(category_second.Id, 10) + ">&#12288;&#8866;" + category_second.Name + "</option>"
			}

			for _, category_last := range category_list[category_second.Id] {
				if category_last.Id == Id {
					Html += "<option value=" + strconv.FormatInt(category_last.Id, 10) + " selected >&#12288;&#12288;&#8866;" + category_last.Name + "</option>"
				} else {
					Html += "<option value=" + strconv.FormatInt(category_last.Id, 10) + ">&#12288;&#12288;&#8866;" + category_last.Name + "</option>"
				}

			}
		}
	}

	return template.HTML(Html)
}

//删除栏目
func (c *Category) DelByID(Id int64) bool {
	category := new(Category)

	has, err := DB_Write.Id(Id).Delete(category)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
