## **Learning Objective: By the end of this module, learners should be able to write and evaluate logical expressions in Go.**

---

## Part 1: Operators in Go (10 minutes)

If you want to make a comparison in Go, you can use the following operators:
* != 
* ==
* <=
* >=

The `&&` (and) operator is used to combine two conditions. Both conditions must be true for the statement to be true. The `||` (or) operator is used to combine two conditions. If either condition is true, the statement is true.

You can also use the `!` (not) operator to negate a condition.  If you have a function that returns a boolean value, you can use the `!` operator to reverse the value. This if a very common pattern in Go. Here is an example of checking if a value exists in a map:

```go

if _, ok := myMap["key"]; !ok {
    fmt.Println("Key does not exist")
}
```

The comparision operators are most often used in `if` or `switch` statements to determine the flow of a program. 

If you want to increment a variable by one, you can use the `++` operator. If you want to decrement a variable by one, you can use the `--` operator. You can also use the `+=` operator to increment a variable by a specific value.

These operators are used in loops to control the flow of the program. 


## Part 2: Control Structures (10 minutes)

In Go, you can define an `if` statement without the else. That means, that if you have some two options and one is the default behavior you can put the optional behavior behind a single if statement.

<!-- This would be an idea place for code. -->

If there are more that two options for behavior instead of using the keyword `else if`, engineers will typically use the `switch` keyword with and provided conditional `cases`. 

## Part 3: Looping Constructs (10 minutes)

There are 3 major syntaxes of `for` loops in Go: c style `for` loops, `range` style `for` loops, and `while` style `for` loops.

C style for loops are the most common and are used when you know the number of iterations you want to perform[1]. This is also one of the only times that you will purposefully use a `;` in Go.[2]

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

In the above example, the loop will run 10 times and print the value of `i` each time. The first statement `i := 0` is the initialization statement. It defines the variable `i` for the scope of the loop. The second statement `i < 10` is the condition statement. The loop runs until the condition is met. The third statement `i++` is the increment that will be executed after each iteration of the loop. We are incrementing the value of `i` by 1 each time.

You can also declare this loop without using all of the three parts. This example comes from the [Go Tour](https://go.dev/tour/flowcontrol/3).

```go 

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

```     

[1] As of go 1.22 you can use `range` over an integer to get the same effect as a c style for loop. If you just want to run a loop a specific numbner of times and will not be using the value of the index, then this `range` syntax is a much more comforable way to express the same thing.

```go
for i := range 10 {
    fmt.Println("Hello")
}
```

[2] The `;` is actually used in Go at the end of everyline, but the compiler will automatically insert it for you. This is why you don't see it in most Go code.

The most common type of loop in Go is the `range` loop. This is used when you want to iterate over a collection of items. The basic syntax just return the index of the item in the collection.

```go
for i := range mySlice {
 fmt.Println(mySlice[i])
 }
 ```

 Access the value of the item in the collection by using the index can feel a little clunky. You can also use the `range` keyword to get the index and the value of the item in the collection.

 ```go 
 for i, v := range mySlice {
     fmt.Println(v)
    }
    ```

The above example performs the same operation as the previous example, but it uses the `v` variable to store the value of the item in the collection. The `i` variable stores the index of the item in the collection. The compiler will get mad in this case becasue we are not using the index. You can use the `_` to ignore the index.

```go
for _, v := range mySlice {
    fmt.Println(v)
}
```

If you want to iterate over a map, you can use the `range` keyword to get the key and value of each item in the map.

```go
for k, v := range myMap {
fmt.Println(k, v)
}
```

The naming of the variables is up to you, but it is common to use `i` for index. If you are iterating over a map, it is common to use `k` for key and `v` for value. If you are iterating over a slice, the slice is often named as a plural of the object and the value is often named as a singular of the object. 

```go
for _, movie := range movies {
    fmt.Printf("The movie is %s\n", movie)
}
``` 

Condtional `for` loops run as long a certain condition is met. This is the most common type of loop in Go. 

```go
for i < 10 {
    fmt.Println(i)
    i++
}
```

```go
for rows.Next() {
    var name string
    var age int
    err := rows.Scan(&name, &age)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name, age)
}
``` 

you can also run an infinite loop by leaving the condition blank. This is often used when you want to run a loop until a certain condition is met. 

```go
for {
    fmt.Println("Hello")
}
```

To exit a `for` loop early, you can use the `break` keyword. This will exit the loop and continue executing the code after the loop. 

```go
for {
    fmt.Println("Hello")
    break
}
```

To skip a single iteration of a loop, you can use the `continue` keyword. This will skip the current iteration and continue with the next iteration of the loop. 

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        continue
    }
    fmt.Println(i)
}
```

Coming in Go 1.24 is the option to range over a function. This a more complex generic pattern that can be very powerful, but doesn't have many practical use cases.

<!-- add link to go 1.24 release notes and video -->
        - range over function
