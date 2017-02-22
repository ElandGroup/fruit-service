package main

import (
	"goApiSample/sample"

	"github.com/labstack/echo"
)

func main() {
	sample.Sampleinit()
	e := echo.New()
	InitApi(e)
	e.Start(":5000")
}
