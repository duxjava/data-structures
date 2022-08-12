package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	head := Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	if head.String() != "1234" {
		t.Errorf("Add = %s; should 1234", head.String())
	}
}

func TestGet(t *testing.T) {
	head := Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	node, _ := head.Get(2)

	_, err := head.Get(-1)

	if err.Error() != "index must be positive" {
		t.Errorf("Get incorrect err: %s; should 'index must be positive'", err)
	}

	if node.Value != 3 {
		t.Errorf("Get incorrect value: %d; should 3", node.Value)
	}

	_, err = head.Get(5)

	if err.Error() != "not found index: 5; max index: 3" {
		t.Errorf("Get incorrect err: %s; should 'not found index: 5; max index: 3'", err)
	}
}

func TestFindBy(t *testing.T) {
	head := Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	node, _ := head.FindBy(Node{Value: 2})

	if node.Value != 2 {
		t.Errorf("Get incorrect node: %d; should 2", node.Value)
	}

	_, err := head.FindBy(Node{Value: 5})

	if err.Error() != "Node not found" {
		t.Errorf("Get incorrect err: %s; should 'Node not found'", err)
	}
}

func TestDeleteAt(t *testing.T) {
	head := Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	_, err := head.DeleteAt(-1)

	if err.Error() != "index must be positive" {
		t.Errorf("Get incorrect err: %s; should 'index must be positive'", err)
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	removeHead, _ := head.DeleteAt(0)

	if removeHead.Value != 1 || head.String() != "234" {
		t.Errorf("DeleteAt = %s; should 234", head.String())
	}

	removeHead, _ = head.DeleteAt(0)

	if removeHead.Value != 2 || head.String() != "34" {
		t.Errorf("DeleteAt = %s; should 34", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	removeHead, _ = head.DeleteAt(2)

	if removeHead.Value != 3 || head.String() != "124" {
		t.Errorf("DeleteAt = %s; should 124", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	removeHead, _ = head.DeleteAt(3)

	if removeHead.Value != 4 || head.String() != "123" {
		t.Errorf("DeleteAt = %s; should 123", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	_, err = head.DeleteAt(5)

	if err.Error() != "not found index: 5; max index: 3" {
		t.Errorf("Get incorrect err: %s; should 'not found index: 5; max index: 3'", err)
	}
}

func TestInsertAt(t *testing.T) {
	head := Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	_, err := head.InsertAt(-1, Node{Value: 10})

	if err.Error() != "index must be positive" {
		t.Errorf("Get incorrect err: %s; should 'index must be positive'", err)
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	insertHead, _ := head.InsertAt(0, Node{Value: 10})

	if insertHead.Value == 10 || head.String() != "101234" {
		t.Errorf("DeleteAt = %s; should 101234", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	insertHead, _ = head.InsertAt(1, Node{Value: 10})

	if insertHead.Value == 10 || head.String() != "110234" {
		t.Errorf("DeleteAt = %s; should 110234", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	insertHead, _ = head.InsertAt(2, Node{Value: 10})

	if insertHead.Value == 10 || head.String() != "121034" {
		t.Errorf("DeleteAt = %s; should 121034", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	insertHead, _ = head.InsertAt(3, Node{Value: 10})

	if insertHead.Value == 10 || head.String() != "123104" {
		t.Errorf("DeleteAt = %s; should 123104", head.String())
	}

	head = Node{Value: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)

	_, err = head.InsertAt(5, Node{Value: 10})

	if err.Error() != "not found index: 5; max index: 3" {
		t.Errorf("Get incorrect err: %s; should 'not found index: 5; max index: 3'", err)
	}
}
