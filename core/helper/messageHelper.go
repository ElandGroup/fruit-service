package helper

import (
	"encoding/json"
	"fmt"
	. "fruit-service/core"
	"reflect"
	"strconv"
	"strings"

	"errors"
)

// Parse accept-language in header to convert it as: tw, en, jp ...
func ParseAcceptLang(acceptLang string) string {
	// 1. Chrome: [zh-TW,zh;q=0.8,en-US;q=0.6,en;q=0.4]
	// 2. Safari: [zh-tw]
	// 3. FF: [zh-TW,zh;q=0.8,en-US;q=0.5,en;q=0.3]
	//
	// Ret: zh or en ...
	tarStrings := strings.Split(acceptLang, ";")
	if len(strings.Split(tarStrings[0], ",")) > 1 {
		return strings.Split(tarStrings[0], ",")[1]
	}
	return strings.Split(tarStrings[0], "-")[0]
}

var Lang []string

func MessageString(resourceKey int, params ...interface{}) string {
	langStr := strings.Join(Lang[:], ";")
	currntLang := ParseAcceptLang(langStr)
	if len(currntLang) != 2 {
		currntLang = "en"
	}
	if len(params) == 0 {
		return MessageMap[currntLang][resourceKey]
	} else {
		return fmt.Sprintf(MessageMap[currntLang][resourceKey], params...)
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
