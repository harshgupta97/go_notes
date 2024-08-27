# Go Language

## Go use case

- For performant application
- Running on scaled, distributed system

## Challenges of Multi-Threading

- Multiple user editing the same document
- Multiple users booking at the same time


NOTE:

- Concurrency is about dealing with lots of things at once

- Developers need to write code to prevent conflicts

    - when task run in parallel
    - accessing shared data


## Go characterstics

- Simple and readable syntax of dynamically typed language like python

- Efficiency adn safety foa lower-level, statically typed language like c++

- Fast build time, start up and run

- Requires fewer resources

- Compiled language

## Common commands

- Command to create a new module

```
go mod init <module_name>
```

- Compiles and run the code

```
go run <file_name>
```



## Variable and Constants

```
var key = "value"
```

```
fmt.Printf()
```

Print formated data


## Data Types in Go

- Strings
- Integers
- Boolans
- Maps
- Arrays


## Array and Slices

### Array

```
var <key> [<array_size>]<array_type> = [<array_size>]<array_type>{<place default values>}
```

### Slices

- Slice is an abstraction of array

- Slices are more flexible and powerful: variable-length or get an sub-array of its own

- Slices are also index-based and have a size, but is resized when needed

```
var <key> []<array_type> = []<array_type>{<place default values>}
```

One common method to add item to slices

```
append(slice_name, value)
```

## Loops

```
for {
    // code
}
```

## For-Each loop

### Range

- Range iterates over elements for different data structures (so not only array and slices)

- For array and slices, range provides the index and value for each element

## Blank Identifier

- To ignore a variable you don't want to use

- So with Go you need to make unused variable explicit

## If-Else

```
if statement {

}
```

## Continue statement

- Causes loop to skip the remainder of its body

- And immediately retesting its condition (in our infinite loop our condition is always true)

## Condition in For-loop



## Package level Variables

- Define at the top outside all functions

- They can be accessed inside any of the functions

- And in all files, which are in the same package

```
// We cannot use the below in package level variable
key := value
```

## Package

- Go programming are organizaed into packages

- A package is a collection of Go files


## 3 Level of Scope

- Declaration within function, can be used only within that function

- Local

- Package

- Global

## Maps

- Maps unique keys to values

- YOu can retrive the value by using its key later

```
var key = make(map[string]string)
userData["key"] = value

strconv.FormatUint(uint64(userTickets), 10)

// initialize a list of maps
var key = make([]map[string]string, 0)
```

## Struct

- stads for structure

- Can holds mixed data type

## Type `type`

- The type keyword creates a new type, with the name you specify

## Concurrency

```
go statement
```

## WaitGroup

- waits for the launched goroutine to finish

- Package "sync" provides basic synchronizaation functionality

- Add: Sets the number of goroutines to wait for (increase the counter by the provided number)

```
var wg = sync.WaitGroup{}

wg.Add(1)
wg.Done()
wg.Wait()
```

Goroutine

- Go is using, what's called a "Green thread"

- Abstraction of an actual thread

- Managed by the go runtime, we are only interacting with these high level goroutines

- Cheap & lightweight

- You can run hundred of thousands or mmillions goroutine without affecting the performance

