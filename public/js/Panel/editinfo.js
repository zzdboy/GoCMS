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
 * 修改密码
 */

/**
 * 提交检测
 */
function form_submit() {
	var realname = $.trim($("#realname").val());
	if (realname == '') {
		$("#realname").addClass("onFocus").focus();
		$("#realnameTip").addClass("onShow onError").html("请输入真实姓名!");
		return false;
	}else{
		$("#realname").removeClass("onFocus");
		$("#realnameTip").removeClass("onError").addClass("onShow onCorrect").html("正确");
	}

	var email = $.trim($("#email").val());
	if (email == '') {
		$("#email").addClass("onFocus").focus();
		$("#emailTip").addClass("onShow onError").html("请输入真实姓名!");
		return false;
	}else{
		$("#email").removeClass("onFocus");
		$("#emailTip").removeClass("onError").addClass("onShow onCorrect").html("正确");
	}

	var lang = $.trim($("#lang").val());
	if (lang == '') {
		$("#lang").addClass("onFocus").focus();
		$("#langTip").addClass("onShow onError").html("请选择语言!");
		return false;
	}else{
		$("#lang").removeClass("onFocus");
		$("#langTip").removeClass("onError").addClass("onShow onCorrect").html("正确");
	}

	return true;
}