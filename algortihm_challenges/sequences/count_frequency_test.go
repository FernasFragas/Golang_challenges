package sequences

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countFrequency(t *testing.T) {
	occurrers := countFrequency("hello")
	const single = 1
	assert.Equal(t, single, occurrers['h'])
	assert.Equal(t, single, occurrers['e'])
	assert.Equal(t, 2, occurrers['l'])
	assert.Equal(t, single, occurrers['o'])
}

func Test_countFrequency_MixedUpperAndLowerCases(t *testing.T) {
	occurrers := countFrequency("HeLlo")
	const single = 1
	assert.Equal(t, single, occurrers['h'])
	assert.Equal(t, single, occurrers['e'])
	assert.Equal(t, 2, occurrers['l'])
	assert.Equal(t, single, occurrers['o'])
}
