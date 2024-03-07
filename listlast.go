package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	value := &NodeL{}

	for l.Head != nil {
		value = l.Head
		l.Head = l.Head.Next
	}

	return value.Data
}
