package main

import "fmt"
import "lxy"

type director struct {
	lxy.Student
	Name string
}

func (di *director) GetName() string {
	fmt.Println("get director name")
	return di.Name
}
func (di *director) SetName(name string) {
	di.Name = name
}
func main() {
	ss := new(lxy.Student)
	ss.SetName("lxy")
	ss.SetAge(20)
	dd := new(director)
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
