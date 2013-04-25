// lxy project lxy.go  
package lxy

import "strconv"

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
