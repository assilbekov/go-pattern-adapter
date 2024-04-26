package data

type JSONDocument struct{}

func (doc JSONDocument) ConvertToXML() string {
	return "<xml>converted from JSON</xml>"
}
