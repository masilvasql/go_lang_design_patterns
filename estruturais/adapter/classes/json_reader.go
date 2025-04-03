package classes

type JSONReader struct {
}

// ReadData Implementação compatível
func (j *JSONReader) ReadData() string {
	return "Reading data from JSON file"
}
