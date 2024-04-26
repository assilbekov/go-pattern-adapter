package data

import "fmt"

type JSONDocument struct{}

func (doc JSONDocument) ConvertToXML(data string) string {
	return fmt.Sprintf("<xml>%s</xml>", data)
}

type JSONDocumentAdapter struct {
	JSONDocument *JSONDocument
}

func (adapter JSONDocumentAdapter) SendXMLData(data string) interface{} {
	return adapter.JSONDocument.ConvertToXML(data)
}
