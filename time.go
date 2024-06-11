package main
import ( "fmt" ; "time" )

func main() { 
t := time.Date(2019, time.November, 3, 13, 0, 0, 0, time.UTC)
chLoc, _ := time.LoadLocation("America/Chicago")
fmt.Println(t.In(chLoc))
}
