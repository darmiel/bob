package bob

import "strings"

type Option func(b *Builder)

func WithLineSeparator(sep string) Option {
	return func(b *Builder) {
		b.newLine = sep
	}
}

func WithDefaultWritesSeparator(sep string) Option {
	return func(b *Builder) {
		b.sep = writeAllOptionSeparator(sep)
	}
}

type Builder struct {
	strings.Builder

	// WriteNewLine
	newLine string

	// Writes defaults
	sep writeAllOptionSeparator
	nl  writeAllOptionNewLine
}

func New(options ...Option) *Builder {
	b := new(Builder)
	for _, o := range options {
		o(b)
	}
	return b
}

///

func Write(data []byte) *Builder {
	return New().Write(data)
}

func WriteRune(r rune) *Builder {
	return New().WriteRune(r)
}

func WriteByte(c byte) *Builder {
	return New().WriteByte(c)
}

func WriteString(data string) *Builder {
	return New().WriteString(data)
}

func WriteNewLine() *Builder {
	return New().WriteNewLine()
}

func WriteAny(data interface{}) *Builder {
	return New().WriteAny(data)
}

func WriteBytesLine(data []byte) *Builder {
	return New().WriteBytesLine(data)
}

func WriteRuneLine(data rune) *Builder {
	return New().WriteRuneLine(data)
}

func WriteByteLine(data byte) *Builder {
	return New().WriteByteLine(data)
}

func WriteStringLine(data string) *Builder {
	return New().WriteStringLine(data)
}

func WriteAnyLine(data interface{}) *Builder {
	return New().WriteAnyLine(data)
}

func Writes(data ...interface{}) *Builder {
	return New().Writes(data...)
}
