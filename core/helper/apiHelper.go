package helper

import (
	"strings"

	reflections "gopkg.in/oleiade/reflections.v1"
)

func CheckFieldParams(fields string, checkObject interface{}) (errParam string) {
	//2.check field format
	if len(fields) != 0 {
		fields := strings.Split(fields, ",")
		for _, v := range fields {
			if ok, err := reflections.HasField(checkObject, v); err != nil || ok == false {
				errParam += "," + v
			}
		}
		if len(errParam) != 0 {
			errParam = errParam[1:]
		}
	}
	return
}
