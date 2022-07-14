package main

import ( 
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {

	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(sd)

	s := strings.Split(whatIsIt, ":")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	for _, v := range s {

		fmt.Print(reverse(v) + " ")
	}
//Join us at LINE MAN Wongnai 

}
func reverse(s string) string {
	rns := []rune(s)  
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
 
		rns[i], rns[j] = rns[j], rns[i]
	} 
	return string(rns)
}
