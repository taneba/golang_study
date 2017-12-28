# OOP in go
By creating a new type with a function that has a receiver, we are adding a 'method' to any value of that type.

```
package main

import "fmt"

type book string

func (b, book) printTitle() {
  fmt.Println(b)
}

func main() {
  var b book = "Harry Potter"
  b.printTitle()
}
```

this example print out "Harry Potter", because book is a value of book type.

# function

go supports for returning multiple values from one function.

```
func getBookInfo() (string, int) {
  return "War and Peace", 1000
}

title, pages := getBookInfo()
```