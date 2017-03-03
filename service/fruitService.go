package service

import (
	. "fruit-service/core"
	. "fruit-service/core/dto"
	"fruit-service/core/helper"
	"fruit-service/dao"
	"net/http"
	"strings"
)

type FruitService struct {
}

func (FruitService) Find(dto *FruitQuery) (fruits []Fruit, err error) {
	return dao.GetFruitDao().Find(dto)
}

func (FruitService) Get(code string) (has bool, fruit *Fruit, err error) {
	return dao.GetFruitDao().Get(code)
}

func (f *FruitService) Post(fruit *[]Fruit) (apiMessage *APIStatusMessage) {
	keys := []string{}
	for _, v := range *fruit {
		keys = append(keys, v.Code)
	}
	dupplicationData, e := f.Exists(keys)
	if e != nil {
		apiMessage = helper.NewApiStatusMessage(http.StatusInternalServerError, 10001, e.Error())
		return
	} else if len(dupplicationData) != 0 {
		apiMessage = helper.NewApiStatusMessage(http.StatusOK, 10013, "", helper.MessageString(10012, strings.Join(dupplicationData[:], ",")))
		return
	}
	affectedRows, err := dao.GetFruitDao().Post(fruit)
	if err != nil {
		apiMessage = helper.NewApiStatusMessage(http.StatusInternalServerError, 10001, "", err.Error())
	} else if affectedRows == 0 {
		apiMessage = helper.NewApiStatusMessage(http.StatusOK, 10013, "", helper.MessageString(10007))
	} else {
		apiMessage = nil
	}
	return
}

func (FruitService) Patch(code string, fruit *Fruit) (affectedrow int64, err error) {
	return dao.GetFruitDao().Patch(code, fruit)
}
func (f *FruitService) Delete(code string) (apiMessage *APIStatusMessage) {
	var affectedrow int64
	var err error
	var has bool
	if has, _, err = f.Get(code); err != nil {
		apiMessage = helper.NewApiStatusMessage(http.StatusInternalServerError, 10001, err.Error())
		return
	} else if has == false {
		apiMessage = helper.NewApiStatusMessage(http.StatusOK, 10013, "", helper.MessageString(10005, "Fruit"))
		return
	}

	if affectedrow, err = dao.GetFruitDao().Delete(code); err != nil {
		apiMessage = helper.NewApiStatusMessage(http.StatusInternalServerError, 10001, err.Error())
		return
	} else if affectedrow == 0 {
		apiMessage = helper.NewApiStatusMessage(http.StatusOK, 10013, "", helper.MessageString(10007))
		return
	}

	return
}

func (FruitService) Exists(keys []string) (existKeys []string, err error) {
	return dao.GetFruitDao().Exists(keys)
}
