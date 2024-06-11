package main

import (
    "fmt"
//    "math/rand"
//    "strings"
//    "time"
//    "unsafe"
)

type Animal struct {
  Name string
  Age  int
}
type Dog struct {
  Animal // embedding
}

type PointXY struct {X,Y int}

func main() {
  var gaga = Dog{}
  pp := &PointXY{}
  fmt.Println(gaga.Age)
  fmt.Println(gaga.Age)
}
