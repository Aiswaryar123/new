// Write a function that accepts a sentence and returns how many words it contains.
package main 
import ("fmt"
        "strings")
func  sentence(s string)int{
  words := strings.Fields(s)
  
  return len(words)
}

func main(){

	fmt.Println(sentence("hai hellow how are tou"))
}