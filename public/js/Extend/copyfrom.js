// +----------------------------------------------------------------------
// | GoCMS 0.1
// +----------------------------------------------------------------------
// | Copyright (c) 2013-2014 http://www.6574.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: zzdboy <zzdboy1616@163.com>
// +----------------------------------------------------------------------

/**
 * 来源管理
 */

 /**
 * 提交检测
 */
function form_submit() {
	var sitename = $.trim($("#sitename").val());
	if (sitename == '') {
		$("#sitename").focus();
		notice_tips("请输入来源名称!");
		return false;
	}

	var siteurl = $.trim($("#siteurl").val());
	if (siteurl == '') {
		$("#siteurl").focus();
		notice_tips("请输入来源链接!");
		return false;
	}

	return true;
}