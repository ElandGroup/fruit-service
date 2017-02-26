package helper

import (
	. "goApiSample/core"
	"reflect"
	"strings"

	reflections "gopkg.in/oleiade/reflections.v1"
)

func CheckQueryCondition(dto *APIParam, checkStruct interface{}) (apiResult *APIResult) {
	var err error
	if len(dto.Fields) != 0 {
		var ok bool
		fields := strings.Split(dto.Fields, ",")
		for _, v := range fields {
			if ok, err = reflections.HasField(checkStruct, v); err != nil || ok == false {
				dto = nil
				apiResult = NewApiMessage(10011, "", v)
				return
			}
		}
	}
	if dto.MaxResultCount == 0 {
		dto.MaxResultCount = 10
	} else if dto.MaxResultCount > 10000 {
		dto.MaxResultCount = 10000
	}

	if len(dto.Sort) != 0 {
		sorts := strings.Split(dto.Sort, ",")
		for _, v := range sorts {
			if strings.HasPrefix(v, "-") {
				v = string(v[1:])
				dto.SortDesc += "," + v
			} else {
				dto.SortAsc += "," + v
			}

			if ok, err := reflections.HasField(checkStruct, v); err != nil || ok == false {
				dto = nil
				apiResult = NewApiMessage(10011, "", v)
				return
			}
		}
		if len(dto.SortDesc) != 0 {
			dto.SortDesc = dto.SortDesc[1:]
		}
		if len(dto.SortAsc) != 0 {
			dto.SortAsc = dto.SortAsc[1:]
		}
	}
	apiResult = nil
	return
}

func FilterFieldsMap(objArray interface{}, fields []string) []map[string]interface{} {
	slice, _ := takeArg(objArray, reflect.Slice)
	returnMaps := []map[string]interface{}{}
	for i := 0; i < slice.Len(); i++ {
		returnMap := make(map[string]interface{})
		for _, k := range fields {
			returnMap[k], _ = reflections.GetField(slice.Index(i).Interface(), k)
		}
		returnMaps = append(returnMaps, returnMap)
	}
	return returnMaps
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
