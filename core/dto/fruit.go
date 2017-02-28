package dto

import . "fruit-service/core"

type Fruit struct {
	Code      string  `xorm:"varchar(20) pk notnull 'Code'"`
	Name      string  `xorm:"'Name'"`
	Price     float32 `xorm:"'Price'"`
	Color     string  `xorm:"'Color'"`
	StoreCode string  `xorm:"'StoreCode'"`
}

type FruitQuery struct {
	Fruit
	APIParam
}
