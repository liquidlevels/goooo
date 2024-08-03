package main 

import (
  "fmt"
  "example.com/greetings"
)

func main() {
  message := greetings.Hello("maballs")
  fmt.Println(message)
}
