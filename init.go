package main

import (
	"fruit-service/api"
	"fruit-service/config"
	"fruit-service/core/helper"
	"net/http"

	"fruit-service/dao"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	config.InitConfig("config/config.json")
	dao.InitDB("mssql", config.Config.SampleMssql.Conn) //mssql,mysql
	//model.InitDB("mysql", config.Config.SampleMysql.Conn) //mssql,mysql
	//model.InitMssql("adodb", config.Config.Sample.Conn)

}

func InitApi(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello api")
	})

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			helper.Lang = c.Request().Header["Accept-Language"]
			return next(c)
		}
	}

	v1 := e.Group("/v1", track)

	v1.GET("/fruits", api.Find)
	v1.GET("/fruits/:Code", api.Get)
	v1.POST("/fruits", api.Post)
	v1.PATCH("/fruits/:Code", api.Patch)
	v1.DELETE("/fruits/:Code", api.Delete)

}
