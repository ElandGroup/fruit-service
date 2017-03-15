package dao

import . "fruit-service/core/dto"

type FruitDao struct{}

func (f FruitDao) Find(fruitQuery *FruitQuery) (fruits []Fruit, err error) {
	session := Filter(&fruitQuery.APIParam)
	err = session.Find(&fruits)
	return
}

func (FruitDao) Get(code string) (has bool, fruit *Fruit, err error) {
	fruit = &Fruit{Code: code}
	has, err = Db.Get(fruit)
	return
}
func (FruitDao) Exists(keys []string) (existKeys []string, err error) {
	err = Db.Select("Code").In("Code", keys).Iterate(new(Fruit), func(i int, bean interface{}) error {
		fruit := bean.(*Fruit)
		existKeys = append(existKeys, fruit.Code)
		return nil
	})
	return
}

func (FruitDao) Post(fruit *[]Fruit) (affectedRows int64, err error) {
	return Db.Insert(fruit)
}

func (FruitDao) Patch(code string, fruit *Fruit) (affectedrow int64, err error) {
	return Db.Cols("Price").Where("code=?", code).Update(fruit)
}

func (FruitDao) Delete(code string) (affectedrow int64, err error) {
	fruit := &Fruit{Code: code}
	return Db.Where("code=?", code).Delete(fruit)
}
