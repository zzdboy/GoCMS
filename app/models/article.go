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

//内容管理
import "os"
import "fmt"
import "time"
import "strconv"
import "strings"
import "admin/utils"
import "html/template"
import "github.com/revel/revel"
import "github.com/revel/config"

type Article struct {
	Id             int64      `xorm:"pk autoincr"`
	Cid            int64      `xorm:"index"`
	Category       *Category  `xorm:"- <- ->"`
	Aid            int64      `xorm:"index"`
	Admin          *Admin     `xorm:"- <- ->"`
	Title          string     `xorm:"varchar(80)"`
	Color          string     `xorm:"char(24)"`
	Font           string     `xorm:"char(24)"`
	Thumb          string     `xorm:"varchar(80)"`
	Content        string     `xorm:"text"`
	Copyfrom       string     `xorm:"varchar(100)"`
	Keywords       string     `xorm:"varchar(255)"`
	Description    string     `xorm:"varchar(255)"`
	Relation       string     `xorm:"varchar(255)"`
	RelationList   []*Article `xorm:"- <- ->"`
	Pagetype       int64      `xorm:"tynyint(1)"`
	Maxcharperpage int64      `xorm:"mediumint(6)"`
	Istop          int64      `xorm:"tynyint(1)"`
	Status         int64      `xorm:"tynyint(1)"`
	Hits           int64      `xorm:"tynyint(5)"`
	Iscomment      int64      `xorm:"tynyint(1)"`
	Createtime     string     `xorm:"DateTime"`
	Updatetime     string     `xorm:"DateTime"`
}

//根据Id获取内容信息
func (a *Article) GetById(Id int64) *Article {
	article := new(Article)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(article)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		admin := new(Admin)
		article.Admin = admin.GetById(article.Aid)

		//所属栏目
		category := new(Category)
		article.Category = category.GetById(article.Cid)

		if article.Thumb != "" {
			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//前台网站地址
			sitedomain, _ := config_conf.String("website", "website.sitedomain")
			article.Thumb = sitedomain + article.Thumb
		}

		//相关文章
		if len(article.Relation) > 0 {
			DB_Write.Where("id in (" + article.Relation + ")").Find(&article.RelationList)
		}

	}

	return article
}

//根据Cid获取栏目单页面内容
func (a *Article) GetByCid(Cid int64) *Article {
	article := new(Article)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Where("`cid`=" + strconv.FormatInt(Cid, 10)).Get(article)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		admin := new(Admin)
		article.Admin = admin.GetById(article.Aid)

		//所属栏目
		category := new(Category)
		article.Category = category.GetById(article.Cid)
	}

	return article
}

//获取内容列表
func (a *Article) GetByAll(search string, Cid int64, Page int64, Perpage int64) (article_arr []*Article, html template.HTML, where map[string]interface{}) {
	article_list := []*Article{}

	//查询条件
	var WhereStr string = " 1 AND `cid`=" + strconv.FormatInt(Cid, 10) + " AND "

	if len(search) > 0 {

		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if where["start_time"] != "" {
			WhereStr += " `createtime`='" + fmt.Sprintf("%s", where["start_time"]) + "' AND "
		}

		if where["start_time"] != "" {
			WhereStr += " `createtime`='" + fmt.Sprintf("%s", where["end_time"]) + "' AND "
		}

		if where["istop"] != "" {
			WhereStr += " `istop`='" + fmt.Sprintf("%s", where["istop"]) + "' AND "
		}

		if where["searchtype"] != "" && where["keyword"] != "" {

			if where["searchtype"] == "1" {
				//标题
				WhereStr += " `title` like '%" + fmt.Sprintf("%s", where["keyword"]) + "%' AND "
			} else if where["searchtype"] == "2" {
				//简介
				WhereStr += " `description` like '%" + fmt.Sprintf("%s", where["keyword"]) + "%' AND "
			} else if where["searchtype"] == "3" {

				//用户名
				keyword := fmt.Sprintf("%s", where["keyword"])
				Keyword, err := strconv.Atoi(keyword)

				revel.WARN.Println(Keyword)

				if err != nil {
					admin := new(Admin)
					admin_info := admin.GetByRealName(keyword)

					if admin_info.Id > 0 {
						WhereStr += " `aid`='" + strconv.FormatInt(admin_info.Id, 10) + "' AND "
					}
				} else {
					revel.WARN.Println(keyword)
					WhereStr += " `aid`='" + keyword + "' AND "
				}

			} else {
				//ID
				WhereStr += " `id`='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			}
		}
	}

	WhereStr += " 1 "

	//查询总数
	article := new(Article)
	Total, err := DB_Read.Where(WhereStr).Count(article)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Content/list/" + strconv.FormatInt(Cid, 10) + "/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&article_list)

	if len(article_list) > 0 {
		admin := new(Admin)
		category := new(Category)

		for i, v := range article_list {
			article_list[i].Admin = admin.GetById(v.Aid)

			//所属栏目
			article_list[i].Category = category.GetById(v.Cid)
		}
	}

	return article_list, pages, where
}

