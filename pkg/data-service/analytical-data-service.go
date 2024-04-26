package data_service

import "fmt"

type AnalyticalDataService interface {
	SendXMLData(data string) interface{}
}

type XMLDocument struct {
	XMLData string
}

func (x *XMLDocument) SendXMLData(data string) interface{} {
	x.XMLData = data
	fmt.Println("Data sent: ", x.XMLData)
	return x.XMLData
}
