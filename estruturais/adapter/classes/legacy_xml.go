package classes

type LegacyXMLReader struct {
}

type XMLReader interface {
	ReadXML() string
}

// ReadData Código legado com interface diferente
func (l *LegacyXMLReader) ReadXML() string {
	return "<data>Reading data from XML file</data>"
}
