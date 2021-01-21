package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

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

}
