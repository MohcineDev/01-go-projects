package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	newNodeL := new(NodeL)

	newNodeL = &NodeL{Data: data}

	if l.Head == nil && l.Tail == nil {
		l.Head = newNodeL
		l.Tail = newNodeL
	} else {
		newNodeL.Next = l.Head
		l.Head = newNodeL
	}
}
