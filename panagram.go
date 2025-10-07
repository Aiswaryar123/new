
    // - Write a function that accepts a string s as input and returns
    //   true if s is a panagram. false otherwise. A panagram is a
    //   sentence which contains all letters of the alphabet (e.g. "the
    //   quick brown fox jumps over the lazy dog")

	package main
	import ("fmt"
	"strings")

	func panagram(s string)bool{
		s = strings.ToLowercase(s)
		letters := make(map[rune]bool)
		for _,ch =range s{
			if ch>= 'a' && ch <= 'z'{
			letter[ch] == true 
			}
		}
		return len(letters)== 26
	}
	func main(){
		fmt.Println(panagram("the
      quick brown fox jumps over the lazy dog")
)

	}