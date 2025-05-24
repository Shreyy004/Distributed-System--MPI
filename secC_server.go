package main
import (
"fmt"
"net/rpc"
)
type StringOperations struct {}

type StringArgs struct {
Str string
Index int
Char string
}

type ConcatArgs struct {

}

func (s *StringOperations) insert (args StringArgs, reply *string) error {
if (args.Index<0 || args.Index>len(args.Str)) {
*reply = "index out of bound"
return nil
}
*reply = args.Str[:args.Index] + args.Char + args.Str[args.Index:]
return nil
}

func main() {
var server = &StringOperations{}
rpc.Register(server)
listener, err := net.Listen("tcp", ":1234")
if (err!=nil) {
fmt.Println("listener error...")
return
}
defer listener.Close()
for {
conn, err := listener.Accept()
if (err!=nil) {
fmt.Println("connection error...")
continue
}
go rpc.ServeConn(conn)
}
}
