package integers

import "fmt"

// Add takes two integers and returns the sum of them.
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(2, 2))
}
