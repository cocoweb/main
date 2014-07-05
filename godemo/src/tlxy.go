package main

import "fmt"
import "lxy"

func main() {
	ss := new(lxy.Student)
	ss.SetName("lxy")
	ss.SetAge(20)
	dd := new(lxy.Director)
	dd.Name = "director"
	dd.Student = *ss
	var ii lxy.IPeople
	ii = dd
	ii.SetName("test")
	fmt.Println(ii.GetName())
	fmt.Println(dd.Student.String())

	oo := new(lxy.OldStudent)
	oo.SetAge(50)
	oo.SetName("Old man")

	fmt.Println(oo.GetName())

}
