package bob

import (
	"fmt"
	"strings"
)

type Builder struct {
	strings.Builder

	// WriteNewLine
	newLine string

	// WriteAll
	sep bobWriteAllOptionSeparator
	nl  bobWriteAllOptionNewLine
}

type Option func(b *Builder)

func WithLineSeparator(sep string) Option {
	return func(b *Builder) {
		b.newLine = sep
	}
}

func New(options ...Option) *Builder {
	b := new(Builder)
	for _, o := range options {
		o(b)
	}
	return b
}

// "normal" methods

func (b *Builder) Write(data []byte) *Builder {
	b.Builder.Write(data)
	return b
}

func (b *Builder) WriteRune(r rune) *Builder {
	b.Builder.WriteRune(r)
	return b
}

func (b *Builder) WriteByte(c byte) *Builder {
	b.Builder.WriteByte(c)
	return b
}

func (b *Builder) WriteString(data string) *Builder {
	b.Builder.WriteString(data)
	return b
}

func (b *Builder) WriteAny(data interface{}) *Builder {
	if s, ok := data.(fmt.Stringer); ok {
		b.Builder.WriteString(s.String())
		return b
	}

	switch t := data.(type) {
	case byte:
		b.Builder.WriteByte(t)
	case []byte:
		b.Builder.Write(t)
	case rune:
		b.Builder.WriteRune(t)
	case string:
		b.Builder.WriteString(t)
	default:
		b.Builder.WriteString(fmt.Sprintf("%+v", t))
	}
	return b
}

// "custom" methods

func (b *Builder) WriteNewLine() *Builder {
	var nl string
	if b.newLine == "" {
		nl = "\n"
	} else {
		nl = b.newLine
	}
	b.Builder.WriteString(nl)
	return b
}

func (b *Builder) WriteBytesLine(data []byte) *Builder {
	return b.Write(data).WriteNewLine()
}

func (b *Builder) WriteRuneLine(data rune) *Builder {
	return b.WriteRune(data).WriteNewLine()
}

func (b *Builder) WriteByteLine(data byte) *Builder {
	return b.WriteByte(data).WriteNewLine()
}

func (b *Builder) WriteStringLine(data string) *Builder {
	return b.WriteString(data).WriteNewLine()
}

func (b *Builder) WriteAnyLine(data interface{}) *Builder {
	return b.WriteAny(data).WriteNewLine()
}

// "all" write

type (
	bobWriteAllOptionSeparator string
	bobWriteAllOptionNewLine   bool
)

func WriteWithSeparator(sep string) bobWriteAllOptionSeparator {
	return bobWriteAllOptionSeparator(sep)
}

func WriteWithNewLine() bobWriteAllOptionNewLine {
	return bobWriteAllOptionNewLine(true)
}

func (b *Builder) WriteAll(data ...interface{}) *Builder {
	var (
		sep bobWriteAllOptionSeparator
		nl  bobWriteAllOptionNewLine
	)
	for _, d := range data {
		// write separator
		if sep != "" && b.Len() > 0 {
			b.Builder.WriteString(string(sep))
		}

		// options
		switch t := d.(type) {
		case bobWriteAllOptionSeparator:
			sep = t
			continue
		case bobWriteAllOptionNewLine:
			nl = t
			continue
		}

		// write data
		b.WriteAny(d)
	}

	if nl {
		b.Builder.WriteRune('\n')
	}
	return b
}
