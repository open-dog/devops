
<script>
    $('a.new-tab-link').on('click', function () {
        listenerForAddNavTab($(this).attr('href'), $(this).attr('data-title'))
    });
</script>

<script src="/admin/assets/dist/js/form.min.js"></script><script src="/admin/assets/dist/js/tree.min.js"></script><script src="/admin/assets/dist/js/datatable.min.js"></script>

<section class="content-header">
    <h1>
        工单
        <small>工单列表</small>
    </h1>
    <ol class="breadcrumb" style="margin-right: 30px;">
        <li><a href="/admin/info/manager"><i class="fa fa-dashboard"></i> 首页</a></li>
        <li>工单</li><li>工单列表</li>
    </ol>
</section>



<section class="content">
    <div>
        <div class="box box-">
            <div class="box-header with-border">
                <div class="pull-right">
                    <div class="dropdown pull-right column-selector" style="margin-right: 10px">
                        <button type="button" class="btn btn-sm btn-instagram dropdown-toggle" data-toggle="dropdown">
                            <i class="fa fa-table"></i>
                            &nbsp;
                            <span class="caret"></span>
                        </button>
                        <ul class="dropdown-menu" role="menu" style="padding: 10px;max-height: 400px;overflow: scroll;">
                            <li>
                                <ul style="padding: 0;">
                                    <li class="checkbox icheck" style="margin: 0;">
                                        <label style="width: 100%;padding: 3px;">
                                            <input type="checkbox" class="column-select-item" data-id="id"
                                                   style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;ID
                                        </label>
                                    </li>
                                    <li class="checkbox icheck" style="margin: 0;">
                                        <label style="width: 100%;padding: 3px;">
                                            <input type="checkbox" class="column-select-item" data-id="username"
                                                   style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;用户名
                                        </label>
                                    </li>
                                    <li class="checkbox icheck" style="margin: 0;">
                                        <label style="width: 100%;padding: 3px;">
                                            <input type="checkbox" class="column-select-item" data-id="name"
                                                   style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;昵称
                                        </label>
                                    </li>
                                    <li class="checkbox icheck" style="margin: 0;">
                                        <label style="width: 100%;padding: 3px;">
                                            <input type="checkbox" class="column-select-item" data-id="goadmin_roles_goadmin_join_name"
                                                   style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;角色
                                        </label>
                                    </li>
                                    <li class="checkbox icheck" style="margin: 0;">
                                        <label style="width: 100%;padding: 3px;">
                                            <input type="checkbox" class="column-select-item" data-id="created_at"
                                                   style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;创建时间
                                        </label>
                                    </li>
                                    <li class="checkbox icheck" style="margin: 0;">
                                        <label style="width: 100%;padding: 3px;">
                                            <input type="checkbox" class="column-select-item" data-id="updated_at"
                                                   style="position: absolute; opacity: 0;">&nbsp;&nbsp;&nbsp;更新时间
                                        </label>
                                    </li>
                                </ul>
                            </li>
                            <li class="divider">
                            </li>
                            <li class="text-right">
                                <button class="btn btn-sm btn-default column-select-all">全部</button>&nbsp;&nbsp;
                                <button class="btn btn-sm btn-primary column-select-submit">提交</button>
                            </li>
                        </ul>
                    </div>
                    <div class="btn-group pull-right" style="margin-right: 10px">
                        <a href="javascript:;" class="btn btn-sm btn-primary" id="filter-btn"><i
                                    class="fa fa-filter"></i>&nbsp;&nbsp;筛选</a>
                    </div>
                    <script>
                        $("#filter-btn").click(function () {
                            $('.filter-area').toggle();
                        });
                    </script>
                    <div class="btn-group pull-right" style="margin-right: 10px">
                        <a href="/admin/info/manager/new?__page=1&amp;__pageSize=10&amp;__sort=id&amp;__sort_type=desc" class="btn btn-sm btn-success">
                            <i class="fa fa-plus"></i>&nbsp;&nbsp;新建
                        </a>
                        <div class="btn-group">
                            <a class="btn btn-sm btn-default">导出</a>
                            <button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
                                <span class="caret"></span>
                                <span class="sr-only">下拉</span>
                            </button>
                            <ul class="dropdown-menu" role="menu">
                                <li><a href="#" id="export-btn-0">当前页</a></li>
                                <li><a href="#" id="export-btn-1">全部</a></li>
                            </ul>
                        </div>
                    </div>
                </div>
                <span>
