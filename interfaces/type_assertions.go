package interfaces

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
