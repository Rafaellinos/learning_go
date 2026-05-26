# Go learning

## packages

- must have
- multi-packages
- `package main` is the entrypoint and its necessary to compiled code
- can use `go run .` to run all files within the same package (must have main func)
- functions can only be imported with upper case (eg: `somepackage.SomeFunc`)
- to import packages, use full name (eg: `"github.com/Rafaellinos/bank/fileops"` `fileops.SomeFuncFromFileOps`)
- go add third-party package: `go get github.com/Pallinder/go-randomdata`

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

- points to an address
- avoid unnecessary copies, allow mutate objects
- `&someVariable` indicates a pointer
- `func someFunction(value *int)` funcs declared as pointers expects a pointer. eg: `someFunction(&value)`
- __copies in go are just shallow copies__
- `anotherInt := &someInt` this creates a pointer variable
- `var anotherInt *int; anotherInt = &someInt` another way to create variable pointers pointers
- must **dereference** a pointer in order to extract the value `fmt.Println("value: ", *someIntPointer)`
- default value for pointers are `nil`. It represents the pointer has no address, no value in memory
- cannot perform operations in pointer, __must extract the value first__ eg: `*somePointer`

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

### for loops

- `for {}` infinity loop
- `for value != 4 {}` do something while value different than 4
- `for i := 0; i < 5; i++ {}` classic for loop
- `continue` and `break` is also availabe
- `for n := range 6 {}` in range, last value excluded, starts with 0

### switchs

- `switch variable { case 1: ... default: ...}`
- `break` gets out of the swith statement, not the loop if its inside a loop

## add dependencies

- `go mod tidy` remove unnused and add missing dependencies
- `go get package@v123` add this dependency to project


## Structs

- `type NameStruct struct {}` defines the struct boby
- to access a field in struct with pointers: `(*someStruct).someField` that the tecnicall correct whey, but can be done without `*` as a syntax sugar for dereference

```go
import (
    "time"
    "fmt"
)

type User struct {
    name string
    createdAt time.Time
}

func main {
    someUser := User{
        name: "Rafael Veloso",
        createdAt: time.Now(),
    }
    // can omit key, following the sequence keys:
    anotherUser := User{
        "Yasmin",
        time.Now(),
    }
    emptyUser := User{} // all nil values
    fmt.Println("Fullname: ", someUser.name)
}
```

### struct methods

- func in a struct is a __method__
- `func (user User) someMethod() {}` has access to "this" with `user User`
- syntax = `userInstance.someMethod()`
- method can mutate its own objects by using pointers `func (user *User) someMethod() {}`
- contructor are just convention eg: `func New(params...) *User {return &User{}}`

### struct fields

- fields with lower case cannot be accessed, the rule for upper case algo applies to fields

```go
package user

type User struct {
    age int // cannot be accessed outside of user package
    Name string // can be accessed outside of user package
}
```

### Struct embedding (like inheritance)

```go
package user

type Admin struct {
    email string
    password string
    user User // or just user
}

func newAdmin(email, password string) Admin {
    return Admin{
        email: email,
        password: password,
        user: User{
            FirstName: "Name"
        }
    }
}
```

### struct tags

- metadata, like annotation in java

## Utils

- `gofmt -w yourfile.go` to format

## Custom types (type alias)

- custom types `type str string`
- can be used with functions. Eg: `func (text str) log() {fmt.Println(text}`
    * usage: `var name str = "Max";name.log()`

