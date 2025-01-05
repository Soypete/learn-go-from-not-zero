# Module 3: Go Idioms and Advanced Patterns (30 minutes)

## **Learning Objective: By the end of this module, learners should be able to understand and implement key Go idioms and advanced patterns including structs, interfaces, inheritance, and channels.**

## Part 1: Structs and Interfaces (10 minutes)

We are going to start by revisiting structs. Structs are one of Go's versatile data types. They are used to group related data together. Just to recap, a struct is defined like this:

```go
type Person struct {
	FirstName string
	LastName string
	BirthDate time.Time
}
```

Structs can be used just like any other type in go, you can have a slice of structs, you can have a map of structs, you can have a struct as a field in another struct. They can be passed to functions and returned from functions. Possibilities are endless. A common pattern in go is to use structs and the object of a method. For example

```go
func (p Person) GetFullName() string {
return p.FirstName + " " + p.LastName
}
```

_NOTE_: getters and setters as they are used in other Object Oriented Paradigms are not idiomatic in Go. If you want to access a field of a struct, you should just access it directly. Outside of initialization, If you want to change a field of a struct, you should just change it directly.

Methods are used to define business logic that requires the values of the struct. It is very common to see methods on structs in imported packages, look at the `net/http` package for examples of this. Here the `http.Request` struct has many methods that can only be called once the struct has been instantiated and populated with values.

```go
req := http.NewRequest("GET", "http://example.com", nil)
req.Header.Add("Content-Type", "application/json")
resp, err := req.Do()
...
```

Go does not allow for overloading methods, but if you want to have different use cased for different types you can use. Interfaces are inheretence constructs in Go. They are used to define a set of methods that a struct must implement in order to be considered an instance of that interface. One common use case of this is the `error` interface. The `error` interface is defined as:

```go
type error interface {
    Error() string
}
```

This allows for any struct that has a method `Error() string` to be used as an error in Go. This is why you can return a `string` as an error in Go. This is also why you can define your own error types in Go.

## Part 2: Inheritance in Go (10 minutes)

There are a couple of tricks to using interfaces in Go. The first one is to require the struct to implement the interface. This is done by defining a type. If I want do define a new type that implements the `error` interface, I can do it like this:

```go
type MyError struct {
    Message string
}

func (e MyError) Error() string {
    return e.Message
}

func main() {
    var err error
    err = &MyError{Message: "This is my error"}
    fmt.Println(err.Error())
}
```

The `var err error` requires that whatever is assigned to `err` must implement the `error` interface. If it does not (meaning the `func (e MyError) Error() string {}` is not defined), the code will not compile. This is a great way to ensure that your code is correct at compile time.

You may have noticed that I used an `&` here. This is because if something is not exactly the type require, but in a derrived type, you can use a pointer to the type. Because `MyError` satisfies the `error` interface, by using a pointer to `MyError` we can can act anywhere the interface is required. Just to be explicit, in the above example, `var err error` is of type `error` which is an interface. `MyError` because it implements the methods required by the `error` interface, can be assigned to `err` and will be treated as an `error` type by the compiler. this is the power of the pointer reference.

This is how inheritance is done in Go. There is no `extends` keyword like in Java. Instead, you define an interface and require that the struct implement that interface. If you have a set of methods that are common to many struct objects, you can define an interface and require that the struct implement that interface.

Interfaces can also extend other interfaces. A great example of this is the `io` package. The `io.Reader` interface is defined as:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

The `io.Writer` interface is defined as:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

These are combined to create the `io.ReadWriter` interface:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

That means that if a struct is to implements the `io.ReadWriter` interface, it must implement the `Read()` and `Write()` methods from the `Reader` and `Writer` interfaces. This is a great way to add a set of methods to a struct without redefine any dependencies.

Notice that in Go, the pattern for naming interfaces is to end the name with `er`. This is not a rule enforced by the compiler, but it is a common pattern in the Go community. This is because interfaces are used to define a set of methods that perform an action. The `er` suffix is a common way to denote that the action associated with the methods.

Another great feature this unlocks is struct mocking for tests, but we will talk about that in the next module.

## Part 3: Advanced Patterns: idioms, Channels, and memory (10 minutes)

### Idomatic Go

Go's idioms are almost as old of as the language. They were presented by Rob Pike as the [Go proverbs](https://go-proverbs.github.io/)

```
Go Proverbs:
Simple, Poetic, Pithy
Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.
```

These idioms have a big impact on how Go code is written. We will go into some important idioms in this module. The most important thing to remember when trying to apply them to your code is to keep it simple. There are always cases where the best choice may not align with the idioms, but they will apply to 95% of all other Go applications, and you should always strive to keep your code simple and close to the idioms.

My favorite idioms are:

