package sample

import (
	"encoding/xml"
	"fmt"
)

func Sampleinit() {
	buf := []byte(`<response>
     <success>true</success>
    <persons>
    <person>
    <name>xiao</name>
    </person>
    </persons>
    </response>`)
	u2 := ResponseDto{}
	err := xml.Unmarshal(buf, &u2)
	fmt.Println(err)
	fmt.Printf("%+v\n", u2)
}
