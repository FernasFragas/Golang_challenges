package sequences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	example0 = []int{4, 1, 3, 2}
	example1 = []int{100, 4, 200, 1, 3, 2}
	example2 = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
)

func Test_longestConsecutive(t *testing.T) {
	ex0 := longestConsecutive(example0)
	assert.Equal(t, 4, ex0)
	ex1 := longestConsecutive(example1)
	assert.Equal(t, 4, ex1)
	ex2 := longestConsecutive(example2)
	assert.Equal(t, 9, ex2)
}

func Test_longestConsecutiveHash(t *testing.T) {
	ex0 := longestConsecutiveHash(example0)
	assert.Equal(t, 4, ex0)
	ex1 := longestConsecutiveHash(example1)
	assert.Equal(t, 4, ex1)
	ex2 := longestConsecutiveHash(example2)
	assert.Equal(t, 9, ex2)
}

/*
func Test_sortArray(t *testing.T) {
	auxEx0 := example0
	auxEx0 = sortArray(auxEx0)
	auxEx1 := example1
	auxEx1 = sortArray(auxEx1)
	auxEx2 := example2
	auxEx2 = sortArray(auxEx2)
	assert.Equal(t, []int{1, 2, 3, 4, 100, 200}, auxEx1)
	assert.Equal(t, []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}, auxEx2)
}*/
