package iteration

func Repeat(character string, count int) string {
	// var is used to declare a variable that can later be set := is just short
	// hand to declare and set a variable in one line.
	// var can also be used to declare functions, which will be covered later.
	var repeated string

	// Unlike other languages like C, Java, or JavaScript there are no parentheses
	// surrounding the three components of the for statement and the braces { }
	// are always required.
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