<div class="btn-group">
<a class="btn btn-sm btn-default">操作</a>
<button type="button" class="btn btn-sm btn-default dropdown-toggle" data-toggle="dropdown">
<span class="caret"></span>
<span class="sr-only">下拉</span>
</button>
<ul class="dropdown-menu" role="menu">
<li><a href="#" class="grid-batch-0">删除</a></li>
<li><a href="#" class="grid-batch-1">导出</a></li>
</ul>
</div>
<a class="btn btn-sm btn-primary grid-refresh">
<i class="fa fa-refresh"></i> 刷新
</a>
</span>
                <script>
                    let toastMsg = '刷新成功 !';
                    $('.grid-refresh').on('click', function () {
                        $.pjax.reload('#pjax-container');
                        toastr.success(toastMsg);
                    });
                    $("#export-btn-0").click(function () {
                        ExportData("false")
                    });
                    $("#export-btn-1").click(function () {
                        ExportData("true")
                    });
                    function ExportData(isAll) {
                        let form = $("<form>");
                        form.attr("style", "display:none");
                        form.attr("target", "");
                        form.attr("method", "post");
                        form.attr("action","/admin/export/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc");
                        let input1 = $("<input>");
                        input1.attr("type", "hidden");
                        input1.attr("name", "time");
                        input1.attr("value", (new Date()).getTime());
                        let input2 = $("<input>");
                        input2.attr("type", "hidden");
                        input2.attr("name", "is_all");
                        input2.attr("value", isAll);
                        $("body").append(form);
                        form.append(input1);
                        form.append(input2);
                        form.submit();
                        form.remove()
                    }
                </script>
            </div>
            <div class="box-header filter-area ">
                <form action="/admin/info/manager?__pageSize=10&amp;__sort=id&amp;__sort_type=desc" method="get" accept-charset="UTF-8" class="form-horizontal" pjax-container
                      style="background-color: white;">
                    <div class="box-body">
                        <div class="box-body">
                            <div class="fields-group">
                                <div class="form-group" >
                                    <label for="username"
                                           class="col-sm-2  control-label">用户名</label>
                                    <div class="col-sm-10">
                                        <div class="input-group">
                                            <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
                                            <input  type="text" id="username" name="username" value=''
                                                    class="form-control json" placeholder="输入 用户名">
                                        </div>
                                    </div>
                                </div>
                                <div class="form-group" >
                                    <label for="name"
                                           class="col-sm-2  control-label">昵称</label>
                                    <div class="col-sm-10">
                                        <div class="input-group">
                                            <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
                                            <input  type="text" id="name" name="name" value=''
                                                    class="form-control json" placeholder="输入 昵称">
                                        </div>
                                    </div>
                                </div>
                                <div class="form-group" >
                                    <label for="goadmin_roles_goadmin_join_name"
                                           class="col-sm-2  control-label">角色</label>
                                    <div class="col-sm-10">
                                        <div class="input-group">
                                            <span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
                                            <input  type="text" id="goadmin_roles_goadmin_join_name" name="goadmin_roles_goadmin_join_name" value=''
                                                    class="form-control json" placeholder="输入 角色">
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="box-footer">
                        <div class="col-md-2 "></div>
                        <div class="col-md-8 ">
                            <div class="btn-group pull-left" >
                                <button type="submit" class="btn btn-sm btn-primary" data-loading-text="&nbsp;搜索">
                                    <i class="icon fa fa-search"></i>&nbsp;&nbsp;搜索
                                </button>
                            </div>
                            <div class="btn-group pull-left" style="margin-left:12px;">
                                <a href="/admin/info/manager" type="reset" class="btn btn-sm btn-default" data-loading-text="&nbsp;保存">
                                    <i class="icon fa fa-undo"></i>&nbsp;&nbsp;重置
                                </a>
                            </div>
                        </div>
                    </div>
                    <input type="hidden" name="__go_admin_no_animation_" value='true'>
                </form>
            </div>

            <div class="box-body" style="overflow-x: scroll;overflow-y: hidden;padding:0;">
                <table class="table table-hover" style="min-width: 1000px;table-layout: auto;">
                    <tbody>
                    <tr>
                        <th style="text-align: center;">
                            <input type="checkbox" class="grid-select-all" style="position: absolute; opacity: 0;">
                        </th>
                        <th>
                            工单ID
                            <a class="fa fa-fw fa-sort" id="sort-id"
                               href="?__sort=id&__sort_type=desc"></a>
                        </th>
                        <th>
                            工单名
                        </th>
                        <th>
                            创建人
                        </th>
                        <th>
                            状态
                        </th>
                        <th>
                            工单描述
                        </th>
                        <th>
                            创建时间
                        </th>
                        <th>
                            更新时间
                        </th>
                        <th style="text-align: center;">操作</th>
                    </tr>
                    <!-- 需要循环的表单数据 -->
                    <tr>
                        <td style="text-align: center;">
                            <input type="checkbox" class="grid-row-checkbox"
                                   data-id="2"
                                   style="position: absolute; opacity: 0;">
                        </td>
                        <td>2</td>
                        <td>operator</td>
                        <td>Operator</td>
                        <td>Operator</td>
                        <td>
                            <span class="label label-success" style="background-color: ;">Operator</span>
                        </td>
                        <td>2019-09-10 00:00:00</td>
                        <td>2019-09-10 00:00:00</td>
                        <td style="text-align: center;">
                            <a href='/admin/info/manager/edit?__page=1&amp;__pageSize=10&amp;__sort=id&amp;__sort_type=desc&__goadmin_edit_pk=2'><i
                                        class="fa fa-edit"></i></a>
                            <a href="javascript:void(0);" data-id='2'
                               class="grid-row-delete"><i class="fa fa-trash"></i></a>
                            <a href="/admin/info/manager/detail?__page=1&amp;__pageSize=10&amp;__sort=id&amp;__sort_type=desc&__goadmin_detail_pk=2"
                               class="grid-row-view">
                                <i class="fa fa-eye"></i>
                            </a>
                        </td>
                    </tr>

                    </tbody>
                </table>
                <script>
                    window.selectedRows = function () {
                        let selected = [];
                        $('.grid-row-checkbox:checked').each(function () {
                            selected.push($(this).data('id'));
                        });
                        return selected;
                    };
                    const selectedAllFieldsRows = function () {
                        let selected = [];
                        $('.column-select-item:checked').each(function () {
                            selected.push($(this).data('id'));
                        });
                        return selected;
                    };
                    const pjaxContainer = "#pjax-container";
                    const noAnimation = "__go_admin_no_animation_";
                    function iCheck(el) {
                        el.iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function () {
                            if (this.checked) {
                                $(this).closest('tr').css('background-color', "#ffffd5");
                            } else {
                                $(this).closest('tr').css('background-color', '');
                            }
                        });
                    }
                    $(function () {
                        $('.grid-select-all').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function (event) {
                            if (this.checked) {
                                $('.grid-row-checkbox').iCheck('check');
                            } else {
                                $('.grid-row-checkbox').iCheck('uncheck');
                            }
                        });
                        let items = $('.column-select-item');
                        iCheck(items);
                        iCheck($('.grid-row-checkbox'));
                        let columns = getQueryVariable("__columns");
                        if (columns === -1) {
                            items.iCheck('check');
                        } else {
                            let columnsArr = columns.split(",");
                            for (let i = 0; i < columnsArr.length; i++) {
                                for (let j = 0; j < items.length; j++) {
                                    if (decodeURI(columnsArr[i]) === $(items[j]).attr("data-id")) {
                                        $(items[j]).iCheck('check');
                                    }
                                }
                            }
                        }
                        $('.filter-area').hide();
                        let lastTd = $("table tr:last td:last div");
                        if (lastTd.hasClass("dropdown")) {
                            let popUpHeight = $("table tr:last td:last div ul").height();
                            let trs = $("table tr");
                            let totalHeight = 0;
                            for (let i = 1; i < trs.length - 1; i++) {
                                totalHeight += $(trs[i]).height();
                            }
                            if (popUpHeight > totalHeight) {
                                let h = popUpHeight + 16;
                                $("table tbody").append("<tr style='height:" + h + "px;'></tr>");
                            }
                            trs = $("table tr");
                            for (let i = trs.length - 1; i > 1; i--) {
                                let td = $(trs[i]).find("td:last-child div");
                                let combineHeight = $(trs[i]).height() / 2 - 20;
                                for (let j = i + 1; j < trs.length; j++) {
                                    combineHeight += $(trs[j]).height();
                                }
                                if (combineHeight < popUpHeight) {
                                    td.removeClass("dropdown");
                                    td.addClass("dropup");
                                }
                            }
                        }
                        let sort = getQueryVariable("__sort");
                        let sort_type = getQueryVariable("__sort_type");
                        if (sort !== -1 && sort_type !== -1) {
                            let sortFa = $('#sort-' + sort);
                            if (sort_type === 'asc') {
                                sortFa.attr('href', '?__sort=' + sort + "&__sort_type=desc")
                            } else {
                                sortFa.attr('href', '?__sort=' + sort + "&__sort_type=asc")
                            }
                            sortFa.removeClass('fa-sort');
                            sortFa.addClass('fa-sort-amount-' + sort_type);
                        }
                    });
                    $('.column-select-all').on('click', function () {
                        if ($(this).data('check') === '') {
                            $('.column-select-item').iCheck('check');
                            $(this).data('check', 'true')
                        } else {
                            $('.column-select-item').iCheck('uncheck');
                            $(this).data('check', '')
                        }
                    });
                    $('.column-select-submit').on('click', function () {
                        let param = new Map();
                        param.set('__columns', selectedAllFieldsRows().join(','));
                        param.set(noAnimation, 'true');
                        $.pjax({
                            url: addParameterToURL(param),
                            container: pjaxContainer
                        });
                        toastr.success('加载成功 !');
                    });
                    $('.grid-batch-1').on('click', function () {
                        let rows = selectedRows();
                        if (rows.length > 0) {
                            ExportAll(rows.join())
                        }
                    });
                    function ExportAll(id) {
                        let form = $("<form>");
                        form.attr("style", "display:none");
                        form.attr("target", "");
                        form.attr("method", "post");
                        form.attr("action","/admin/export/manager?__page=1\u0026__pageSize=10\u0026__sort=id\u0026__sort_type=desc");
                        let input1 = $("<input>");
                        input1.attr("type", "hidden");
                        input1.attr("name","id");
                        input1.attr("value", id);
                        $("body").append(form);
                        form.append(input1);
                        form.submit();
                        form.remove()
                    }
                    $('.grid-row-delete').unbind('click').click(function () {
                        DeletePost($(this).data('id'))
                    });
                    $('.grid-batch-0').on('click', function () {
                        let rows = selectedRows();
                        if (rows.length > 0) {
                            DeletePost(rows.join())
                        }
                    });
                    function DeletePost(id) {
                        swal({
                                    title: "你确定要删除吗",
                                    type: "warning",
                                    showCancelButton: true,
                                    confirmButtonColor: "#DD6B55",
                                    confirmButtonText: "确定",
                                    closeOnConfirm: false,
                                    cancelButtonText: "取消",
                                },
                                function () {
                                    $.ajax({
                                        method: 'post',
                                        url: "/admin/delete/manager",
                                        data: {
                                            id: id
                                        },
                                        success: function (data) {
                                            let param = new Map();
                                            param.set(noAnimation, "true");
                                            $.pjax({
                                                url: addParameterToURL(param),
                                                container: pjaxContainer
                                            });
                                            if (typeof (data) === "string") {
                                                data = JSON.parse(data);
                                            }
                                            if (data.code === 200) {
                                                $('#_TOKEN').val(data.data);
                                                let lastTd = $("table tr:last td:last div");
                                                if (lastTd.hasClass("dropdown")) {
                                                    let popUpHeight = $("table tr:last td:last div ul").height();
                                                    let trs = $("table tr");
                                                    let totalHeight = 0;
                                                    for (let i = 1; i < trs.length - 1; i++) {
                                                        totalHeight += $(trs[i]).height();
                                                    }
                                                    if (popUpHeight > totalHeight) {
                                                        let h = popUpHeight + 16;
                                                        $("table tbody").append("<tr style='height:" + h + "px;'></tr>");
                                                    }
                                                }
                                                swal(data.msg, '', 'success');
                                            } else {
                                                swal(data.msg, '', 'error');
                                            }
                                        },
                                        error: function (data) {
                                            if (data.responseText !== "") {
                                                swal(data.responseJSON.msg, '', 'error');
                                            } else {
                                                swal("错误", '', 'error');
                                            }
                                        },
                                    });
                                });
                    }
                    function getQueryVariable(variable) {
                        let query = window.location.search.substring(1);
                        let vars = query.split("&");
                        for (let i = 0; i < vars.length; i++) {
                            let pair = vars[i].split("=");
                            if (pair[0] === variable) {
                                return pair[1];
                            }
                        }
                        return -1;
                    }
                    function addParameterToURL(params) {
                        let newUrl = location.href.replace("#", "");
                        for (let [field, value] of params) {
                            if (getQueryVariable(field) !== -1) {
                                newUrl = replaceParamVal(newUrl, field, value);
                            } else {
                                if (newUrl.indexOf("?") > 0) {
                                    newUrl = newUrl + "&" + field + "=" + value;
                                } else {
                                    newUrl = newUrl + "?" + field + "=" + value;
                                }
                            }
                        }
                        return newUrl
                    }
                    function replaceParamVal(oUrl, paramName, replaceWith) {
                        let re = eval('/(' + paramName + '=)([^&]*)/gi');
                        return oUrl.replace(re, paramName + '=' + replaceWith);
                    }
                    $(function () {
                        $('.editable-td-select').editable({
                            "type": "select",
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>"
                        });
                        $('.editable-td-text').editable({
                            emptytext: "<i class=\"fa fa-pencil\"><\/i>",
                            type: "text"
                        });
                        $('.editable-td-datetime').editable({
                            "type": "combodate",
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                            "format": "YYYY-MM-DD HH:mm:ss",
                            "viewformat": "YYYY-MM-DD HH:mm:ss",
                            "template": "YYYY-MM-DD HH:mm:ss",
                            "combodate": {"maxYear": 2035}
                        });
                        $('.editable-td-date').editable({
                            "type": "combodate",
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                            "format": "YYYY-MM-DD",
                            "viewformat": "YYYY-MM-DD",
                            "template": "YYYY-MM-DD",
                            "combodate": {"maxYear": 2035}
                        });
                        $('.editable-td-year').editable({
                            "type": "combodate",
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                            "format": "YYYY",
                            "viewformat": "YYYY",
                            "template": "YYYY",
                            "combodate": {"maxYear": 2035}
                        });
                        $('.editable-td-month').editable({
                            "type": "combodate",
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                            "format": "MM",
                            "viewformat": "MM",
                            "template": "MM",
                            "combodate": {"maxYear": 2035}
                        });
                        $('.editable-td-day').editable({
                            "type": "combodate",
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>",
                            "format": "DD",
                            "viewformat": "DD",
                            "template": "DD",
                            "combodate": {"maxYear": 2035}
                        });
                        $('.editable-td-textarea').editable({
                            "type": "textarea",
                            "rows": 10,
                            "emptytext": "<i class=\"fa fa-pencil\"><\/i>"
                        });
                        $(".info_edit_switch").bootstrapSwitch({
                            onSwitchChange: function (event, state) {
                                let obejct = $(event.target);
                                let val = "";
                                if (state) {
                                    val = obejct.closest('.bootstrap-switch').next().val();
                                } else {
                                    val = obejct.closest('.bootstrap-switch').next().next().val()
                                }
                                $.ajax({
                                    method: 'post',
                                    url: obejct.data("updateurl"),
                                    data: {
                                        name: obejct.data("field"),
                                        value: val,
                                        pk: obejct.data("pk")
                                    },
                                    success: function (data) {
                                        if (typeof (data) === "string") {
                                            data = JSON.parse(data);
                                        }
                                        if (data.code !== 200) {
                                            swal(data.msg, '', 'error');
                                        }
                                    },
                                    error: function (data) {
                                        if (data.responseText !== "") {
                                            swal(data.responseJSON.msg, '', 'error');
                                        } else {
                                            swal("错误", '', 'error');
                                        }
                                    },
                                });
                            }
                        })
                    });
                </script>
                <style>
                    table tbody tr td {
                        word-wrap: break-word;
                        word-break: break-all;
                    }
                </style>
            </div>
            <div class="box-footer clearfix">
                <div style="float: left;margin-top: 21px;">showing <b>0</b> to
                    <b>10</b> of <b>2</b> entries &nbsp;&nbsp;&nbsp;<b>query time: </b>2.548ms
                </div>
                <ul class="pagination pagination-sm no-margin pull-right">
                    <li class="page-item disabled">
                        <span class="page-link">«</span>
                    </li>
                    <li class="page-item active"><span class="page-link">1</span></li>
                    <li class='page-item disabled'>
                        <span class="page-link">»</span>
                    </li>
                </ul>
                <label class="control-label pull-right" style="margin-right: 10px; font-weight: 100;">
                    <small>show</small>&nbsp;
                    <select class="input-sm grid-per-pager" name="per-page">
                        <option value="/admin/info/manager?__is_all=false&amp;__page=1&amp;__sort=id&amp;__sort_type=desc&amp;__go_admin_no_animation_=true&__pageSize=10" selected>
                            10
                        </option>
                        <option value="/admin/info/manager?__is_all=false&amp;__page=1&amp;__sort=id&amp;__sort_type=desc&amp;__go_admin_no_animation_=true&__pageSize=20" ZgotmplZ>
                            20
                        </option>
                        <option value="/admin/info/manager?__is_all=false&amp;__page=1&amp;__sort=id&amp;__sort_type=desc&amp;__go_admin_no_animation_=true&__pageSize=30" ZgotmplZ>
                            30
                        </option>
                        <option value="/admin/info/manager?__is_all=false&amp;__page=1&amp;__sort=id&amp;__sort_type=desc&amp;__go_admin_no_animation_=true&__pageSize=50" ZgotmplZ>
                            50
                        </option>
                        <option value="/admin/info/manager?__is_all=false&amp;__page=1&amp;__sort=id&amp;__sort_type=desc&amp;__go_admin_no_animation_=true&__pageSize=100" ZgotmplZ>
                            100
                        </option>
                    </select>
                    <small>entries</small>
                </label>
                <script>
                    let gridPerPaper = $('.grid-per-pager');
                    gridPerPaper.on('change', function () {
                        $.pjax({url: this.value, container: '#pjax-container'});
                    });
                </script>
            </div>
        </div>
    </div>
</section>

