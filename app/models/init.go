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

//初始化模型
import "os"
import "fmt"
import "strings"
import _ "github.com/go-sql-driver/mysql"
import "github.com/go-xorm/xorm"
import "github.com/revel/revel"
import "github.com/revel/config"

//数据库链接

//读数据
var DB_Read *xorm.Engine

//写数据
var DB_Write *xorm.Engine

func init() {
	revel.OnAppStart(InitDB)
}

//设置数据库
func InitDB() {

	//判断是否是系统的分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	config_file := (revel.BasePath + "/conf/databases.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	c, _ := config.ReadDefault(config_file)

	read_driver, _ := c.String("database", "db.read.driver")
	read_dbname, _ := c.String("database", "db.read.dbname")
	read_user, _ := c.String("database", "db.read.user")
	read_password, _ := c.String("database", "db.read.password")
	read_host, _ := c.String("database", "db.read.host")
	//read_prefix, _ := c.String("database", "db.read.prefix")

	//数据库链接
	var err error
	DB_Read, err = xorm.NewEngine(read_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", read_user, read_password, read_host, read_dbname))
	if err != nil {
		revel.WARN.Printf("DB_Read错误: %v", err)
	}

	write_driver, _ := c.String("database", "db.write.driver")
	write_dbname, _ := c.String("database", "db.write.dbname")
	write_user, _ := c.String("database", "db.write.user")
	write_password, _ := c.String("database", "db.write.password")
	write_host, _ := c.String("database", "db.write.host")
	//write_prefix, _ := c.String("database", "db.write.prefix")

	DB_Write, err = xorm.NewEngine(write_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", write_user, write_password, write_host, write_dbname))
	if err != nil {
		revel.WARN.Printf("DB_Write错误: %v", err)
	}

	//缓存方式是存放到内存中，缓存struct的记录数为1000条
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//DB_Read.SetDefaultCacher(cacher)
	//DB_Write.SetDefaultCacher(cacher)

	//控制台打印SQL语句
	//DB_Read.ShowSQL = true
	//DB_Write.ShowSQL = true

	//控制台打印调试信息
	//DB_Read.ShowDebug = true
	//DB_Write.ShowDebug = true

	//控制台打印错误信息
	//DB_Read.ShowErr = true
	//DB_Write.ShowErr = true

	//控制台打印警告信息
	//DB_Read.ShowWarn = true
	//DB_Read.ShowWarn = true
}
