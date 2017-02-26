package dao

import (
	. "goApiSample/core/dto"

	"github.com/go-xorm/xorm"
)

type FruitDao struct{}

func (f FruitDao) Find(fruitQuery *FruitQuery) (fruits []Fruit, err error) {
	session := f.Query(fruitQuery)
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

func (FruitDao) Query(dto *FruitQuery) (session *xorm.Session) {
	sqlText := ""
	if len(dto.Fields) != 0 {
		// fields := strings.Split(dto.Fields, ",")
		// newFields := ""
		// for _, v := range fields {
		// 	newFields += v + ","
		// }
		// if len(newFields) != 0 {
		// 	newFields = newFields[0 : len(newFields)-1]
		// }
		sqlText += dto.Fields
	} else {
		sqlText += `
			Code     
			,Name     
			,Price     
			,Color    
			,StoreCode
		`
	}
	session = Db.Select(sqlText).Limit(dto.MaxResultCount, dto.SkipCount)
	if len(dto.SortAsc) != 0 {
		session.Asc(dto.SortAsc)
	}
	if len(dto.SortDesc) != 0 {
		session.Desc(dto.SortDesc)
	}
	return
}
