package main
import (
"fmt"
"sync"
)

func take_input(abc chan int, x *sync.WaitGroup) {
defer x.Done()
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

func greetings(def chan string, y *sync.WaitGroup) {
defer y.Done()
def <- "welcome!"
}

func main() {
var wg sync.WaitGroup
myChannel := make(chan int)
anotherChannel := make(chan string)
wg.Add(2)
go take_input(myChannel, &wg)
go greetings(anotherChannel, &wg)
msgFrom_myChannel := <- myChannel
fmt.Printf("result: %v", msgFrom_myChannel)
msgFrom_anotherChannel := <- anotherChannel
fmt.Println(msgFrom_anotherChannel)
wg.Wait()

}
