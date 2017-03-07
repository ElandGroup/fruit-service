package dto

import . "fruit-service/core"

type Fruit struct {
	Code      string  `json:"code",xorm:"varchar(20) pk notnull 'Code'"`
	Name      string  `json:"name",xorm:"'Name'"`
	Price     float32 `json:"price",xorm:"'Price'"`
	Color     string  `json:"color",xorm:"'Color'"`
	StoreCode string  `json:"storeCode",xorm:"'StoreCode'"`
}

type FruitQuery struct {
	Fruit
	APIParam
}
