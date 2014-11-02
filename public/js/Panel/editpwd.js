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
	var old_password = $.trim($("#old_password").val());
	if (old_password == '') {
		$("#old_password").addClass("onFocus").focus();
		$("#old_passwordTip").addClass("onError").html("请输入旧密码!");
		return false;
	}else{
		$("#old_password").removeClass("onFocus");
		$("#old_passwordTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var new_password = $.trim($("#new_password").val());
	if (new_password == '') {
		$("#new_password").addClass("onFocus").focus();
		$("#new_passwordTip").addClass("onError").html("请输入新密码!");
		return false;
	}else{
		$("#new_password").removeClass("onFocus");
		$("#new_passwordTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var new_pwdconfirm = $.trim($("#new_pwdconfirm").val());
	if (new_pwdconfirm == '') {
		$("#new_pwdconfirm").addClass("onFocus").focus();
		$("#new_pwdconfirmTip").addClass("onError").html("重复新密码不能为空!");
		return false;
	}else{
		$("#new_pwdconfirm").removeClass("onFocus");
		$("#new_pwdconfirmTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	if (new_password != new_pwdconfirm) {
		$("#new_pwdconfirm").addClass("onFocus").focus();
		$("#new_pwdconfirmTip").removeClass("onCorrect").addClass("onError").html("两次输入密码不一致!");
		return false;
	}else{
		$("#new_pwdconfirm").removeClass("onFocus");
		$("#new_pwdconfirmTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	return true;
}