package api

import (
	"goApiSample/core/helper"
	"net/http"

	. "goApiSample/core"
	. "goApiSample/core/dto"
	"goApiSample/service"

	"strings"

	"github.com/labstack/echo"
)

func Find(c echo.Context) error {
	dto := FruitQuery{}
	//1.check request format
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10004, err.Error(), "Object"))
	}
	//1.check condition
	if apiMessage := helper.CheckQueryCondition(&dto.APIParam, Fruit{}); apiMessage != nil {
		return c.JSON(http.StatusBadRequest, apiMessage)
	}

	//3.Query data
	if fruits, err := service.GetFruitService().Find(&dto); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.SystemMessage(err.Error()))
	} else {
		if len(fruits) == 0 {
			return c.JSON(http.StatusOK, APIResult{Success: true, Result: QueryResult{TotalCount: len(fruits), Items: make([]interface{}, 0)}})
		}
		if len(dto.Fields) != 0 {
			fields := strings.Split(dto.Fields, ",")
			return c.JSON(http.StatusOK, APIResult{Success: true, Result: QueryResult{TotalCount: len(fruits), Items: helper.FilterFieldsMap(fruits, fields)}})
		} else {
			return c.JSON(http.StatusOK, APIResult{Success: true, Result: QueryResult{TotalCount: len(fruits), Items: fruits}})
		}
	}
}

func Get(c echo.Context) error {
	if code := c.Param("Code"); len(code) == 0 {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10009, "", "Code"))

	} else if _, fruits, err := service.GetFruitService().Get(code); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.SystemMessage(err.Error()))
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: fruits})
	}
}

func Post(c echo.Context) error {
	fruitsParam := new([]Fruit)
	if err := c.Bind(fruitsParam); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10004, err.Error(), "Array"))
	} else {
		if apiError := service.GetFruitService().Post(fruitsParam); apiError != nil {
			if apiError.Code == 10001 {
				return c.JSON(http.StatusInternalServerError, apiError)
			} else {
				return c.JSON(http.StatusOK, apiError)
			}
		} else {
			return c.JSON(http.StatusCreated, APIResult{Success: true, Result: nil})
		}
	}
}

func Patch(c echo.Context) error {
	var code string
	if code = c.Param("Code"); len(code) == 0 {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10009, "", "Code"))
	}
	fruitsParam := new(Fruit)
	if err := c.Bind(fruitsParam); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10004, err.Error(), "Object"))
	} else {

		if affectedRows, e := service.GetFruitService().Patch(code, fruitsParam); e != nil {
			return c.JSON(http.StatusInternalServerError, helper.SystemMessage(e.Error()))
		} else if affectedRows == 0 {
			return c.JSON(http.StatusNotFound, helper.NewApiMessage(10010, "", "Code:"+code))
		} else {
			return c.JSON(http.StatusNoContent, APIResult{Success: true, Result: nil})
		}
	}
}

func Delete(c echo.Context) error {
	var code string
	if code = c.Param("Code"); len(code) == 0 {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10009, "", "Code"))
	} else {
		if apiError := service.GetFruitService().Delete(code); apiError != nil {
			if apiError.Code == 10001 {
				return c.JSON(http.StatusInternalServerError, apiError)
			} else if apiError.Code == 10010 {
				return c.JSON(http.StatusNotFound, apiError)
			} else {
				return c.JSON(http.StatusOK, apiError)
			}
		} else {
			return c.JSON(http.StatusNoContent, APIResult{Success: true, Result: nil})
		}
	}
}
