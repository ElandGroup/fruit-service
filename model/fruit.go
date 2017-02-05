package model

import "fmt"
import "sync"

type Fruit struct {
	Code      string  `xorm:"varchar(20) pk notnull 'Code'"`
	Name      string  `xorm:"'Name'"`
	Price     float32 `xorm:"'Price'"`
	Color     string  `xorm:"'Color'"`
	StoreCode string  `xorm:"'StoreCode'"`
}

func (Fruit) Find() (fruits []*Fruit, err error) {
	fmt.Println("======", Db)
	err = Db.Find(&fruits)
	return
}

func (Fruit) Get(code string) (fruit *Fruit, err error) {
	f := &Fruit{Code: code}
	_, err = Db.Get(f)
	fruit = f
	return
}

func (Fruit) Post(fruit *[]Fruit) (affectedRows int64, err error) {
	return Db.Insert(fruit)
}

func (Fruit) Put(fruit *Fruit) (affectedrow int64, err error) {
	code := fruit.Code
	fruit.Code = ""
	return Db.Id(code).Update(fruit)
}

func (Fruit) Patch(code string, fruit *Fruit) (affectedrow interface{}, err error) {
	sql := "update `fruit` set Color=? where code=?"
	return Db.Exec(sql, fruit.Color, code)

	//return model.Db.Id(fruit.Code).Update(fruit)
}

var fruitSingleton *Fruit
var once sync.Once

func GetFruit() *Fruit {
	once.Do(func() {
		fruitSingleton = &Fruit{}
	})
	return fruitSingleton
}
