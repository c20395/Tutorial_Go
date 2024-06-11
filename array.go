
package main

import "fmt"

func main() {
	a:=[10]int{}
	for i:=0;i<len(a);i++ {
		a[i]=i+1
	}

	fmt.Println(a)
}
