# Learning Go for Software Engineers

## Module 1: Introduction to Go (60 minutes)
**Learning Objective: By the end of this module, learners should be able to set up the Go development environment and understand the basics of Go syntax and structure.**
- Part 1: Setting Up Go Development Environment (10 minutes)
- Part 2: Go Program Structure (15 minutes)
- Part 3: Basic Syntax and Variables (15 minutes)
- Part 4: Functions in Go (10 minutes)

## Module 2: Logical Expressions in Go (30 minutes)
**Learning Objective: By the end of this module, learners should be able to write and evaluate logical expressions in Go.**
- Part 1: Operators in Go (10 minutes)
    - Arithmetic, comparison, logical operators
- Part 2: Control Structures (10 minutes)
    - If-else statements
        - In Go it is possible to have an if statement without an else
    - Switch statements
- Part 3: Looping Constructs (10 minutes)
    - For loops
        - C style
        - While style
    - Range loops
        - rang over num
        - range over function
    
## Module 3: Go Idioms and Advanced Patterns (30 minutes) 
**Learning Objective: By the end of this module, learners should be able to understand and implement key Go idioms and advanced patterns including structs, interfaces, inheritance, and channels.**
- Part 1: Structs and Interfaces (10 minutes)
    - Defining structs
    - Accessing and modifying struct fields
    - Using interfaces in functions
- Part 2: Inheritance in Go (10 minutes)
    - Using the interface and methods
- Part 3: Advanced Patterns: idioms, Channels, and memory (10 minutes)
    - Idiomatic Go - How to detect and avoid code smells
    - Linters, checks, and best practices
    - Channels - Sending and receiving data through channels
        - Goroutines
        - Select statement
    - Memory model -How to share memory?

## Module 4: Testing Go Programs (30 minutes)
**Learning Objective: By the end of this module, learners should be able to write and execute tests in Go, use table tests, measure test coverage, and understand the basics of the Go testing framework.**
- Part 1: Introduction to Testing in Go (10 minutes)
    - Overview of the testing package
    - Writing simple test functions
- Part 2: Table-Driven Tests (10 minutes)
    - Structuring table-driven tests
    - Advantages of table-driven tests
- Part 3: Test Coverage and Other Testing Tools (10 minutes)
    - Measuring test coverage
    - Additional testing tools and practices (e.g., benchmarks, mock testing)

## Module 5: Building a CLI Application with Go (30 minutes)
**Learning Objective: By the end of this module, learners should be able to build and run a basic CLI application using Go.**
- Part 1: Introduction to CLI Applications (10 minutes)
    - Basics of CLI applications
    - Setting up a CLI project
- Part 2: Building the CLI Application (10 minutes)
    - Handling user input
        - Reading from the command line
        - Command-line flags vs. arguments
    - Executing commands
        - Using the `os/exec` package
        - Running external commands
- Part 3: Packaging and Distributing the CLI Application (10 minutes)
    - Compiling and distributing binaries
    - Cross-compilation
    - GoReleaser - intro
