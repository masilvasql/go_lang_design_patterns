package main

import "sync"

type Singleton struct {
	url string
}

var instance *Singleton
var once sync.Once // Garante que a inicialização ocorra apenas uma vez

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {

	s1 := GetInstance()
	s1.url = "http://www.google.com"

	s2 := GetInstance()
	println(s2.url)

	s2.url = "http://www.facebook.com"

	println(s1.url)
	println(s2.url)

}
