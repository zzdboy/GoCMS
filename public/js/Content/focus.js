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
	var cid = $.trim($("#cid").val());
	if (cid == '') {
		$("#cid").addClass("onFocus").focus();
		$("#cidTip").addClass("onError").html("请选择所属分类!");
		return false;
	}else{
		$("#cid").removeClass("onFocus");
		$("#cidTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	if (title == '') {
		$("#title").addClass("onFocus").focus();
		$("#titleTip").addClass("onError").html("请输入标题!");
		return false;
	}else{
		$("#title").removeClass("onFocus");
		$("#titleTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	if (url == '') {
		$("#url").addClass("onFocus").focus();
		$("#urlTip").addClass("onError").html("请输入地址!");
		return false;
	}else{
		$("#url").removeClass("onFocus");
		$("#urlTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	if (img == '') {
		$("#img").addClass("onFocus").focus();
		$("#imgTip").addClass("onError").html("请输入图片!");
		return false;
	}else{
		$("#img").removeClass("onFocus");
		$("#imgTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	if (content == '') {
		$("#content").addClass("onFocus").focus();
		$("#contentTip").addClass("onError").html("请输入摘要!");
		return false;
	}else{
		$("#content").removeClass("onFocus");
		$("#contentTip").removeClass("onError").addClass("onCorrect").html("正确");
	}
	
	return true;
}