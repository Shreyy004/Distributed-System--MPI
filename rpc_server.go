package main
import (
"fmt"
"net"
"net/rpc"
)
type API struct {}
type Item struct {
Num int
Value string
}
type Args struct {
ABC int
XYZ string
}
type Args1 struct {
Num int
}

func (a *API) AddItem (A Args, reply *string) error {
database = append(database, Item{A.ABC, A.XYZ})
*reply = "successfully added item"
return nil
}

func (a *API) DisplayContents(_ struct{}, reply *[]Item) error{
*reply = database
return nil
}

func (a *API) DeleteItem (A Args, reply *string) error {
for idx, contents := range (database) {
if (contents.Num == A.ABC) {
database = append(database[:idx], database[idx+1:]...)
*reply = "successfully deleted item"
return nil
}
}
*reply = "item not found"
return nil

}

func (a *API) EditItem (A Args, reply *string) error {
for idx, contents := range (database) {
if (contents.Num == A.ABC) {
database[idx].Value = A.XYZ
*reply = "successfully edited"
}
}
*reply = "item not found"
return nil

}

func (a *API) GetByNum(A Args1, Resp *string) error {
for _, contents := range (database) {
if (contents.Num == A.Num) {
*Resp = contents.Value
return nil
}
}
*Resp = "invalid"
return nil
}

var database []Item

func main() {
var api = &API{}
rpc.Register(api)
listener, err := net.Listen("tcp", ":1234")
if err != nil {
fmt.Println("listener error")
return
}
defer listener.Close()
for {
conn, err := listener.Accept()
if err != nil {
fmt.Println("connection error")
continue
}
go rpc.ServeConn(conn)
}

}
