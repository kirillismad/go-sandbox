package examples

import (
	"fmt"

	"github.com/google/uuid"
)

func DemoStructs() {
	p1 := NewPerson(uuid4(), "Kirill", 27)
	fmt.Println(p1)
	p2 := NewPerson(uuid4(), "Tata", 32)
	fmt.Println(p2)

	(*p1).setAge(rint(100))
	fmt.Println(p1)
	p1.setAge(rint(100))
	fmt.Println(p1)

	fmt.Println((*p1).getName())
	fmt.Println(p1.getName())
}

type Person struct {
	id   uuid.UUID
	name string
	age  int
}

func NewPerson(id uuid.UUID, name string, age int) *Person {
	p := new(Person)
	p.id = id
	p.name = name
	p.age = age
	return p
}

func (p Person) String() string { // p is a copy object
	return fmt.Sprintf("Person(name:%v, age:%v)", p.name, p.age)
}

func (p *Person) setAge(value int) { // p is the pointer to given object
	p.age = value
}

func (p Person) getName() string {
	return p.name
}
