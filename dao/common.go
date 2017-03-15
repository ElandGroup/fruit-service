package dao

import (
	. "fruit-service/core"

	"github.com/go-xorm/xorm"
)

func Filter(dto *APIParam) (session *xorm.Session) {
	sqlText := ""
	if len(dto.Fields) != 0 {
		sqlText += dto.Fields
	} else {
		sqlText += `
			*
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
