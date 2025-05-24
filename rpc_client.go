package main
import (
"fmt"
"net/rpc"
)
type Item struct {
Num int
Value string
}
type Args struct {
ABC int
XYZ string
}

func main() {
var db []Item
client, err := rpc.Dial("tcp", "localhost:1234")
if err!= nil {
fmt.Println("connection error")
return
}
defer client.Close()
a := Item{1, "first"}
b := Item{2, "second"}
c := Item{3, "third"}
var reply string
client.Call("API.AddItem", Args{a.Num, a.Value}, &reply)
client.Call("API.AddItem", Args{b.Num, b.Value}, &reply)
client.Call("API.AddItem", Args{c.Num, c.Value}, &reply)
client.Call("API.DisplayContents", struct{}{}, &db)
for _, contents := range(db) {
fmt.Printf("%v, %v\n", contents.Num, contents.Value)
}
}
