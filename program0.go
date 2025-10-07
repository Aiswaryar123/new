package main 
import "fmt"
func add(a,b int)(int,int,int){
	return a+b,a-b,a*b	
}
func main(){
fmt.Println(add(10,20))
}