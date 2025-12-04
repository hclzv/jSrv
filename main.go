package main

import (
	"fmt"
	"os"
)


func main() {
	fmt.Println("hello go !")
	fmt.Printf("os.Getenv(\"HJS\"): %v\n", os.Getenv("HJS"))

}



func ui()  {
	fmt.Println("this ui")
	fmt.Println("this vscode edit!!")
}
