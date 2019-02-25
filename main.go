package main

import (
	"os"

	"github.com/bakito/commdb/controller"
	"github.com/bakito/commdb/service"
	"github.com/bakito/commdb/types"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	_ "github.com/lib/pq"
)

func main() {
	app := iris.New()

	orm, err := xorm.NewEngine("postgres", os.Getenv("DB_URL"))
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}

	iris.RegisterOnInterrupt(func() {
		orm.Close()
	})

	err = orm.Sync2(new(types.Command))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized Command table: %v", err)
	}

	mvc.Configure(app.Party("/api/command"), func(app *mvc.Application) {
		app.Register(service.NewCommandService(orm))
		app.Handle(controller.New())
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
