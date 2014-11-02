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
 * tips提示消息
 * 
 * @param content
 */
function notice_tips(content) {
	art.dialog.tips(content, 1.5);
}

/**
 * 页面跳转
 * 
 * @param url
 */
function redirect(url) {
	location.href = url;
}

/**
 * 正确提示对话框
 * 
 * @param content
 */
function notice(content) {
	art.dialog({
		lock : true, // 锁屏
		icon : 'succeed',
		title : '提示',
		content : content
	}).time(3);
}

/**
 * 错误提示对话框
 * 
 * @param content
 */
function notice_error(content) {
	art.dialog({
		lock : true, // 锁屏
		icon : 'error',
		title : '提示',
		content : content
	}).time(3);
}

/**
 * 后台地图
 */
function OpenMap() {
	art.dialog.open('/Map/', {
		id : 'map',
		title : '后台地图',
		width : 700,
		height : 500,
		lock : true
	});
}

/**
 * 右侧内容页刷新
 */
function right_refresh() {
	art.dialog.top.$('#rightMain').attr('src',art.dialog.top.$('#rightMain').attr("src"));
}

/**
 * 设置iframe
 * @param id
 * @param src
 */
function set_iframe(id, src) {
	$("#" + id).attr("src", src);
}

// 滚动条
$(function() {
	$(":text").addClass('input-text');
});

/**
 * 全选checkbox,注意：标识checkbox id固定为为check_box
 * 
 * @param string
 *            name 列表check名称,如 uid[]
 * 
 */
function selectall(name) {
	if ($("#check_box").is(":checked")) {
		$("input[name='" + name + "']").each(function() {
			this.checked = true;
		});
	} else {
		$("input[name='" + name + "']").each(function() {
			this.checked = false;
		});
	}
}

function openwinx(url, name, w, h) {
	if (!w)
		w = screen.width - 4;
	if (!h)
		h = screen.height - 95;

	window.open(url,name,"top=100,left=400,width="
		+ w
		+ ",height="
		+ h
		+ ",toolbar=no,menubar=no,scrollbars=yes,resizable=yes,location=no,status=no");
}

// 添加快捷方式
function add_panel() {
	var mid = $("#mid").val();

	$.post("/ajax/AddPanel/",{'mid':mid,'csrf_token':csrf_token}, function(data){
		$("#panellist").html(data);
	});
}

// 删除快捷方式
function delete_panel() {
	var mid = $("#mid").val();

	$.post("/ajax/DelPanel/",{'mid':mid,'csrf_token':csrf_token}, function(data){
		if (data.status == 0) {

			//快捷方式
			$.post("/ajax/GetPanel/",{'mid':mid,'csrf_token':csrf_token}, function(data){
				$("#panellist").html(data);
			});

			notice_tips("取消成功!");
			return false;
		}else{
			notice_tips(data.message);
			return false;
		}
	});
}

// 弹出对话框
function omnipotent(id, linkurl, title, close_type, w, h) {
	if (!w) {
		w = 700;
	}

	if (!h) {
		h = 500;
	}

	art.dialog.open(linkurl, {
		id : id,
		title : title,
		width : w,
		height : h,
		lock : true
	});
}