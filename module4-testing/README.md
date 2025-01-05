# Module 4: Testing Go Programs (30 minutes)

## **Learning Objective: By the end of this module, learners should be able to write and execute tests in Go, use table tests, measure test coverage, and understand the basics of the Go testing framework.**

## Part 1: Introduction to Testing in Go (10 minutes)

Testing is an essential part of software development. Go has a testing package built into the Standard Library. This is the de facto testing framework for unit tests, integration tests, fuzzing, and benchmarking. In some code bases, you will see tests organized in separate testing packages. Theses files will start with `package {name}_test`. This was originally to let the compiler know not to add the functions into the compiled binary. The more contemporary paradigm in Go, is that tests are written in the same package as the code they are testing in separate test files are named `*_test.go`. These files are not included in the compiled binary.

```go
package framework_test

import (
    "testing"
)

func TestGetFramework(t *testing.T) {
    // test code
}
```

### Writing a Test

Here is a simple function that we want to test:

```go
func Add(a, b int) int {
    return a + b
}
```

To write a test for this function, we create a new file named `add_test.go` in the same package:

```go
func TestAdd(t *testing.T) {
    result := Add(1, 2)
    if result != 3 {
        t.Errorf("Add(1, 2) = %d; want 3", result)
    }
}
```

To run the tests we use the `go test` tool. This tool autodetects test functions in the provided path.

```bash
go test ./add-example
```

For test functions to be run by the `go test` command, they must have a name that begins with `Test`. These are called "capital-T" tests.

The `go test` tool uses the type `testing.T` type to track is a test passes or fails. This type is passed into the function signature are a parameter. All "captal-T" tests need to have the funtion signature `func(t *testing.T)`. In the example above, we use the `Errorf` method to report a test failure.

To run the tests, you can use the `go test` command in the package directory:

```bash
go test {path-to-package}
```

## Part 2: Table-Driven Tests (10 minutes)

Table-driven tests are a common pattern for writing testing in Go. They allow for creating variable input tests in a structured way an allow you to follow a more DRY(Do not Repeat Yourself). Instead of writing multiple test functions for different test cases, you can use a single test function with a table of values that define the test cases.

Here is an example of a table-driven test for the `Add` function:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }
    tt := tests
    for _, tc := range tt {
        t.Run(fmt.Sprintf("Add(%d, %d)", tc.a, tc.b), func(t *testing.T) {
            got := Add(tc.a, tc.b)
            if got != tc.want {
                t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
            }
        })
    }
}
```

The `t.Run` method allows you to run multiple sub-tests for each "capital-T" test. This is useful for better test output and organization. The `t.Run(name, func(t *testing.T))` method takes a test name and a test function as arguments. If we run the above test, we will see the output for each test case. If we run the test with the `-v` flag, we will see the output for each sub-test.

This command will test all the functions in the working directory and all subdirectories in verbose mode:

```bash
go test -v ./...
```

## Part 3: Test Coverage and Other Testing Tools (10 minutes)

Go has built-in support for measuring test coverage. You can use the `go test` command with the `-cover` flag to generate a coverage report. For example:

```bash
go test -cover {path-to-package}
```

This can be useful to identify areas of your code that are not covered by tests. It can also be run with the `-coverprofile` flag to generate a coverage profile that can be used for further analysis. This allows you to explore test coverage in the browser using the `go tool cover` command.

Let's run the test coverage for the `Add` function:

```bash
go test -coverprofile=coverage.out {path-to-package}
go tool cover -html=coverage.out
```

### Benchmarking and Main

There are two other types in the `testing` package. Benchmarks are used to measure the performance of your code. They are written in the same way as tests but the type `*testing.B` is passed to the benchmark function. Here is an example of a benchmark from the [davecheney fibonaci example](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go). This is the function that calculates the nth fibonaci number:

```go
func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}
```

This is the benchmark for that function:

```go
func BenchmarkFib10(b *testing.B) {
        // run the Fib function b.N times
        for n := 0; n < b.N; n++ {
                Fib(10)
        }
}
```

You can run it with this flag:

```bash
go test -bench=. -benchmem=true -benchtime=10s
```

Benchmarks are great for testing the speed, cpu usage, and memory of your code. It runs the function multiple times and averages the results. I specified to run this function as many times as can in 1 second or it maxes out at 10,000,000. This is a great tool for optimizing your code.

There are two additional ways of testing that I consider to be "Expert Mode". Fuzz testing is newer to go, it was realesed in 1.18. It is a way to test your code with random inputs. This is great for finding edge cases and bugs in your code. Main is a way to test your code as a standalone program when extra setup is needed or teardown is needed. This is great when you need to to run a server or a database to test your code.
