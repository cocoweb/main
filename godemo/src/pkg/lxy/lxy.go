package lxy

// lxy project lxy.go

import "strconv"
import "fmt"

type IPeople interface {
	SetName(string)
	GetName() string
}
type Student struct {
	Name string
	Age  int
}
type Teacher struct {
	Name   string
	Course string
}

type OldStudent struct { //匿名组合，完成继承
	Student
	year int
}

func (s *Student) SetName(name string) {
	s.Name = name
}
func (t *Teacher) SetName(name string) {
	t.Name = name
}
func (s *Student) GetName() string {
	return s.Name
}
func (t *Teacher) GetName() string {
	return t.Name
}
func (s *Student) SetAge(age int) {
	s.Age = age
}
func (s *Student) GetAge() int {
	return s.Age
}
func (s *Student) String() string {
	return "name is " + s.Name + ",age is " + strconv.Itoa(s.Age)
}
func (t *Teacher) SetCourse(course string) {
	t.Course = course
}
func (t *Teacher) GetCourse() string {
	return t.Course
}

type Director struct {
	 Student
	Name string
}

func (di *Director) GetName() string {
	fmt.Println("get director name")
	return di.Name
}
func (di *Director) SetName(name string) {
	di.Name = name
}
