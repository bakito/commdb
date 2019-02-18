package main

import (
	"github.com/bakito/commdb/controller"
	"github.com/bakito/commdb/types"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := iris.New()

	orm, err := xorm.NewEngine("sqlite3", "./commdb.db")
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

	mvc.Configure(app.Party("/command"), func(app *mvc.Application) {
		app.Handle(controller.New(orm))
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
