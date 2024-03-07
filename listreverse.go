package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	current := l.Head
	l.Tail = current
	var prev *NodeL

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
