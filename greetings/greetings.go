package greetings

import (
  "fmt"
  "errors"
)

func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("empty name, bro what?? u have no name")
  }

  message := fmt.Sprintf("Hi, %v. te apesta el culo!", name)
  return message, nil
}