//获取内容列表
func (a *Article) GetByList(Cid int64, Search string, Page int64, Perpage int64) (article_arr []*Article, html template.HTML, where map[string]interface{}) {
	article_list := []*Article{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(Search) > 0 {

		//解码
		where = utils.DecodeSegment(Search)

		revel.WARN.Println(where)

		if where["cid"] != "0" {
			WhereStr += " `cid`='" + fmt.Sprintf("%s", where["cid"]) + "' AND "
		}

		if where["field"] != "" && where["keyword"] != "" {

			if where["field"] == "title" {
				//标题
				WhereStr += " `title` like %'" + fmt.Sprintf("%s", where["keyword"]) + "'% AND "

			} else if where["field"] == "keywords" {
				//关键词
				WhereStr += " `keywords` like %'" + fmt.Sprintf("%s", where["keyword"]) + "'% AND "
			} else if where["field"] == "description" {
				//描述
				WhereStr += " `description` like %'" + fmt.Sprintf("%s", where["keyword"]) + "'% AND "
			} else if where["field"] == "id" {
				//ID
				WhereStr += " `id`='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			}

		}
	}

	WhereStr += " 1 "

	//查询总数
	article := new(Article)
	Total, err := DB_Read.Where(WhereStr).Count(article)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Content/relationlist/" + strconv.FormatInt(Cid, 10) + "/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&article_list)

	if len(article_list) > 0 {

		admin := new(Admin)
		category := new(Category)

		for i, v := range article_list {
			article_list[i].Admin = admin.GetById(v.Aid)

			//所属栏目
			article_list[i].Category = category.GetById(v.Cid)
		}
	}

	return article_list, pages, where
}

//添加内容
func (a *Article) Save() bool {

	article := new(Article)
	article.Cid = a.Cid
	article.Aid = a.Aid
	article.Title = a.Title
	article.Color = a.Color
	article.Font = a.Font
	article.Thumb = a.Thumb
	article.Content = a.Content
	article.Copyfrom = a.Copyfrom
	article.Keywords = a.Keywords
	article.Description = a.Description
	article.Relation = a.Relation
	article.Pagetype = a.Pagetype
	article.Maxcharperpage = a.Maxcharperpage
	article.Istop = a.Istop
	article.Status = a.Status
	article.Hits = 0
	article.Iscomment = a.Iscomment
	article.Createtime = time.Now().Format("2006-01-02 15:04:04")
	article.Updatetime = "0000-00-00 00:00:00"

	has, err := DB_Write.Insert(article)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑内容
func (a *Article) Edit(Id int64) bool {

	article := new(Article)

	article.Title = a.Title
	article.Color = a.Color
	article.Font = a.Font
	article.Thumb = a.Thumb
	article.Content = a.Content
	article.Copyfrom = a.Copyfrom
	article.Keywords = a.Keywords
	article.Description = a.Description
	article.Relation = a.Relation
	article.Pagetype = a.Pagetype
	article.Maxcharperpage = a.Maxcharperpage
	article.Istop = a.Istop
	article.Status = a.Status
	article.Iscomment = a.Iscomment
	article.Updatetime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Id(Id).Cols("title", "color", "font", "thumb", "content", "copyfrom", "keywords", "description", "relation", "pagetype", "Maxcharperpage", "istop", "status", "iscomment", "updatetime").Update(article)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//批量移动
func (a *Article) Remove(Cid int64, ids string) bool {
	article := new(Article)

	article.Cid = Cid

	has, err := DB_Write.Where("id in (" + ids + ")").Cols("cid").Update(article)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//删除内容
func (a *Article) DelByID(Id int64) bool {
	article := new(Article)

	has, err := DB_Write.Id(Id).Delete(article)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//标题是否存在
func (a *Article) HasTitle() bool {

	article := new(Article)
	has, err := DB_Read.Where("title=?", a.Title).Get(article)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if article.Id > 0 {
		return true
	}
	return false
}
