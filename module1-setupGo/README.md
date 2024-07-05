# Module 1: Introduction to Go (60 minutes)
**Learning Objective: By the end of this module, learners should be able to set up the Go development environment and understand the basics of Go syntax and structure.**

---

## Part 1: Setting Up Go Development Environment (10 minutes)

###Installing Go
**Install via website: https://golang.org/dl/**
    We are glad you chose to learn Go. In this section, we will guide you through the installation process of Go on your system. Go is distributed as a binary package, and you can download the installer from the official Go website. This is the recomended way to install Go on your system if you are unfamiliar with package managers.
    [Download Go](https://golang.org/dl/)
    Add image of the download page with the download button highlighted
    - Click on the download button to download the installer for your operating system.
    - Run the installer and follow the on-screen instructions to install Go on your system.

** install via package manager**
    Unix-based systems like Linux and macOS have package managers that can be used to install Go. If you are using a package manager like Homebrew on macOS, you can install Go using the following command:
    ```bash
    brew install go
    ```
    This will download and install Go on your system.

** install via  webi ** 
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

### Workspaces
If you have a project wiht many contributors, many modules, or is a monorepo, you should consider looking at [workspaces](https://go.dev/doc/tutorial/workspace](https://go.dev/doc/tutorial/workspaces)

## Part 2: Go Program Structure (15 minutes)
### Understanding Go file structure
<!-- The add a go repo file structure image here -->

The executable Go programs are made up of packages. A package is a collection of Go source files that are compiled together. Each package has a unique name, and the name of the package is the same as the name of the directory that contains the package source files. The `main` package is a special package that contains the `main` function, which is the entry point of the program. The main file can me named anything, but it is common to name it `main.go`. When compiling a Go program, the `main.go` file becomes a `main` binary file that can be executed on your system. 


        - Modules, imports, and workspaces
### Writing and running your first Go program
        - Create a simple `hello.go` program
        - Run the program using `go run hello.go`
        - Explain the output and the `main` function
### Go tooling: `go build`, `go install`, `go fmt`, `go vet`, `go get`
        - explain the go build main.go vs go build .

## Part 3: Basic Syntax and Variables (15 minutes)
### Variable declaration
        - `var`, `const`, `:=`
        - Naming conventions
        - scope of variables and memory sharing
### Basic data types
        - `int`, `float`, `string`, `bool`
        - Type inference
        - Zero values
        - custom types
        - generics
### Basics of pointers in Go
        - `&` and `*` operators
        - Passing by value vs. passing by reference
## Part 4: Functions in Go (10 minutes)
### Defining and calling functions
        - method vs function
### functions as a type
