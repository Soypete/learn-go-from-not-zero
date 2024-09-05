**Learning Objective: By the end of this module, learners should be able to set up the Go development environment and understand the basics of Go syntax and structure.**

---

## Part 1: Setting Up Go Development Environment (10 minutes)

###Installing Go
**Install via website: https://golang.org/dl/**
We are glad you chose to learn Go. In this section, we will guide you through the installation process of Go on your system. Go is distributed as a binary package, and you can download the installer from the official Go website. This is the recomended way to install Go on your system if you are unfamiliar with package managers.
[Download Go](https://golang.org/dl/)

<!-- Add image of the download page with the download button highlighted -->

* Click on the download button to download the installer for your operating system.
* Run the installer and follow the on-screen instructions to install Go on your system.

** install via package manager**
Unix-based systems like Linux and macOS have package managers that can be used to install Go. If you are using a package manager like Homebrew on macOS, you can install Go using the following command:

```
$ apt-get install go
* brew install go
```

This will download and install Go on your system.

** install via webi **
I have a third option for you that works on all operating systems. It is called [webi](https://webi.dev). Webi will install the binary for you and set up the environment variables for you. It is a great way to get started with Go without having to worry about the installation process.

Once you have installed Go, run the command `go version` to verify the installation

### Working With Go

The best way to interact with the Go tooling is through the command line. The Go tooling provides a set of commands that help you manage your Go projects, build and run your code, and manage dependencies. Here are some helpful commands do familiarize yourself with your local Go development environment:

** make sure we know how to update go version and to check the path**

if you have tried installing Go via various methods, it is possible that you have multiple versions of Go installed on your system. To check the version of Go you are using, run the following command:

```bash
    which go -a
```

If you have multiple versions of Go installed, you can switch between them by setting the `GOROOT` environment variable to the path of the Go version you want to use, or just simply delete the version you don't want to use.

#### The Go CLI

The Go CLI is a powerful tool that allows you to manage your Go projects, build and run your code, and manage dependencies. To learn more about the Go CLI, run the following command:

```bash
go help
```

This will display a list of available commands and their descriptions. You can run `go help <command>` to get more information about a specific command.

The most common commands you will use are `go build`, `go run`, and `go test`. The `go build` command is used to compile your Go program. The `go run` command is used to compile and run your Go program. The `go test` command is used to run tests for your Go program. `go build` and `go test` are commonly used in CI/CD pipelines to build and test your Go programs.

We will go through some of the other commands in the next section.

**GOPATH and GOROOT**

`GOPATH` is an environment variable that specifies the location of your Go binaries and dependencies are stored. By default, the Go workspace is located at `$HOME/go` on Unix-based systems and `%USERPROFILE%\go` on Windows. You can change the location of the workspace by setting the `GOPATH` environment variable. For a cleaner setup, I recommend setting the `GOPATH` to a directory of your choice in your bash profile or zshrc file.

`GOROOT` is an environment variable that specifies the location where the Go binary is installed on your system. This is set automatically when you install Go, and you don't need to change it unless you have multiple versions of Go installed. [reference](https://go.dev/wiki/InstallTroubleshooting#goroot-vs-gopath). To check the value of `GOROOT`, run the following command:

```bash
go env
```

### Setup Your First Go Project

To create a new Go project, you need to create a new directory in your workspace and initialize it as a Go module. A Go module is a collection of Go packages that are versioned together. To create a new Go module, run the following command in your project directory:

```bash
go mod init <module-path>
```

I almost always use my public github repo as the module path. So if I was creating a new project called `myproject`, I would run the following command:

```bash
go mod init github.com/Soypete/myproject
```

The next step is to run the command `go mod tidy` to download the dependencies for your project. Now you can validate that your project was setup as a go module by accessing the `go.mod` file. For an empty go project it should look like this:

[]() <!--- add image of go.mod file -->

Here is a project with a many dependencies [project]()

The `go.mod` file here had several dependencies that were added by running the `go get` command. This command will download the dependencies for your project and add them to the `go.mod` file. The `replace` directive in the `go.mod` file allows you to replace a dependency with a local version of the package. There are many uses for the `replace` directive, but one common use case is to replace a remote dependency with a local version of the package for testing or development purposes.

When you run `go mod tidy`, the Go tooling will download the dependencies for your project and add them to the `go.mod` file. It will also generate a `go.sum` file that contains the checksums of the dependencies. This file is used to ensure that the dependencies are downloaded securely and have not been tampered with.

<!-- - explain workspaces and modules for internal and external dependency changes -->

From the project directory, open a main.go file and add the following code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Save and close the file.

To run the program, use the `go run` command followed by the name of the file:

```bash
go run main.go
```

You have now built your first Go program!

<!-- ### Writing and running your first Go program -->

<!--         - Create a simple `hello.go` program -->

<!--         - Run the program using `go run hello.go` -->

<!--         - Explain the output and the `main` function -->

### Workspaces

If you have a project wiht many contributors, many modules, or is a monorepo, you should consider looking at [workspaces](https://go.dev/doc/tutorial/workspace](https://go.dev/doc/tutorial/workspaces)

Workspaces allow you to manage depency changes across repo with multiple modules. This is a great way to manage internal and external dependencies.

to create a workspace use the command

```bash
go work init {module-path}
```

## Part 2: Go Program Structure (15 minutes)

### Understanding Go file structure

<!-- The add a go repo file structure image here -->

The executable Go programs are made up of packages. A package is a collection of Go source files that are compiled together. Each package has a unique name, and the name of the package is the same as the name of the directory that contains the package source files. The `main` package is a special package that contains the `main` function, which is the entry point of the program. The main file can me named anything, but it is common to name it `main.go`. When compiling a Go program, the `main.go` file becomes a `main` binary file that can be executed on your system.

In Go, each directory is a self contained module. All of the files in that package belong to the module this is denoted on the first line of the file.

```go
package {module-name}
```

The module name is the name of the module that the file belongs to. This is used to import the module into other files outside of the module. The main function belongs in the main module denoted by the `main` package name.

```go
package main
```

you can use the go doc tool to see the documentation for a module. This is a great way to see the documentation for a module and see what functions are available in the module.

```bash
go doc
```

https://pkg.go.dev/github.com/soypete/Meetup-Go-Graphql-Scraper@v0.1.1/auth

If you want to import an external module, you can do so by using the `import` keyword followed by the module path. This is a common cli package that is used to parse command line arguments.

```go
import "github.com/urfave/cli/v2"
```

If you have not used that module before, you can run the `go get` command to download the module and add it to your `go.mod` file.

```bash
go get github.com/urfave/cli/v2
```

The `go get` command will download the module into your go path and add it to your `go.mod` file. this allows the compiler to access the module when building your project.

If you want to import a package from the standard lib, you will only specify the module name. For example, to import the `fmt` package, you would write:

```go
import "fmt"
```

The `go mod` command is used to manage Go modules. To import modules you first need to setup your project as a go module. This is done by running the `go mod init` command.

```bash
go mod init github.com/Soypete/{myproject}
```

The `go mod tidy` command to download the dependencies for your project. Running `go mod tidy` will create a `go.sum` file that contains all the information about modules subdependencies and what hash/version of those subdepemdencies. The `go mod vendor` command is used to create a `vendor` directory that contains the dependencies for your project. This is useful if you want to distribute your project with its dependencies.

Modules is go contain all the logic needed to run a specific task or set of tasks. For example in this scraper projectwe have the following modules:

```bash
.
├── auth
│   ├── auth.go
│   └── kolla.go
├── meetup
│   ├── analytics.go
│   ├── api.go
│   ├── api_test.go
│   ├── meetup.go
│   └── meetup_test.go
├── go.mod
├── go.sum
└── main.go
```

The `auth` module contains all the logic needed to authenticate with the meetup api. The `csv` module contains all the logic needed to read and write csv files. The `meetup` module contains all the logic needed to interact with the meetup api. The `main.go` file is the entry point of the program and contains the `main` function. The directory name is often the same as the module name, but it doesn't have to be.

Every directory you great in you project is exportable as a module by default. This means that you can import the module into files in other directories and projects by using the module path. There are two main ways to stop your modulesand modules resources (functions, variables, etc) from being exported. The first is to use a lowercase letter at the beginning of the module name. This will make the module private and only accessible within the directory. The second way is to use an `/internal` directory. And module defined in the `/internal` directory is only accessible within the project it is defined in and not in any other project.

A lot of programers in the tend to user `/interal` directories like `/src` or `/lib` directories for all their project specific code. Personally, I don't care for this convention. I will talk about some other conventions that I use in the next section.

Using a `/cmd` directory is a common convention for Go projects. This directory is used to store types and methods that are meant be called by a user (for instance via a CLI). In a monorepo you will find that many different cli tools are stored in the `/cmd` directory. In microservice architecture, you will find it is just configuration and any additional functions or executables that are needed to run the service.

Using a `/test` directory is another common convention for Go projects. This directory is used to store test files that are used to test the code in the project. I also don't use this convention as I can use the `*_test.go` file naming convention to store my test files in the same directory as the code I am testing, but it can be helpful for larger test suites such as integration tests or end-to-end tests. If you do that, make sure to use the \\+build:ignore directive to exclude the test files from the build and only run in the correct environment.

### Go tooling: `go build`, `go install`, `go fmt`, `go vet`, `go get`

The go binary comes with a set of tools that help you manage your Go projects. The `go build` command is used to compile your Go program. The `go install` command is used to compile and install Go programs as binaries. The `go fmt` command is used to format your Go source code. The `go vet` command is used to check your Go source code for errors. The `go get` command is used to download and install packages and dependencies for your Go project.

To explore the Go tooling, run the following commands in your project directory:

```bash
go help
```

This will display a list of available commands and their descriptions. You can run `go help <command>` to get more information about a specific command.

The most common commands you will use are `go build`, `go run`, and `go test`. The `go build` command is used to compile your Go program. The `go run` command is used to compile and run your Go program. The `go test` command is used to run tests for your Go program. `go build` and `go tests` are commonly used in CI/CD pipelines to build and test your Go programs.

## Part 3: Basic Syntax and Variables (15 minutes)

### Variable declaration

In go there 2 ways to store variables in memory. `var` is for variatic values. Values that have the potential to change. `const` is for variables that are never going to change after created. Typically `const` values are global to their respecive module. If a `const` is declared in the `main` mondule it can only be access in the `main` module, but all other const values can be access outside their values if they are exported (capitalized) but including their modules import statement. The type can be specified with out a value when you use the `var` and `const` keywords. If you use the `:=` expression the type is always infered. (I think we can detect type)

In Go, you can declare the type or let the type be inferred by (reflection)[]. Go is a staticly typed language, so once the variable type is set by the compiler, it cannot be changed.

```go
const githubPath = "github.com/Soypete"
const Repo = "learning-go"

var name string = "Pete"
var age int = 30
var isCool bool = true
height := 6.2
```

(CamelCase)[https://en.wikipedia.org/wiki/Camel_case] is the typical convention used when the naming things in Go. Names in Go are typically short, but descriptive. Interface names should be actions like `Reader` or `Writer`. Names should not be stuttery on import, for example the package `client` should not have a struct type client because then import your be `client.Client`. It would be better to be specific with the package name specifing the type of client in the package name so the import would be `http.Client` or `grpc.Client`.

In opensource go code, you name see a lot of single letter name variables. The best practice for this is a practice of scoping. If is ok tp use a single name variable in a `for` loop or a `select` statement because the scope is localized and it is only extended a few lines. It is also common to see single letter variables names as the method delimiter.
if a single name variable used in one methods it should be used across all methods for readability.

Code should be self documenting. In go, go docs are generated by default for all exported variables. Including descriptiptive docstrings and package descriptions are a great way to keep code self documenting. The other ways it to use clear variable names that should what data is being used and what is being acted upon.

Variables are scoped to the functions, loops, and packages that they are defined in. If you want to share data outside of it's scope it needs to be passed as a paramenter in a function ot shared across a go channel.

<!-- I don't know what wanted above, but it is rought. I think it would make more sense with code -->

### Basic data types

In Go, everything can be thought of as a type. We will start with the primitive types and then move on to custom types.

from [go docs](https://go.dev/ref/spec#Types)

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE 754 32-bit floating-point numbers
float64     the set of all IEEE 754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
uint        either 32 or 64 bits
int         same size as uint

bool       true or false
```

These types always contain a value in the allocated memory. If you do not set the value at allocation time, that value is set to zero.

Complex types are a combination of one or more types. These are strings, maps, structs, slices, and arrays. A slice is a grouping of many variables of the same type. It has a backing array where the data is stored. A struct is an object that is comprised of many types. A map is key-value store. The key and value can be of any type. Functions are also type in go, they can be passed as parameters and returned as functions. This allows for a more functional style approach to programming.

Custom types are a way to define type inheretance. A good example of this is the `time.Duration` type in the standard langauge. This is a type that is a `int64` but has a specific set of methods that can be called on it. This way it can have interfaces that are applied to it to enforce behavior. Using custom types allows your code to be more readable and maintainable. It also allows for more complex logic to be applied to the type.

There are times when a type will not be known before runtime. In these cases we have the tools like reflection, `any` and generics. The `any` type is an placeholder type that can be used to store any type. This is useful when you don't know the type of the data that you are going to be storing. The `any` type is used in the `json` package to store the data that is being unmarshalled. You have to pass in the pointer to the type as your parameter and then the type is infered via reflection.

Before generics, you would have to write a function for each type that you wanted to sum. In this example we are talking the sum of all these numeric values in a map. We can do this on a number of different static types so it makes sense to implement on one function that can run on all these numeric types.

```go
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```

With generics you can write a single function that can sum any type that satisfies a set of constraints. The `comparable` keyword is a constraint that is used to make sure that the type can be compared with the `==`, and `!=` operators. So any type that is comparable can be used as `K`. In this function we are only using `int64` and `float64` used as `V`. Make sure that if you choose to you generics make sure you continue to understand their implementation and some of the memory constraints.

```go
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

### Basics of pointers in Go

Pointers are byte addresses to where a value is stored. Pointers are the best way to update a value across functions. `*` is the pointer dereference. This let's us know that they type is defined as a pointer. The `&` is a reference. It allows you to pass a type as it's pointer.

Pointers have a binary state but a variable value. Pointer are either allocated or nil. Anytime a pointer is called, but not allocated you will get a nil pointer dereference error. This is a runtime err

## Part 4: Functions in Go (10 minutes)

### Defining and calling functions

A function in Go is all the logic contained with in the `{}` following the `func` keyword. A method is a function that is called on a struct object. Methods can be used to satisfy an interface. Interfaces is Go are a set of methods that define a perscribed set of generic functionality. For example, the `Reader` interface has a `Read()` method. Any type that implements this interface must also implement their own `Read()` function. This is a great way to enforce behavior in go and allows for engineers to add custom logic for more complex use cases. - method vs function

```go
type MyError struct {
    When time.Time
    What string
}

func FormatError(e MyError) string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
```

### functions as a type

Functions are types in Go. This means you can pass functions around as types, define functions anonymously, and allow for generic use cases by allowing functions to be passed. Go is not a strictly Object Oriented, niether is it a strictly functional language. My mixing the two, you get really dynamic functionality in your go programs and can tackle complex problems.

```go
handler := func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

### Anonymous functions

we can define the above function as a anonymous function. This is a function that is defined without a name. This is useful when you want to pass a function as a parameter to another function. This is a common pattern in go and is used in the `http` package to define the handler for a route.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

```
