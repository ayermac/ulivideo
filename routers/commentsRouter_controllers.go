package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"],
        beego.ControllerComments{
            Method: "CreateUploadVideo",
            Router: "/aliyun/create/upload/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"],
        beego.ControllerComments{
            Method: "RefreshUploadVideo",
            Router: "/aliyun/refresh/upload/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"],
        beego.ControllerComments{
            Method: "VideoCallback",
            Router: "/aliyun/video/callback",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:AliyunController"],
        beego.ControllerComments{
            Method: "GetPlayAuth",
            Router: "/aliyun/video/play/auth",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/barrage/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:BarrageController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "BarrageWs",
            Router: "/barrage/ws",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:BaseController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: "/channel/region",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:BaseController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: "/channel/type",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/comment/list",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/comment/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/es/add",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"],
        beego.ControllerComments{
            Method: "CreateIndex",
            Router: "/es/createIndex",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/es/delete",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: "/es/edit",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:EsDemoController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/es/search",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"],
        beego.ControllerComments{
            Method: "ChannelDemo",
            Router: "/go/channel/demo",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"],
        beego.ControllerComments{
            Method: "Demo",
            Router: "/go/demo",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"],
        beego.ControllerComments{
            Method: "SelectDemo",
            Router: "/go/select/demo",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:GoDemoController"],
        beego.ControllerComments{
            Method: "TaskDemo",
            Router: "/go/task/demo",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"],
        beego.ControllerComments{
            Method: "GetDirect",
            Router: "/mq/direct/push",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"],
        beego.ControllerComments{
            Method: "GetDlx",
            Router: "/mq/dlx/push",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"],
        beego.ControllerComments{
            Method: "GetTwoDlx",
            Router: "/mq/dlx/two/push",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"],
        beego.ControllerComments{
            Method: "GetFanout",
            Router: "/mq/fanout/push",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"],
        beego.ControllerComments{
            Method: "GetMq",
            Router: "/mq/push",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:MqDemoController"],
        beego.ControllerComments{
            Method: "GetTopic",
            Router: "/mq/topic/push",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:TopController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: "/channel/top",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:TopController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: "/type/top",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginDo",
            Router: "/login/do",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveRegister",
            Router: "/register/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:UserController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendMessageDo",
            Router: "/send/message",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: "/channel/advert",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelHotList",
            Router: "/channel/hot",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRecommendList",
            Router: "/channel/recommend/region",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetChannelRecommendTypeList",
            Router: "/channel/recommend/type",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: "/channel/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "UserVideo",
            Router: "/user/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: "/video/episodes/list",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: "/video/info",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoSave",
            Router: "/video/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/video/search",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"] = append(beego.GlobalControllerRouter["ulivideoapi/controllers:VideoController"],
        beego.ControllerComments{
            Method: "SendEs",
            Router: "/video/send/es",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
