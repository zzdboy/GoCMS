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
		$("#sitename").addClass("onFocus").focus();
		notice_tips("请输入来源名称!");
		return false;
	}else{
		$("#sitename").removeClass("onFocus");
	}

	var siteurl = $.trim($("#siteurl").val());
	if (siteurl == '') {
		$("#siteurl").addClass("onFocus").focus();
		notice_tips("请输入来源链接!");
		return false;
	}else{
		$("#siteurl").removeClass("onFocus");
	}

	return true;
}

function delete_copyfrom(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.post("/Copyfrom/Delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				window.location.reload();
				notice_tips("删除成功!");
			} else {
				notice_tips(tmp.message);
			}
		});
	}, function() {
		notice_tips("你取消了删除分类操作!");
	});
}