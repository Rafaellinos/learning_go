// rune = alias for int32 that represents a letter
// string is a series bytes | String are not null terminated
// "" <- string
// ” <- rune
// `let's code\n` <- raw string

// go commands:
// build -> distribute
// build -race -> check for concurrency problems
// run -> runs directly, no output executable
// mod -> manage modules and dependencies , mod tidy -> update dependencies
// test -> runs the projects test
// fmt -> format all source files
// go mod init live <- creates a local package
// go mod tidy <- download dependencies

//------------

// go frameworks

// echo framework, http

// variables -> Name, data, etc

### Go characteristics

- Not Object-Oriented
- Has built-ins like net/http (like python)
- Command: panic(err) -> stops the application
- defer db.Close() -> starts a context, when the code ends, executes the defer func


### Go lang Advantages:


- Garbage collector
- Go routines (green thread, thread in userland or light threads) -> threads managed by golang
- Avoid concurrent problems with variables by using channels to share variables
- Only 25 keywords to learn
- Easy to learn
- Static code analysis
- Can use C libs
- Compiled to machine code
- Easy to manage concurrency


### Go lang Disadvantages


- Young language -> absence of some SDKs
- Can lose "time-to-market" compared to languages like Python (Time Consuming)


#### Compile to windows

```shell
GOOS=windows go build main.go
```

<p>The command above generates a main.exe</p>


#### Run local server

```go
package main

import (
	"fmt"
	"net/http"
) //built-in

type Product struct { // like classes
	ID    string
	Name  string
	Price float64
}

func main() {
	product := Product{
		ID:    "1",
		Name:  "Product 1",
		Price: 100,
	}
	fmt.Println(product.Name)
	http.HandleFunc("/hello", homeHandler)
	http.ListenAndServe(":3000", nil) // Go routine
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello World"))
}
```

#### Notes

<p>When u have a letter in capital in a function/package, this function will be exported.
</p>

<h3>e.g.:</h3>


```go
package test

import "fmt"

func Test(msg string) { // will be exported (public) otherwise, private
	fmt.Println(msg)
}
```
