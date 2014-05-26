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
 * 商品分类
 */

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

function goods_submit() {
	var cid = $.trim($("#cid").val());
	if (cid == 0) {
		$("#cid").focus();
		notice_tips("请选择所属分类!");
		return false;
	}
	var title = $.trim($("#title").val());
	if (title == '') {
		$("#title").focus();
		notice_tips("请输入商品名称!");
		return false;
	}
	var img = $.trim($("#img").val());
	if (img == '') {
		$("#thumb_preview").focus();
		notice_tips("请选择商品图片!");
		return false;
	}
	var url = $.trim($("#url").val());
	if (url == '') {
		$("#url").focus();
		notice_tips("请输入链接地址!");
		return false;
	}
	var price = $.trim($("#price").val());
	if (price == '') {
		$("#price").focus();
		notice_tips("请输入价格!");
		return false;
	}
	return true;
}

function set_goods_color(color) {
	$('#color').val(color);
}

function set_title_color(color) {
	$('#title_color').val(color);
}

function close_window() {
	if ($('#title').val() != '') {
		art.dialog({
			content : '商品已经录入，确定离开将不保存数据！',
			fixed : true,
			yesText : '我要关闭',
			noText : '返回保存数据',
			style : 'confirm',
			id : 'bnt4_test'
		}, function() {
			window.close();
		}, function() {

		});
	} else {
		window.close();
	}
	return false;
}

/**
 * 删除商品分类
 */
function cate_delete(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Goods/Delete/",
			data : "id=" + id,
			success : function(tmp) {
				if (tmp.status == 0) {
					window.location.reload();
					notice_tips("删除成功!");
				} else {
					notice_tips(tmp.message);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除分类操作!");
	});
}

$(document).ready(function() {
	// 选择基础分类
	$("#brand_cate").change(function() {
		// setBasic($(this).val());

		// setBrand($("#basic_cate").val());
	});

	// 选择品牌
	$("#basic_cate").change(function() {
		// setBrand($(this).val());
	});

	// 马甲专辑
	$("#uid").change(function() {
		setAlbum($(this).val());
	});

	setAlbum($("#uid").val());
});

/**
 * 设置马甲专辑
 * 
 * @param uid
 */
function setAlbum(uid) {
	if (uid == '' || uid == 0) {
		$("#aid").html("<option value=\"0\">请选择上级分类</option>");
		return;
	}

	$.ajax({
		type : "POST",
		url : "/Ajax/getUserAlbum/",
		data : "uid=" + uid,
		success : function(html) {
			$("#aid").html(html);
		}
	});
}

/**
 * 设置基础分类
 * 
 * @param brand_cate
 */
function setBasic(brand_cate) {
	if (brand_cate == '' || brand_cate == 0) {
		$("#basic_cate").html("<option value=\"0\">请选择上级分类</option>");
		return;
	}

	$.ajax({
		type : "POST",
		url : "/Ajax/getBasic/",
		data : "brand_cate=" + brand_cate,
		success : function(html) {
			$("#basic_cate").html(html);
		}
	});
}

/**
 * 设置品牌
 * 
 * @param basic_cate
 */
function setBrand(basic_cate) {

	if (basic_cate == '' || basic_cate == 0) {
		$("#bid").html("<option value=\"0\">请选择上级分类</option>");
		return;
	}
	$.ajax({
		type : "POST",
		url : "/Ajax/getBrand/",
		data : "basic_cate=" + basic_cate,
		success : function(html) {
			$("#bid").html(html);
		}
	});
}