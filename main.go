package main

import "fmt"

func main() {
 var input uint
 
 fmt.Scanln(&input)
 
 if input % 2 == 0 {
    fmt.Print("Bob")
 } else {
     fmt.Print("Alice")
 }
    
}
