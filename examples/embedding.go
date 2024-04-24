package examples

import "fmt"

func DemoEmbedding() {
	runnersRun()
	baseSlice()
}

func baseSlice() {
	fmt.Println(newline + fname(baseSlice))
	h := Human{
		Being:  Being{age: 26},
		Runner: Runner{maxSpeed: 44},
		job:    "Software Engeneer",
	}
	ch := Cheetah{
		Being:    Being{age: 3},
		Runner:   Runner{maxSpeed: 120},
		cubCount: 3,
	}

	s := []Beingable{&h, &ch}

	for _, item := range s {
		item.printAge()
	}
}

func runnersRun() {
	fmt.Println(newline + fname(runnersRun))
	s := []Runnable{
		&Human{
			Being:  Being{age: 26},
			Runner: Runner{maxSpeed: 44},
			job:    "Software Engeneer",
		},
		&Cheetah{
			Being:    Being{age: 3},
			Runner:   Runner{maxSpeed: 120},
			cubCount: 3,
		},
	}

	for _, r := range s {
		r.run()
	}
}

type Runnable interface {
	run()
	printSpeed()
}

type Beingable interface {
	printAge()
}
type Being struct {
	age int
}

func (b Being) printAge() {
	fmt.Println("Age:", b.age)
}

type Runner struct {
	maxSpeed int
}

func (r Runner) printSpeed() {
	fmt.Println("Max speed:", r.maxSpeed)
}

type Human struct {
	Being
	Runner
	job string
}

func (h *Human) run() {
	fmt.Println("Run like a human with speed:", h.maxSpeed)
}

type Cheetah struct {
	Being
	Runner
	cubCount int
}

func (c *Cheetah) run() {
	fmt.Println("Run like a cheetah with speed:", c.maxSpeed)
}
