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
 * 菜单管理
 */

/**
 * 提交检测
 */
function form_submit() {
	var name = $.trim($("#name").val());
	if (name == '') {
		$("#name").addClass("onFocus").focus();
		$("#nameTip").addClass("onError").html("请输入中文语言名称!");
		return false;
	}else{
		$("#name").removeClass("onFocus");
		$("#nameTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var enname = $.trim($("#enname").val());
	if (enname == '') {
		$("#enname").addClass("onFocus").focus();
		$("#ennameTip").addClass("onError").html("请输入英文语言名称!");
		return false;
	}else{
		$("#enname").removeClass("onFocus");
		$("#ennameTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var url = $.trim($("#url").val());
	if (url == '') {
		$("#url").addClass("onFocus").focus();
		$("#urlTip").addClass("onError").html("请输入功能地址!");
		return false;
	}else{
		$("#url").removeClass("onFocus");
		$("#urlTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	var order = $.trim($("#order").val());
	if (order == '') {
		$("#order").addClass("onFocus").focus();
		$("#orderTip").addClass("onError").html("请输入排序!");
		return false;
	}else{
		$("#order").removeClass("onFocus");
		$("#orderTip").removeClass("onError").addClass("onCorrect").html("正确");
	}

	return true;
}

/**
 * 删除菜单
 * 
 * @param id
 */
function delete_menu(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.post("/Menu/Delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 1) {
				notice_tips("删除菜单成功!");
				right_refresh();
			} else {
				notice_tips(tmp.message);
			}
		});

	}, function() {
		notice_tips("你取消了删除菜单操作!");
	});
}