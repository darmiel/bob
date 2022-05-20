package bob

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWrite(t *testing.T) {
	var b Builder
	b.Write([]byte("Hello"))
	assert.Equal(t, "Hello", b.String(), "Write []byte")

	b.WriteRune('!')
	assert.Equal(t, "Hello!", b.String(), "Write rune")

	b.WriteString(" World")
	assert.Equal(t, "Hello! World", b.String(), "Write string")

	b.WriteByte('!')
	assert.Equal(t, "Hello! World!", b.String(), "Write byte")

	b.Reset()
	assert.Equal(t, "", b.String(), "Reset")

	// line
	assert.Equal(t, "\n", b.WriteNewLine().String(), "New Line")
	b.Reset()

	b.WriteBytesLine([]byte("Hello"))
	assert.Equal(t, "Hello\n", b.String(), "WriteBytesLine")

	b.WriteRuneLine('!')
	assert.Equal(t, "Hello\n!\n", b.String(), "WriteRuneLine")

	b.WriteStringLine("World")
	assert.Equal(t, "Hello\n!\nWorld\n", b.String(), "WriteStringLine")

	b.WriteByteLine('!')
	assert.Equal(t, "Hello\n!\nWorld\n!\n", b.String(), "WriteByteLine")
}

type S1 struct {
	Name string
	Age  int
}

type S2 struct {
	Name string
	Age  int
}

func (s S2) String() string {
	return s.Name + " is " + strconv.Itoa(s.Age)
}

func TestAny(t *testing.T) {
	var b Builder
	b.WriteAny(byte('H'))
	b.WriteAny([]byte("el"))
	b.WriteAny("lo")
	b.WriteAny('!')
	b.WriteAny(1)
	b.WriteAny(nil)
	b.WriteAnyLine(true)
	assert.Equal(t, "Hello!1<nil>true\n", b.String(), "WriteAny")
	b.Reset()

	b.WriteAny(S1{Name: "Bob", Age: 18})
	b.WriteAny(" -- ")
	b.WriteAny(S2{Name: "Alice", Age: 18})
	assert.Equal(t, "{Name:Bob Age:18} -- Alice is 18", string(b.Bytes()), "WriteAny")
}

func TestAnyAll(t *testing.T) {
	var b *Builder
	b = new(Builder)
	b.Writes(byte('H'), []byte("el"), "lo", '!', 1, nil, true)
	assert.Equal(t, "Hello!1<nil>true", b.String(), "WriteAll")

	b.Reset()
	b.Writes(
		WriteWithSeparator(" "), byte('H'), []byte("el"), "lo",
		WriteWithSeparator("+"), '!', 1, nil, true,
	)
	assert.Equal(t, "H el lo+!+1+<nil>+true", b.String(), "WriteAll")

	b.Reset()
	b.Writes("Hello", " World")
	assert.Equal(t, "Hello World", b.String(), "WriteAll")

	b.Writes(WriteWithNewLine(), " Hi!")
	assert.Equal(t, "Hello World Hi!\n", b.String(), "WriteAll")
	b.Reset()

	b = New(WithDefaultWritesSeparator(" "), WithLineSeparator("\n\n"))
	b.Writes(WriteWithNewLine(), "Hello", "World", "What's", "up?", b.WriteNewLine)
	assert.Equal(t, "Hello World What's up?\n\n\n\n", b.String(), "WriteAll")
}
