package main
import (
"fmt"
)

func take_input(abc chan int) {
var op string
var a, b, res int
fmt.Println("enter operation:(addition/subtraction) ")
fmt.Scanf("%s", &op)
fmt.Println("enter a: ")
fmt.Scanf("%d", &a)
fmt.Println("enter b: ")
fmt.Scanf("%d", &b)
if op == "addition" {
res = a+b
} else if op == "subtraction" {
res = a-b
} else {
fmt.Println("invalid input")
}
abc <- res
}

func greetings() {
fmt.Println("welcome")
}

func main() {

myChannel := make(chan int)
go take_input(myChannel)
go greetings()
var result int
result =<- myChannel
fmt.Println(result)
}
