package stack

import "testing"
import "fmt"
type t1 struct {
	num int
}

func TestBehaviour(t *testing.T){
	//push var, change outside var, check
	fmt.Println()
	tmp1 := t1{1}
	stack := NewStack(3)
	stack.Push(tmp1)
	tmp1.num = 10
	var check t1;
	result, err := stack.Pop()
	check = result.(t1)
	if check.num != 1 {
		t.Errorf("changes in outside behaviour unexcepted")
	}

	//push var, change inside var, check
	tmp1 = t1{1}
	if stack.Len() != 0 {
		t.Errorf("stack should be empty")
	}
	stack.Push(tmp1)
	result, _ = stack.Peek()
	result = t1{10}
	result, _ = stack.Pop()
	check = result.(t1)
	if check.num != 1 {
		t.Errorf("changes in inside behaviour unexcepted")
	}

	//check error
	result, err = stack.Pop()
	if result != nil && err != ErrEmptyStack {
		t.Errorf("empty stack behaviour unexcepted")	
	}

	//check adding more items
	for i:=0; i!=5; i++ {
		stack.Push(123)
	}
	for result, err := stack.Pop(); err != ErrEmptyStack; result, err = stack.Pop() {
		check2 := result.(int)
		if (check2 != 123) {
			t.Errorf("a mistake")
		}
	}
}
