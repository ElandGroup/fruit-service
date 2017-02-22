package service

import (
	"fmt"
	. "goApiSample/core"
	. "goApiSample/core/dto"
	"goApiSample/core/helper"
	"goApiSample/dao"
	"strings"
)

type FruitService struct {
}

func (FruitService) Find(dto *FruitQuery) (fruits []Fruit, err error) {
	//2.parse apiParam
	if dto.MaxResultCount == 0 {
		dto.MaxResultCount = 10
	} else if dto.MaxResultCount > 100 {
		dto.MaxResultCount = 100
	}
	if len(dto.Sort) != 0 {
		sortTemp := strings.Split(dto.Sort, ",")
		for _, v := range sortTemp {
			if v[0:1] == "-" {
				dto.SortDesc += "," + v[1:]
			} else {
				dto.SortAsc += "," + v
			}
		}
		if len(dto.SortDesc) > 0 {
			dto.SortDesc = dto.SortDesc[1:]
		}
		if len(dto.SortAsc) > 0 {
			dto.SortAsc = dto.SortAsc[1:]
		}
	}
	return dao.GetFruitDao().Find(dto)
}

func (FruitService) Get(code string) (has bool, fruit *Fruit, err error) {
	return dao.GetFruitDao().Get(code)
}

func (f *FruitService) Post(fruit *[]Fruit) (affectedRows int64, apiError *APIError) {
	keys := []string{}
	for _, v := range *fruit {
		keys = append(keys, v.Code)
	}
	dupplicationData, e := f.Exists(keys)
	if e != nil {
		return 0, helper.NewApiError(10001, e.Error())
	} else if len(dupplicationData) != 0 {
		fmt.Println(strings.Join(dupplicationData[:], ","))
		return 0, helper.NewApiError(10012, "", strings.Join(dupplicationData[:], ","))
	}
	var err error
	affectedRows, err = dao.GetFruitDao().Post(fruit)
	if err != nil {
		affectedRows = 0
		apiError = helper.NewApiError(10001, "", err.Error())
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
