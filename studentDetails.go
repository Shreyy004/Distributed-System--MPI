package main
import "fmt"

type Student struct {
name string
age int
cgpa float64
isEnrolled bool
}

func main() {
var greetings string = "Welcome!"
fmt.Println(greetings)
student1:=Student{"a", 20, 9.4, true} 
student2:=Student{"b", 20, 9.7, true}
student3:=Student{"c", 20, 8.9, true}
var studentDetails[] Student
studentDetails = append(studentDetails, student1, student2, student3)
studentDetails = append(studentDetails, Student{"d", 20, 9.5, true})

for _, student := range studentDetails {
fmt.Printf("name: %v, age: %v, cgpa: %v, enrolled: %v \n", student.name, student.age, student.cgpa, student.isEnrolled)
}

}
