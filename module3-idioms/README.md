# Module 3: Go Idioms and Advanced Patterns (30 minutes) 
**Learning Objective: By the end of this module, learners should be able to understand and implement key Go idioms and advanced patterns including structs, interfaces, inheritance, and channels.**
---
## Part 1: Structs and Interfaces (10 minutes)

Structs are one of Go's complex data types. They are used to group related data together. You can define a struct using the `type` keyword followed by the name of the struct and the fields that make up the struct. Here is an example of a struct that represents a person:

```go
type Person struct {
FirstName string
LastName string
BirthDate time.Time
}
```

There are a couples of ways to instantiate and populate the values of a struct. You can straight declare a struct and populate the values like this:

```go

brother := Person{
FirstName: "John",
LastName: "Doe",
BirthDate: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
}
```

Or you can make and empty struct and populate the values like this:

```go
brother := Person{} // var brother Person
brother.FirstName = "John"
brother.LastName = "Doe"
brother.BirthDate = time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
```

Structs can be used just like any other type in go, you can have a slice of structs, you can have a map of structs, you can have a struct as a field in another struct. They can be passed to functions and returned from functions. Possibilities are endless. One of the most common way to use structs is as an object in a method. For example

```go
func (p Person) GetFullName() string {
return p.FirstName + " " + p.LastName
}
``` 

*NOTE*: getters and setters are not idiomatic in Go. If you want to access a field of a struct, you should just access it directly. If you want to change a field of a struct, you should just change it directly. If you want to do some logic when you change a field, you should use a method on the struct. 

These methods are used to define business logic that requires the values of the struct. It is very common to see methods on structs in imported packages, look at the `net/http` package for examples of this. Here the `http.Request` struct has many methods that can only be called once the struct has been instantiated and populated with values.

```go
req := http.NewRequest("GET", "http://example.com", nil)
req.Header.Add("Content-Type", "application/json")
resp, err := req.Do()
...
``` 

Interfaces are inheretence constructs in Go. They are used to define a set of methods that a struct must implement in order to be considered an instance of that interface. One common use case of this is the `error` interface. The `error` interface is defined as:

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

You may have noticed that I used an `&` here. This is because if something is not exactly the type require, but in a derrived type, you can use a pointer to the type. This is a common pattern in Go.

This is how inheritance is done in Go. There is no `extends` keyword like in Java. Instead, you define an interface and require that the struct implement that interface. If you have a set of methods that are common to many struct objects, you can define an interface and require that the struct implement that interface.

Another great feature this unlocks is struct mocking, but we will talk about that in the next module.

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

These idioms have a big impact on how Go code is written. We will go into some important idioms in this module. The most important thing to remember when trying to apply them to your code is to keep it simple. They will apply to 95% of all software written in Go, and you should always strive to keep your code simple and close to the idioms.

My favorite idioms are 
    - Don't communicate by sharing memory, share memory by communicating - This helps to avoid overusing channels and go routines. It is very easy to get carried away with channels and go routines, but remember that they are not free. They take up memory and CPU cycles. Go is fast and often times you can get away with not using them and just writing a simple program.
    - The bigger the interface, the weaker the abstraction - Interfaces add abstraction to your code. Abstraction always add complexity and we need it to be necessary complexity. Small interfaces add value by making code more compartimentalized and testable. Small interfaces are easy to implement and easy to understand. A good example of this is the `io` package. The interfaces have one ot two methods and are used everywhere in the standard lib.
    - Clear is better than clever - This emphasized readable and debugable code. Simplicity is best, let the compiler and the garbage collector handle the "hard parts". 
    - Dont just check errors, handle them gracefully - Returning an error up the stack is the best practice, but that is just the first step. If your can check error types and handle them differently, you should. If you can log the error and continue, you should. If you can retry the operation, you should.

    As you can see, these idioms are simple and easy to understand. They are not hard and fast rules, but they are good guidelines to follow when writing Go code.

### Go and Concurrency

When choosing Go for the firs time, many people are drawn to concurrency. Go does have a great concurrency model, but it has some rough edges. If you are interested in writting concurrent code you should start with the Rob Pike talk, [Concurrency is not parallelism]().

You we can do an entire course on concurrency and how to do it well, but here we are just going to talk about some of the basics of goroutines and channels.

<!-- pull the slides from Production go -->

### Linters

Go has always had linters and formatters built into the standard library. This is a great feature of Go. The linters are very strict and will catch many common errors. The formatters are very opinionated and will format your code in a way that is easy to read and understand.I highly recommend adding them to you editor and having them run on save. This will save you a lot of time and effort in the long run. This even can help with rewriting code to be more idiomatic.

Here are some tools tp try out:
    - gofmt
    - go vet
    - gofumpt
    - golangci-lint