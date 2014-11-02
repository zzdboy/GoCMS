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

//安全类
import "fmt"
import "time"
import "reflect"
import "strings"
import "strconv"
import "crypto/md5"
import "encoding/base64"
import "encoding/hex"
import "math/rand"
import "bytes"

//返回MD5加密
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//是否Map类型
func IsMap(v interface{}) bool {
	return reflect.ValueOf(&v).Elem().Elem().Kind() == reflect.Map
}

//base64加密
//例如:str := utils.Base64Encode([]byte("Hello, playground"))
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//base64解密
func Base64Decode(src string) string {
	code, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		fmt.Println("Base64解码失败!" + err.Error())
	}
	return string(code)
}

//将搜索条件数组编码成一个字符串
//where 搜索条件Map
//例如:map转换成格式 name:张三|age:20 然后Base64
func EncodeSegment(where map[string]interface{}) string {
	if !IsMap(where) {
		return ""
	}

	whereStr := ""
	for key, val := range where {
		whereStr += key + ":" + fmt.Sprintf("%d", val) + "|"
	}

	whereStr = Base64Encode([]byte(strings.Trim(whereStr, "|")))

	return whereStr
}

//将编码的搜索条件，解码成搜索条件Map数组
//where 搜索条件的编码字符串
func DecodeSegment(where string) map[string]interface{} {
	where_map := make(map[string]interface{})

	where = Base64Decode(where)

	Search := strings.Split(where, "|")

	if len(Search) > 0 {

		for _, val := range Search {
			arr := strings.Split(val, ":")

			if arr[0] == "id" || arr[0] == "cid" && len(arr[1]) > 0 {
				Number, err := strconv.ParseInt(arr[1], 10, 64)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					where_map[arr[0]] = Number
				}
			} else {
				where_map[arr[0]] = arr[1]
			}
		}
	}

	return where_map
}

//生成随机字符串
func RandomString(num int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < num; {
		if string(RandomInt(65, 90)) != temp {
			temp = string(RandomInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

//生成随机数字
func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
