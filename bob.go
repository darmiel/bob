package bob

import (
	"fmt"
)

var (
	WriteWithNone  = WriteWithSeparator("")
	WriteWithSpace = WriteWithSeparator(" ")
	WriteWithComma = WriteWithSeparator(",")
	WriteWithTab   = WriteWithSeparator("\t")
)

func (b *Builder) WriteBytes(data []byte) *Builder {
	b.Builder.Write(data)
	return b
}

func (b *Builder) Write(data []byte) *Builder {
	return b.WriteBytes(data)
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

func (b *Builder) Writes(data ...interface{}) *Builder {
	var (
		sep = b.sep
		nl  = b.nl
	)
	a := false
	for _, d := range data {
		// options
		switch t := d.(type) {
		case writeAllOptionSeparator:
			sep = t
			continue
		case writeAllOptionNewLine:
			nl = t
			continue
		case func() *Builder:
			t()
			continue
		}

		// write separator
		if sep != "" && a {
			b.Builder.WriteString(string(sep))
		}
		a = true

		// write data
		b.WriteAny(d)
	}

	if nl {
		b.WriteNewLine()
	}
	return b
}

func (b *Builder) Bytes() []byte {
	return []byte(b.String())
}
