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
