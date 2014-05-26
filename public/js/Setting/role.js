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
 * 角色管理
 */

/**
 * 表单提交
 * 
 * @returns {Boolean}
 */
function form_submit() {

	var rolename = $.trim($("#rolename").val());
	if (rolename == '') {
		$("#rolename").focus();
		notice_tips("请输入角色名称!");
		return false;
	}

	var desc = $.trim($("#desc").val());
	if (desc == '') {
		$("#desc").focus();
		notice_tips("请输入角色描述!");
		return false;
	}

	var tree = GetTreeCheckedAll();
	if(tree=='') {
		notice_tips("请选择所属权限!");
		return false;
	}else{
		$("#data").val(tree);
	}

	return true;
}

 //获取所有选中节点的值
function GetTreeCheckedAll() {
    var treeObj = $.fn.zTree.getZTreeObj("tree");
    var nodes = treeObj.getCheckedNodes(true);
    var msg = "";
    for (var i = 0; i < nodes.length; i++) {
        msg += nodes[i].id+",";
    }
    msg = msg.substring(0,msg.length-1)
    return msg;
}

/**
 * 删除角色
 * 
 * @param roleid
 */
function delete_role(roleid) {
	if (roleid == '') {
		notice_tips("参数错误!");
		return false;
	}

	art.dialog.confirm('你确定要删除这个角色吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Role/delete/",
			data : "id=" + roleid,
			success : function(tmp) {
				if (tmp.status == 0) {
					notice_tips("删除成功!");
					window.location.reload();
				} else {
					notice_tips(tmp.message);
				}
			}
		});
	}, function() {
		notice_tips("你取消了删除角色操作!");
	});
}

/**
 * 设置状态
 * 
 * @param roleid
 * @param status
 */
function setStatus(roleid, status) {
	if (roleid == '') {
		notice_tips("参数错误!");
		return false;
	}

	if (status == 1) {
		var message = '你确定要启用这个角色及用户吗?';
	} else {
		var message = '你确定要禁用这个角色及用户吗?';
	}

	art.dialog.confirm(message, function() {
		$.ajax({
			type : "POST",
			url : "/Role/setStatus/",
			data : "id=" + roleid + "&status=" + status,
			success : function(tmp) {
				if (tmp.status == 1) {
					notice_tips("设置成功!");
					window.location.reload();
				} else {
					notice_tips(tmp.message);
				}
			}
		});
	}, function() {
		notice_tips("你取消了设置状态操作!");
	});
}