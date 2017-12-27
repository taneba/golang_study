# go

## Commands
### go build
build only

### go run
build + execute


### go fmt
Formats all the code in each file in the cuurnt directory

### go install
install 

### go get

### go test

## Types of Packages

- Executable
- Reusable

### Executable
Generates a file that we can run

### Reusable
Code used as 'helpers'.
Good place to put reusable logic

## What does 'package main' mean?

*'main' is a Executable package*

`package main` defines a package that can be compiled and then executed. *Must have a func called 'main'*

## How is the foo.go file organized?

1. Package declaration

```
package main
```

2. Import other packages that we need

```
import "fmt"
```

3. Declare function, tell Go to do things

```
func main() {
  fmt.PrintLn("Hi there")
}
```

---------
spit out
吐き出す

It's actually the name of the package that you use that determines whether you are making an executable or dependency type package
