// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/base64"
	"fmt"
)

func reverseString(text string) string {
	chars := []rune(text)
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, err := base64.StdEncoding.DecodeString(secret)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(string(sd))
	whatIsIt = reverseString(string(sd))
	fmt.Println(whatIsIt)
}
