package go_even_odd

func EvenOdd(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

func ListEvenOdd(num int) []string {
	var list []string
	for i := 0; i < num; i++ {
		list = append(list, EvenOdd(i))
	}
	return list
}
