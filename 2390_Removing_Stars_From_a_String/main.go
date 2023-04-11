package main

func removeStars(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		b := s[i]
		if b == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}
	return string(stack)
}

func main() {
	s := "leet**cod*e"
	s = removeStars(s)
	println(s)
}
