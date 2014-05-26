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
		$("#name").focus();
		notice_tips("请输入中文语言名称!");
		return false;
	}

	var enname = $.trim($("#enname").val());
	if (enname == '') {
		$("#enname").focus();
		notice_tips("请输入英文语言名称!");
		return false;
	}

	var url = $.trim($("#url").val());
	if (url == '') {
		$("#url").focus();
		notice_tips("请输入功能地址!");
		return false;
	}

	var order = $.trim($("#order").val());
	if (order == '') {
		$("#order").focus();
		notice_tips("请输入排序!");
		return false;
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
		$.ajax({
			type : "POST",
			url : "/Menu/delete/",
			data : "id=" + id,
			success : function(tmp) {
				if (tmp.status == 1) {
					notice_tips("删除菜单成功!");
					right_refresh();
				} else {
					notice_tips(tmp.message);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除菜单操作!");
	});
}