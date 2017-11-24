package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func multiply(x, y int) int {
  return x * y
}

func main(){
  fmt.Println(add( multiply(6,7) , 13))
}
