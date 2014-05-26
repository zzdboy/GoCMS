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
 * 编辑焦点图
 */
function cate_edit(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.open('/Focus/editCate/' + id + '/', {
		id : 'cate_edit',
		title : '编辑分类',
		width : 300,
		height : 120,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;

			var id = iframe.$('#id').val();
			var name = iframe.$('#name').val(); // 分类名称
			var width = iframe.$('#width').val(); // 宽度
			var height = iframe.$('#height').val(); // 高度

			if (id == '' || id == 'undefined') {
				art.dialog.alert('参数错误!');
				return false;
			}

			if (name == '') {
				art.dialog.alert('请输入分类名称!');
				return false;
			}
			if (width == '') {
				art.dialog.alert('请输入宽度!');
				return false;
			}
			if (height == '') {
				art.dialog.alert('请输入高度!');
				return false;
			}

			var par = [];
			var pars = "id=" + id;
			par.push(pars);
			pars = "name=" + name;
			par.push(pars);
			pars = "width=" + width;
			par.push(pars);
			pars = "height=" + height;
			par.push(pars);
			pars = par.join("&");

			$.ajax({
				type : "POST",
				url : "/Focus/editCate/",
				data : pars,
				success : function(html) {
					var tmp = jQuery.parseJSON(html);
					if (tmp.rtn_code == 0) {
						window.location.reload();
						notice_tips("编辑分类成功!");
					} else {
						notice_tips(tmp.content);
					}
				}
			});
		},
		cancel : true
	});
}

/**
 * 删除焦点图分类
 */
function cate_delete(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Focus/deleteCate/",
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

/**
 * 删除焦点图
 */
function delete_Focus(id) {
	if (id == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Focus/delete/",
			data : "id=" + id,
			success : function(html) {
				var tmp = jQuery.parseJSON(html);
				if (tmp.rtn_code == 0) {
					window.location.reload();
					notice_tips("删除成功!");
				} else {
					notice_tips(tmp.content);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除操作!");
	});
}