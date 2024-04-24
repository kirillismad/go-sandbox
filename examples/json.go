package examples

import (
	"encoding/json"
	"fmt"
)

func DemoJson() {
	demoStructsJson()
}

type Entity struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

func NewEntity(id int, name string, salary int) *Entity {
	entity := new(Entity)
	entity.Id = id
	entity.Name = name
	entity.Salary = salary
	return entity
}

type Entity2 struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
}

func demoStructsJson() {
	entity := NewEntity(1, "Kirill", 1000)
	entityMarshalled, _ := json.Marshal(entity)
	fmt.Printf("%v\n", string(entityMarshalled))

	entityUnmarshaled := new(Entity2)
	_ = json.Unmarshal(entityMarshalled, entityUnmarshaled)
	fmt.Printf("%+v\n", entityUnmarshaled)
}
