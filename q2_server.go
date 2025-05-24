package main
import (
"fmt"
"net/rpc"
)
type room struct {
total int
booked int
price int
}

func bookRoom(int type_room .....) {

fmt.Println("booking successful!")
}

func cancelRoom(int type_room ....) {

fmt.Println("booking cancelled")
}

func checkAvailability(int type_room ....) int {


return room_count
}

func main() {
var Rooms[] room
room1 := room{10, 0, 1000}
room2 := room{20, 0, 1500}
room3 := room{5, 0, 2000}
room4 := room{3, 0, 3000}
room5 := room{2, 0, 5000} 
Rooms = append(Rooms, room1, room2, room3, room4, room5)
rpc.Register(Rooms)
listener, err := net.listen("tcp", ":1234")
if (err!=nil) {
fmt.Println("error")
}
defer listener.Close()
for {
conn, err = listener.Accept()
if (err!=nil) {
fmt.Println("error")
continue
}


}
}

}


