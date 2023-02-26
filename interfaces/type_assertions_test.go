package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetDeveloper(t *testing.T) {
	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	assert.Equal(t, name.(string), dynamicDev.Name)
	assert.Equal(t, age.(int), dynamicDev.Age)
}

func Test_InterfaceImplementation(t *testing.T) {

	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	programmers = append(programmers, elliot)

	assert.Equal(t, "Elliot programs in Go", programmers[0].Language())
	assert.NotNil(t, programmers[0].Age())
	assert.NotEmpty(t, programmers)
	assert.Equal(t, len(programmers), 1)
}
