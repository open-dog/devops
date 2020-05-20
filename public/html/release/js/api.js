let host='http://10.1.2.8:8089';
let version='/v1/'
host=host+version;
let api={
    devops_release_common:host+'devops/release/common',// 下拉;
    devops_git_branch:host+'devops/git/branch',//获取分支信息
    devops_release_add:host+'devops/release/add',//添加工单
    devops_release_info:host+'devops/release/info',//获取工单信息
    devops_release_edit:host+'devops/release/edit',//编辑工单
}
