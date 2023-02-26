package interfaces

import "math/rand"

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	dev := Developer{
		Name: name.(string),
		Age:  age.(int),
	}
	return dev
}

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return int(rand.Int63())
}
