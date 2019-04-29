package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ClusterConfigController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"],
        beego.ControllerComments{
            Method: "GetByPage",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:InstanceResourceGroupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:OrderMqMessageController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProductInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ProjectSafeInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportDailyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportHdfsDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnDetailController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ReportYarnResourceLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceCommonController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:ResourceEndPointController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeInstanceResourceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeOrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeProductInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:TradeSubOrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductGroupProjectInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"] = append(beego.GlobalControllerRouter["rs_web_go/controllers:UserProductOrderInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
