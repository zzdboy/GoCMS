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

import "testing"

func Test_beemap(t *testing.T) {
	bm := NewBeeMap()
	if !bm.Set("astaxie", 1) {
		t.Error("set Error")
	}
	if !bm.Check("astaxie") {
		t.Error("check err")
	}

	if v := bm.Get("astaxie"); v.(int) != 1 {
		t.Error("get err")
	}

	bm.Delete("astaxie")
	if bm.Check("astaxie") {
		t.Error("delete err")
	}
}
