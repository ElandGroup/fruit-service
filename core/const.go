package core

//resource

var MessageMap = map[string]map[int]string{
	"en": {
		10001: "System error",
		10002: "A required parameter is missing or doesn't have the right format:%v",
		10003: "%v is required",
		10004: "The parameter format should be %v",
		10005: "%v: Not Exist",
		10006: "%v failure",
		10007: "No row is affected",
		10008: "%v has a wrong format",
		10009: "Routing is missing parameter: %v",

		10011: "Please check the parameters contained in Fields:%v",
		10012: "There is duplicate data:%v",
		10013: "Save failed because:%v"},
	"zh": {
		10001: "系统错误",
		10002: "缺少必要的参数，或者参数格式不正确：%v",
		10003: "%v 不能为空",
		10004: "参数格式应该是 %v",
		10005: "%v：不存在",
		10006: "%v 失败",
		10007: "没有数据被改变",
		10008: "%v 有一个错误的格式",
		10009: "路由缺少参数：%v",

		10011: "请检查Fields所包含的参数：%v",
		10012: "存在重复的数据：%v",
		10013: "保存失败，原因是：%v"},
	"ko": {
		10001: "시스템 오류입니다.",
		10002: "필요되는 파라미터가 없거나 파라미터 포맷이 정확하지 않습니다:%v",
		10003: "%v 가 빈값이면 안됩니다.",
		10004: "파마미터 포맷은 %v 입니다.",
		10005: "%v：존재하지 않습니다.",
		10006: "%v 실패하였습니다.",
		10007: "변경된 데이터가 없습니다.",
		10008: "%v 에 오류 포맷이 존재합니다.",
		10009: "루팅에 파라미터가 부족합니다. %v",

		10011: "Fields에 포함된 파라미터 %v를 점검하세요.",
		10012: "중복된 데이터가 존재합니다. %v",
		10013: "저장실패하였습니다. 원인:%v"},
}
