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

//测试
import "time"
import "github.com/revel/revel"

type Test struct {
	Id         int64  `xorm:"pk autoincr"`
	Content    string `xorm:"varchar(255)"`
	Createtime string `xorm:"DateTime"`
}

func (c *Test) Save() bool {

	test := new(Test)
	test.Content = "测试"
	test.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(test)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}
