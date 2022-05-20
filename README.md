# bob

A simple wrapper for [strings.Builder](https://pkg.go.dev/strings#Builder) with some (more or less) useful methods.

```go
name := "bob"
age := 42
fmt.Println(bob.Writes(bob.WriteWithSpace, "My name is", name, "and I am", age, "years old").String())
```

## Differences from `strings.Builder`

### Method Chaining

Bob allows you to call multiple methods on top of each other:

```go
fmt.Println(bob.WriteString("Hello").
    WriteRune(' ').
    WriteString("World").
    String()) // Hello World
```

**NOTE:** Because of this change, `bob` doesn't implement the [io.StringWriter](https://pkg.go.dev/io#StringWriter)
or [io.Writer](https://pkg.go.dev/io#Writer) interface anymore.


### `WriteAny(any)`

Accepts any object and tries to convert to string using one of the following methods:

| Type         | Method                                |
|--------------|---------------------------------------|
| fmt.Stringer | `fmt.Stringer.String()`               |
| byte         | `strings.Builder.WriteByte(byte)`     |
| byte[]       | `strings.Builder.WriteBytes([]byte)`  |
| rune         | `strings.Builder.WriteRune(rune)`     |
| string       | `strings.Builder.WriteString(string)` |
| any          | `fmt.Sprintf("%+v", any)`             |

### `Writes(...interface{})`

```go
name := "bob"
age := 42
fmt.Println(bob.Writes(bob.WriteWithSpace, "My name is", name, "and I am", age, "years old").String())
```

#### Modify

By default, there is no separator between arguments. 
You can change this behavior using `bob.WriteWithSeparator(<separator>)` or using one of the predefined separators:

- `bob.WriteWithNone`: abc
- `bob.WriteWithSpace`: a b c
- `bob.WriteWithComma`: a,b,c
- `bob.WriteWithTab`: a\tb\tc

```go
bob.Writes(
    bob.WriteWithSeparator("+"), "This", "Is",
    bob.WriteWithSeparator(" "), "A", "Test",
) // This+Is A Test
```

---

### `WriteNewLine()`

Adds a new line.

```go
fmt.Println(bob.WriteString("Hello").WriteNewLine().WriteString("World").String()) // Hello\nWorld
```

The default line separator (`\n`) can be changed using `WithLineSeparator`:

```go
bob.New(bob.WithLineSeparator("\n\n")).
	WriteString("Hello").
	WriteNewLine().
	WriteString("World") // Hello\n\nWorld
```

---

## Aliases

### `WriteBytesLine([]byte)`

> Alias to `b.Write([]byte).WriteNewLine()`

### `WriteRuneLine(rune)`

> Alias to `b.WriteRune(rune).WriteNewLine()`

### `WriteByteLine(byte)`

> Alias to `b.WriteByte(byte).WriteNewLine()`

### `WriteStringLine(string)`

> Alias to `b.WriteString(string).WriteNewLine()`

### `WriteAnyLine(interface{})`

> Alias to `b.WriteAny(interface{}).WriteNewLine()`
