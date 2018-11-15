package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		return c.JSON(200, m)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
