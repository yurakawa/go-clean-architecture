package infrastructure

import (
	"cleanarchitecture/adapter/controller"
	"cleanarchitecture/adapter/interfaces"
	"cleanarchitecture/infrastructure/driver/mysql"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
}

func Router() {
	e := echo.New()

	logger := &Logger{}

	conn := mysql.Connect()

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

	userController := controller.NewUserController(conn, logger)


	POST("/users", userController.Create)

	err := e.Start(":8000")
	if err != nil {
		panic(err)
	}

}

