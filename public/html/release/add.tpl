
<script>
    $('a.new-tab-link').on('click', function () {
        listenerForAddNavTab($(this).attr('href'), $(this).attr('data-title'))
    });
</script>

<script src="/admin/assets/dist/js/tree.min.js"></script>
<script src="/admin/assets/dist/js/datatable.min.js">
</script><script src="/admin/assets/dist/js/form.min.js"></script>

<section class="content-header">
	<h1>
		添加工单
		<small>添加工单</small>
	</h1>
	<ol class="breadcrumb" style="margin-right: 30px;">
		<li><a href="/admin/info/manager"><i class="fa fa-dashboard"></i> 首页</a></li>
	</ol>
</section>



<section class="content">
	<div>
		<div class="box box-">
			<!-- 工单的头信息 -->
			<div class="box-header with-border">
				<h3 class="box-title">新建</h3>
				<div class="box-tools">
					<div class="btn-group pull-right" style="margin-right: 10px">
						<a href='/devops/release/list' class="btn btn-sm btn-default form-history-back"><i
									class="fa fa-arrow-left"></i> 返回</a>
					</div>
				</div>
			</div>


			<div class="box-body" style=" ">
				<form action="/admin/new/permission" method="post" accept-charset="UTF-8" class="form-horizontal" pjax-container
				      style="background-color: white;">
					<div class="box-body">
						<div class="box-body">
							<div class="fields-group card-body">


								<div class="form-group" >
									<label for="name"
									       class="col-sm-2 asterisk control-label">工单名</label>
									<div class="col-sm-8">
										<div class="input-group">
											<span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
											<input required="1" type="text" id="title" name="title" value=''
											       class="form-control json" placeholder="输入 工单名">
										</div>
									</div>
								</div>


								<div class="form-group" >
									<label for="slug"
									       class="col-sm-2 asterisk control-label">负责人</label>
									<div class="col-sm-8">
										<div class="input-group">
											<span class="input-group-addon"><i class="fa fa-pencil fa-fw"></i></span>
											<input required="1" type="text" id="slug" name="slug" value=''
											       class="form-control json" placeholder="输入 项目负责人">
										</div>
									</div>
								</div>


								<div class="form-group" >
									<label for="http_method"
									       class="col-sm-2  control-label">添加sql</label>
									<div class="col-sm-8">
										<input type="button" class="btn btn-primary" value="添加">
									</div>
								</div>
								<!-- 对应的sql区域 -->
								<div class="form-group " style="border: 1px solid #cccccc" >
									<div class="row" >
										<div class="col-md-6">
											<div class="form-group">
												<label class="col-sm-4  control-label">DB名</label>
												<select class="form-control col-sm-2" style="width: 20%"  data-select2-id="1" tabindex="-1" aria-hidden="true">
													<option selected="selected" data-select2-id="3">haiji_goods</option>
													<option data-select2-id="30">haiji_tapin</option>
													<option data-select2-id="31">haiji_order</option>
													<option data-select2-id="32">haiji_activity</option>
													<option data-select2-id="33">haiji_idc</option>
													<option data-select2-id="34">haiji_user</option>
													<option data-select2-id="35">haiji_admin</option>
												</select>
											</div>
										</div>
										<!-- /.col -->
										<div class="col-md-6">
											<div class="form-group">
												<label class="col-sm-2  control-label">备注</label>
												<input class="form-control sql_remark " style="width: 30%" type="text" name="remark"value="" placeholder="对应sql解决的问题">
											</div>
										</div>

										<div class="col-md-3">
											<div class="form-group">
												<input type="button" name="" value="删除" class="btn btn-warning" >
											</div>
										</div>
										<!-- /.col -->
									</div>
									<div class="row">
										<textarea class="form-control" name="" id="" rows="3" placeholder="输入 对应的sql语句"></textarea>
									</div>
								</div>



								<!-- package扩展区域 -->
								<div class="form-group" >
									<label for="http_method"
									       class="col-sm-2  control-label">添加package</label>
									<div class="col-sm-8">
										<input type="button" class="btn btn-primary" value="添加">
									</div>
								</div>
								<div class="form-group " style="border: 1px solid #cccccc" >
									<div class="row" >
										<div class="col-md-6">
											<div class="form-group">
												<label class="col-sm-4  control-label">package名</label>
												<select class="form-control col-sm-2" style="width: 20%"  data-select2-id="1" tabindex="-1" aria-hidden="true">
													<option selected="selected" data-select2-id="3">goods-package</option>
													<option data-select2-id="30">order-package</option>
													<option data-select2-id="31">tapin-package</option>
													<option data-select2-id="32">epet-center</option>
													<option data-select2-id="33">user-package</option>
													<option data-select2-id="34">common</option>
													<option data-select2-id="35">common-config</option>
												</select>
											</div>
										</div>
										<!-- /.col -->
										<div class="col-md-6">
											<div class="form-group">
												<label class="col-sm-2  control-label">分支名</label>
												<select class="form-control col-sm-2" style="width: 20%"  data-select2-id="1" tabindex="-1" aria-hidden="true">
													<option selected="selected" data-select2-id="3">master</option>
													<option data-select2-id="30">test-master</option>
													<option data-select2-id="31">prod-master</option>
													<option data-select2-id="32">feature_dzx_test</option>
													<option data-select2-id="33">counpon_20191129</option>
													<option data-select2-id="34">counpon_20191129</option>
													<option data-select2-id="35">counpon_20191129</option>
												</select>
											</div>
										</div>

										<div class="col-md-3">
											<div class="form-group">
												<input type="button" name="" value="删除" class="btn btn-warning" >
											</div>
										</div>
										<!-- /.col -->
									</div>
								</div>

								<!-- 服务扩展区域 -->
								<div class="form-group" >
									<label for="http_method"
									       class="col-sm-2  control-label">添加service</label>
									<div class="col-sm-8">
										<input type="button" class="btn btn-primary" value="添加">
									</div>
								</div>
								<div class="form-group " style="border: 1px solid #cccccc" >
									<div class="row" >
										<div class="col-md-6">
											<div class="form-group">
												<label class="col-sm-4  control-label">package名</label>
												<select class="form-control col-sm-2" style="width: 20%"  data-select2-id="1" tabindex="-1" aria-hidden="true">
													<option selected="selected" data-select2-id="3">goods-package</option>
													<option data-select2-id="30">order-package</option>
													<option data-select2-id="31">tapin-package</option>
													<option data-select2-id="32">epet-center</option>
													<option data-select2-id="33">user-package</option>
													<option data-select2-id="34">common</option>
													<option data-select2-id="35">common-config</option>
												</select>
											</div>
										</div>
										<!-- /.col -->
										<div class="col-md-6">
											<div class="form-group">
												<label class="col-sm-2  control-label">分支名</label>
												<select class="form-control col-sm-2" style="width: 20%"  data-select2-id="1" tabindex="-1" aria-hidden="true">
													<option selected="selected" data-select2-id="3">master</option>
													<option data-select2-id="30">test-master</option>
													<option data-select2-id="31">prod-master</option>
													<option data-select2-id="32">feature_dzx_test</option>
													<option data-select2-id="33">counpon_20191129</option>
													<option data-select2-id="34">counpon_20191129</option>
													<option data-select2-id="35">counpon_20191129</option>
												</select>
											</div>
										</div>

										<div class="col-md-3">
											<div class="form-group">
												<input type="button" name="" value="删除" class="btn btn-warning" >
											</div>
										</div>
										<!-- /.col -->



									</div>
									<div class="row">
										<div class="col-md-6">
											<textarea class="form-control" name="" id="" rows="3" placeholder="输入 env修改内容"  ></textarea>
										</div>
										<div class="col-md-6">
											<textarea class="form-control" name="" id="" rows="3" placeholder="输入 脚本执行命令"  ></textarea>
										</div>
									</div>
								</div>


							</div>
						</div>
					</div>


					<div class="box-footer">
						<div class="col-md-2 "></div>
						<div class="col-md-8 ">
							<div class="btn-group pull-right" >
								<button type="submit" class="btn  btn-primary" data-loading-text="&nbsp;保存">
									保存
								</button>
							</div>
							<label class="pull-right" style="margin: 5px 10px 0 0;">
								<input type="checkbox" class="continue_new" style="position: absolute; opacity: 0;"> 继续新增
							</label>
							<div class="btn-group pull-left" >
								<button type="reset" class="btn  btn-warning" data-loading-text="&nbsp;保存">
									重置
								</button>
							</div>
							<script>
                                let previous_url_goadmin = $('input[name="__go_admin_previous_"]').attr("value")
                                $('.continue_new').iCheck({checkboxClass: 'icheckbox_minimal-blue'}).on('ifChanged', function (event) {
                                    if (this.checked) {
                                        $('input[name="__go_admin_previous_"]').val(location.href)
                                    } else {
                                        $('input[name="__go_admin_previous_"]').val(previous_url_goadmin)
                                    }
                                });
							</script>
						</div>
					</div>
					<input type="hidden" name="__go_admin_previous_" value='/admin/info/permission?__page=1&amp;__pageSize=10&amp;__sort=id&amp;__sort_type=desc'>
					<input type="hidden" name="__go_admin_t_" value='9fff925b-2157-4ce1-b97c-06453c9f1387'>
				</form>
			</div>
		</div>
	</div>
</section>


    
