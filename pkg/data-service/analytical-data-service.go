package data_service

type AnalyticalDataService interface {
	SendXMLData(data string) interface{}
}

type XMLDocument struct {
	XMLData string
}

func (x *XMLDocument) SendXMLData(data string) interface{} {
	x.XMLData = data
	return x.XMLData
}
