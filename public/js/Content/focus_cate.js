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
 * 焦点图管理
 */

/**
 * 提交检测
 */
function form_submit() {
	var name = $.trim($("#name").val());
	if (name == '') {
		$("#name").addClass("onFocus").focus();
		$("#nameTip").addClass("onError").html("请输入分类名称!");
		return false;
	}else{
		$("#name").removeClass("onFocus");
		$("#nameTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var width = $.trim($("#width").val());

	if (width == '') {
		$("#width").addClass("onFocus").focus();
		$("#widthTip").addClass("onError").html("请输入宽度!");
		return false;
	}else{
		$("#width").removeClass("onFocus");
		$("#widthTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var height = $.trim($("#height").val());

	if (height == '') {
		$("#height").addClass("onFocus").focus();
		$("#heightTip").addClass("onError").html("请输入宽度!");
		return false;
	}else{
		$("#height").removeClass("onFocus");
		$("#heightTip").removeClass("onError").addClass("onCorrect").html("正确");
	}
	
	return true;
}

 /**
 * 删除分类
 */
function delete_cate(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.post("/FocusCate/Delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
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