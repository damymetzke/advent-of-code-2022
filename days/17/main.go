package main

import (
	"fmt"
	"log"
	"os"
)

func GetInput() <-chan string {
  result := make(chan string, 1)
  go func(){
    data, err := os.ReadFile("input/17")
    if err != nil {
      log.Fatal("Could not read file 'input/17':\n  * ", err)
    }

    result <- string(data)
    close(result)
  }()
  return result
}

func main() {
  fmt.Println(<-GetInput())
	fmt.Println("=-= PART 1 =-=")
	fmt.Println("=-= PART 2 =-=")
}
