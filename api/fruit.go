package api

import (
	"fruit-service/core/helper"
	"net/http"

	. "fruit-service/core"
	. "fruit-service/core/dto"
	"fruit-service/service"

	"strings"

	"github.com/labstack/echo"
)

func Find(c echo.Context) error {
	dto := FruitQuery{}
	//1.check request format
	if err := helper.Bind(&dto, c); err != nil {
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
	if code := helper.Param("Code", c); len(code) == 0 {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10009, "", "Code"))

	} else if has, fruits, err := service.GetFruitService().Get(code); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.SystemMessage(err.Error()))
	} else {
		if has == false {
			fruits = nil
		}
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: fruits})
	}
}

func Post(c echo.Context) error {
	fruitsParam := new([]Fruit)
	if err := helper.Bind(fruitsParam, c); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10004, err.Error(), "Array"))
	} else {
		if apiMessage := service.GetFruitService().Post(fruitsParam); apiMessage != nil {
			return c.JSON(apiMessage.StatusCode, apiMessage.APIResult)
		} else {
			return c.JSON(http.StatusCreated, APIResult{Success: true, Result: nil})
		}
	}
}

func Patch(c echo.Context) error {
	var code string
	if code = helper.Param("Code", c); len(code) == 0 {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10009, "", "Code"))
	}
	fruitsParam := new(Fruit)
	if err := helper.Bind(fruitsParam, c); err != nil {
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
	if code := helper.Param("Code", c); len(code) == 0 {
		return c.JSON(http.StatusBadRequest, helper.NewApiMessage(10009, "", "Code"))
	} else {
		if apiMessage := service.GetFruitService().Delete(code); apiMessage != nil {
			return c.JSON(apiMessage.StatusCode, apiMessage.APIResult)
		} else {
			return c.JSON(http.StatusNoContent, APIResult{Success: true, Result: nil})
		}
	}
}
