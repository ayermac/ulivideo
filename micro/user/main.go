package main

import (
	"time"
	_ "ulivideoapi/routers"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"


	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	defaultdb, _ := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", defaultdb)
	if err != nil {
		panic(err)
	}

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.ulivideo.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	proto.RegisterUserServiceHandler(service.Server(), new(controllers.UserRpcController))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
