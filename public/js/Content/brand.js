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
 * 品牌分类
 */

$(document).ready(function() {
	// 选择基础分类
	$("#brand_cate").change(function() {
		setBasic($(this).val());
	});
});

/**
 * 设置基础分类
 * 
 * @param brand_cate
 */
function setBasic(brand_cate) {
	if (brand_cate == '' || brand_cate == 0) {
		$("#parentid").html("<option value=\"0\">请选择上级分类</option>");
		return;
	}

	$.post("/Ajax/getBasic/",{'brand_cate':brand_cate,'csrf_token':csrf_token}, function(data){
		$("#parentid").html(html);
	});
}

/**
 * 提交检测
 */
function form_submit() {
	var name = $.trim($("#name").val());
	if (name == '') {
		$("#name").focus();
		notice_tips("请输入分类名称!");
		return false;
	}
	return true;
}

function set_Brand_color(color) {
	$('#color').val(color);
}

function set_title_color(color) {
	$('#title_color').val(color);
}

/**
 * 删除品牌分类
 */
function del(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.post("/Brand/delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
			if (tmp.status == 0) {
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