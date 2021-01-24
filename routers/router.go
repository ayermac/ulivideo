// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"ulivideoapi/controllers"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.ObjectController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)

	beego.Include(
		&controllers.UserController{},
	)
	beego.Include(
		&controllers.VideoController{},
	)
	beego.Include(
		&controllers.BaseController{},
	)
	beego.Include(
		&controllers.CommentController{},
	)
	beego.Include(
		&controllers.TopController{},
	)
	beego.Include(
		&controllers.BarrageController{},
	)
	beego.Include(
		&controllers.AliyunController{},
	)
	beego.Include(
		&controllers.MqController{},
	)
}
