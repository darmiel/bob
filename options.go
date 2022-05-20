package bob

type (
	writeAllOptionSeparator string
	writeAllOptionNewLine   bool
)

//goland:noinspection ALL
func WriteWithSeparator(sep string) writeAllOptionSeparator {
	return writeAllOptionSeparator(sep)
}

//goland:noinspection ALL
func WriteWithNewLine() writeAllOptionNewLine {
	return true
}
