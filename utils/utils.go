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

//公共方法
import "fmt"
import "net"
import "strings"
import "regexp"
import "encoding/json"
import "github.com/PuerkitoBio/goquery"

//获取服务器IP
func GetServerIP() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

//获取客户端IP
func GetClientIP() string {
	// if ip := c.Request.Header.Get("X-Forwarded-For"); ip != "" {
	// 	ips := strings.Split(ip, ",")
	// 	if len(ips) > 0 && ips[0] != "" {
	// 		rip := strings.Split(ips[0], ":")
	// 		return rip[0]
	// 	}
	// }

	// ip := strings.Split(c.Request.RemoteAddr, ":")
	// if len(ip) > 0 {
	// 	if ip[0] != "[" {
	// 		return ip[0]
	// 	}
	// }
	return "127.0.0.1"
}

//获取Ip地址详细信息
func GetIpAddress(ip string) map[string]interface{} {
	doc, err := goquery.NewDocument("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	if err != nil {
		fmt.Println(err)
	}

	content := doc.Find("body").Text()

	var IpAddress map[string]interface{}

	if err := json.Unmarshal([]byte(content), &IpAddress); err == nil {
		return IpAddress
	}

	return IpAddress
}

// 是否是email
func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[0-9a-zA-Z]{2,3}$`, email)
	return ok
}

//Html过滤
func Html2str(html string) string {
	src := string(html)

	//替换HTML的空白字符为空格
	re := regexp.MustCompile(`\s`) //ns*r
	src = re.ReplaceAllString(src, " ")

	//将HTML标签全转换成小写
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

// 按字节截取字符串 utf-8不乱码
func Substr(str string, length int64) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}
