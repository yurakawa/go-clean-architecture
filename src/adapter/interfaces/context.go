package interfaces

import "github.com/labstack/echo"

type IContext interface {
	echo.Context
	// Param(string) string
	// Bind(interface{}) error
	// Status(int)
	// JSON(int, interface{})
}
