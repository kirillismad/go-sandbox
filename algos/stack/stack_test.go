package stack

import (
	"errors"
	"fmt"
	"sandbox/gof"
	"testing"
)

func TestStackPush(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack[int]
		value int
	}{
		{
			name:  "FixedStack",
			stack: NewFixedStack[int](8),
			value: 1111,
		},
		{
			name:  "InfStack",
			stack: NewInfStack[int](),
			value: 2222,
		},
		{
			name:  "InfStack2",
			stack: NewInfStack2[int](),
			value: 3333,
		},
		{
			name:  "InfStack3",
			stack: NewInfStack3[int](),
			value: 4444,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.stack.Push(test.value)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestStackPopEmpty(t *testing.T) {
	tests := []struct {
		name  string
		stack Stack[int]
	}{
		{
			name:  "FixedStack",
			stack: NewFixedStack[int](8),
		},
		{
			name:  "InfStack",
			stack: NewInfStack[int](),
		},
		{
			name:  "InfStack2",
			stack: NewInfStack2[int](),
		},
		{
			name:  "InfStack3",
			stack: NewInfStack3[int](),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item, err := test.stack.Pop()
			if !errors.Is(err, ErrStackIsEmpty) {
				t.Error(item, err)
			}
		})
	}
}

func TestStackPushPop64Items(t *testing.T) {
	const count = 64
	tests := []struct {
		name  string
		stack Stack[int]
	}{
		{
			name:  "FixedStack",
			stack: NewFixedStack[int](count),
		},
		{
			name:  "InfStack",
			stack: NewInfStack[int](),
		},
		{
			name:  "InfStack2",
			stack: NewInfStack2[int](),
		},
		{
			name:  "InfStack3",
			stack: NewInfStack3[int](),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < count; i++ {
				err := test.stack.Push(i)
				if err != nil {
					t.Error(err)
				}
			}
			for i := 0; i < count; i++ {
				item, err := test.stack.Pop()
				if err != nil {
					t.Error(err)
				}
				fmt.Printf("item: %v\n", item)
			}
		})
	}
}

func TestStackIterator(t *testing.T) {
	const count = 4
	cases := []struct {
		name  string
		stack Stack[int]
	}{
		{
			name:  "FixedStack",
			stack: NewFixedStack[int](count),
		},
		{
			name:  "InfStack",
			stack: NewInfStack[int](),
		},
		{
			name:  "InfStack2",
			stack: NewInfStack2[int](),
		},
		{
			name:  "InfStack3",
			stack: NewInfStack3[int](),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for i := 0; i < count; i++ {
				err := c.stack.Push(i)
				if err != nil {
					t.Error(err)
				}
			}
			var iter gof.Iterator[int] = c.stack.Iterator()
			for {
				next, err := iter.Next()
				if errors.Is(err, gof.ErrIteratorNext) {
					break
				}
				if err != nil {
					t.Error(err)
				}
				fmt.Printf("next: %v\n", next)
			}
		})
	}
}
