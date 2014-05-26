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
 * 管理员管理
 */

/**
 * 提交检测
 */
function form_submit() {
	var username = $.trim($("#username").val());
	if (username == '') {
		$("#username").focus();
		notice_tips("请输入用户名!");
		return false;
	}

	var password = $.trim($("#password").val());
	if (password == '') {
		$("#password").focus();
		notice_tips("请输入密码!");
		return false;
	}

	var pwdconfirm = $.trim($("#pwdconfirm").val());
	if (pwdconfirm == '') {
		$("#pwdconfirm").focus();
		notice_tips("请输入确认密码!");
		return false;
	}

	if(password != pwdconfirm) {
		$("#pwdconfirm").focus();
		notice_tips("两次输入密码不一致!");
		return false;
	}

	var email = $.trim($("#email").val());
	if (email == '') {
		$("#email").focus();
		notice_tips("请输入E-mail!");
		return false;
	}

	var realname = $.trim($("#realname").val());
	if (realname == '') {
		$("#realname").focus();
		notice_tips("请输入真实姓名!");
		return false;
	}

	var roleid = $.trim($("#roleid").val());
	if (roleid == '') {
		$("#roleid").focus();
		notice_tips("请选择所属角色!");
		return false;
	}

	var lang = $.trim($("#lang").val());
	if (lang == '') {
		$("#lang").focus();
		notice_tips("请选择语言!");
		return false;
	}

	return true;
}

/**
 * 编辑提交检测
 */
function form_edit_submit() {
	var username = $.trim($("#username").val());
	if (username == '') {
		$("#username").focus();
		notice_tips("请输入用户名!");
		return false;
	}

	var email = $.trim($("#email").val());
	if (email == '') {
		$("#email").focus();
		notice_tips("请输入E-mail!");
		return false;
	}

	var realname = $.trim($("#realname").val());
	if (realname == '') {
		$("#realname").focus();
		notice_tips("请输入真实姓名!");
		return false;
	}

	var roleid = $.trim($("#roleid").val());
	if (roleid == '') {
		$("#roleid").focus();
		notice_tips("请选择所属角色!");
		return false;
	}

	var lang = $.trim($("#lang").val());
	if (lang == '') {
		$("#lang").focus();
		notice_tips("请选择语言!");
		return false;
	}

	return true;
}

/**
 * 删除管理员
 */
function delete_admin(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个管理员吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Admin/delete/",
			data : "id=" + id,
			success : function(tmp) {
				if (tmp.status == 1) {
					window.location.reload();
					notice_tips("删除成功!");
				} else {
					notice_tips(tmp.message);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除管理员操作!");
	});
}