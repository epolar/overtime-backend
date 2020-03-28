package user

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"orderStatistics/config"
)

func Run(addr string) error {
	app := iris.New()
	app.Logger().SetLevel("debug")

	corsConfig := config.Config().CorsConfig()
	app.Use(cors.New(cors.Options{
		AllowedOrigins: corsConfig.Origins,
		AllowedMethods: corsConfig.Methods,
		AllowedHeaders: corsConfig.Headers,
	}))
	// enable options methods
	mvc.Configure(app.AllowMethods(iris.MethodOptions))

	app.Use(recover2.New())
	app.Use(logger.New())

	mvcApp := mvc.New(app)

	mvcApp.Party("user").Configure(userControl)
	mvcApp.Party("overtime").Configure(overtimeControl)

	return app.Run(iris.Addr(addr))
}

func userControl(app *mvc.Application) {
	app.Handle(new(PersonController))
}

func overtimeControl(app *mvc.Application) {
	app.Handle(new(OvertimeController))
}

func JSON(ctx iris.Context, v interface{}) error {
	_, err := ctx.JSON(v)
	if err != nil {
		return err
	}
	return nil
}
