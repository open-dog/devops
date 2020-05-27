var hostname='http://10.1.2.37:8089';
var version='/v1/';
host=hostname+version;
let api={
    devops_release_common:host+'devops/release/common',// 下拉;
    devops_git_branch:host+'devops/git/branch',//获取分支信息
    devops_release_add:host+'devops/release/add',//添加工单
    devops_release_info:host+'devops/release/info',//获取工单信息
    devops_release_edit:host+'devops/release/edit',//编辑工单
    devops_release_list:host+'devops/release/list',//编辑工单
    devops_release_export:host+'devops/release/export',//编辑工单
};
admin_host = hostname + '/admin/';
let admin={
    admin_devops_release_info:  admin_host + 'devops/release/info',
    admin_devops_release_add:  admin_host + 'devops/release/add',
};
