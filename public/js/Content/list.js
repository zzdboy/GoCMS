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
 * 内容列表
 */

$(document).ready(function() {

	setInterval(closeParent, 5000);
	$.cookie('refersh_time', 0);
	setInterval("refersh_window()", 3000);
	
	$(".preview").preview();

	Calendar.setup({
		weekNumbers : false,
		inputField : "start_time",
		trigger : "start_time",
		dateFormat : "%Y-%m-%d",
		showTime : false,
		minuteStep : 1,
		onSelect : function() {
			this.hide();
		}
	});

	Calendar.setup({
		weekNumbers : false,
		inputField : "end_time",
		trigger : "end_time",
		dateFormat : "%Y-%m-%d",
		showTime : false,
		minuteStep : 1,
		onSelect : function() {
			this.hide();
		}
	});
});

//推送
function push() {
	var str = 0;
	var id = tag = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).attr('checked') == 'checked') {
			str = 1;
			id += tag + $(this).val();
			tag = ',';
		}
	});
	if (str == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	window.top.art.dialog.open('/Content/push/' + id + '/',{
		id : 'push',
		title : '推送',
		id : 'push',
		width : 800,
		height : 500,
		lock : true
	});
}

//移动
function move() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	var cid = $("#cid").val();
	if(cid == "0") {
		window.top.art.dialog({
			icon : 'error',
			content : '请选择栏目!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	$.ajax({
        type : "POST",
        url : "/Content/remove/",
        data: "cid="+cid+"&ids="+ids+"&csrf_token"+csrf_token,
        success : function(data) {
            if(data.status==1){
                window.location.reload();
				notice_tips("移动成功!");
            } else {
                window.top.art.dialog({
					icon : 'error',
					content : data.message
				});
				return false;
            }
        }
    });
}

function show_remove() {

	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	if ($("#show_move").css("display") == "none") {
		$("#show_move").fadeIn("slow");
	}else{
		$("#show_move").fadeOut("slow");
	}
}


//批量移动
function batch_remove() {
	var num = 0;
	var ids = '';
	$("input[name='ids[]']").each(function() {
		if ($(this).is(":checked")) {
			num = 1;
			ids += $(this).val() + ',';
		}
	});
	if (num == 0) {
		window.top.art.dialog({
			icon : 'error',
			content : '您没有勾选信息!'
		});
		return false;
	}

	ids = ids.substring(0,ids.length - 1);

	window.top.art.dialog.open('/Content/remove/' + ids + '/',{
		title : '批量移动',
		id : 'remove',
		width : 800,
		height : 450,
		lock : true
	});
}

//删除新闻
function confirm_delete() {
	window.top.art.dialog({
		content:'确认删除吗？',
		ok:function(){
			myform.submit();
			return true;
		},
		canceVal:'取消',
		cancel:true
	});
}

//查看评论
function view_comment(id, name) {
	window.top.art.dialog.open('/Content/comment/' + id + '/',{
		yesText : '关闭',
		title : '查看评论：' + name,
		id : 'view_comment',
		width : 800,
		height : 500,
		lock : true
	});
}

function refersh_window() {
	var refersh_time = $.cookie('refersh_time');
	if (refersh_time == 1) {
		window.location.reload();
	}
}

function closeParent() {
	if ($('#closeParentTime').html() == '') {
		window.top.$(".left_menu").addClass("left_menu_on");
		window.top.$("#openClose").addClass("close");
		window.top.$("html").addClass("on");
		$('#closeParentTime').html('1');
		window.top.$("#openClose").data('clicknum', 1);
	}
}