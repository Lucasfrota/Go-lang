package main

import(
  "bufio"
  "fmt"
  "os"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter texte here: ")

  text, _ := reader.ReadString('\n')
  fmt.Println(text)
}
