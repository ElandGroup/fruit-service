package dao

import "sync"

var fruitDaoSingleton *FruitDao
var onceFruitDao sync.Once

func GetFruitDao() *FruitDao {
	onceFruitDao.Do(func() {
		fruitDaoSingleton = &FruitDao{}
	})
	return fruitDaoSingleton
}
