// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/base64"
	"fmt"
	"strings"
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
	// reverse string and replace : with space
	whatIsIt = strings.Replace(reverseString(string(sd)), ":", " ", -1)
	fmt.Println(whatIsIt)
}
