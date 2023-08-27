package mystrings

// reverses a string
// we need to capitilize the first letter of the function
// if we don't then we won't be able to access this function outside of the mystrings package

func Reverse(str string) string {
	result := ""
	for _, char := range str {
		result = string(char) + result
	}
	return result
}
