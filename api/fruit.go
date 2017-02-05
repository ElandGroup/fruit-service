package api

import (
	"fmt"
	"net/http"

	. "goApiSample/model"
	"goApiSample/service"

	"github.com/labstack/echo"
)

func Find(c echo.Context) error {
	//fruit := new(Fruit)
	if fruits, err := service.GetFruitService().Find(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10002, Message: err.Error()}})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: QueryResult{TotalCount: len(fruits), Items: fruits}})
	}
}

func Get(c echo.Context) error {
	//fruit := new(Fruit)
	fmt.Println(c.Param("id"))
	if code := c.Param("id"); len(code) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10001, Message: "A required parameter is missing or doesn't have the right format:" + "code"}})
	} else if fruits, err := service.GetFruitService().Get(code); err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10002, Message: err.Error()}})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: fruits})
	}
}

func Post(c echo.Context) error {
	//fruit := new(Fruit)
	fruitsParam := new([]Fruit)
	if err := c.Bind(fruitsParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10001, Message: "A required parameter is missing or doesn't have the right format"}})
	} else if affectedRows, err := service.GetFruitService().Post(fruitsParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10002, Message: err.Error()}})
	} else if affectedRows == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10003, Message: "no row is affected"}})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: affectedRows})
	}
}

func Put(c echo.Context) error {
	//fruit := new(Fruit)
	fruitsParam := new(Fruit)
	if err := c.Bind(fruitsParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10001, Message: "A required parameter is missing or doesn't have the right format"}})
	} else if affectedRows, err := service.GetFruitService().Put(fruitsParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10002, Message: err.Error()}})
	} else if affectedRows == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10003, Message: "no row is affected"}})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: affectedRows})
	}
}

func Patch(c echo.Context) error {
	//fruit := new(Fruit)
	fruitsParam := new(Fruit)
	if err := c.Bind(fruitsParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10001, Message: "A required parameter is missing or doesn't have the right format"}})
	} else if affectedRows, err := service.GetFruitService().Patch(fruitsParam.Code, fruitsParam); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10002, Message: err.Error()}})
	} else if affectedRows == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, APIResult{Success: false, Error: APIError{Code: 10003, Message: "no row is affected"}})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: affectedRows})
	}
}
