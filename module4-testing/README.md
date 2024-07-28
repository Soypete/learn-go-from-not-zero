# Module 4: Testing Go Programs (30 minutes)
**Learning Objective: By the end of this module, learners should be able to write and execute tests in Go, use table tests, measure test coverage, and understand the basics of the Go testing framework.**
---

## Part 1: Introduction to Testing in Go (10 minutes)

Testing is an essential part of software development. Go has a  testing package built into the Standard Library. This is the defacto testing framework for unit tests, integration tests, and benchmarking. In Go, tests are written in the same package as the code they are testing, and test files are named `*_test.go`.

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

For test functions to be run by the `go test` command, they must have a name that begins with `Test`. These are called "capital-T" tests. The core feature of the testing package is the `testing.T` type that is passed to the test function. This type provides methods for reporting test failures and logging test output. In the example above, we use the `Errorf` method to report a test failure.

To run the tests, you can use the `go test` command in the package directory:

```bash
go test {path-to-package}
```

## Part 2: Table-Driven Tests (10 minutes)

Table-driven tests are a common pattern for testing in Go. They allow for creating variable input tests in a structured way. Instead of writing multiple test functions for different test cases, you can use a single test function with a table of values that define the test cases.

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

The `t.Run` method allows you to create subtests for each test case. This is useful for better test output and organization. The `t.Run(name, func(t *testing.T))` method takes a test name and a test function as arguments.

## Part 3: Test Coverage and Other Testing Tools (10 minutes)

Go has built-in support for measuring test coverage. You can use the `go test` command with the `-cover` flag to generate a coverage report. For example:

```bash
go test -cover {path-to-package}
```

This can be useful to identify areas of your code that are not covered by tests. It can also be run with the `-coverprofile` flag to generate a coverage profile that can be used for further analysis. This allows you to explore test coverage in the browser using the `go tool cover` command.

### Benchmarking and Main

There are two other types in the `testing` package: benchmarks and Main. Benchmarks are used to measure the performance of your code. They are written in the same way as tests but the type `*testing.B` is passed to the benchmark function. Main tests are user-defined tests that can be run using the `go test` command against a `main()` function.

I have never use the `Main` tests, but I have used benchmarks. Here is an example of a benchmark from the davecheney fibonaci example. This is the function that calculates the nth fibonaci number:

```go

```

This is the benchmark for that function:

```go

```

You can run it with this flag:

```bash

```

Benchmarks are great for testing the speed, cpu usage, and memory of your code. It runs the function multiple times and averages the results. I specified to run this function as many times as can in 1 second or it maxes out at 10,000,000. This is a great tool for optimizing your code.