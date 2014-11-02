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
 * 会员组管理
 */


//添加会员组
function add() {
	art.dialog.open('/Group/add/', {
		id : 'group_add',
		title : '添加用户组',
		width : 880,
		height : 500,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var name = iframe.$('#name').val();
			if (name == '') {
				iframe.$('#name').addClass("onFocus").focus();
				iframe.$('#nametip').addClass("onError").html('请输入会员组名称!');
				return false;
			}else{
				iframe.$('#name').removeClass("onFocus");
				iframe.$('#nametip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "name=" + name;
				par.push(pars);
			}

			var point = iframe.$('#point').val();
			if (point == '') {
				iframe.$('#point').addClass("onFocus").focus();
				iframe.$('#pointtip').addClass("onError").html('请输入积分!');
				return false;
			}else{
				iframe.$('#point').removeClass("onFocus");
				iframe.$('#pointtip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "point=" + point;
				par.push(pars);
			}

			var star = iframe.$('#star').val();
			if (star == '') {
				iframe.$('#star').addClass("onFocus").focus();
				iframe.$('#startip').addClass("onError").html('请输入星星数!');
				return false;
			}else{
				iframe.$('#star').removeClass("onFocus");
				iframe.$('#startip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "star=" + star;
				par.push(pars);
			}

			if (iframe.$('#allowpost').is(":checked")) {
				var pars = "allowpost=1";
				par.push(pars);
			}else{
				var pars = "allowpost=0";
				par.push(pars);
			}

			if (iframe.$('#allowpostverify').is(":checked")) {
				var pars = "allowpostverify=1";
				par.push(pars);
			}else{
				var pars = "allowpostverify=0";
				par.push(pars);
			}

			if (iframe.$('#allowupgrade').is(":checked")) {
				var pars = "allowupgrade=1";
				par.push(pars);
			}else{
				var pars = "allowupgrade=0";
				par.push(pars);
			}

			if (iframe.$('#allowsendmessage').is(":checked")) {
				var pars = "allowsendmessage=1";
				par.push(pars);
			}else{
				var pars = "allowsendmessage=0";
				par.push(pars);
			}

			if (iframe.$('#allowattachment').is(":checked")) {
				var pars = "allowattachment=1";
				par.push(pars);
			}else{
				var pars = "allowattachment=0";
				par.push(pars);
			}

			if (iframe.$('#allowsearch').is(":checked")) {
				var pars = "allowsearch=1";
				par.push(pars);
			}else{
				var pars = "allowsearch=0";
				par.push(pars);
			}

			var priceday = iframe.$('#priceday').val();
			if (priceday == '') {
				var pars = "priceday=0";
				par.push(pars);
			}else{
				var pars = "priceday=" + priceday;
				par.push(pars);
			}

			var pricemonth = iframe.$('#pricemonth').val();
			if (pricemonth == '') {
				var pars = "pricemonth=0";
				par.push(pars);
			}else{
				var pars = "pricemonth=" + pricemonth;
				par.push(pars);
			}

			var priceyear = iframe.$('#priceyear').val();
			if (priceyear == '') {
				var pars = "priceyear=0";
				par.push(pars);
			}else{
				var pars = "priceyear=" + priceyear;
				par.push(pars);
			}

			var allowmessage = iframe.$('#allowmessage').val();
			if (allowmessage == '') {
				iframe.$('#allowmessage').addClass("onFocus").focus();
				iframe.$('#allowmessagetip').addClass("onError").html('请输入最大短消息数!');
				return false;
			}else{
				iframe.$('#allowmessage').removeClass("onFocus");
				var pars = "allowmessage=" + allowmessage;
				par.push(pars);
			}

			var allowpostnum = iframe.$('#allowpostnum').val();
			if (allowpostnum == '') {
				iframe.$('#allowpostnum').addClass("onFocus").focus();
				iframe.$('#allowpostnumtip').addClass("onError").html('请输入日最大投稿数!');
				return false;
			}else{
				iframe.$('#allowpostnum').removeClass("onFocus");
				var pars = "allowpostnum=" + allowpostnum;
				par.push(pars);
			}

			var usernamecolor = iframe.$('#usernamecolor').val();
			var pars = "usernamecolor=" + usernamecolor;
			par.push(pars);

			var icon = iframe.$('#icon').val();
			var pars = "icon=" + icon;
			par.push(pars);

			var desc = iframe.$('#desc').val();
			var pars = "desc=" + desc;
			par.push(pars);

			var status = iframe.$("input[name='status']:checked").val();
			var pars = "status=" + status;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Group/add/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						right_refresh();
						notice_tips("添加成功!");
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
 * 编辑用户组
 */
function edit(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Group/edit/' + id + '/', {
		id : 'group_edit',
		title : '编辑用户',
		width : 880,
		height : 500,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var pars = "id=" + id;
			par.push(pars);

			var name = iframe.$('#name').val();
			if (name == '') {
				iframe.$('#nametip').removeClass("onShow").addClass("onError").html('请输入会员组名称!');
				return false;
			}else{
				iframe.$('#nametip').removeClass("onError").addClass("onCorrect").html('正确');
				var pars = "name=" + name;
				par.push(pars);
			}

			var point = iframe.$('#point').val();
			if (point == '') {
				iframe.$('#pointtip').removeClass("onShow").addClass("onError").html('请输入积分!');
				return false;
			}else{
				var pars = "point=" + point;
				par.push(pars);
			}

			var star = iframe.$('#star').val();
			if (star == '') {
				iframe.$('#startip').removeClass("onShow").addClass("onError").html('请输入星星数!');
				return false;
			}else{
				var pars = "star=" + star;
				par.push(pars);
			}

			if (iframe.$('#allowpost').is(":checked")) {
				var pars = "allowpost=1";
				par.push(pars);
			}else{
				var pars = "allowpost=0";
				par.push(pars);
			}

			if (iframe.$('#allowpostverify').is(":checked")) {
				var pars = "allowpostverify=1";
				par.push(pars);
			}else{
				var pars = "allowpostverify=0";
				par.push(pars);
			}

			if (iframe.$('#allowupgrade').is(":checked")) {
				var pars = "allowupgrade=1";
				par.push(pars);
			}else{
				var pars = "allowupgrade=0";
				par.push(pars);
			}

			if (iframe.$('#allowsendmessage').is(":checked")) {
				var pars = "allowsendmessage=1";
				par.push(pars);
			}else{
				var pars = "allowsendmessage=0";
				par.push(pars);
			}

			if (iframe.$('#allowattachment').is(":checked")) {
				var pars = "allowattachment=1";
				par.push(pars);
			}else{
				var pars = "allowattachment=0";
				par.push(pars);
			}

			if (iframe.$('#allowsearch').is(":checked")) {
				var pars = "allowsearch=1";
				par.push(pars);
			}else{
				var pars = "allowsearch=0";
				par.push(pars);
			}

			var priceday = iframe.$('#priceday').val();
			if (priceday == '') {
				var pars = "priceday=0";
				par.push(pars);
			}else{
				var pars = "priceday=" + priceday;
				par.push(pars);
			}

			var pricemonth = iframe.$('#pricemonth').val();
			if (pricemonth == '') {
				var pars = "pricemonth=0";
				par.push(pars);
			}else{
				var pars = "pricemonth=" + pricemonth;
				par.push(pars);
			}

			var priceyear = iframe.$('#priceyear').val();
			if (priceyear == '') {
				var pars = "priceyear=0";
				par.push(pars);
			}else{
				var pars = "priceyear=" + priceyear;
				par.push(pars);
			}

			var allowmessage = iframe.$('#allowmessage').val();
			if (allowmessage == '') {
				iframe.$('#allowmessagetip').removeClass("onShow").addClass("onError").html('请输入最大短消息数!');
				return false;
			}else{
				var pars = "allowmessage=" + allowmessage;
				par.push(pars);
			}

			var allowpostnum = iframe.$('#allowpostnum').val();
			if (allowpostnum == '') {
				iframe.$('#allowpostnumip').removeClass("onShow").addClass("onError").html('请输入日最大投稿数!');
				return false;
			}else{
				var pars = "allowpostnum=" + allowpostnum;
				par.push(pars);
			}

			var usernamecolor = iframe.$('#usernamecolor').val();
			var pars = "usernamecolor=" + usernamecolor;
			par.push(pars);

			var icon = iframe.$('#icon').val();
			var pars = "icon=" + icon;
			par.push(pars);

			var desc = iframe.$('#desc').val();
			var pars = "desc=" + desc;
			par.push(pars);

			var status = iframe.$("input[name='status']:checked").val();
			var pars = "status=" + status;
			par.push(pars);

			var pars = "csrf_token=" + csrf_token;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Group/edit/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						right_refresh();
						notice_tips("编辑成功!");
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
 * 删除用户组
 */
function del(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个用户吗?', function() {
		$.post("/Group/delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				right_refresh();
				notice_tips("删除成功!");
			} else {
				notice_tips(tmp.content);
			}
		});

	}, function() {
		notice_tips("你取消了删除用户操作!");
	});
}

function set_username_color(color) {
	$('#usernamecolor').val(color);
}