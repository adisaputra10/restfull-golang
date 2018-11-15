package page

import (
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World! hai ")
	return c.JSON(200, "GET Tasks")
}
