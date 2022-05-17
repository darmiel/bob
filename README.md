# bob

A simple wrapper for [strings.Builder](https://pkg.go.dev/strings#Builder) with some (more or less) useful methods.

## Differences from strings.Builder

### Method Chaining Pattern

**⚠️ NOTE:** Doesn't implement [io.StringWriter](https://pkg.go.dev/io#StringWriter)
or [io.Writer](https://pkg.go.dev/io#Writer) interface anymore.

**bob.Builder**

```go
var b bob.Builder
b.WriteString("Hello").WriteString(" ").WriteString("World")
fmt.Println(b.String()) // Hello World
```

> **strings.Builder**
> ```go
> var b strings.Builder
> b.WriteString("Hello")
> b.WriteString(" ")
> b.WriteString("World")
> fmt.Println(b.String()) // Hello World
> ```

### `WriteAny(any)`

Accepts any data and tries to convert to string:

| Type         | Method                                |
|--------------|---------------------------------------|
| fmt.Stringer | `fmt.Stringer.String()`               |
| byte         | `strings.Builder.WriteByte(byte)`     |
| byte[]       | `strings.Builder.WriteBytes([]byte)`  |
| rune         | `strings.Builder.WriteRune(rune)`     |
| string       | `strings.Builder.WriteString(string)` |
| any          | `fmt.Sprintf("%+v", any)`             |

### `WriteAll(...interface{})`

```go
var (
    Name = "Bob"
    Age = 1337
	
    b bob.Builder
)
b.WriteAll(bob.WriteWithSeparator(" "), "Hello my name is", Name, "and I'm", Age, "years old")
```

#### Modify

By default, there is no separator between arguments. You can change this behavior using `bob.WriteWithSeparator(" ")`:

```go
var b bob.Builder
b.WriteAll(
	bob.WriteWithSeparator("+"), "This", "Is", 
	bob.WriteWithSeparator(" "), "A", "Test", 
)
// This+Is A Test
```

> same example using **strings.Builder**
> ```go
> var (
>     Name = "Bob"
>     Age = 1337
>	
>     b strings.Builder
> )
> b.WriteString("Hello my name is ")
> b.WriteString(Name)
> b.WriteString(" and I'm ")
> b.WriteString(strconv.Itoa(Age))
> b.WriteString(" years old")
> ```

---

### `WriteNewLine()`

> Alias to `b.WriteString("\n")`

The line separator (`\n`) can be changed using `WithLineSeparator`:

```go
b := bob.New(bob.WithLineSeparator("\n\n"))
b.WriteNewLine()
```

---

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
