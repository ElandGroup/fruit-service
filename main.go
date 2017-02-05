package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	InitApi(e)
	e.Start(":1111")
}
