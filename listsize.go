package piscine

func ListSize(l *List) int {
	count := 1
	k := l.Head
	for k.Next != nil {
		count++
		k = k.Next
	}
	return count
}
