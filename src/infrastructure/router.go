package infrastructure

import (
	"cleanarchitecture/adapter/controller"
	"cleanarchitecture/adapter/interfaces"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type CustomContext struct {
	echo.Context
}

func Router(dbConn *gorm.DB) {
	e := echo.New()

	logger := &Logger{}

	e.Use( func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &CustomContext{c}
			return h(ctx)
		}
	})

	type ControllerFunc func(c interfaces.IContext) error
	POST := func(path string, f ControllerFunc) *echo.Route {
		return e.POST(path, func(c echo.Context) error {
			return f(c.(*CustomContext))
		})
	}

	userController := controller.NewUserController(dbConn, logger)


	POST("/users", userController.Create)

	err := e.Start(viper.GetString(`server.address`))
	if err != nil {
		panic(err)
	}

}

