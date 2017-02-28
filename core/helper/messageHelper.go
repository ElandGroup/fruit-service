package helper

import (
	"encoding/json"
	"fmt"
	. "fruit-service/core"
	"reflect"
	"strconv"

	"errors"
)

//private,don't change public
func MessageString(resourceKey int, params ...interface{}) string {
	if len(params) == 0 {
		return MessageMap[resourceKey]
	} else {
		return fmt.Sprintf(MessageMap[resourceKey], params...)
	}
}

//success:false,details:detail,message:10001
func SystemMessage(detail string) *APIResult {
	return NewApiMessage(10001, detail)
}

func NewApiError(resourceKey int, details string, params ...interface{}) (apiError *APIError) {
	return &APIError{Code: resourceKey, Message: MessageString(resourceKey, params...), Details: details}
}

func NewApiMessage(resourceKey int, details string, params ...interface{}) *APIResult {
	return &APIResult{Success: false, Error: NewApiError(resourceKey, details, params...)}
}

func ConvJson(anyObject interface{}) (result string, err error) {

	val := reflect.ValueOf(anyObject).Elem()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		result = val.String()
	case reflect.Slice, reflect.Map, reflect.Struct, reflect.Array:
		var bytevv []byte
		bytevv, err = json.Marshal(val)
		result = string(bytevv)
	default:
		result = ""
		err = errors.New("Type is not recognized")
	}
	return
}
