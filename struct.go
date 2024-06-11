package main
import "fmt"

type name string

type location struct {
x int
y int
}

type size struct {
width int
height int
}

type dot struct {
name
location
size
}

func getDots() []dot {

// Our first dot uses the var notation. This will result in all the fields having a zero value:

  var dot1 dot

// With dot2, we're also initializing with zero values:

  dot2 := dot{}

// To set the name, we use the type's name as if it were a field:

  dot2.name = "A"

// For size and location, we'll use the promoted fields to set their value:

  dot2.x = 5
  dot2.y = 6
  dot2.width = 10
  dot2.height = 20

// When initializing embedded types, you can't use promotion. For name, the result is the same but for location and size, you need to put more work into this:

  dot3 := dot{ 
  name: "B", 
  location: location{ 
    x: 13, 
    y: 27, 
  }, 
  size: size{ 
    width: 5, 
    height: 7, 
  }, 
} 

// For dot4, we'll use the type names to set data:

    dot4 := dot{} 
    dot4.name = "C" 
    dot4.location.x = 101 
    dot4.location.y = 209 
    dot4.size.width = 87 
    dot4.size.height = 43 

// Return all the dots in a slice and then close the function:

  return []dot{dot1, dot2, dot3, dot4}
}

// In main, call the function. Then, loop over the slice and print it to the console:

func main() { 
  dots := getDots() 
  for i := 0; i < len(dots); i++ { 
    fmt.Printf("dot%v: %#v\n", i+1, dots[i]) 
  } 
} 
