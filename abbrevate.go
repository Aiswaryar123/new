//  Write a function which takes a string s as input and returns the
//       short form of s by taking the first characters of all the words
//       which are in caps. e.g. if the input is "Indian Institute of
//       Management", it will return "IIM" (the "of" is skipped because
//       it doesn't start with a capital letter)
package main
import ("fmt"
        "strings")
func abbrevatestring(s string)string{
	word := strings.Split(s," ")
	result := ""
	for _,w:= range word{
		if w !="" && w[0]>='A' && w[0] <='Z'{	
		result += string(w[0])
	}}
	return result
	
}
func main(){

str := "Hamon Technologies"
fmt.Println("abbrevate string:",abbrevatestring(str))
}