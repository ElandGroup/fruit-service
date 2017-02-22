package service

import "sync"

var fruitServiceSingleton *FruitService
var onceFruitService sync.Once

func GetFruitService() *FruitService {
	onceFruitService.Do(func() {
		fruitServiceSingleton = &FruitService{}
	})
	return fruitServiceSingleton
}
