package service

import (
	. "goApiSample/model"
	"sync"
)

type FruitService struct {
}

func (FruitService) Find() (fruits []*Fruit, err error) {
	return GetFruit().Find()
}

func (FruitService) Get(code string) (fruit *Fruit, err error) {
	return GetFruit().Get(code)
}

func (FruitService) Post(fruit *[]Fruit) (affectedRows int64, err error) {
	return GetFruit().Post(fruit)
}

func (FruitService) Put(fruit *Fruit) (affectedrow int64, err error) {

	return GetFruit().Put(fruit)
}

func (FruitService) Patch(code string, fruit *Fruit) (affectedrow interface{}, err error) {
	return GetFruit().Patch(code, fruit)

	//return model.Db.Id(fruit.Code).Update(fruit)
}

var fruitServiceSingleton *FruitService
var once sync.Once

func GetFruitService() *FruitService {
	once.Do(func() {
		fruitServiceSingleton = &FruitService{}
	})
	return fruitServiceSingleton
}
