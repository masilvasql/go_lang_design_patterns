package classes

import "fmt"

func UseReader(reader DataReader) {
	fmt.Println("Reading data with DataReader interface")
	data := reader.ReadData()
	fmt.Println(data)
}
