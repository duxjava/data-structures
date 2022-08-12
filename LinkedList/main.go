package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Value int
	Next  *Node
}

func (head *Node) Print() {
	for {
		if head == nil {
			return
		}

		fmt.Printf("%v\n", head.Value)
		head = head.Next
	}
}

func (head *Node) String() string {
	response := bytes.NewBufferString("")

	for {
		if head == nil {
			return response.String()
		}

		response.WriteString(strconv.Itoa(head.Value))
		head = head.Next
	}
}

func (head *Node) Add(value int) {
	for {
		if head.Next == nil {
			head.Next = &Node{Value: value}
			return
		}
		head = head.Next
	}
}

func (head *Node) Get(index int) (*Node, error) {
	i := 0

	if index < 0 {
		return nil, errors.New("index must be positive")
	}

	for {
		if i == index {
			return head, nil
		}

		if head.Next == nil {
			return nil, fmt.Errorf("not found index: %d; max index: %d", index, i)
		}

		i++
		head = head.Next
	}
}

func (head *Node) FindBy(value Node) (*Node, error) {
	for {
		if head.Next == nil {
			return nil, errors.New("Node not found")
		}

		if head.Value == value.Value {
			return head, nil
		}

		head = head.Next
	}
}

func (head *Node) DeleteAt(index int) (*Node, error) {

	if index < 0 {
		return nil, errors.New("index must be positive")
	}

	if index == 0 {
		removeHead := *head
		for {
			head.Value = head.Next.Value

			if head.Next.Next == nil {
				head.Next = nil
				break
			}

			head = head.Next
		}
		return &removeHead, nil
	}

	i := 0

	for {
		if head.Next == nil {
			return nil, fmt.Errorf("not found index: %d; max index: %d", index, i)
		}

		if i+1 == index {
			removeHead := head.Next
			head.Next = head.Next.Next
			return removeHead, nil
		}

		head = head.Next
		i++
	}
}

func (head *Node) InsertAt(index int, value Node) (*Node, error) {
	if index < 0 {
		return nil, errors.New("index must be positive")
	}

	if index == 0 {
		insertHead := *head
		for {
			oldValue := head.Value
			head.Value = value.Value
			value.Value = oldValue

			if head.Next == nil {
				head.Next = &Node{Value: value.Value}
				break
			}

			head = head.Next
		}
		return &insertHead, nil
	}

	i := 0

	for {
		if head.Next == nil {
			return nil, fmt.Errorf("not found index: %d; max index: %d", index, i)
		}

		if i+1 == index {
			value.Next = head.Next
			head.Next = &value
			return head, nil
		}

		head = head.Next
		i++
	}
}

func main() {

	head := Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	fmt.Println(head.String())

}
