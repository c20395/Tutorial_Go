package main 
import ( 
  "fmt"  // ; "unicode/utf8"
)

func nums(i ...int) { 
  fmt.Println(i) 
}

func sum(nums ...int) int { 
  total := 0 
  for _, num := range nums { 
    total += num 
  } 
  return total 
} 

func main() { 
nums(99,100,101)  
  fmt.Println(sum(99, 100,101))
} 
