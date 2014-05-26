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
 * 管理栏目
 */

// Tab切换
function SwapTab(name, cls_show, cls_hide, cnt, cur) {
	for ( var i = 1; i <= cnt; i++) {
		if (i == cur) {
			$('#div_' + name + '_' + i).show();
			$('#tab_' + name + '_' + i).attr('class', cls_show);
		} else {
			$('#div_' + name + '_' + i).hide();
			$('#tab_' + name + '_' + i).attr('class', cls_hide);
		}
	}
}

/**
 * 提交检测
 */
function form_submit() {
	var name = $.trim($("#name").val());
	if (name == '') {
		$("#name").focus();
		notice_tips("请输入栏目名称!");
		return false;
	}

	var url = $.trim($("#url").val());
	if (url == '') {
		$("#url").focus();
		notice_tips("请输入栏目地址!");
		return false;
	}

	var desc = $.trim($("#desc").val());
	if (desc == '') {
		$("#desc").focus();
		notice_tips("请输入描述!");
		return false;
	}

	return true;
}

/**
 * 删除栏目
 * 
 * @param id
 */
function delete_cate(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Category/Delete/",
			data : "id=" + id,
			success : function(tmp) {
				if (tmp.status == 1) {
					notice_tips("删除栏目成功!");
					right_refresh();
				} else {
					notice_tips(tmp.message);
					right_refresh();
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除栏目操作!");
	});
}