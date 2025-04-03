package classes

type XMLAdapter struct {
	legacyXML *LegacyXMLReader
}

func NewXMLAdapter(legacyXML *LegacyXMLReader) *XMLAdapter {
	return &XMLAdapter{legacyXML: legacyXML}
}

func (x *XMLAdapter) ReadData() string {
	return x.legacyXML.ReadXML()
}
