package sample

import "encoding/xml"

type Person struct {
	Name string `xml:"name"`
}

type ResponseDtoNew struct {
	Person Person `xml:"person"`
}

type ResponseDto struct {
	XMLName xml.Name         `xml:"response"`
	Success bool             `xml:"success"`
	Persons []ResponseDtoNew `xml:"persons"`
}
