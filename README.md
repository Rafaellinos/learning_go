# Go learning

## packages

- must have
- multi-packages
- `package main` is the entrypoint and its necessary to compiled code

## modules

- 1 modules have multiple packages
- `go mod init github.com/first-app` to create a module
- `go.mod` module info and version (like maven/requirements.txt)


## conventions

- files with underscore instead
- variables in camelCase

## Types

- Integers: int, int8, int16, int32, int64
- Floating-Point: float32, float64 (100.0 inferred) default value 0.0
- Complex Numebers: complex64, complex128
- String: string
- Booleans: bool (true, false)
- Byte: byte is an alias for uint8 (uint strictly non-negative)
- Char: rune is an alias for int32

## assignment

- `var someAmount float64 = 1000` explicitly converting
- `someAmount := 1000` implicity assignment, inferred type
- `var someAmount, salary float64 = 1000, 200`
- `const constantAmount = 2.5` const cannot be re-assigned

- PS: when using `var` keyword, mandatory specify the type. eg: `var years int64`


## terminal interaction

- `fmt.Scan(&someVariable)` get input from user and inserts into the pointer (does not accept multi word input)
- `fmt.Print()` stdout without breakline


## Pointers

- `&someVariable` indicates a pointer
- `func someFunction(value *int)` funcs declared as pointers expects a pointer. eg: `someFunction(&value)`

## Functions

- `func someFunc(text string) {}` keyword for creating funcs
- `func calculate() {return 1,2,3} (int64, int64, int64)` can return many values with comma

```go
package main

import "fmt"

func main() {
    result := someCalculation(1, 3)
    fmt.Println(result)
}

func someCalculation(number1, number2 int64) (result int64) {
    result = number1 + number2
    return result
}
```

## Scopes

## Control structures

### for loops

- `for {}` infinity loop
- `for value != 4 {}` do something while value different than 4
- `for i := 0; i < 5; i++ {}` classic for loop
- `continue` and `break` is also availabe
- `for n := range 6 {}` in range, last value excluded, starts with 0

### switchs

- `switch variable { case 1: ... default: ...}`

## add dependencies

- `go mod tidy` remove unnused and add missing dependencies
- `go get package@v123` add this dependency to project

## Utils

- `gofmt -w yourfile.go` to format


