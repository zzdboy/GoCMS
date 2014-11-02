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

//计划任务
import "strings"
import "github.com/revel/revel"
import "github.com/PuerkitoBio/goquery"

type Task struct {
}

//抓取cnBeta
func (c *Task) GrabCnbeta() {
	doc, err := goquery.NewDocument("http://www.cnbeta.com/")
	if err != nil {
		revel.WARN.Println(err)
	}

	doc.Find(".all_news_wildlist .alllist .items_area dl").Each(func(i int, s *goquery.Selection) {

		revel.WARN.Println(i)

		title := strings.Trim(s.Find("dt a").Text(), "")
		url, _ := s.Find("dt a").Attr("href")
		thumb_img, _ := s.Find("dd .pic img").Attr("src")
		description := strings.Trim(s.Find("dd .newsinfo p").Text(), "")

		doc, err := goquery.NewDocument("http://www.cnbeta.com" + url)
		if err != nil {
			revel.WARN.Println(err)
		}
		content, _ := doc.Find(".content").Html()
		content = strings.Trim(content, "")

		//插入数据库
		article := new(Article)

		article.Cid = 11
		article.Aid = 0
		article.Title = title
		article.Thumb = thumb_img
		article.Description = description
		article.Content = content

		if !article.HasTitle() {
			article.Save()
		}
	})
}
