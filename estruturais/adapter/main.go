package main

import "github.com/masilvasq/adapter/classes"

func main() {
	// Leitor JSON direto
	jsonReader := &classes.JSONReader{}
	classes.UseReader(jsonReader)
	// Leitor XML adaptado
	legacyXML := &classes.LegacyXMLReader{}
	xmlAdapter := classes.NewXMLAdapter(legacyXML)
	classes.UseReader(xmlAdapter)
}
