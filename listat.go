package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i <= pos; i++ {
		if i == pos {
			return l
		}
		if l != nil {
			l = l.Next
		}
	}

	return nil
}
