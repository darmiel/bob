package bob

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert.Equal(t, "Hello", WriteString("Hello").String())
	assert.Equal(t, "Hello", Write([]byte("Hello")).String())
	assert.Equal(t, "H", WriteRune('H').String())
	assert.Equal(t, "H", WriteByte('H').String())
	assert.Equal(t, "World", WriteAny("World").String())
	assert.Equal(t, "\n", WriteNewLine().String())
	assert.Equal(t, "Hello\n", WriteBytesLine([]byte("Hello")).String())
	assert.Equal(t, "Hello\n", WriteStringLine("Hello").String())
	assert.Equal(t, "H\n", WriteRuneLine('H').String())
	assert.Equal(t, "H\n", WriteByteLine('H').String())
	assert.Equal(t, "Hello\n", WriteAnyLine("Hello").String())
	assert.Equal(t, "Hello Mom!", Writes("He", 'l', 'l', []byte("o"), ' ', "Mom!").String())
}
