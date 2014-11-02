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
 * 用户管理
 */

//添加会员
function add() {
	art.dialog.open('/User/add/', {
		id : 'user_add',
		title : '添加用户',
		width : 700,
		height : 500,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var username = iframe.$('#username').val();
			if (username == '') {
				iframe.$('#usernameTip').removeClass("onShow").addClass("onError").html('请输入用户名!');
				return false;
			}else{
				iframe.$('#usernameTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "username=" + username;
				par.push(pars);
			}

			var password = iframe.$('#password').val();
			if (password == '') {
				iframe.$('#passwordTip').removeClass("onShow").addClass("onError").html('请输入密码!');
				return false;
			}else{
				iframe.$('#passwordTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "password=" + password;
				par.push(pars);
			}

			var pwdconfirm = iframe.$('#pwdconfirm').val();
			if (pwdconfirm == '') {
				iframe.$('#pwdconfirmTip').removeClass("onShow").addClass("onError").html('请输入确认密码!');
				return false;
			}else{
				if(password != pwdconfirm) {
					iframe.$('#pwdconfirmTip').removeClass("onShow").addClass("onError").html('两次输入密码不一致!');
					return false;
				}
				iframe.$('#pwdconfirmTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "pwdconfirm=" + pwdconfirm;
				par.push(pars);
			}

			var nickname = iframe.$('#nickname').val();
			if (nickname == '') {
				iframe.$('#nicknameTip').removeClass("onShow").addClass("onError").html('请输入昵称!');
				return false;
			}else{
				iframe.$('#nicknameTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "nickname=" + nickname;
				par.push(pars);
			}

			var email = iframe.$('#email').val();
			if (email == '') {
				iframe.$('#emailTip').removeClass("onShow").addClass("onError").html('请输入邮箱!');
				return false;
			}else{
				iframe.$('#emailTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "email=" + email;
				par.push(pars);
			}

			var mobile = iframe.$('#mobile').val();
			if (mobile == '') {
				iframe.$('#mobileTip').removeClass("onShow").addClass("onError").html('请输入手机号码!');
				return false;
			}else{
				iframe.$('#mobileTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "mobile=" + mobile;
				par.push(pars);
			}

			var groupid = iframe.$('#groupid').val();
			if (groupid == '') {
				iframe.$('#groupidTip').removeClass("onShow").addClass("onError").html('请选择会员组!');
				return false;
			}else{
				iframe.$('#groupidTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "groupid=" + groupid;
				par.push(pars);
			}

			var islock = iframe.$('#islock').val();
			if (islock == '') {
				iframe.$('#islockTip').removeClass("onShow").addClass("onError").html('请选择是否定锁!');
				return false;
			}else{
				iframe.$('#islockTip').removeClass("onError").addClass("onCorrect").html('选择正确');
				var pars = "islock=" + islock;
				par.push(pars);
			}

			var point = iframe.$('#point').val();
			if (point == '') {
				iframe.$('#pointtip').removeClass("onShow").addClass("onError").html('请输入积分点数!');
				return false;
			}else{
				iframe.$('#pointtip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "point=" + point;
				par.push(pars);
			}

			if (iframe.$("#vip").is(":checked")) {

				var pars = "vip=1";
				par.push(pars);

				var overduedate = iframe.$('#overduedate').val();

				if (overduedate == "") {
					iframe.$('#overduedateTip').removeClass("onShow").addClass("onError").html('请选择过期时间!');
					return false;
				}else{
					var pars = "overduedate=" + overduedate;
					par.push(pars);
				}				
			}else{
				var pars = "vip=0";
				par.push(pars);
			}

			var birthday = iframe.$('#birthday').val();
			var pars = "birthday=" + birthday;
			par.push(pars);

			var pars = "csrf_token=" + csrf_token;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/User/add/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						window.location.reload();
						notice_tips("添加用户成功!");
					} else {
						notice_tips(tmp.message);
					}
				}
			});
		},
		cancel : true
	});
}

/**
 * 编辑用户
 */
function edit(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/User/edit/' + uid + '/', {
		id : 'user_edit',
		title : '编辑用户',
		width : 700,
		height : 500,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var pars = "id=" + uid;
			par.push(pars);

			var password = iframe.$('#password').val();
			if (password == '') {
				
			}else{
				iframe.$('#passwordTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "password=" + password;
				par.push(pars);
			}

			var pwdconfirm = iframe.$('#pwdconfirm').val();
			if (pwdconfirm == '') {
				
			}else{
				if(password != pwdconfirm) {
					iframe.$('#pwdconfirmTip').removeClass("onShow").addClass("onError").html('两次输入密码不一致!');
					return false;
				}
				iframe.$('#pwdconfirmTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "pwdconfirm=" + pwdconfirm;
				par.push(pars);
			}

			var nickname = iframe.$('#nickname').val();
			if (nickname == '') {
				iframe.$('#nicknameTip').removeClass("onShow").addClass("onError").html('请输入昵称!');
				return false;
			}else{
				iframe.$('#nicknameTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "nickname=" + nickname;
				par.push(pars);
			}

			var email = iframe.$('#email').val();
			if (email == '') {
				iframe.$('#emailTip').removeClass("onShow").addClass("onError").html('请输入邮箱!');
				return false;
			}else{
				iframe.$('#emailTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "email=" + email;
				par.push(pars);
			}

			var mobile = iframe.$('#mobile').val();
			if (mobile == '') {
				iframe.$('#mobileTip').removeClass("onShow").addClass("onError").html('请输入邮箱!');
				return false;
			}else{
				iframe.$('#mobileTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "mobile=" + mobile;
				par.push(pars);
			}

			var groupid = iframe.$('#groupid').val();
			if (groupid == '') {
				iframe.$('#groupidTip').removeClass("onShow").addClass("onError").html('请选择会员组!');
				return false;
			}else{
				iframe.$('#groupidTip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "groupid=" + groupid;
				par.push(pars);
			}

			var islock = iframe.$('#islock').val();
			if (islock == '') {
				iframe.$('#islockTip').removeClass("onShow").addClass("onError").html('请选择是否定锁!');
				return false;
			}else{
				iframe.$('#islockTip').removeClass("onError").addClass("onCorrect").html('选择正确');
				var pars = "islock=" + islock;
				par.push(pars);
			}

			var point = iframe.$('#point').val();
			if (point == '') {
				iframe.$('#pointtip').removeClass("onShow").addClass("onError").html('请输入积分点数!');
				return false;
			}else{
				iframe.$('#pointtip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "point=" + point;
				par.push(pars);
			}

			if (iframe.$("#vip").is(":checked")) {

				var pars = "vip=1";
				par.push(pars);

				var overduedate = iframe.$('#overduedate').val();

				if (overduedate == "") {
					iframe.$('#overduedateTip').removeClass("onShow").addClass("onError").html('请选择过期时间!');
					return false;
				}else{
					var pars = "overduedate=" + overduedate;
					par.push(pars);
				}				
			}else{
				var pars = "vip=0";
				par.push(pars);
			}

			var birthday = iframe.$('#birthday').val();
			var pars = "birthday=" + birthday;
			par.push(pars);

			var pars = "csrf_token=" + csrf_token;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/User/edit/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						window.location.reload();
						notice_tips("编辑用户成功!");
					} else {
						notice_tips(tmp.message);
					}
				}
			});
		},
		cancel : true
	});
}

/**
 * 用户详细
 */
function user_info(uid) {
	if (uid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/User/userinfo/' + uid + '/', {
		id : 'user_info',
		title : '个人信息',
		width : 700,
		height : 500,
		lock : true,
		ok : function() {},
		cancel : true
	});
}

//锁定
function lock() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	art.dialog.confirm('确定要锁定用户吗?', function() {
		$.post("/User/lock/",{'ids':ids,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("锁定成功!");
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("取消了锁定用户操作!");
	});
}

//解锁
function unlock() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	art.dialog.confirm('确定要解锁用户吗?', function() {
		$.post("/User/unlock/",{'ids':ids,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("解锁成功!");
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("取消了解锁用户操作!");
	});
}

function show_move() {

	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	if ($("#show_move").css("display") == "none") {
		$("#show_move").fadeIn("slow");
	}else{
		$("#show_move").fadeOut("slow");
	}
}

//移动
function move() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	var groupid = $('#groupid').val();

	art.dialog.confirm('确定要移动用户吗?', function() {
		$.post("/User/move/",{'ids':ids,'groupid':groupid,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("移动成功!");
			} else {
				notice_tips(tmp.message);
			}
		});
	}, function() {
		notice_tips("取消了移动用户操作!");
	});
}

/**
 * 删除用户
 */
function confirm_delete() {

	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	art.dialog.confirm('确定要删除用户吗?', function() {
		$.post("/User/delete/",{'ids':ids,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("删除成功!");
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("取消了删除用户操作!");
	});
}

//搜索
function Search() {
	url = '/User/';

	search = "";

	search += 'islock:' + $('#islock').val() + '|';
	search += 'type:' + $('#type').val() + '|';
	search += 'keyword:' + $('#keyword').val() + '|';
	search += 'start_time:' + $('#start_time').val() + '|';
	search += 'end_time:' + $('#end_time').val() + '';

	$.base64.encode(search)

	url += $.base64.encode(search)+'/1/';

	redirect(url);
};

//重置
function Reset() {
	redirect("/User/");
}