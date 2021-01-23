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

}
