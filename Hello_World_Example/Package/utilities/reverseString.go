// Package utilities defines the name of the package that is going to be the container for these methods
package utilities

// Reverse and returns its argument string reversed rune-wise left to right
/*
Takes in as input an argument of type string and returns a string
*/
func Reverse(inputString string) string {
	// Create an array of integers from the input string
	// The *rune* key is able to interchange between int32 and character values
	stringRune := []rune(inputString)
	// Reverse the array
	for i, j := 0, len(stringRune)-1; i < len(stringRune)/2; i, j = i+1, j-1 {
		stringRune[i], stringRune[j] = stringRune[j], stringRune[i]
	}

	return string(stringRune)
}