* Don't communicate by sharing memory, share memory by communicating: This helps to avoid overusing channels and go routines. It is very easy to get carried away with channels and go routines, but remember that they are not free. They take up memory and CPU cycles. Go is fast and often times you can get away with not using them and just writing a simple program.
* The bigger the interface, the weaker the abstraction: Interfaces naturally complicate code by abstracting the logic from the flow. Abstraction always add complexity we should all strive to only add necessary complexity. Small interfaces add value by making code more compartimentalized and testable. Small interfaces are easy to implement and easy to understand. A good example of this is the `io` package. The interfaces have one or two methods and are used everywhere in the standard lib.
* Clear is better than clever: This emphasized readable and debuggable code. Simplicity is best, let the compiler and the garbage collector handle the "hard parts". If you have to choose between clear easy to understand code and clever intricate code, always choose the clear code.
* Don't just check errors, handle them gracefully: Returning an error up the stack is the best practice, but that is just the first step. If your can check error types and handle them differently, you should. Panicing on an error is not handling it gracefully and should be avoided as it is only intended when the program is in an unrecoverable state and needs to be shut down.

  As you can see, these idioms are simple and easy to understand. They are not hard and fast rules, but they are good guidelines to follow when writing Go code.

### Go and Concurrency

When choosing Go for the first time, many people are drawn to concurrency. Go does have a great concurrency model, but it has some rough edges. Concurrency a balancing act between Memory sharing, Orchestration, Garbage Collection. If you are interested in writing concurrent code you should start with the Rob Pike talk, [Concurrency is not parallelism](https://blog.heroku.com/concurrency_is_not_parallelism).

You we can do an entire course on concurrency and how to do it well, but here we are just going to talk about some of the basics of goroutines and channels.

#### Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`
Channels can be instantiated in a struct for global variable, but like maps they are `nil`. Channels must be allocated with `make(chan int)`

```go
func main() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    fmt.Println(<-ch)
}
```

Channels enable concurrency by allowing you to send and receive values between goroutines, but they are not the only way to manage concurrency. This is where the Go proverb "Don't communicate by sharing memory, share memory by communicating" comes in. To avoid locks and race conditions make sure you follow the memory sharing best practices across goroutines even if you are using channels.

There are 2 main types of channels `buffered` and `unbuffered` channels. Buffered channels are created with a `buffer` that allows for a certain number of values to be stored in the channel. Unbuffered channels are created with no buffer and will block until a value is sent or received.

You can explore this example in the Tour of Go section on [channels](https://go.dev/tour/concurrency/3)

Make sure you always run `close()` when a channel is no longer needed. This will prevent memory leaks and other issues.

#### Goroutines

A goroutine is a lightweight thread of execution. It is a virtual thread that is managed by the Go runtime and does not necessary map to a physical thread. The exit of a goroutine is not guaranteed to happen before any event in the program. Goroutines are started with the `go` keyword.

```go
func hello() {
    var a string
	go func() { a = "hello" }()
	print(a)
}
```

There are two options for running goroutines: anonymous or literal functions. Anonymous functions are functions that are defined in the same line as they are called. Literal functions are functions that are defined elsewhere in the code and are called by name.

```go
func printHello() {
    fmt.Println("Hello")
}

func main() {
    go printHello()
}
```

sync.WaitGroup are used to block exit until all goroutines have finished execution. This is the best way to orchestate the execution of distinct processes in your program.

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("Hello")
    }()
    wg.Wait()
    fmt.Println("Goodbye")
}
```

In addition the `sync` package has a `Mutex` type that can be used to lock and unlock access to shared memory. This is a great way to avoid race conditions and other issues that can arise from shared memory. This is used to lock and unlock access to shared memory so that only one goroutine can access it at a time and the rest are blocked until the lock is released.

```go
var mu sync.Mutex
var balance int
func addToBalance(increment int) {
    mu.Lock()
    balance += increment
    mu.Unlock()
}
```

The last major feature of Go's concurrency is the `select` statement. The `select` statement is used to choose between multiple send and receive operations. It is a blocking call and will wait until one of its cases can run, then it executes that case.

This is the select statement from the [Tour of Go](https://go.dev/tour/concurrency/5)

### Linters

Go has always had linters and formatters built into the standard library. This is a great feature of Go. The linters are very strict and will catch many common errors. The formatters are very opinionated and will format your code in a way that is easy to read and understand.I highly recommend adding them to you editor and having them run on save. This will save you a lot of time and effort in the long run. This even can help with rewriting code to be more idiomatic.

Here are some tools to try out:

* go fmt : formatter shipped with Go
* go vet : linter shipped with Go
* [gofumpt](https://github.com/mvdan/gofumpt): a stricter version of gofmt
* [golangci-lint](https://github.com/golangci/golangci-lint): a linter that combines many linters into one
