#GoCMS 

基于Go语言和Revel框架的内容管理系统


演示地址：
[http://admin.6574.com.cn/](http://admin.6574.com.cn/)

GoCMS QQ交流群：
[<a target="_blank" href="http://shang.qq.com/wpa/qunwpa?idkey=3421374909556d550942819ac01a48339fc70130ebfea330015dee89abb540c2"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="Revel&nbsp;框架交流" title="Revel&nbsp;框架交流"></a>](345304040)

##编译安装说明：

设置GOPATH(安装目录)

	$ export GOPATH=/path/src/admin
	$ cd /path/src/admin

注：把下载的代码复制到src/admin目录下

获取源代码，下载完成后会自动编译为GoCMS可执行文件
	
	$ go get github.com/go-xorm/xorm
	$ go get github.com/revel/revel
	$ go get github.com/cbonello/revel-csrf
	$ go get github.com/PuerkitoBio/goquery
	$ go get github.com/zzdboy/GoCMS

修改数据库配置
	
	admin/conf/databases.conf

导入MySQL

	doc目录下gocms.sql

运行
	
	$ nohup revel run admin &
	设为后台运行

访问： 

http://localhost:9001

后台地址：

http://localhost:9000

帐号：admin
密码：123456

