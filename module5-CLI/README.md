# Module 5: Building a CLI Application with Go (30 minutes)
**Learning Objective: By the end of this module, learners should be able to build and run a basic CLI application using Go.**
---

## Part 1: Introduction to CLI Applications (10 minutes)

If you have been writing software int he past few years, then you have probably used a CLI application written in Go. Docker, Kubernetes, Terraform, and many other popular tools are written in Go. 

In this module we will write a simple CLI application. I encourage you to follow along and write the code as we go. This will help you understand the concepts better.

Most CLI applications are built using open-source libraries. Some of the popular libraries are:
- Cobra
- Viper
- Kingpin
- URFaveCLI
- Charm

For this example we will be keeping it simple and not using any libraries, but I encourage you to explore these libraries and use them to improve your CLI applications.

## Part 2: Building the CLI Application (10 minutes)

Our CLI application will be a password generator. There are a few requirements we have for this application:
- The user should be able to specify the length of the password
- the user should be able to specify special characters
- the user should be able to specify using numbers
- the user should be able to specify using uppercase letters

Let's start by creating a new directory for our project and creating a new file called `main.go`.

There are many ways to get user input in Go. When building a CLI application the two libraries needed are `os` and `flag`. The `os` package provides a way to order based arguments from the command line and the `flag` package provides a way to access flag denoted command-line arguments. You can combine os.Args and flag.Sets to create sub commands and flags. Subcommands are super powerful for robust CLI applications.

This example of a robust CLI application is from go By Example:
```go
https://gobyexample.com/command-line-subcommands
```

Flags are also used for any go program that needs to accept command line arguments.

This is a simple cli application that creates a password:

```go

```

### Wraping bash commands in Go

It is really easy to use the os/exec package to run bash commands in Go. Here is an example of how to run a bash command in Go:

```go
```

If you want to make a set to tools to improve developer productivity, you can wrap bash commands in Go and create a CLI application. This is a great way to improve your productivity and learn Go at the same time.


## Part 3: Packaging and Distributing the CLI Application (10 minutes)

### Compiling and distributing binaries

Once you have written the code for a cli it can be easily compiled into a binary via `go build`. If this binary lives in your `$PATH` you can run it from anywhere in your terminal. For the same of continuing the example of building a tool to for developer productivity, let's explore some ways to distribute the binary. 

`go install` allows you go install the binary in your `$GOPATH/bin` directory. If you have a go module setup and your code is in a git repository, you can use `go install` to install the binary from the repository.

I enjoy using the tool `GoReleaser` to distribute my binaries. GoReleaser is a tool that allows you to release Go projects as fast and easily as possible. It creates releases on GitHub, GitLab, and Gitea. It also supports creating binaries for Linux, Windows, and MacOS.

Here is a simple example of how to use GoReleaser:

```bash

```