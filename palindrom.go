package main 
import ("fmt")
//  Write a function that accepts a string s as input and return
//       true if the string is a palindrome. false otherwise. A palindrome
//       is a string that reads the same from left to right and right to
//       left (e.g. "dad", "malayalam" etc.)
func palindrome(s string)bool{
	reversed := ""
	for i := len(s) -1;i>=0; i-- {
    reversed += string(s[i])
	}
	return reversed == s
	
}
func main (){
	str := "malayalam"
	if palindrome(str){
	fmt.Println(str, "this is palindrom")
	}else{
		fmt.Println(str,"not a palindrom")
	}

}