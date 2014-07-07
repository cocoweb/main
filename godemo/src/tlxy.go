package main

import "fmt"
import . "lxy"

func main() {
	ss := new(Student)
	ss.SetName("lxy")
	ss.SetAge(20)
	dd := new(Director)
	dd.Name = "director"
	dd.Student = *ss
	var ii IPeople
	ii = dd
	ii.SetName("test")
	fmt.Println(ii.GetName())
	fmt.Println(dd.Student.String())

	oo := new(OldStudent)
	oo.SetAge(50)
	oo.SetName("Old man")

	fmt.Println(oo.GetName())

}
