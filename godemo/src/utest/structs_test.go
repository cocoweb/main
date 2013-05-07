package utest

import (
	_ "fmt"
	"testing"
	_ "os"
)

type Person struct {
	Id      int
	Name    string
	Contact struct { // 匿名结构成员
		phone    string
		address  string
		postcode string
	}
}

type Staff struct {
	Person
	StaffID int
}

func TestXYZ(t *testing.T) {

	user := Staff{
		Person:Person{Id: 1,Name: "Jhon"},
		StaffID: 20,
	}
	user.Contact.address = "asldfjlaksjdf"
	user.Contact.phone = "123123"
	user.Contact.postcode = "400000"

	//fmt.Println(user)
	t.Log(user)
	
}

func main(){

   TestXYZ(nil)
   
}