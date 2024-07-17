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
    - Idiomatic Go - How to detect and avoid code smells
    - Linters, checks, and best practices
    - Channels - Sending and receiving data through channels
        - Goroutines
        - Select statement
    - Memory model -How to share memory?
