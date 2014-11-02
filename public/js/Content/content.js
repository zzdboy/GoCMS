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
 * 内容管理
 */

function image_priview(img) {
	window.top.art.dialog({title:'图片查看',fixed:true, content:'<img src="'+img+'" />',id:'image_priview',time:5});
}
function remove_div(id) {
	$('#'+id).remove();
}

function input_font_bold() {
	if($('#title').css('font-weight') == '700' || $('#title').css('font-weight')=='bold') {
		$('#title').css('font-weight','normal');
		$('#style_font_weight').val('');
	} else {
		$('#title').css('font-weight','bold');
		$('#style_font_weight').val('bold');
	}
}
function ruselinkurl() {
	if($('#islink').attr('checked')==true) {
		$('#linkurl').attr('disabled','');
		var oEditor = CKEDITOR.instances.content;
		oEditor.insertHtml('　');
		return false;
	} else {
		$('#linkurl').attr('disabled','true');
	}
}
function close_window() {
	if($('#title').val() !='') {
	art.dialog({content:'内容已经录入，确定离开将不保存数据！', fixed:true,yesText:'我要关闭',noText:'返回保存数据',style:'confirm', id:'bnt4_test'}, function(){
				window.close();
			}, function(){
				
				});
	} else {
		window.close();
	}
	return false;
}


function ChangeInput (objSelect,objInput) {
	if (!objInput) return;
	var str = objInput.value;
	var arr = str.split(",");
	for (var i=0; i<arr.length; i++){
	  if(objSelect.value==arr[i])return;
	}
	if(objInput.value=='' || objInput.value==0 || objSelect.value==0){
	   objInput.value=objSelect.value
	}else{
	   objInput.value+=','+objSelect.value
	}
}

//移除相关文章
function remove_relation(sid,id) {
	var relation_ids = $('#relation').val();
	if(relation_ids !='' ) {
		$('#'+sid).remove();
		var r_arr = relation_ids.split(',');
		var newrelation_ids = '';
		$.each(r_arr, function(i, n){
			if(n!=id) {
				if(i==0) {
					newrelation_ids = n;
				} else {
				 newrelation_ids = newrelation_ids+','+n;
				}
			}
		});
		$('#relation').val(newrelation_ids);
	}
}

//移除ID
function remove_id(id) {
	$('#'+id).remove();
}

function strlen_verify(obj, checklen, maxlen) {
	var charset = 'utf-8'
	var v = obj.value, charlen = 0, maxlen = !maxlen ? 200 : maxlen, curlen = maxlen, len = strlen(v);
	for(var i = 0; i < v.length; i++) {
		if(v.charCodeAt(i) < 0 || v.charCodeAt(i) > 255) {
			curlen -= charset == 'utf-8' ? 2 : 1;
		}
	}
	if(curlen >= len) {
		$('#'+checklen).html(curlen - len);
	} else {
		obj.value = mb_cutstr(v, maxlen, true);
	}
}
function mb_cutstr(str, maxlen, dot) {
	var len = 0;
	var ret = '';
	var dot = !dot ? '...' : '';
	maxlen = maxlen - dot.length;
	for(var i = 0; i < str.length; i++) {
		len += str.charCodeAt(i) < 0 || str.charCodeAt(i) > 255 ? (charset == 'utf-8' ? 3 : 2) : 1;
		if(len > maxlen) {
			ret += dot;
			break;
		}
		ret += str.substr(i, 1);
	}
	return ret;
}
function strlen(str) {
	return (str.indexOf('\n') != -1) ? str.replace(/\r?\n/g, '_').length : str.length;
}

function set_title_color(color) {
	$('#title').css('color',color);
	$('#style_color').val(color);
}

function input_font_bold() {
	if($('#title').css('font-weight') == '700' || $('#title').css('font-weight')=='bold') {
		$('#title').css('font-weight','normal');
		$('#style_font_weight').val('');
	} else {
		$('#title').css('font-weight','bold');
		$('#style_font_weight').val('bold');
	}
}


/**
 * 提交检测
 */
function form_addcontent_submit() {
	var title = $.trim($("#title").val());
	if (title == '') {
		$("#title").addClass("onFocus").focus();
		$("#titleTip").addClass("onError").html("请输入标题!");
		return false;
	}else{
		$("#title").removeClass("onFocus");
	}

	var content = KE.html('content');
	
	if (content == '') {
		$("#content").addClass("onFocus").focus();
		$("#contentTip").addClass("onError").html("请输入内容!");
		return false;
	}else{
		$("#content").removeClass("onFocus");
	}

	return true;
}

/**
 * 提交检测
 */
function form_submit() {
	var title = $.trim($("#title").val());
	if (title == '') {
		$("#title").addClass("onFocus").focus();
		notice_tips("请输入标题!");
		return false;
	}else{
		$("#title").removeClass("onFocus");
	}

	var content = KE.html('content');
	
	if (content == '') {
		$("#content").addClass("onFocus").focus();
		notice_tips("请输入内容!");
		return false;
	}else{
		$("#content").removeClass("onFocus");
	}

	return true;
}
