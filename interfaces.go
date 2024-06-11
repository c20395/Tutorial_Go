package main 

import ( "fmt" )

type Sayer interface {SayHello() interface{ String() string } } 
type Hello struct {} 
func (h *Hello) SayHello() interface{ String() string } {return h} 
func (h *Hello) String() string {return "Hello"} 

func main() { 
  var si Sayer = &Hello{} 
  fmt.Println(si.SayHello()) 
} 
