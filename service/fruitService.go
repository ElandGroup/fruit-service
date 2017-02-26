package service

import (
	. "goApiSample/core"
	. "goApiSample/core/dto"
	"goApiSample/core/helper"
	"goApiSample/dao"
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

func (f *FruitService) Post(fruit *[]Fruit) (apiError *APIError) {
	keys := []string{}
	for _, v := range *fruit {
		keys = append(keys, v.Code)
	}
	dupplicationData, e := f.Exists(keys)
	if e != nil {
		apiError = helper.NewApiError(10001, e.Error())
		return
	} else if len(dupplicationData) != 0 {
		apiError = helper.NewApiError(10012, "", strings.Join(dupplicationData[:], ","))
		return
	}
	affectedRows, err := dao.GetFruitDao().Post(fruit)
	if err != nil {
		apiError = helper.NewApiError(10001, "", err.Error())
	} else if affectedRows == 0 {
		apiError = helper.NewApiError(10007, "")
	} else {
		apiError = nil
	}
	return
}

func (FruitService) Patch(code string, fruit *Fruit) (affectedrow int64, err error) {
	return dao.GetFruitDao().Patch(code, fruit)

	//return model.Db.Id(fruit.Code).Update(fruit)
}

func (FruitService) Exists(keys []string) (existKeys []string, err error) {
	return dao.GetFruitDao().Exists(keys)
}
