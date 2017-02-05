package main

import (
	"goApiSample/api"
	"goApiSample/config"
	"goApiSample/model"
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	config.InitConfig("config/config.json")
	model.InitDB("mssql", config.Config.SampleMssql.Conn) //mssql,mysql
	//model.InitDB("mysql", config.Config.SampleMysql.Conn) //mssql,mysql
	//model.InitMssql("adodb", config.Config.Sample.Conn)

}

func InitApi(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello api")
	})

	v1 := e.Group("/v1")

	// GET /v1/fruits HTTP/1.1
	// Host: localhost:1111
	v1.GET("/fruits", api.Find)

	v1.GET("/fruits/:id", api.Get)
	// 	POST /v1/fruits HTTP/1.1
	// Host: localhost:1111
	// Content-Type: application/json
	// [{
	//     "Code":"A12",
	//     "Price":1,
	//     "Color":"Red",
	//     "Name":"Apple",
	//     "StoreCode":"AA01"
	// }]
	v1.POST("/fruits", api.Post)
	// PUT /v1/fruits HTTP/1.1
	// Host: localhost:1111
	// Content-Type: application/json
	// {
	//     "Code":"A1",
	//     "Price":2,
	//     "Color":"Red",
	//     "Name":"Apple",
	//     "StoreCode":"AA02"
	// }
	v1.PUT("/fruits", api.Put)
	v1.PATCH("/fruits", api.Patch)

}
