package main 

import "fmt" 

// say Hi to someone
func SayHi(name string) string {
   return fmt.Sprintf("Hi + v1, %s", name)
}

func main() {
	fmt.Println("hello world")
}
