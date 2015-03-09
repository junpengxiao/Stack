#Introduction

This is a simple stack based on slice written in go.

`NewStack(number uint) *stack` will return a stack. You should estimate `number`,  the max-number of item in stack. If items number exceeds that estimation, Stack will allocate more space automatically. `Len() int` will return the item number in stack. `Push(value interface{})` will save the item into stack. `Pop() (interface{}, error)` will pop the top item. If the stack is empty, it will return `ErrEmptyStack` error, which is also defined in this package. `Peek() (interface{}, error)` is similar with Pop.

**Notice** Since all return values are interface{}, you may need to use [type assertion](https://golang.org/doc/effective_go.html#interface_conversions).