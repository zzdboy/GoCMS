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

//文件类
import "fmt"
import "os"

//文件大小
func FileSize(size int) string {
	s := float32(size)
	if s > 1024*1024 {
		return fmt.Sprintf("%.1f M", s/(1024*1024))
	}
	if s > 1024 {
		return fmt.Sprintf("%.1f K", s/1024)
	}
	return fmt.Sprintf("%f B", s)
}

//是否文件
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}

//是否目录
func IsDir(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return f.IsDir()
}
