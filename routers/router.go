// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rs_web_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/cluster_config",
			beego.NSInclude(
				&controllers.ClusterConfigController{},
			),
		),

		beego.NSNamespace("/instance_resource_group",
			beego.NSInclude(
				&controllers.InstanceResourceGroupController{},
			),
		),

		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),

		beego.NSNamespace("/order_mq_message",
			beego.NSInclude(
				&controllers.OrderMqMessageController{},
			),
		),

		beego.NSNamespace("/product_info",
			beego.NSInclude(
				&controllers.ProductInfoController{},
			),
		),

		beego.NSNamespace("/project_info",
			beego.NSInclude(
				&controllers.ProjectInfoController{},
			),
		),

		beego.NSNamespace("/project_safe_info",
			beego.NSInclude(
				&controllers.ProjectSafeInfoController{},
			),
		),

		beego.NSNamespace("/report_daily",
			beego.NSInclude(
				&controllers.ReportDailyController{},
			),
		),

		beego.NSNamespace("/report_hdfs_detail",
			beego.NSInclude(
				&controllers.ReportHdfsDetailController{},
			),
		),

		beego.NSNamespace("/report_yarn_detail",
			beego.NSInclude(
				&controllers.ReportYarnDetailController{},
			),
		),

		beego.NSNamespace("/report_yarn_resource_log",
			beego.NSInclude(
				&controllers.ReportYarnResourceLogController{},
			),
		),

		beego.NSNamespace("/resource_common",
			beego.NSInclude(
				&controllers.ResourceCommonController{},
			),
		),

		beego.NSNamespace("/resource_end_point",
			beego.NSInclude(
				&controllers.ResourceEndPointController{},
			),
		),

		beego.NSNamespace("/trade_instance_resource",
			beego.NSInclude(
				&controllers.TradeInstanceResourceController{},
			),
		),

		beego.NSNamespace("/trade_order",
			beego.NSInclude(
				&controllers.TradeOrderController{},
			),
		),

		beego.NSNamespace("/trade_product_info",
			beego.NSInclude(
				&controllers.TradeProductInfoController{},
			),
		),

		beego.NSNamespace("/trade_sub_order",
			beego.NSInclude(
				&controllers.TradeSubOrderController{},
			),
		),

		beego.NSNamespace("/user_product_group_info",
			beego.NSInclude(
				&controllers.UserProductGroupInfoController{},
			),
		),

		beego.NSNamespace("/user_product_group_project_info",
			beego.NSInclude(
				&controllers.UserProductGroupProjectInfoController{},
			),
		),

		beego.NSNamespace("/user_product_order_info",
			beego.NSInclude(
				&controllers.UserProductOrderInfoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
