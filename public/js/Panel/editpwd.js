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
		$("#old_password").focus();
		notice_tips("请输入旧密码!");
		return false;
	}

	var new_password = $.trim($("#new_password").val());
	if (new_password == '') {
		$("#new_password").focus();
		notice_tips("请输入新密码!");
		return false;
	}

	var new_pwdconfirm = $.trim($("#new_pwdconfirm").val());
	if (new_pwdconfirm == '') {
		$("#new_pwdconfirm").focus();
		notice_tips("重复新密码不能为空!");
		return false;
	}

	return true;
}