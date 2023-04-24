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

//------------

// variables -> Name, data, etc


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
- Can lose "time-to-market" compared to languages like Python
