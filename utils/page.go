// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

package utils

//分页
import "html/template"
import "strconv"
import "math"

/*
 * Perpage 每页显示的条目数
 * nums 总条目数
 * current_num 当前被选中的页
 * sub_pages 每次显示的页数
 * subPage_link 每个分页的链接
 * subPage_type 显示分页的类型
 * subPage_type=1的时候为普通分页模式 共4523条记录,每页显示10条,当前第1/453页 [首页] [上页] [下页] [尾页]
 * subPage_type=2的时候为经典分页样式 当前第1/453页 [首页] [上页] 1 2 3 4 5 6 7 8 9 10 [下页] [尾页]
 */

type Page struct {
	Perpage      int64       //每页显示的条目数
	Nums         int64       //总条目数
	Current_page int64       //当前被选中的页
	Sub_pages    int64       //每次显示的页数
	PageNums     int64       //总页数
	Page_array   map[int]int //用来构造分页的数组
	SubPage_link string      //每个分页的链接
	SubPage_type int8        //显示分页的类型
}

/*
 * 显示分页
 */
func (P *Page) Show() template.HTML {
	if P.Perpage <= 0 {
		P.Perpage = 2
	}

	if P.Nums <= 0 {
		P.Nums = 0
	}

	if P.Current_page <= 0 {
		P.Current_page = 1
	}

	if P.Sub_pages <= 0 {
		P.Sub_pages = 5
	}

	P.PageNums = int64(math.Ceil(float64(P.Nums) / float64(P.Perpage)))
	if P.PageNums <= 0 {
		P.PageNums = 1
	}

	if P.SubPage_type <= 0 {
		P.SubPage_type = 1
	}

	if P.SubPage_type == 1 {
		return P.subPageCss1()
	} else if P.SubPage_type == 2 {
		return P.subPageCss2()
	} else {
		return P.subPageCss1()
	}
}

//构造普通模式的分页 共4523条记录,每页显示10条,当前第1/453页 [首页] [上页] [下页] [尾页]
func (P *Page) subPageCss1() template.HTML {
	Html := ""

	Html += "共" + strconv.FormatInt(P.Nums, 10) + "条记录，"
	Html += "每页显示" + strconv.FormatInt(P.Perpage, 10) + "条，"
	Html += "当前第" + strconv.FormatInt(P.Current_page, 10) + "/" + strconv.FormatInt(P.PageNums, 10) + "页 "

	if P.Current_page > 1 {
		firstPageUrl := P.SubPage_link + "1" + "/"
		prewPageUrl := P.SubPage_link + strconv.FormatInt((P.Current_page-1), 10) + "/"
		Html += "&nbsp;<a href='" + firstPageUrl + "'>首页</a>"
		Html += "&nbsp;<a href='" + prewPageUrl + "'>上一页</a>"
	} else {
		Html += "&nbsp;<a href='javascript:;'>首页</a>"
		Html += "&nbsp;<a href='javascript:;'>上一页</a>"
	}

	if P.Current_page < P.PageNums {
		lastPageUrl := P.SubPage_link + strconv.FormatInt(P.PageNums, 10) + "/"
		nextPageUrl := P.SubPage_link + strconv.FormatInt(P.Current_page+1, 10) + "/"
		Html += "&nbsp;<a href='" + nextPageUrl + "'>下一页</a>"
		Html += "&nbsp;<a href='" + lastPageUrl + "'>尾页</a>"
	} else {
		Html += "&nbsp;<a href='javascript:;'>下一页</a>"
		Html += "&nbsp;<a href='javascript:;'>尾页</a>"
	}

	return template.HTML(Html)
}

//构造经典模式的分页 当前第1/453页 [首页] [上页] 1 2 3 4 5 6 7 8 9 10 [下页] [尾页]
func (P *Page) subPageCss2() template.HTML {
	Html := ""

	Html += "总共" + strconv.FormatInt(P.Nums, 10) + "条&nbsp;&nbsp;"
	Html += "当前第" + strconv.FormatInt(P.Current_page, 10) + "/" + strconv.FormatInt(P.PageNums, 10) + "页 "

	if P.Current_page > 1 {
		firstPageUrl := P.SubPage_link + "1" + "/"
		prewPageUrl := P.SubPage_link + strconv.FormatInt((P.Current_page-1), 10) + "/"
		Html += "&nbsp;<a href='" + firstPageUrl + "'>首页</a>"
		Html += "&nbsp;<a href='" + prewPageUrl + "'>上一页</a>"
	} else {
		Html += "&nbsp;<a href='javascript:;'>首页</a>"
		Html += "&nbsp;<a href='javascript:;'>上一页</a>"
	}

	if P.PageNums <= 5 {
		var i int64 = 1
		for ; i <= P.PageNums; i++ {
			if i == P.Current_page {
				Html += " &nbsp;<span class='current'>" + strconv.FormatInt(i, 10) + "</span>"
			} else {
				url := P.SubPage_link + strconv.FormatInt(i, 10) + "/"
				Html += "&nbsp;<a href='" + url + "'>" + strconv.FormatInt(i, 10) + "</a>"
			}
		}
	} else {
		var i int64 = 1
		for ; i < P.Sub_pages; i++ {
			i := P.Current_page - 1 + i

			if i > P.PageNums {
				break
			}

			if i == P.Current_page {
				Html += " &nbsp;<span class='current'>" + strconv.FormatInt(i, 10) + "</span>"
			} else {
				url := P.SubPage_link + strconv.FormatInt(i, 10) + "/"
				Html += " &nbsp;<a href='" + url + "'>" + strconv.FormatInt(i, 10) + "</a>"
			}
		}
	}

	if P.Current_page < P.PageNums {
		lastPageUrl := P.SubPage_link + strconv.FormatInt(P.PageNums, 10) + "/"
		nextPageUrl := P.SubPage_link + strconv.FormatInt(P.Current_page+1, 10) + "/"
		Html += "&nbsp;<a href='" + nextPageUrl + "'>下一页</a>"
		Html += "&nbsp;<a href='" + lastPageUrl + "'>尾页</a>"
	} else {
		Html += "&nbsp;<a href='javascript:;'>下一页</a>"
		Html += "&nbsp;<a href='javascript:;'>尾页</a>"
	}

	return template.HTML(Html)
}
