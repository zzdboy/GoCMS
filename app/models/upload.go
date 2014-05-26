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

//上传文件
import "github.com/revel/revel"

type Upload struct {
}

//编辑器文件上传
func (u Upload) EditorUpload() string {
	revel.WARN.Println("编辑器上传文件")
	return "1.jpg"
}
