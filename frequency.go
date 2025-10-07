// Problem 4 - Frequency
//     - Write a function which accepts a string s as input and returns a
//       map whose keys are the characters in the string and whose values
//       are the number of time the character occurs. e.g. If "bobby" is
//       given as input, it will return a map with these values {b:3,
//       o:1, y:1}
package main
import ("fmt")
func Frequency(s string)map[string]int{
	freq := make(map[string]int)
	for i:=0; i< len(s);i++{
		ch := string(s[i])
		freq[ch]++
	}
	return freq

}
func main(){
	fmt.Println(Frequency("hubby"))

}