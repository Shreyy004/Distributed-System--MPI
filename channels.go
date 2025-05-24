package main
import (
"fmt"
)
//channels help to communicate information between go routines; writing into same location of memory(i.e. channel)
/*
func main can communicate with go routines by reading from the same channel that all go routines can write into
*/
func main() {
myChannel := make(chan string)
go func() {
myChannel <- "data"
}()

msg := <-myChannel //waits for channel to close or a message to be received from this channel
fmt.Println(msg)

}
