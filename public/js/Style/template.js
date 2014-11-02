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
 * 模板风格
 */

$(document).ready(function() {
	$("input[type='radio'][name='type']").click(function(){
		if ($(this).val()==1) {
			$('#upfile').show();
			$('#code').hide();
		} else{
			$('#code').show();
			$('#upfile').hide();
		}	
	})
})

//添加风格
function Import() {
	art.dialog.open('/Style/import/', {
		id : 'import',
		title : '添加风格',
		width : 500,
		height : 200,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var identity = iframe.$('#identity').val();

			if (identity == '') {
				iframe.$('#identitytip').removeClass("onShow").addClass("onError").html('风格标识应该为2-8位英文!');
				return false;
			}else{
				iframe.$('#identitytip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "identity=" + identity;
				par.push(pars);
			}

			var name = iframe.$('#name').val();

			if (name == '') {
				iframe.$('#nametip').removeClass("onShow").addClass("onError").html('风格中文名应该为2-8位!');
				return false;
			}else{
				iframe.$('#nametip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "name=" + name;
				par.push(pars);
			}

			var author = iframe.$('#author').val();

			if (author == '') {
				iframe.$('#authortip').removeClass("onShow").addClass("onError").html('风格作者不能为空!');
				return false;
			}else{
				iframe.$('#authortip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "author=" + author;
				par.push(pars);
			}

			var version = iframe.$('#version').val();

			if (version == '') {
				iframe.$('#versiontip').removeClass("onShow").addClass("onError").html('风格版本不能为空!');
				return false;
			}else{
				iframe.$('#versiontip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "version=" + version;
				par.push(pars);
			}

			var status = iframe.$("input[name='status']:checked").val();
			var pars = "status=" + status;
			par.push(pars);

			var pars = "csrf_token=" + csrf_token;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Style/import/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						window.location.reload();
						notice_tips("添加风格成功!");
					} else {
						notice_tips(tmp.message);
					}
				}
			});
		},
		cancel : true
	});
}

//编辑风格
function Edit(id) {

	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Style/edit/' + id + '/', {
		id : 'edit',
		title : '编辑风格',
		width : 500,
		height : 200,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var par = [];

			var pars = "id=" + id;
			par.push(pars);

			var identity = iframe.$('#identity').val();

			if (identity == '') {
				iframe.$('#identitytip').removeClass("onShow").addClass("onError").html('风格标识应该为2-8位英文!');
				return false;
			}else{
				iframe.$('#identitytip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "identity=" + identity;
				par.push(pars);
			}

			var name = iframe.$('#name').val();

			if (name == '') {
				iframe.$('#nametip').removeClass("onShow").addClass("onError").html('风格中文名应该为2-8位!');
				return false;
			}else{
				iframe.$('#nametip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "name=" + name;
				par.push(pars);
			}

			var author = iframe.$('#author').val();

			if (author == '') {
				iframe.$('#authortip').removeClass("onShow").addClass("onError").html('风格作者不能为空!');
				return false;
			}else{
				iframe.$('#authortip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "author=" + author;
				par.push(pars);
			}

			var version = iframe.$('#version').val();

			if (version == '') {
				iframe.$('#versiontip').removeClass("onShow").addClass("onError").html('风格版本不能为空!');
				return false;
			}else{
				iframe.$('#versiontip').removeClass("onError").addClass("onCorrect").html('输入正确');
				var pars = "version=" + version;
				par.push(pars);
			}

			var status = iframe.$("input[name='status']:checked").val();
			var pars = "status=" + status;
			par.push(pars);

			var pars = "csrf_token=" + csrf_token;
			par.push(pars);

			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Style/edit/",
				data : pars,
				success : function(tmp) {
					if (tmp.status == 1) {
						window.location.reload();
						notice_tips("编辑风格成功!");
					} else {
						notice_tips(tmp.message);
					}
				}
			});
		},
		cancel : true
	});
}

//设置状态
function Setstatus(id,status) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	if (status == '') {
		notice_tips("参数错误!");
		return false;
	}

	var par = [];

	var pars = "id=" + id;
	par.push(pars);

	var pars = "status=" + status;
	par.push(pars);

	pars = par.join("&");

	$.ajax({
		type : "POST",
		url : "/Style/setstatus/",
		data : pars,
		success : function(tmp) {
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("设置成功!");
			} else {
				notice_tips(tmp.message);
			}
		}
	});
}